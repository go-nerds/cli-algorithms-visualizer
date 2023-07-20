package main

func oddEvenSortVisualizer(arr []int) {
	n := len(arr)
	sorted := false

	for !sorted {
		sorted = true
		for i := 1; i <= n-2; i += 2 {
			visualizeIteration(White, arr, i, i+1)
			if arr[i] > arr[i+1] {
				visualizeIteration(Red, arr, i, i+1)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				visualizeIteration(Green, arr, i, i+1)
				sorted = false
			}
		}

		for i := 0; i <= n-2; i += 2 {
			visualizeIteration(White, arr, i, i+1)
			if arr[i] > arr[i+1] {
				visualizeIteration(Red, arr, i, i+1)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				visualizeIteration(Green, arr, i, i+1)
				sorted = false
			}
		}
	}
}
