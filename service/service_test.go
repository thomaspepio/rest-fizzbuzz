package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomaspepio/rest-fizzbuzz/constant"
)

func Test_FizzBuzzService_ShouldReturneError_WhenRequestIsNil(t *testing.T) {
	_, err := ComputeFizzBuzz(nil)
	assert.Error(t, err, "A nil request should result in an error")
}

func Test_FizzBuzzService_ShouldReturnError_WhenNumberIsLesThanOrEqualZero(t *testing.T) {
	_, err := ComputeFizzBuzz(&FizzBuzzRequest{0, 1, 2, constant.Fizz, constant.Buzz})
	assert.Error(t, err, "A FizzBuzzRequest with number = 0 should return an error")

	_, err = ComputeFizzBuzz(&FizzBuzzRequest{-1, 1, 2, constant.Fizz, constant.Buzz})
	assert.Error(t, err, "A FizzBuzzRequest with number < 0 should return an error")
}

func Test_FizzBuzzService_ShouldReturnAResponse_WhenNumberIsValid(t *testing.T) {
	actual, _ := ComputeFizzBuzz(&FizzBuzzRequest{15, 3, 5, constant.Fizz, constant.Buzz})

	assert.Equal(t,
		[]string{"1", "2", constant.Fizz, "4", constant.Buzz, constant.Fizz, "7", "8", constant.Fizz, constant.Buzz, "11", constant.Fizz, "13", "14", constant.Fizzbuzz},
		actual,
		"wrong fizzbuzz computation")
}
