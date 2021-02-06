package endpoint

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/thomaspepio/rest-fizzbuzz/service"

	"github.com/gin-gonic/gin"
)

const (
	// URL for the fizzbuzz API in V1
	v1FizzBuzz = "/1/fizzbuzz"

	// URL for the stats API in V1
	v1Stats = "/1/stats"

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

// Router : return the endpoints of the application
func SetupRouter(urlCounter URLCounter) *gin.Engine {
	router := gin.Default()

	router.GET(v1FizzBuzz, func(context *gin.Context) {
		limit := context.Query(LimitParam)
		fizzer := context.Query(FizzerParam)
		buzzer := context.Query(BuzzerParam)
		fizz := context.Query(FizzParam)
		buzz := context.Query(BuzzParam)

		if (limit == "") || (fizzer == "") || (buzzer == "") || (fizz == "") || (buzz == "") {
			context.JSON(http.StatusBadRequest, gin.H{"error": "at least one mandatory parameter is absent"})
		} else if !isANumber(limit) {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(paramNotANumber, LimitParam)})
		} else if !isANumber(fizzer) {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(paramNotANumber, FizzerParam)})
		} else if !isANumber(buzzer) {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(paramNotANumber, BuzzerParam)})
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
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"result": result})
			}
		}
	})

	router.GET(v1Stats, func(context *gin.Context) {
		var topRequest service.FizzBuzzRequest
		count := 0

		for key, value := range urlCounter {
			if value > count {
				count = value
				topRequest = key
			}
		}

		context.JSON(http.StatusOK, gin.H{"mostRequested": topRequest, "count": count})
	})

	return router
}

func isANumber(allgedNumber string) bool {
	_, err := strconv.Atoi(allgedNumber)
	return err == nil
}
