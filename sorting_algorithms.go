package main

type SortingAlgorithm uint8

const (
	BubbleSort SortingAlgorithm = iota
	SelectionSort
	InsertionSort
	GnomeSort
	CocktailShakerSort
	CombSort
	OddEvenSort
)

func (a SortingAlgorithm) String() string {
	switch a {
	case BubbleSort:
		return "Bubble Sort"
	case SelectionSort:
		return "Selection Sort"
	case InsertionSort:
		return "Insertion Sort"
	case GnomeSort:
		return "Gnome Sort"
	case CocktailShakerSort:
		return "Cocktail Shaker Sort"
	case CombSort:
		return "Comb Sort"
	case OddEvenSort:
		return "Odd-Even Sort"
	default:
		return "Invalid selection"
	}
}
