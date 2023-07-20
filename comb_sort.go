package main

import "time"

func combSortVisualizer(arr []int, delay time.Duration) {
	n := len(arr)
	gap := n
	swapped := true

	for gap > 1 || swapped {
		gap = getNextGap(gap)
		swapped = false

		for i := 0; i < n-gap; i++ {
			visualizeIteration(White, arr, i, i+gap, delay)
			if arr[i] > arr[i+gap] {
				visualizeIteration(LightYellow, arr, i, i+gap, delay)
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				visualizeIteration(LightBlue, arr, i, i+gap, delay)
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
