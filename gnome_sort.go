package main

import "time"

func gnomeSortVisualizer(arr []int, displayType string, delay time.Duration) {
	i := 1

	for i < len(arr) {
		if i == 0 || arr[i] >= arr[i-1] {
			visualizeIteration(arr, i-1, i, displayType, delay)
			i++
		} else {
			visualizeIteration(arr, i-1, i, displayType, delay)
			arr[i], arr[i-1] = arr[i-1], arr[i]
			visualizeIteration(arr, i-1, i, displayType, delay)
			i--
			if i == 0 {
				i = 1
			}
		}
	}
}
