package endpoint

import (
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/thomaspepio/rest-fizzbuzz/constant"
	"github.com/thomaspepio/rest-fizzbuzz/service"
)

const (
	correctURLPattern = "/1/fizzbuzz?" + LimitParam + "=%d&" + FizzerParam + "=%d&" + BuzzerParam + "=%d&" + FizzParam + "=%s&" + BuzzParam + "=%s"
)

func Test_FizzBuzzEndpoint_OK(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := fmt.Sprintf(correctURLPattern, 15, 3, 5, "fizz", "buzz")
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"result\":[\"1\",\"2\",\"fizz\",\"4\",\"buzz\",\"fizz\",\"7\",\"8\",\"fizz\",\"buzz\",\"11\",\"fizz\",\"13\",\"14\",\"fizzbuzz\"]}", w.Body.String())
	assert.Equal(t, 1, urlCounter[service.FizzBuzzRequest{15, 3, 5, constant.Fizz, constant.Buzz}], fmt.Sprintf("%s should have been visited once", url))
}

func Test_FizzBuzzEndpoint_ShouldCorrectlyCountURLs(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()

	url1 := fmt.Sprintf(correctURLPattern, 15, 3, 5, "fizz", "buzz")
	fizzBuzzRequest, _ := http.NewRequest("GET", url1, nil)
	router.ServeHTTP(w, fizzBuzzRequest)
	assert.Equal(t, 1, urlCounter[service.FizzBuzzRequest{15, 3, 5, constant.Fizz, constant.Buzz}], fmt.Sprintf("%s should have been visited once", url1))
	router.ServeHTTP(w, fizzBuzzRequest)
	assert.Equal(t, 2, urlCounter[service.FizzBuzzRequest{15, 3, 5, constant.Fizz, constant.Buzz}], fmt.Sprintf("%s should have been visited once", url1))

	url2 := fmt.Sprintf(correctURLPattern, 5, 2, 3, "fizz", "buzz")
	fizzBuzzRequest, _ = http.NewRequest("GET", url2, nil)
	router.ServeHTTP(w, fizzBuzzRequest)
	assert.Equal(t, 1, urlCounter[service.FizzBuzzRequest{5, 2, 3, constant.Fizz, constant.Buzz}], fmt.Sprintf("%s should have been visited once", url2))

	w = httptest.NewRecorder()
	statsRequest, _ := http.NewRequest("GET", "/1/stats", nil)
	router.ServeHTTP(w, statsRequest)
	assert.Equal(t, "{\"count\":2,\"mostRequested\":{\"Limit\":15,\"Fizzer\":3,\"Buzzer\":5,\"Fizz\":\"fizz\",\"Buzz\":\"buzz\"}}", w.Body.String(), "uh oh")
}

func Test_FizzBuzzEndpoint_BadRequest_WhenLimitParameterIsAbsent(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := "/1/fizzbuzz?fizzer=fizzer&buzzer=buzzer&fizz=fizz&buzz=buzz"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
	assert.Equal(t, 0, len(urlCounter), fmt.Sprintf("%s should note have been visited", url))
}

func Test_FizzBuzzEndpoint_BadRequest_WhenFizzerParameterIsAbsent(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := "/1/fizzbuzz?limit=limit&buzzer=buzzer&fizz=fizz&buzz=buzz"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
	assert.Equal(t, 0, len(urlCounter), fmt.Sprintf("%s should not have been visited", url))
}

func Test_FizzBuzzEndpoint_BadRequest_WhenBuzzerParameterIsAbsent(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := "/1/fizzbuzz?limit=limit&fizzer=fizzer&fizz=fizz&buzz=buzz"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
	assert.Equal(t, 0, len(urlCounter), fmt.Sprintf("%s should not have been visited", url))
}

func Test_FizzBuzzEndpoint_BadRequest_WhenFizzParameterIsAbsent(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := "/1/fizzbuzz?limit=limit&fizzer=fizzer&buzzer=buzzer&buzz=buzz"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
	assert.Equal(t, 0, len(urlCounter), fmt.Sprintf("%s should not have been visited", url))
}

func Test_FizzBuzzEndpoint_BadRequest_WhenBuzzParameterIsAbsent(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := "/1/fizzbuzz?limit=limit&fizzer=fizzer&buzzer=buzzer&fizz=fizz"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
	assert.Equal(t, 0, len(urlCounter), fmt.Sprintf("%s should not have been visited", url))
}

func Test_FizzBuzzEndpoint_BadRequest_WhenLimitIsNotANumber(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := "/1/fizzbuzz?limit=not_a_number&fizzer=3&buzzer=5&fizz=fizz&buzz=buzz"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"could not convert parameter limit to a number\"}", w.Body.String())
	assert.Equal(t, 0, len(urlCounter), fmt.Sprintf("%s should not have been visited", url))
}

func Test_FizzBuzzEndpoint_BadRequest_WhenFizzerIsNotANumber(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := "/1/fizzbuzz?limit=10&fizzer=not_a_number&buzzer=5&fizz=fizz&buzz=buzz"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"could not convert parameter fizzer to a number\"}", w.Body.String())
	assert.Equal(t, 0, len(urlCounter), fmt.Sprintf("%s should not have been visited", url))
}

func Test_FizzBuzzEndpoint_BadRequest_WhenBuzzerIsNotANumber(t *testing.T) {
	urlCounter := emptyURLCounter()
	router := SetupRouter(urlCounter)
	w := httptest.NewRecorder()
	url := "/1/fizzbuzz?limit=10&fizzer=3&buzzer=not_a_number&fizz=fizz&buzz=buzz"
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"could not convert parameter buzzer to a number\"}", w.Body.String())
	assert.Equal(t, 0, len(urlCounter), fmt.Sprintf("%s should not have been visited", url))
}

func emptyURLCounter() URLCounter {
	return map[service.FizzBuzzRequest]int{}
}
