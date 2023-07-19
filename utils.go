package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"github.com/gookit/color"
)

type Color int

const (
	White Color = iota
	Red
	Green
)

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func clearConsole() {
	// TODO: Add platform-specific code to clear the console or scroll up
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printColoredArray(c Color, array []int, idx1, idx2 int) {
	switch c {
	case Red:
		var firstSlice string
		if len(array[:idx1]) > 0 {
			firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
			firstSlice = strings.Replace(firstSlice, "]", "", -1)
		}
		secondSlice := strings.Replace(fmt.Sprintf(" %v", array[idx2+1:]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)
		str := firstSlice +
			color.OpUnderscore.Sprint(color.Red.Sprintf("%v", array[idx1])) +
			" " +
			color.OpUnderscore.Sprint(color.Red.Sprintf("%v", array[idx2])) +
			secondSlice
		fmt.Println(str)
	case Green:
		var firstSlice string
		if len(array[:idx1]) > 0 {
			firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
			firstSlice = strings.Replace(firstSlice, "]", "", -1)
		}
		secondSlice := strings.Replace(fmt.Sprintf(" %v", array[idx2+1:]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)
		str := firstSlice +
			color.OpUnderscore.Sprint(color.Green.Sprintf("%v", array[idx1])) +
			" " +
			color.OpUnderscore.Sprint(color.Green.Sprintf("%v", array[idx2])) +
			secondSlice
		fmt.Println(str)
	case White:
		var firstSlice string
		if len(array[:idx1]) > 0 {
			firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
			firstSlice = strings.Replace(firstSlice, "]", "", -1)
		}
		secondSlice := strings.Replace(fmt.Sprintf(" %v", array[idx2+1:]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)
		str := firstSlice +
			color.OpUnderscore.Sprint(array[idx1]) +
			" " +
			color.OpUnderscore.Sprint(array[idx2]) +
			secondSlice
		fmt.Println(str)
	}
}
