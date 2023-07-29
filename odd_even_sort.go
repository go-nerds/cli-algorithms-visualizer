package main

import "time"

func oddEvenSortVisualizer(arr []int, displayType string, delay time.Duration) {
	n := len(arr)
	sorted := false

	for !sorted {
		sorted = true
		for i := 1; i <= n-2; i += 2 {
			visualizeIteration(arr, i, i+1, displayType, delay)
			if arr[i] > arr[i+1] {
				visualizeIteration(arr, i, i+1, displayType, delay)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				visualizeIteration(arr, i, i+1, displayType, delay)
				sorted = false
			}
		}

		for i := 0; i <= n-2; i += 2 {
			visualizeIteration(arr, i, i+1, displayType, delay)
			if arr[i] > arr[i+1] {
				visualizeIteration(arr, i, i+1, displayType, delay)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				visualizeIteration(arr, i, i+1, displayType, delay)
				sorted = false
			}
		}
	}
}
