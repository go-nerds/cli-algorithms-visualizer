package main

import "time"

func combSortVisualizer(arr []int, displayType string, delay time.Duration) {
	n := len(arr)
	gap := n
	swapped := true

	for gap > 1 || swapped {
		gap = getNextGap(gap)
		swapped = false

		for i := 0; i < n-gap; i++ {
			visualizeIteration(arr, i, i+gap, displayType, delay)
			if arr[i] > arr[i+gap] {
				visualizeIteration(arr, i, i+gap, displayType, delay)
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				visualizeIteration(arr, i, i+gap, displayType, delay)
				swapped = true
			}
		}
	}
}

func getNextGap(gap int) int {
	gap = (gap * 10) / 13
	if gap < 1 {
		return 1
	}
	return gap
}
