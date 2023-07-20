package main

func combSortVisualizer(arr []int) {
	n := len(arr)
	gap := n
	swapped := true

	for gap > 1 || swapped {
		gap = getNextGap(gap)
		swapped = false

		for i := 0; i < n-gap; i++ {
			visualizeIteration(White, arr, i, i+gap)
			if arr[i] > arr[i+gap] {
				visualizeIteration(Red, arr, i, i+gap)
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				visualizeIteration(Green, arr, i, i+gap)
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
