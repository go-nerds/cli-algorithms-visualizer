package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

type options struct {
	Name SortingAlgorithm
}

func main() {
	clearConsole()

	r := bufio.NewReader(os.Stdin)
	n := getSliceSize(r)
	arr := make([]int, 0, n)
	fillSliceWithRandElems(&arr)

	fmt.Println("Initial array:", arr)

	algorithms := []options{
		{Name: BubbleSort},
		{Name: SelectionSort},
		{Name: InsertionSort},
		{Name: GnomeSort},
		{Name: CocktailShakerSort},
		{Name: CombSort},
		{Name: OddEvenSort},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "--> {{ .Name | cyan }}",
		Inactive: "    {{ .Name | cyan }}",
		Selected: "--> {{ .Name | cyan }}",
		Details: `
--------- Algorithm ----------
{{ "Name:" | faint }}	{{ .Name }}
`,
	}

	searcher := func(input string, index int) bool {
		algorithm := algorithms[index]
		name := strings.Replace(strings.ToLower(algorithm.Name.String()), " ", "", -1)
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
		fmt.Printf("Prompt failed: %s\n", err.Error())
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

	clearConsole()

	switch action {
	case "Description":
		printAlgorithmDescription(algorithms[i].Name)
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
			delay = time.Millisecond * 1500
		case "Medium":
			delay = time.Second
		case "Fast":
			delay = time.Second / 2
		case "Very Fast":
			delay = time.Second / 4
		default:
			delay = time.Second
		}

		clearConsole()
		runAlgorithm(algorithms[i].Name, arr, delay)
	default:
		fmt.Println("Invalid choice")
		return
	}
}

func runAlgorithm(algorithm SortingAlgorithm, arr []int, delay time.Duration) {
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
