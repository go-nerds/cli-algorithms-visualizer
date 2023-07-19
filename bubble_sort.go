package main

func bubbleSortVisualizer(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			visualizeIteration(White, arr, j, j+1)
			if arr[j] > arr[j+1] {
				visualizeIteration(Red, arr, j, j+1)
				arr[j], arr[j+1] = arr[j+1], arr[j]
				visualizeIteration(Green, arr, j, j+1)
			}
		}
	}
}
