package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum_BasicAddition(t *testing.T) {
	result := Sum(2, 3)
	assert.Equal(t, 5, result)
}

func TestSum_ZeroValues(t *testing.T) {
	result := Sum(0, 0)
	assert.Equal(t, 0, result)
}

func TestSum_NegativeNumbers(t *testing.T) {
	result := Sum(-5, -3)
	assert.Equal(t, -8, result)
}

func TestSum_PositiveAndNegative(t *testing.T) {
	result := Sum(10, -7)
	assert.Equal(t, 3, result)
}

func TestSum_LargeNumbers(t *testing.T) {
	result := Sum(1000000, 2000000)
	assert.Equal(t, 3000000, result)
}

func TestSum_ZeroWithPositive(t *testing.T) {
	result := Sum(0, 5)
	assert.Equal(t, 5, result)
}

func TestSum_ZeroWithNegative(t *testing.T) {
	result := Sum(0, -5)
	assert.Equal(t, -5, result)
}

func TestSum_Commutative(t *testing.T) {
	result1 := Sum(7, 3)
	result2 := Sum(3, 7)
	assert.Equal(t, result1, result2)
}

func TestSum_MaxInt(t *testing.T) {
	result := Sum(2147483647, 0)
	assert.Equal(t, 2147483647, result)
}

func TestSum_MinInt(t *testing.T) {
	result := Sum(-2147483648, 0)
	assert.Equal(t, -2147483648, result)
}
