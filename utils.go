package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gookit/color"
	"github.com/pterm/pterm"
)

type Color int

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func clearConsole() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported platform. Cannot clear console.")
	}
}

func printColoredArray(array []int, idx1, idx2 int) {
	fmt.Print("[")
	for i := range array {
		if idx1 == i {
			fmt.Print(color.OpUnderscore.Sprint(color.LightYellow.Sprint(array[idx1]), " "))
		} else if idx2 == i {
			fmt.Print(color.OpUnderscore.Sprint(color.LightYellow.Sprint(array[idx2]), " "))
		} else {
			fmt.Print(array[i], " ")
		}
	}
	fmt.Println("]")
	fmt.Println()

}

func handleInterrupt() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		clearConsole()
		fmt.Println("Exiting.")
		os.Exit(0)
	}()
}

func visualizeIteration(array []int, idx1 int, idx2 int, displayType string, delay time.Duration) {
	if displayType == "Graph" {
		printGraph(array)
	} else if displayType == "Array" {
		printColoredArray(array, idx1, idx2)
	}
	time.Sleep(delay)
	clearConsole()
}

func printGraph(array []int) {
	area, _ := pterm.DefaultArea.Start()
	defer area.Stop()

	bars := pterm.Bars{}

	for i, data := range array {
		bar := pterm.Bar{Label: strconv.Itoa(i), Value: data}
		bars = append(bars, bar)
	}

	barchart := pterm.DefaultBarChart.WithBars(bars)

	content, _ := barchart.WithWidth(35).WithShowValue().Srender()
	area.Update(content)
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

func getSliceSize(r *bufio.Reader) int {
	var n int
	for {
		fmt.Print("Enter the size of the slice ( len(s) < 10 // recommended ): ")
		str, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}
		n, err = strconv.Atoi(strings.Trim(str, "\n"))
		if err != nil {
			fmt.Println("Input must be a number")
			continue
		}
		break
	}
	return n
}

func generateSlice(r *bufio.Reader, arrSize int) []int {
	slc := make([]int, arrSize)

	for i := 0; i < arrSize; {
		fmt.Printf("Enter element of index %d: ", i+1)
		str, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading element: ", err)
			continue
		}
		num, err := strconv.Atoi(strings.Trim(str, "\n"))
		if err != nil {
			fmt.Println("Element must be a number")
			continue
		}
		slc[i] = num
		i++
	}

	return slc
}
