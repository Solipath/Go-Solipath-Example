package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3ReturnsFizz(t *testing.T) {
	assert.Equal(t, "Fizz", fizzbuzz(3))
}

func Test5ReturnsBuzz(t *testing.T) {
	assert.Equal(t, "Buzz", fizzbuzz(5))
}

func Test6ReturnsFizz(t *testing.T) {
	assert.Equal(t, "Fizz", fizzbuzz(6))
}

func Test1Returns1(t *testing.T) {
	assert.Equal(t, "1", fizzbuzz(1))
}

func Test10ReturnsBuzz(t *testing.T) {
	assert.Equal(t, "Buzz", fizzbuzz(10))
}

func Test15ReturnsFizzBuzz(t *testing.T) {
	assert.Equal(t, "FizzBuzz", fizzbuzz(15))
}

func Test30ReturnsFizzBuzz(t *testing.T) {
	assert.Equal(t, "FizzBuzz", fizzbuzz(30))
}
