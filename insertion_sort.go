package main

import (
	"time"
)

func insertionSortVisualizer(arr []int, delay time.Duration) {
	for i := 1; i < len(arr); i++ {
		j := i
		visualizeIteration(arr, j-1, j, delay)
		for j > 0 && arr[j] < arr[j-1] {
			visualizeIteration(arr, j-1, j, delay)
			arr[j], arr[j-1] = arr[j-1], arr[j]
			visualizeIteration(arr, j-1, j, delay)
			j--
		}
	}
}
