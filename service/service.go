package service

import (
	"errors"

	"github.com/thomaspepio/rest-fizzbuzz/domain"
)

const (
	unableToComputeFizzBuzzError = "unable to compute fizzbuzz : "
)

// FizzBuzzRequest : a request for a FizzBuzz computation.
type FizzBuzzRequest struct {
	Limit  int    `json:"limit"`
	Fizzer int    `json:"fizzer"`
	Buzzer int    `json:"buzzer"`
	Fizz   string `json:"fizz"`
	Buzz   string `json:"buzz"`
}

// FizzBuzzResponse : type alias for a list of string.
type FizzBuzzResponse = []string

// ComputeFizzBuzz : given a request to compute FizzBuzz result on,
// either the request is well formed and we return a FizzBuzzResponse,
// or the request is incorrect and we return an error.
func ComputeFizzBuzz(request *FizzBuzzRequest) (FizzBuzzResponse, error) {
	if request == nil {
		return nil, errors.New(unableToComputeFizzBuzzError + "request is nil")
	}

	if request.Limit <= 0 {
		return nil, errors.New(unableToComputeFizzBuzzError + "limit is less than or equal to zero")
	}

	var result FizzBuzzResponse
	for i := 1; i <= request.Limit; i++ {
		res, err := domain.FizzBuzz(i, request.Fizzer, request.Buzzer, request.Fizz, request.Buzz)

		if err != nil {
			return nil, err
		}

		result = append(result, res)
	}

	return result, nil
}
