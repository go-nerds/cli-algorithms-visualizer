package main

import "time"

func bubbleSortVisualizer(arr []int, delay time.Duration) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			visualizeIteration(arr, j, j+1, delay)
			if arr[j] > arr[j+1] {
				visualizeIteration(arr, j, j+1, delay)
				arr[j], arr[j+1] = arr[j+1], arr[j]
				visualizeIteration(arr, j, j+1, delay)
			}
		}
	}
}
