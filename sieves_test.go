package primes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var primesUnder100 = []int{0, 1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

func TestEratosthenes(t *testing.T) {
	assert.Equal(t, primesUnder100, Eratosthenes(100))
}
