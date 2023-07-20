package main

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	arr := generateRandomArray(10)
	fmt.Println("Initial array:", arr)

	algorithmPrompt := promptui.Select{
		Label: "Select Sorting Algorithm",
		Items: []string{
			"Bubble Sort",
			"Selection Sort",
			"Insertion Sort",
			"Gnome Sort",
			"Cocktail Shaker Sort",
			"Comb Sort",
			"Odd-Even Sort",
		},
	}

	_, algorithm, err := algorithmPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	speedPrompt := promptui.Select{
		Label: "Select Visualization Speed",
		Items: []string{"Slow", "Medium", "Fast", "Very Fast"},
	}

	_, speed, err := speedPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var delay time.Duration
	switch speed {
	case "Slow":
		delay = time.Second * 2
	case "Medium":
		delay = time.Second
	case "Fast":
		delay = time.Second / 2
	case "Very Fast":
		delay = time.Second / 5
	default:
		delay = time.Second
	}

	clearConsole()

	switch algorithm {
	case "Bubble Sort":
		bubbleSortVisualizer(arr, delay)
	case "Selection Sort":
		selectionSortVisualizer(arr, delay)
	case "Insertion Sort":
		insertionSortVisualizer(arr, delay)
	case "Gnome Sort":
		gnomeSortVisualizer(arr, delay)
	case "Cocktail Shaker Sort":
		cocktailShakerSortVisualizer(arr, delay)
	case "Comb Sort":
		combSortVisualizer(arr, delay)
	case "Odd-Even Sort":
		oddEvenSortVisualizer(arr, delay)
	default:
		fmt.Println("Invalid selection")
		return
	}

	fmt.Println("Sorted array:", arr)
}
