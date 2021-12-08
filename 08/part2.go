package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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

func ReadData() [][][][]int {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		PrintHelp()
		os.Exit(1)
	}
	s := bufio.NewScanner(file)

	res := [][][][]int{}
	for s.Scan() {
		l := strings.Split(s.Text(), " | ")
		p0 := [][]int{}
		p1 := [][]int{}

		for _, s := range strings.Split(l[0], " ") {
			p0 = append(p0, Rune(s))
		}
		for _, s := range strings.Split(l[1], " ") {
			p1 = append(p1, Rune(s))
		}

		res = append(res, [][][]int{p0, p1})
	}

	file.Close()

	return res
}

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

func main() {
	var data [][][][]int = ReadData()

	res := 0
	for _, row := range data {
		res += Solve(row)
	}

	fmt.Println(res)
}

func Solve(data [][][]int) int {
	one, input := FindByLength(data[0], 2)
	seven, input := FindByLength(input, 3)
	four, input := FindByLength(input, 4)
	eight, input := FindByLength(input, 7)
	three, input := FindByDiff(input, seven, 2, 5)
	nine, input := FindByDiff(input, Combine(three, four), 0, 6)
	five, input := FindByDiff(input, nine, 1, 5)
	two, input := FindByLength(input, 5)
	zero, input := FindByDiff(input, one, 4, 6)

	solution := [][]int{
		zero,
		one,
		two,
		three,
		four,
		five,
		input[0],
		seven,
		eight,
		nine,
	}

	digits := []int{}

	for _, n := range data[1] {
		for i, _ := range solution {
			if CompareSlice(n, solution[i]) {
				digits = append(digits, i)
			}
		}
	}

	for i := 0; i < len(digits); i++ {
		value += digits[i] * Raise(10, len(digits)-i)
	}

	value := 0

	return value
}

func FindByDiff(input [][]int, r []int, d, l int) ([]int, [][]int) {
	res := []int{}

	for _, n := range input {
		if len(n) != l {
			continue
		}
		if Difference(n, r) == d {
			res = n
			break
		}
	}

	return res, Remove(input, res)
}

func FindByLength(input [][]int, l int) ([]int, [][]int) {
	res := []int{}

	for _, n := range input {
		if len(n) == l {
			res = n
			break
		}
	}

	return res, Remove(input, res)
}

func Combine(a ...[]int) []int {
	res := []int{}

	for _, n := range a[0] {
		res = append(res, n)
	}

	for i := 1; i < len(a); i++ {
		for _, n := range a[i] {
			ii := Index(res, n)
			if ii < 0 {
				res = append(res, n)
			}
		}
	}

	sort.Ints(res)
	return res
}

func CompareSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, n := range a {
		if n != b[i] {
			return false
		}
	}

	return true
}

func Difference(a, b []int) int {
	var min, max []int
	res := 0

	if len(a) < len(b) {
		min, max = a, b
	} else {
		min, max = b, a
	}

	for _, n := range max {
		i := Index(min, n)
		if i >= 0 {
			res++
		}
	}

	return len(max) - res
}

func Index(a []int, e int) int {
	for i, r := range a {
		if r == e {
			return i
		}
	}
	return -1
}

func IndexArray(a [][]int, e []int) int {
	for i, r := range a {
		if CompareSlice(r, e) {
			return i
		}
	}
	return -1
}

func Raise(n, r int) int {
	res := 1
	for i := 1; i < r; i++ {
		res *= n
	}
	return res
}

func Remove(a [][]int, b ...[]int) [][]int {
	res := [][]int{}
	for _, e := range a {
		i := IndexArray(b, e)
		if i < 0 {
			res = append(res, e)
		}
	}
	return res
}

func Rune(s string) []int {
	a := []rune(s)
	res := []int{}

	for _, n := range a {
		res = append(res, int(n))
	}

	sort.Ints(res)
	return res
}
