package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func ReadData() []int8 {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		PrintHelp()
		os.Exit(1)
	}
	s := bufio.NewScanner(file)

	res := []int8{}
	for s.Scan() {
		for _, f := range strings.Split(s.Text(), ",") {
			n, _ := strconv.Atoi(f)
			res = append(res, int8(n))
		}
	}

	file.Close()

	return res
}

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

func main() {
	days := 80
	newFish := []int8{}
	var fish []int8 = ReadData()

	for i := 0; i < days; i++ {
		for i, _ := range fish {
			fish[i]--
			if fish[i] < 0 {
				fish[i] = 6
				newFish = append(newFish, 8)
			}
		}

		fish = append(fish, newFish...)
		newFish = []int8{}
	}

	fmt.Println(len(fish))
}
