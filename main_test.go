package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		actLen int
		expLen int
		isNil  bool
	}{
		{1, 1, false},
		{3, 3, false},
		{5, 5, false},
		{0, 0, true},
		{-1, 0, true},
	}
	for _, v := range tests {
		array := generateRandomElements(v.actLen)
		assert.Equal(t, v.expLen, len(array))
		assert.Equal(t, v.isNil, array == nil)

	}
}

func TestMaximum(t *testing.T) {
	testsOk := []struct {
		array  []int
		answer int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{234}, 234},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, -5, -3}, -1},
		{[]int{10, 5, 7, 9}, 10},
	}
	for _, v := range testsOk {
		assert.Equal(t, v.answer, maximum(v.array))
	}
}
