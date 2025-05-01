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
	tests := []struct {
		array  []int
		answer int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{234}, 234},
		{[]int{0, 0, 0}, 0},
		{[]int{10, 5, 7, 9}, 10},
	}
	for _, v := range tests {
		assert.Equal(t, v.answer, maximum(v.array))
	}
}

func TestMaxChunkWithTestTable(t *testing.T) {
	tests := []struct {
		array  []int
		answer int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{234}, 234},
		{[]int{0, 0, 0}, 0},
		{[]int{11, 5, 7, 9}, 11},
	}
	for _, v := range tests {
		assert.Equal(t, v.answer, maxChunks(v.array))
	}
}

func maxChunksTEST(data []int, ch int) int {
	if len(data) <= 0 {
		return 0
	}
	wg.Add(ch)
	lenOfSlices := len(data) / ch
	shift := len(data) % ch
	arrayOfMax := make([]int, ch)
	right := 0
	for i := range ch {
		left := right
		right = (i+1)*lenOfSlices + shift
		go func(left int, right int, arrayOfMax []int) {
			defer wg.Done()
			max := maximum(data[left:right])
			arrayOfMax[i] = max
		}(left, right, arrayOfMax)
	}
	wg.Wait()
	return maximum(arrayOfMax)
}

func TestMaxChunkWithRandomElements(t *testing.T) {
	chunks := 2
	for i := 1000; i <= 100_000_000; i = i * 10 {
		array := generateRandomElements(i)
		assert.Equal(t, maxChunksTEST(array, chunks), slices.Max(array))
		chunks++
	}
}
