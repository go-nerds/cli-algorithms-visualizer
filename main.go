package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func main() {
	arr := generateRandomArray(10)
	fmt.Println("Initial array:", arr)

	// Prompt user to select the sorting algorithm
	prompt := promptui.Select{
		Label: "Select Sorting Algorithm",
		Items: []string{"Bubble Sort", "Selection Sort", "Merge Sort"},
	}

	_, algorithm, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	clearConsole()

	switch algorithm {
	case "Bubble Sort":
		bubbleSortVisualizer(arr)
	case "Selection Sort":
		// selectionSortVisualizer(arr)
	case "Merge Sort":
		// mergeSortVisualizer(arr)
	default:
		fmt.Println("Invalid selection")
		return
	}

	fmt.Println("Sorted array:", arr)
}
