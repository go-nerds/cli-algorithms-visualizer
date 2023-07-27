package main

import (
	"time"
)

func selectionSortVisualizer(arr []int, delay time.Duration) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			visualizeIteration(arr, i, j, delay)
			if arr[j] < arr[min] {
				visualizeIteration(arr, i, j, delay)
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		visualizeIteration(arr, i, min, delay)
	}
}
