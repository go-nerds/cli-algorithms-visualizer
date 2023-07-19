package main

import (
	"fmt"
	"time"
)

func selectionSortVisualizer(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			printColoredArray(White, arr, i, j)
			fmt.Printf(" // min = %d\n", arr[min])
			time.Sleep(time.Second)
			clearConsole()
			if arr[j] < arr[min] {
				visualizeIteration(Red, arr, i, j)
				min = j
			}
		}
		swapStr := fmt.Sprintf(" // swap arr[i]: %d and arr[min]: %d\n", arr[i], arr[min])
		arr[i], arr[min] = arr[min], arr[i]
		printColoredArray(Green, arr, i, min)
		fmt.Print(swapStr)
		time.Sleep(time.Second * 2)
		clearConsole()
	}
}
