package domain

import (
	"errors"
	"strconv"
)

// FizzBuzz : a generalized FizzBuzz implementation
// In the original problem a number is tested to be divisible by either 3, 5, both or none ;
// and as a result fizzbuzz returns respectively "fizz", "buzz", "fizzbuzz" or the number itself as a string.
//
// This functions does the same, except the fixed values (divisors and corresponding results) are parameters.
func FizzBuzz(number, fizzer, buzzer int, fizz, buzz string) (string, error) {
	if number <= 0 {
		return "", errors.New("number cannot be less than or equal to zero")
	}

	if (number%fizzer == 0) && (number%buzzer == 0) {
		return fizz + buzz, nil
	}

	if number%fizzer == 0 {
		return fizz, nil
	}

	if number%buzzer == 0 {
		return buzz, nil
	}

	return strconv.Itoa(number), nil
}
