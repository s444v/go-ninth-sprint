package main

import (
	"slices"
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

func TestMaxChunk(t *testing.T) {
	for i := 1000; i <= 100_000_000; i = i * 10 {
		array := generateRandomElements(i)
		assert.Equal(t, maxChunks(array), slices.Max(array))
	}
	testsOk := []struct {
		array  []int
		answer int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{234}, 234},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, -5, -3}, -1},
		{[]int{11, 5, 7, 9}, 11},
	}
	for _, v := range testsOk {
		assert.Equal(t, v.answer, maxChunks(v.array))
	}
}
