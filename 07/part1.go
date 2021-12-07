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

func ReadData() []int {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		PrintHelp()
		os.Exit(1)
	}
	s := bufio.NewScanner(file)

	res := []int{}
	for s.Scan() {
		for _, f := range strings.Split(s.Text(), ",") {
			n, _ := strconv.Atoi(f)
			res = append(res, int(n))
		}
	}

	file.Close()

	return res
}

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

func main() {
	var positions []int = ReadData()

	P := Min(positions...)
	max := Max(positions...)
	cost := CalcCost(positions, P)
	for i := P + 1; i <= max; i++ {
		current := CalcCost(positions, i)
		if current < cost {
			cost = current
			P = i
		}
	}

	fmt.Println(P, "-", cost)
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func CalcCost(a []int, t int) int {
	cost := 0

	for _, p := range a {
		cost += Abs(t - p)
	}

	return cost
}

func Max(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > res {
			res = a[i]
		}
	}
	return res
}

func Min(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < res {
			res = a[i]
		}
	}
	return res
}
