package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gookit/color"
)

type Color int

const (
	White Color = iota
	LightYellow
	LightBlue
)

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

func printColoredArray(c Color, array []int, idx1, idx2 int) {
	var firstSlice string
	var secondSlice string
	var thirdSlice string

	if len(array[:idx1]) > 0 {
		firstSlice = strings.Replace(fmt.Sprintf("%v ", array[:idx1]), "[", "", -1)
		firstSlice = strings.Replace(firstSlice, "]", "", -1)
	}
	if idx1 != idx2 && len(array[idx1+1:idx2]) > 0 {
		secondSlice = strings.Replace(fmt.Sprintf("%v ", array[idx1+1:idx2]), "[", "", -1)
		secondSlice = strings.Replace(secondSlice, "]", "", -1)
	}
	if len(array[idx2+1:]) > 0 {
		thirdSlice = strings.Replace(fmt.Sprintf(" %v", array[idx2+1:]), "[", "", -1)
		thirdSlice = strings.Replace(thirdSlice, "]", "", -1)
	}

	var str string
	switch c {
	case LightYellow:
		str = firstSlice +
			color.OpUnderscore.Sprint(color.LightYellow.Sprint(array[idx1])) +
			" " +
			secondSlice +
			color.OpUnderscore.Sprint(color.LightYellow.Sprint(array[idx2])) +
			thirdSlice
	case LightBlue:
		if idx1 != idx2 {
			str = firstSlice +
				color.OpUnderscore.Sprint(color.LightBlue.Sprint(array[idx1])) +
				" " +
				secondSlice +
				color.OpUnderscore.Sprint(color.LightBlue.Sprint(array[idx2])) +
				thirdSlice
		} else {
			str = firstSlice +
				color.OpUnderscore.Sprint(color.LightBlue.Sprint(array[idx1])) +
				" " +
				thirdSlice
		}
	case White:
		str = firstSlice +
			color.OpUnderscore.Sprint(array[idx1]) +
			" " +
			secondSlice +
			color.OpUnderscore.Sprint(array[idx2]) +
			thirdSlice
	}

	fmt.Printf("[ %v ]", str)
}

func visualizeIteration(c Color, array []int, idx1, idx2 int, delay time.Duration) {
	printColoredArray(c, array, idx1, idx2)
	time.Sleep(delay)
	clearConsole()
}
