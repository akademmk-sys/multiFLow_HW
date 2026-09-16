package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}
	sl := make([]int, size)
	for i := 0; i < size; i++ {
		sl[i] = int(rand.Int63())
	}
	return sl
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	subSlLen := len(data) / CHUNKS
	maxes := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		openSlice := i * subSlLen
		closeSlice := subSlLen * (i + 1)

		if i == CHUNKS-1 {
			closeSlice = len(data)
		}

		sl := data[openSlice:closeSlice]

		wg.Add(1)
		go func(idx int, slice []int) {
			defer wg.Done()
			maxes[idx] = maximum(slice)
		}(i, sl)
	}
	wg.Wait()

	goal := maximum(maxes)
	return goal
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	holyCrapSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(holyCrapSlice)
	elapsed := time.Duration(time.Since(start).Microseconds())

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(holyCrapSlice)
	elapsed = time.Duration(time.Since(start).Microseconds())

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
