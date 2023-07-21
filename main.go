package main

import (
	"fmt"
	"time"

	"github.com/gookit/color"
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
			"Quick Sort",
		},
	}

	_, algorithm, err := algorithmPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	clearConsole()

	actionPrompt := promptui.Select{
		Label: "Select Action",
		Items: []string{"Description", "Run"},
	}

	_, action, err := actionPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch action {
	case "Description":
		printAlgorithmDescription(algorithm)
	case "Run":
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
			delay = time.Second / 3
		case "Very Fast":
			delay = time.Second / 6
		default:
			delay = time.Second
		}

		clearConsole()

		runAlgorithm(algorithm, arr, delay)
	default:
		fmt.Println("Invalid choice")
		return
	}
}

func printAlgorithmDescription(algorithm string) {
	switch algorithm {
	case "Bubble Sort":
		color.Bold.Println(color.Yellow.Sprint("Bubble Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Bubble Sort is a simple comparison-based sorting algorithm. It works by repeatedly stepping through the list to be sorted,"))
		fmt.Println(color.Cyan.Sprint("compares adjacent elements, and swaps them if they are in the wrong order. The pass through the list is repeated until the list"))
		fmt.Println(color.Cyan.Sprint("is sorted. The algorithm gets its name from the way smaller elements 'bubble' to the top of the list."))
	case "Selection Sort":
		color.Bold.Println(color.Yellow.Sprint("Selection Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Selection Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into two parts:"))
		fmt.Println(color.Cyan.Sprint("the sorted part and the unsorted part. The algorithm repeatedly selects the minimum element from the unsorted part"))
		fmt.Println(color.Cyan.Sprint("and moves it to the end of the sorted part. This process continues until the entire list is sorted."))
	case "Insertion Sort":
		color.Bold.Println(color.Yellow.Sprint("Insertion Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Insertion Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into two parts:"))
		fmt.Println(color.Cyan.Sprint("the sorted part and the unsorted part. The algorithm repeatedly picks an element from the unsorted part and inserts"))
		fmt.Println(color.Cyan.Sprint("it into the correct position in the sorted part. This process continues until the entire list is sorted."))
	case "Gnome Sort":
		color.Bold.Println(color.Yellow.Sprint("Gnome Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Gnome Sort is an in-place comparison-based sorting algorithm. It works by repeatedly comparing adjacent elements."))
		fmt.Println(color.Cyan.Sprint("If the two elements are in the wrong order, it swaps them and moves one step backward. If the two elements"))
		fmt.Println(color.Cyan.Sprint("are in the correct order, it moves one step forward. This process continues until the entire list is sorted."))
	case "Cocktail Shaker Sort":
		color.Bold.Println(color.Yellow.Sprint("Cocktail Shaker Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Cocktail Shaker Sort, also known as Bidirectional Bubble Sort, is a variation of Bubble Sort. It works by repeatedly"))
		fmt.Println(color.Cyan.Sprint("sorting the list in both directions, first from the beginning to the end (like Bubble Sort), and then from the end"))
		fmt.Println(color.Cyan.Sprint("to the beginning. The algorithm stops when the list becomes sorted."))
	case "Comb Sort":
		color.Bold.Println(color.Yellow.Sprint("Comb Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Comb Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into a series of"))
		fmt.Println(color.Cyan.Sprint("gaps and repeatedly sorting the list with a specific shrink factor. The shrink factor reduces the gaps until it becomes 1."))
		fmt.Println(color.Cyan.Sprint("At this point, the algorithm behaves similar to Bubble Sort. Comb Sort is an improvement over Bubble Sort for large lists."))
	case "Odd-Even Sort":
		color.Bold.Println(color.Yellow.Sprint("Odd-Even Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Odd-Even Sort is an in-place comparison-based sorting algorithm. It works by repeatedly comparing and swapping"))
		fmt.Println(color.Cyan.Sprint("adjacent elements at even and odd indices. The process continues until the list is sorted. Odd-Even Sort is known"))
		fmt.Println(color.Cyan.Sprint("for its simplicity but is not very efficient for large lists."))
	case "Quick Sort":
		color.Bold.Println(color.Yellow.Sprint("Quick Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("The quick sort algorithm falls under the divide and conquer class of algorithms, where we break (divide) a problem into smaller chunks that are much simpler to solve (conquer). In this case, an unsorted array is broken into sub-arrays that are partially sorted, until all elements in the list are in the right position, by which time our unsorted list will have become sorted."))
	default:
		fmt.Println("Invalid selection")
	}
}

func runAlgorithm(algorithm string, arr []int, delay time.Duration) {
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
	case "Quick Sort":
		quickSort(arr, delay)
	default:
		fmt.Println("Invalid selection")
	}
	fmt.Println("Sorted array:", arr)
}
