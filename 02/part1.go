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

func ReadData() [][]string {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := [][]string{}
	s := bufio.NewScanner(file)
	for s.Scan() {
		res = append(res, strings.Split(s.Text(), " "))
	}

	file.Close()

	return res
}

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

func main() {
	data := ReadData()

	d := 0
	f := 0
	for _, r := range data {
		n, _ := strconv.Atoi(r[1])
		if r[0] == "down" {
			d += n
		} else if r[0] == "up" {
			d -= n
		} else {
			f += n
		}
	}

	fmt.Println(d * f)
}
