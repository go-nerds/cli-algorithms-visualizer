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
	var firstSlice string
	var secondSlice string
	var thirdSlice string

	if idx1 < idx2 {
		if len(array[:idx1]) > 0 {
			firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
			firstSlice = strings.Replace(firstSlice, "]", "", -1)
		}
		if len(array[idx1+1:idx2]) > 0 {
			secondSlice = strings.Replace(fmt.Sprintf("%v ", array[idx1+1:idx2]), "[", "", -1)
			secondSlice = strings.Replace(secondSlice, "]", "", -1)
		}

		thirdSlice = strings.Replace(fmt.Sprintf(" %v", array[idx2+1:]), "[", "", -1)
		thirdSlice = strings.Replace(thirdSlice, "]", "", -1)
	} else {
		if len(array[:idx2]) > 0 {
			firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx2]), "[", "", -1)
			firstSlice = strings.Replace(firstSlice, "]", "", -1)
		}

		secondSlice = strings.Replace(fmt.Sprintf(" %v ", array[idx2+1:idx1]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)

		thirdSlice = strings.Replace(fmt.Sprintf(" %v", array[idx1:]), "[", "", -1)
		thirdSlice = strings.Replace(thirdSlice, "]", "", -1)
	}

	var str string
	switch c {
	case Red:
		str = firstSlice +
			color.OpUnderscore.Sprint(color.Red.Sprint(array[idx1])) +
			" " +
			secondSlice +
			color.OpUnderscore.Sprint(color.Red.Sprint(array[idx2])) +
			thirdSlice
	case Green:
		str = firstSlice +
			color.OpUnderscore.Sprint(color.Green.Sprint(array[idx1])) +
			" " +
			secondSlice +
			color.OpUnderscore.Sprint(color.Green.Sprint(array[idx2])) +
			thirdSlice
	case White:
		str = firstSlice +
			color.OpUnderscore.Sprint(array[idx1]) +
			" " +
			secondSlice +
			color.OpUnderscore.Sprint(array[idx2]) +
			thirdSlice
	}

	fmt.Print(str)
}

func visualizeIteration(c Color, array []int, idx1, idx2 int) {
	printColoredArray(c, array, idx1, idx2)
	time.Sleep(time.Second)
	clearConsole()
}
