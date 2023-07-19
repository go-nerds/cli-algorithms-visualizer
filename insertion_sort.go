package main

import (
	"fmt"
	"time"
)

func insertionSortVisualizer(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		visualizeIteration(White, arr, j-1, j)
		for j > 0 && arr[j] < arr[j-1] {
			printColoredArray(Red, arr, j-1, j)
			fmt.Printf(" // %d is smaller than %d\n", arr[j], arr[j-1])
			time.Sleep(time.Second * 1)
			clearConsole()
			str := fmt.Sprintf(" // swap %d and %d\n", arr[j], arr[j-1])
			arr[j], arr[j-1] = arr[j-1], arr[j]
			printColoredArray(Green, arr, j-1, j)
			fmt.Print(str)
			time.Sleep(time.Second * 1)
			clearConsole()
			j--
		}
	}
}
