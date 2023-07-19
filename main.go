package main

import (
	"fmt"
	"time"
)

func main() {
	arr := generateRandomArray(10)
	fmt.Println("Initial array:", arr)
	time.Sleep(time.Second) // Pause for visualization
	clearConsole()
	bubbleSortVisualizer(arr)
	fmt.Println("Sorted array:", arr)
}
