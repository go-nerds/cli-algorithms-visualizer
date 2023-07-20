package main

func gnomeSortVisualizer(arr []int) {
	i := 1

	for i < len(arr) {
		if i == 0 || arr[i] >= arr[i-1] {
			visualizeIteration(White, arr, i-1, i)
			i++
		} else {
			visualizeIteration(Red, arr, i-1, i)
			arr[i], arr[i-1] = arr[i-1], arr[i]
			visualizeIteration(Green, arr, i-1, i)
			i--
			if i == 0 {
				i = 1
			}
		}
	}
}
