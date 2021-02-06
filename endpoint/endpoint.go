package endpoint

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/thomaspepio/rest-fizzbuzz/service"

	"github.com/gin-gonic/gin"
)

const (
	// URL for the fizzbuzz in V1
	v1FizzBuzz = "/1/fizzbuzz"

	// query parameter that sets the limit up to which we compute a fizzbuzz
	limitParam = "limit"

	// query parameter for the fizz divisor
	fizzerParam = "fizzer"

	// query parameter for the buzz divisor
	buzzerParam = "buzzer"

	// query parameter for the fizz parameter
	fizzParam = "fizz"

	// query parameter for the buzz parameter
	buzzParam = "buzz"

	// not a number error message
	paramNotANumber = "could not convert parameter %s to a number"
)

// Router : return the endpoints of the application
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET(v1FizzBuzz, func(context *gin.Context) {
		limit := context.Query(limitParam)
		fizzer := context.Query(fizzerParam)
		buzzer := context.Query(buzzerParam)
		fizz := context.Query(fizzParam)
		buzz := context.Query(buzzParam)

		if (limit == "") || (fizzer == "") || (buzzer == "") || (fizz == "") || (buzz == "") {
			context.JSON(http.StatusBadRequest, gin.H{"error": "at least one mandatory parameter is absent"})
		} else if !isANumber(limit) {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(paramNotANumber, limitParam)})
		} else if !isANumber(fizzer) {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(paramNotANumber, fizzerParam)})
		} else if !isANumber(buzzer) {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(paramNotANumber, buzzerParam)})
		} else {
			limitToInt, _ := strconv.Atoi(limit)
			fizzerToInt, _ := strconv.Atoi(fizzer)
			buzzerToInt, _ := strconv.Atoi(buzzer)
			fizzBuzzRequest := service.FizzBuzzRequest{limitToInt, fizzerToInt, buzzerToInt, fizz, buzz}
			result, err := service.ComputeFizzBuzz(&fizzBuzzRequest)

			if err != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"result": result})
			}
		}
	})

	return router
}

func isANumber(allgedNumber string) bool {
	_, err := strconv.Atoi(allgedNumber)
	return err == nil
}
