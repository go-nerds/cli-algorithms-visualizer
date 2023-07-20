package main

import (
	"fmt"
	"time"
)

func selectionSortVisualizer(arr []int, delay time.Duration) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			printColoLightYellowArray(White, arr, i, j)
			fmt.Printf(" // min = %d\n", arr[min])
			time.Sleep(delay)
			clearConsole()
			if arr[j] < arr[min] {
				visualizeIteration(LightYellow, arr, i, j, delay)
				min = j
			}
		}
		swapStr := fmt.Sprintf(" // swap arr[i]: %d and arr[min]: %d\n", arr[i], arr[min])
		arr[i], arr[min] = arr[min], arr[i]
		printColoLightYellowArray(LightBlue, arr, i, min)
		fmt.Print(swapStr)
		time.Sleep(delay)
		clearConsole()
	}
}
