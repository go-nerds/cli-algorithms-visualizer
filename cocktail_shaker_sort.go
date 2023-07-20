package main

func cocktailShakerSortVisualizer(arr []int) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		for i := left; i < right; i++ {
			visualizeIteration(White, arr, i, i+1)
			if arr[i] > arr[i+1] {
				visualizeIteration(Red, arr, i, i+1)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				visualizeIteration(Green, arr, i, i+1)
			}
		}
		right--

		for i := right; i > left; i-- {
			visualizeIteration(White, arr, i-1, i)
			if arr[i-1] > arr[i] {
				visualizeIteration(Red, arr, i-1, i)
				arr[i], arr[i-1] = arr[i-1], arr[i]
				visualizeIteration(Green, arr, i-1, i)
			}
		}
		left++
	}
}
