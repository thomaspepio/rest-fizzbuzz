package endpoint

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzzEndpoint_OK(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?limit=15&fizzer=3&buzzer=15&fizz=fizz&buzz=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"result\":[\"1\",\"2\",\"fizz\",\"4\",\"5\",\"fizz\",\"7\",\"8\",\"fizz\",\"10\",\"11\",\"fizz\",\"13\",\"14\",\"fizzbuzz\"]}", w.Body.String())
}

func Test_FizzBuzzEndpoint_BadRequest_WhenLimitParameterIsAbsent(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?fizzer=fizzer&buzzer=buzzer&fizz=fizz&buzz=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
}

func Test_FizzBuzzEndpoint_BadRequest_WhenFizzerParameterIsAbsent(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?limit=limit&buzzer=buzzer&fizz=fizz&buzz=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
}

func Test_FizzBuzzEndpoint_BadRequest_WhenBuzzerParameterIsAbsent(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?limit=limit&fizzer=fizzer&fizz=fizz&buzz=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
}

func Test_FizzBuzzEndpoint_BadRequest_WhenFizzParameterIsAbsent(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?limit=limit&fizzer=fizzer&buzzer=buzzer&buzz=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
}

func Test_FizzBuzzEndpoint_BadRequest_WhenBuzzParameterIsAbsent(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?limit=limit&fizzer=fizzer&buzzer=buzzer&fizz=fizz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"at least one mandatory parameter is absent\"}", w.Body.String())
}

func Test_FizzBuzzEndpoint_BadRequest_WhenLimitIsNotANumber(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?limit=not_a_number&fizzer=3&buzzer=5&fizz=fizz&buzz=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"could not convert parameter limit to a number\"}", w.Body.String())
}

func Test_FizzBuzzEndpoint_BadRequest_WhenFizzerIsNotANumber(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?limit=10&fizzer=not_a_number&buzzer=5&fizz=fizz&buzz=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"could not convert parameter fizzer to a number\"}", w.Body.String())
}

func Test_FizzBuzzEndpoint_BadRequest_WhenBuzzerIsNotANumber(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/1/fizzbuzz?limit=10&fizzer=3&buzzer=not_a_number&fizz=fizz&buzz=buzz", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "{\"error\":\"could not convert parameter buzzer to a number\"}", w.Body.String())
}
