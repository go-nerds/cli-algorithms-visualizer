package main

import "time"

func cocktailShakerSortVisualizer(arr []int, displayType string, delay time.Duration) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		for i := left; i < right; i++ {
			visualizeIteration(arr, i, i+1, displayType, delay)
			if arr[i] > arr[i+1] {
				visualizeIteration(arr, i, i+1, displayType, delay)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				visualizeIteration(arr, i, i+1, displayType, delay)
			}
		}
		right--

		for i := right; i > left; i-- {
			visualizeIteration(arr, i-1, i, displayType, delay)
			if arr[i-1] > arr[i] {
				visualizeIteration(arr, i-1, i, displayType, delay)
				arr[i], arr[i-1] = arr[i-1], arr[i]
				visualizeIteration(arr, i-1, i, displayType, delay)
			}
		}
		left++
	}
}
