package main

import (
	"fmt"
	"strings"
	"time"

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
	handleInterrupt()

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
