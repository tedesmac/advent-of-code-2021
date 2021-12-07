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
	days := 256
	counter := []int64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var fish []int8 = ReadData()

	for _, f := range fish {
		counter[f]++
	}

	var newFish int64 = 0
	for i := 0; i < days; i++ {
		for j, _ := range counter {
			if j == 0 {
				newFish = counter[j]
			} else if j == 8 {
				break
			}
			counter[j] = counter[j+1]
		}
		counter[6] += newFish
		counter[8] = newFish
		newFish = 0
	}

	var sum int64 = 0
	for _, n := range counter {
		sum += n
	}

	fmt.Println(sum)
}
