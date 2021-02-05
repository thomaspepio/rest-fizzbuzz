package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomaspepio/rest-fizzbuzz/constant"
)

func Test_FizzBuzz_ShouldReturnError_WithNLessThanOrEqualToZero(t *testing.T) {
	_, err := FizzBuzz(0, 0, 0, constant.Fizz, constant.Buzz)
	assert.Error(t, err, "FizzBuzz should not work on 0")

	_, err = FizzBuzz(-1, 0, 0, constant.Fizz, constant.Buzz)
	assert.Error(t, err, "FizzBuzz should not work on n < 0")
}

func Test_FizzBuzz_ShouldReturnFizz_WhenNDivisibleOnlyByFizzzer(t *testing.T) {
	actual, _ := FizzBuzz(9, 3, 5, constant.Fizz, constant.Buzz)
	assert.Equal(t, constant.Fizz, actual, "in this context : 9 is only divisible by 3, FizzBuzz should've returned \"fizz\"")
}

func Test_FizzBuzz_ShouldReturnBuzz_WhenNDivisibleOnlyByBuzzer(t *testing.T) {
	actual, _ := FizzBuzz(25, 3, 5, constant.Fizz, constant.Buzz)
	assert.Equal(t, constant.Buzz, actual, "in this context : 25 is only divisible by 5. FizzBuzz should've returned \"buzz\"")
}

func Test_FizzBuzz_ShouldReturnFizzBuzz_WhenNDivisibleByBoth(t *testing.T) {
	actual, _ := FizzBuzz(15, 3, 5, constant.Fizz, constant.Buzz)
	assert.Equal(t, constant.Fizzbuzz, actual, "15 is divisible by both 3 and 5. FizzBuzz should've returned \"fizzbuzz\"")
}

func Test_FizzBuzz_ShouldReturnN_WhenDivisibleByNeither(t *testing.T) {
	actual, _ := FizzBuzz(1, 3, 5, constant.Fizz, constant.Buzz)
	assert.Equal(t, "1", actual, "1 is neither divisible by 3 nor by 5. FizzBuzz should've returned \"1\"")
}
