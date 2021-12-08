package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	// "strconv"
	"strings"
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

func ReadData() [][]string {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		PrintHelp()
		os.Exit(1)
	}
	s := bufio.NewScanner(file)

	res := [][]string{}
	for s.Scan() {
		l := strings.Split(s.Text(), " | ")
		res = append(res, strings.Split(l[1], " "))
	}

	file.Close()

	return res
}

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

func main() {
	var data [][]string = ReadData()

	count := 0
	for _, row := range data {
		for _, digit := range row {
			switch len(digit) {
			case 2, 3, 4, 7:
				count += 1
			}
		}
	}

	fmt.Println(count)
}
