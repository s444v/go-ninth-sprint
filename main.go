package main

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int, size)
	for i, _ := range array {
		array[i] = int(r.Int63())
	}
	return array
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) <= 0 {
		return 0
	}
	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	wg.Add(CHUNKS)
	lenOfSlices := len(data) / CHUNKS
	shift := len(data) - lenOfSlices*CHUNKS
	max := math.MinInt
	arrayOfMax := make([]int, CHUNKS)
	for i := range CHUNKS {
		left := i * lenOfSlices
		right := (i+1)*lenOfSlices + shift
		go func(left int, right int, arrayOfMax []int) {
			defer wg.Done()
			max = maximum(data[left:right])
			arrayOfMax[i] = max
		}(left, right, arrayOfMax)
	}
	wg.Wait()
	return maximum(arrayOfMax)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	array := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	t := time.Now()
	max := maximum(array)
	elapsed := time.Since(t)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())

	t = time.Now()
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	max = maxChunks(array)
	elapsed = time.Since(t)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Microseconds())
}
