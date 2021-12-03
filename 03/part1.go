package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strings"
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

func ReadData() []string {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		PrintHelp()
		os.Exit(1)
	}

	res := []string{}
	s := bufio.NewScanner(file)
	for s.Scan() {
		res = append(res, s.Text())
	}

	file.Close()

	return res
}

/* -------------------------------------------------------------------------------------------------
 * Solution
 * -----------------------------------------------------------------------------------------------*/

func GetEpsilon(ones, zeros []int) int64 {
	s := ""
	for i, n := range ones {
		if n > zeros[i] {
			s += "0"
		} else {
			s += "1"
		}
	}

	r, _ := strconv.ParseInt(s, 2, 32)

	return r
}

func GetGamma(ones, zeros []int) int64 {
	s := ""
	for i, n := range ones {
		if n > zeros[i] {
			s += "1"
		} else {
			s += "0"
		}
	}

	r, _ := strconv.ParseInt(s, 2, 32)

	return r
}

func main() {
	data := ReadData()

	ones := []int{}
	zeros := []int{}

	for i := 0; i < len(data[0]); i++ {
		ones = append(ones, 0)
		zeros = append(zeros, 0)
	}

	for _, l := range data {
		for i, r := range l {
			if r == '0' {
				zeros[i] += 1
			} else {
				ones[i] += 1
			}
		}
	}

	e := GetEpsilon(ones, zeros)
	g := GetGamma(ones, zeros)

	fmt.Println(e * g)
}
