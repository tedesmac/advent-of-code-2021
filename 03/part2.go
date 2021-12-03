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

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

type Method func([]string, int) byte

func Count(a *[]string, i int) (int, int) {
	o := 0
	z := 0
	for _, l := range *a {
		if l[i] == 48 {
			z += 1
		} else {
			o += 1
		}
	}
	return o, z
}

func GetReading(data []string, f Method) int64 {
	current := data
	filtered := []string{}
	i := 0

	for {
		c := f(current, i)
		for _, l := range current {
			if l[i] == c {
				filtered = append(filtered, l)
			}
		}

		current = filtered
		filtered = []string{}
		i++

		if len(current) == 1 {
			break
		}
	}

	r, _ := strconv.ParseInt(current[0], 2, 32)
	return r
}

func GetLeastCommon(a []string, i int) byte {
	o, z := Count(&a, i)

	if o < z {
		return 49
	}
	return 48
}

func GetMostCommon(a []string, i int) byte {
	o, z := Count(&a, i)

	if z > o {
		return 48
	}
	return 49
}

func main() {
	data := ReadData()

	csr := GetReading(data, GetLeastCommon)
	ogr := GetReading(data, GetMostCommon)

	fmt.Println(csr * ogr)
}
