package endpoint

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/thomaspepio/rest-fizzbuzz/endpoint/docs" // import of swagger docs : see README.md and make doc
	"github.com/thomaspepio/rest-fizzbuzz/service"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	v1 = "/1"

	// URL for the fizzbuzz API in V1
	fizzBuzzURL = "/fizzbuzz"

	// URL for the stats API in V1
	statsURL = "/stats"

	// LimitParam : query parameter that sets the limit up to which we compute a fizzbuzz
	LimitParam = "limit"

	// FizzerParam : query parameter for the fizz divisor
	FizzerParam = "fizzer"

	// BuzzerParam : query parameter for the buzz divisor
	BuzzerParam = "buzzer"

	// FizzParam : query parameter for the fizz parameter
	FizzParam = "fizz"

	// BuzzParam : query parameter for the buzz parameter
	BuzzParam = "buzz"

	// not a number error message
	paramNotANumber = "could not convert parameter %s to a number"
)

// URLCounter : type alias for a map service.FizzBuzzRequest => int
type URLCounter = map[service.FizzBuzzRequest]int

type EndpointError struct {
	ErrorMessage string `json:"errorMessage"`
}

// Router : return the endpoints of the application
// @title Fizzbuzz API
// @version 1.0
// @description An API that computes fizzbuzzes
// BasePath /1
func SetupRouter(port string, urlCounter URLCounter) *gin.Engine {
	router := gin.Default()

	v1 := router.Group(v1)
	{
		v1.GET(fizzBuzzURL, V1FizzBuzz(urlCounter))
		v1.GET(statsURL, V1Stats(urlCounter))
	}

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", port))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

// V1FizzBuzz : returns a gin.HandlerFunc handler for FizzBuzz API calls
// @Summary Computes Fizzbuzz
// @Description get fizzbuzz computation
// @Accept json
// @Produce json
// @Param limit query int true "Limit up to which fizzbuzz is computed"
// @Param fizzer query int true "First divisor : all its multiples will be replaced by @fizz"
// @Param buzzer query int true "Second divisor: all its multiples will be replaced by @buzz"
// @Param fizz query string true "Output for all numbers divisible by @fizzer (numbers divisible by both will output as the concatenation of @fizz@uzz)"
// @Param buzz query string true "Output for all numbers divisible by @buzzer (numbers divisible by both will output as the concatenation of @fizz@uzz)"
// @Success 200 {array} string
// @Failure 400 {object} EndpointError
// @Failure 500 {object} EndpointError
// @Router /1/fizzbuzz [get]
func V1FizzBuzz(urlCounter URLCounter) gin.HandlerFunc {
	return func(context *gin.Context) {
		limit := context.Query(LimitParam)
		fizzer := context.Query(FizzerParam)
		buzzer := context.Query(BuzzerParam)
		fizz := context.Query(FizzParam)
		buzz := context.Query(BuzzParam)

		if (limit == "") || (fizzer == "") || (buzzer == "") || (fizz == "") || (buzz == "") {
			context.JSON(http.StatusBadRequest, EndpointError{"at least one mandatory parameter is absent"})
		} else if !isANumber(limit) {
			context.JSON(http.StatusBadRequest, EndpointError{fmt.Sprintf(paramNotANumber, LimitParam)})
		} else if !isANumber(fizzer) {
			context.JSON(http.StatusBadRequest, EndpointError{fmt.Sprintf(paramNotANumber, FizzerParam)})
		} else if !isANumber(buzzer) {
			context.JSON(http.StatusBadRequest, EndpointError{fmt.Sprintf(paramNotANumber, BuzzerParam)})
		} else {
			limitToInt, _ := strconv.Atoi(limit)
			fizzerToInt, _ := strconv.Atoi(fizzer)
			buzzerToInt, _ := strconv.Atoi(buzzer)
			fizzBuzzRequest := service.FizzBuzzRequest{limitToInt, fizzerToInt, buzzerToInt, fizz, buzz}

			count, urlExists := urlCounter[fizzBuzzRequest]
			if urlExists {
				urlCounter[fizzBuzzRequest] = count + 1
			} else {
				urlCounter[fizzBuzzRequest] = 1
			}

			result, err := service.ComputeFizzBuzz(&fizzBuzzRequest)

			if err != nil {
				context.JSON(http.StatusInternalServerError, EndpointError{err.Error()})
			} else {
				context.JSON(http.StatusOK, result)
			}
		}
	}
}

type StatsResponse struct {
	Request service.FizzBuzzRequest `json:"request"`
	Count   int                     `json:"count"`
}

// V1Stats : returns statistics about the most popular searches
// @Summary Statistics endpoint
// @Description returns statistics about the most popular searches
// @Produce json
// @Success 200 {object} StatsResponse
// @Router /1/stats [get]
func V1Stats(urlCounter URLCounter) gin.HandlerFunc {
	return func(context *gin.Context) {
		var topRequest service.FizzBuzzRequest
		count := 0

		for key, value := range urlCounter {
			if value > count {
				count = value
				topRequest = key
			}
		}

		context.JSON(http.StatusOK, StatsResponse{topRequest, count})
	}
}

func isANumber(allgedNumber string) bool {
	_, err := strconv.Atoi(allgedNumber)
	return err == nil
}
