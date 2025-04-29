package main

import (
	"fmt"
	"math"
	"math/rand"
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
	max := math.MinInt
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	wg.Add(CHUNKS)
	lenOfSlices := len(data) / CHUNKS
	max := math.MinInt
	arrayOfMax := make([]int, CHUNKS)
	for i := range CHUNKS {
		go func(i int, arrayOfMax []int) {
			defer wg.Done()
			max = maximum(data[i*lenOfSlices : (i+1)*lenOfSlices])
			arrayOfMax[i] = max
		}(i, arrayOfMax)
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
