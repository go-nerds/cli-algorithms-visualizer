package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gookit/color"
)

type Color int

const (
	White Color = iota
	Red
	Green
)

func main() {
	arr := generateRandomArray(10)
	fmt.Println("Initial array:", arr)
	time.Sleep(time.Second) // Pause for visualization
	clearConsole()
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
			printColoredArray(White, arr, j, j+1)
			time.Sleep(time.Second)
			clearConsole()
			if arr[j] > arr[j+1] {
				printColoredArray(Red, arr, j, j+1)
				arr[j], arr[j+1] = arr[j+1], arr[j]
				time.Sleep(time.Second) // Pause for visualization
				clearConsole()
				printColoredArray(Green, arr, j, j+1)
				time.Sleep(time.Second) // Pause for visualization
				clearConsole()
			}
		}
	}
}

func printColoredArray(c Color, array []int, idx1, idx2 int) {
	switch c {
	case Red:
		var firstSlice string
		if len(array[:idx1]) > 0 {
			firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
			firstSlice = strings.Replace(firstSlice, "]", "", -1)
		}
		secondSlice := strings.Replace(fmt.Sprintf(" %v", array[idx2:]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)
		str := firstSlice +
			color.OpUnderscore.Sprint(color.Red.Sprintf("%v ", array[idx1])) +
			color.OpUnderscore.Sprint(color.Red.Sprintf("%v ", array[idx2])) +
			secondSlice
		fmt.Println(str)
	case Green:
		var firstSlice string
		if len(array[:idx1]) > 0 {
			firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
			firstSlice = strings.Replace(firstSlice, "]", "", -1)
		}
		secondSlice := strings.Replace(fmt.Sprintf(" %v", array[idx2:]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)
		str := firstSlice +
			color.OpUnderscore.Sprint(color.Green.Sprintf("%v ", array[idx1])) +
			color.OpUnderscore.Sprint(color.Green.Sprintf("%v ", array[idx2])) +
			secondSlice
		fmt.Println(str)
	case White:
		var firstSlice string
		if len(array[:idx1]) > 0 {
			firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
			firstSlice = strings.Replace(firstSlice, "]", "", -1)
		}
		secondSlice := strings.Replace(fmt.Sprintf(" %v", array[idx2:]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)
		str := firstSlice +
			color.OpUnderscore.Sprint(array[idx1]) +
			" " +
			color.OpUnderscore.Sprint(array[idx2]) +
			secondSlice
		fmt.Println(str)
	}
}

func clearConsole() {
	// TODO: Add platform-specific code to clear the console or scroll up
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
