package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
)

const (
	BubbleSort         int = 1
	SelectionSort      int = 2
	InsertionSort      int = 3
	GnomeSort          int = 4
	CocktailShakerSort int = 5
	CombSort           int = 6
	OddEvenSort        int = 7
)

type options struct {
	Name string
	Uuid int
}

func main() {

	clearConsole()

	arr := generateRandomArray(10)
	fmt.Println("Initial array:", arr)

	algorithms := []options{
		{Name: "Bubble Sort", Uuid: 1},
		{Name: "Selection Sort", Uuid: 2},
		{Name: "Insertion Sort", Uuid: 3},
		{Name: "Gnome Sort", Uuid: 4},
		{Name: "Cocktail Shaker Sort", Uuid: 5},
		{Name: "Comb Sort", Uuid: 6},
		{Name: "Odd-Even Sort", Uuid: 7},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U00002728 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "\U00002728 {{ .Name | red | cyan }}",
		Details: `
--------- Algorithm ----------
{{ "Name:" | faint }}	{{ .Name }}`,
	}

	searcher := func(input string, index int) bool {
		algorithm := algorithms[index]
		name := strings.Replace(strings.ToLower(algorithm.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	algorithmPrompt := promptui.Select{
		Label:     "Select a sorting algorithm!",
		Items:     algorithms,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := algorithmPrompt.Run()

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
		printAlgorithmDescription(algorithms[i].Uuid)
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

		runAlgorithm(algorithms[i].Uuid, arr, delay)
	default:
		fmt.Println("Invalid choice")
		return
	}
}

func printAlgorithmDescription(algorithm int) {
	switch algorithm {
	case BubbleSort:
		color.Bold.Println(color.Yellow.Sprint("Bubble Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Bubble Sort is a simple comparison-based sorting algorithm. It works by repeatedly stepping through the list to be sorted,"))
		fmt.Println(color.Cyan.Sprint("compares adjacent elements, and swaps them if they are in the wrong order. The pass through the list is repeated until the list"))
		fmt.Println(color.Cyan.Sprint("is sorted. The algorithm gets its name from the way smaller elements 'bubble' to the top of the list."))
	case SelectionSort:
		color.Bold.Println(color.Yellow.Sprint("Selection Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Selection Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into two parts:"))
		fmt.Println(color.Cyan.Sprint("the sorted part and the unsorted part. The algorithm repeatedly selects the minimum element from the unsorted part"))
		fmt.Println(color.Cyan.Sprint("and moves it to the end of the sorted part. This process continues until the entire list is sorted."))
	case InsertionSort:
		color.Bold.Println(color.Yellow.Sprint("Insertion Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Insertion Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into two parts:"))
		fmt.Println(color.Cyan.Sprint("the sorted part and the unsorted part. The algorithm repeatedly picks an element from the unsorted part and inserts"))
		fmt.Println(color.Cyan.Sprint("it into the correct position in the sorted part. This process continues until the entire list is sorted."))
	case GnomeSort:
		color.Bold.Println(color.Yellow.Sprint("Gnome Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Gnome Sort is an in-place comparison-based sorting algorithm. It works by repeatedly comparing adjacent elements."))
		fmt.Println(color.Cyan.Sprint("If the two elements are in the wrong order, it swaps them and moves one step backward. If the two elements"))
		fmt.Println(color.Cyan.Sprint("are in the correct order, it moves one step forward. This process continues until the entire list is sorted."))
	case CocktailShakerSort:
		color.Bold.Println(color.Yellow.Sprint("Cocktail Shaker Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Cocktail Shaker Sort, also known as Bidirectional Bubble Sort, is a variation of Bubble Sort. It works by repeatedly"))
		fmt.Println(color.Cyan.Sprint("sorting the list in both directions, first from the beginning to the end (like Bubble Sort), and then from the end"))
		fmt.Println(color.Cyan.Sprint("to the beginning. The algorithm stops when the list becomes sorted."))
	case CombSort:
		color.Bold.Println(color.Yellow.Sprint("Comb Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Comb Sort is an in-place comparison-based sorting algorithm. It works by dividing the input list into a series of"))
		fmt.Println(color.Cyan.Sprint("gaps and repeatedly sorting the list with a specific shrink factor. The shrink factor reduces the gaps until it becomes 1."))
		fmt.Println(color.Cyan.Sprint("At this point, the algorithm behaves similar to Bubble Sort. Comb Sort is an improvement over Bubble Sort for large lists."))
	case OddEvenSort:
		color.Bold.Println(color.Yellow.Sprint("Odd-Even Sort"))
		fmt.Println()
		fmt.Println(color.Cyan.Sprint("Odd-Even Sort is an in-place comparison-based sorting algorithm. It works by repeatedly comparing and swapping"))
		fmt.Println(color.Cyan.Sprint("adjacent elements at even and odd indices. The process continues until the list is sorted. Odd-Even Sort is known"))
		fmt.Println(color.Cyan.Sprint("for its simplicity but is not very efficient for large lists."))
	default:
		fmt.Println("Invalid selection")
	}
}

func runAlgorithm(algorithm int, arr []int, delay time.Duration) {
	switch algorithm {
	case BubbleSort:
		bubbleSortVisualizer(arr, delay)
	case SelectionSort:
		selectionSortVisualizer(arr, delay)
	case InsertionSort:
		insertionSortVisualizer(arr, delay)
	case GnomeSort:
		gnomeSortVisualizer(arr, delay)
	case CocktailShakerSort:
		cocktailShakerSortVisualizer(arr, delay)
	case CombSort:
		combSortVisualizer(arr, delay)
	case OddEvenSort:
		oddEvenSortVisualizer(arr, delay)
	default:
		fmt.Println("Invalid selection")
	}
	fmt.Println("Sorted array:", arr)
}
