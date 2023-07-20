package main

import "time"

func gnomeSortVisualizer(arr []int, delay time.Duration) {
	i := 1

	for i < len(arr) {
		if i == 0 || arr[i] >= arr[i-1] {
			visualizeIteration(White, arr, i-1, i, delay)
			i++
		} else {
			visualizeIteration(Red, arr, i-1, i, delay)
			arr[i], arr[i-1] = arr[i-1], arr[i]
			visualizeIteration(Green, arr, i-1, i, delay)
			i--
			if i == 0 {
				i = 1
			}
		}
	}
}
