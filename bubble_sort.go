package main

import "time"

func bubbleSortVisualizer(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			printColoredArray(White, arr, j, j+1)
			time.Sleep(time.Second)
			clearConsole()
			if arr[j] > arr[j+1] {
				printColoredArray(Red, arr, j, j+1)
				arr[j], arr[j+1] = arr[j+1], arr[j]
				time.Sleep(time.Second) // Pause for visualization
				clearConsole()
				printColoredArray(Green, arr, j, j+1)
				time.Sleep(time.Second) // Pause for visualization
				clearConsole()
			}
		}
	}
}
