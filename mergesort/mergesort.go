package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(sortedSlice1 []int, sortedSlice2 []int) []int {
	mergedSlice := make([]int, 0, len(sortedSlice1)+len(sortedSlice2))
	var index1, index2 int
	for index1 < len(sortedSlice1) && index2 < len(sortedSlice2) {
		if sortedSlice1[index1] < sortedSlice2[index2] {
			mergedSlice = append(mergedSlice, sortedSlice1[index1])
			index1++
		} else {
			mergedSlice = append(mergedSlice, sortedSlice2[index2])
			index2++
		}
	}
	mergedSlice = append(mergedSlice, sortedSlice1[index1:]...)
	mergedSlice = append(mergedSlice, sortedSlice2[index2:]...)
	return mergedSlice
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	mid := len(items) / 2
	first := mergeSort(items[:mid])
	second := mergeSort(items[mid:])
	return merge(first, second)
}

func main() {
	const nElements = 10000
	unsortedSlice := make([]int, nElements)

	// generate numbers
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := 0; i < nElements; i++ {
		unsortedSlice[i] = rng.Intn(10000)
	}

	sorted := mergeSort(unsortedSlice)

	fmt.Println(sorted[:100])
}
