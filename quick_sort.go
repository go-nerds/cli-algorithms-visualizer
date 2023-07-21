package main

import (
	"time"
)

func partition(arr []int, low, high int, delay time.Duration) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		visualizeIteration(White, arr, j, j+1, delay)
		if arr[j] < pivot {
			visualizeIteration(LightYellow, arr, j, j+1, delay)
			arr[i], arr[j] = arr[j], arr[i]
			i++
			visualizeIteration(LightBlue, arr, j, j+1, delay)
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func processQuickSort(arr []int, low, high int, delay time.Duration) []int {
	low = 0
	high = len(arr) - 1
	if low < high {
		var p int
		arr, p = partition(arr, low, high, delay)
		arr = processQuickSort(arr, low, p-1, delay)
		arr = processQuickSort(arr, p+1, high, delay)
	}
	return arr
}

func quickSort(arr []int, delay time.Duration) []int {
	return processQuickSort(arr, 0, len(arr)-1, delay)
}
