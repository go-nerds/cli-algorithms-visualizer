package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/gookit/color"
)

func main() {
	arr := generateRandomArray(10)
	fmt.Println("Initial array:", arr)

	bubbleSortVisualizer(arr)
	fmt.Println("Sorted array:", arr)
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func bubbleSortVisualizer(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			printArrayWithCursor(arr, j, j+1)
			if arr[j] > arr[j+1] {
				color.Green.Printf("Swapping %d and %d\n", arr[j], arr[j+1])
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

			// Visualization
			time.Sleep(time.Second) // Pause for visualization
			clearConsole()
		}
	}
}

func printArrayWithCursor(arr []int, cursorIndex, swapIndex int) {
	for i, num := range arr {
		if i == cursorIndex || i == swapIndex {
			color.New(color.OpUnderscore.Light()).Print(fmt.Sprintf("%d ", num)) // Underline swapping elements
		} else {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()
}

func clearConsole() {
	// TODO: Add platform-specific code to clear the console or scroll up
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
