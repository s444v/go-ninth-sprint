package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	testsOk := []struct {
		actLen int
		expLen int
	}{
		{1, 1},
		{3, 3},
		{101, 101},
	}
	for _, v := range testsOk {
		assert.Equal(t, v.expLen, len(generateRandomElements(v.actLen)))
	}
	assert.Nil(t, generateRandomElements(0))
	assert.Nil(t, generateRandomElements(-1))

}

func TestMaximum(t *testing.T) {
	testsOk := []struct {
		array  []int
		answer int
	}{
		{nil, math.MinInt},
		{[]int{}, math.MinInt},
		{[]int{234}, 234},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, -5, -3}, -1},
		{[]int{10, 5, 7, 9}, 10},
	}
	for _, v := range testsOk {
		assert.Equal(t, v.answer, maximum(v.array))
	}
}
