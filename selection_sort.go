package main

import (
	"time"
)

func selectionSortVisualizer(arr []int, displayType string, delay time.Duration) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			visualizeIteration(arr, i, j, displayType, delay)
			if arr[j] < arr[min] {
				visualizeIteration(arr, i, j, displayType, delay)
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		visualizeIteration(arr, i, min, displayType, delay)
	}
}
