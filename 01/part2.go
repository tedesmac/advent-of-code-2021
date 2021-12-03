package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ParseArgs() string {
	if len(os.Args) < 2 {
		fmt.Println("Mising path to data")
		PrintHelp()
		os.Exit(1)
	}

	isHelp, _ := regexp.MatchString("^(-?-h(elp)?|h(elp)?)$", os.Args[1])
	if isHelp {
		PrintHelp()
		os.Exit(0)
	}

	return os.Args[1]
}

func PrintHelp() {
	fmt.Println("\nUsage:")
	fmt.Println("\tprogram <PATH>")
	fmt.Println("")
}

func ReadData() []int {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		PrintHelp()
		os.Exit(1)
	}

	res := []int{}
	s := bufio.NewScanner(file)
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		res = append(res, n)
	}

	file.Close()

	return res
}

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

func GetIncrements(a *[]int) int {
	count := 0
	prev := 2147483647
	for _, n := range *a {
		if prev < n {
			count += 1
		}
		prev = n
	}
	return count
}

func ProcessWindows(data *[]int) []int {
	windows := []int{}
	buffer := [][]int{}

	for _, n := range *data {
		for i, b := range buffer {
			buffer[i] = append(b, n)
		}

		buffer = append(buffer, []int{n})

		if len(buffer[0]) == 3 {
			windows = append(windows, SumSlice(&buffer[0]))
			buffer = buffer[1:]
		}
	}

	return windows
}

func SumSlice(s *[]int) int {
	res := 0
	for _, n := range *s {
		res += n
	}
	return res
}

func main() {
	d := ReadData()
	w := ProcessWindows(&d)
	n := GetIncrements(&w)
	fmt.Println(n)
}
