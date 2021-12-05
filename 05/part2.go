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

func ReadData() []Line {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		PrintHelp()
		os.Exit(1)
	}
	s := bufio.NewScanner(file)

	res := []Line{}
	for s.Scan() {
		m := re.FindAllStringSubmatch(s.Text(), -1)
		p0 := Point{X: Atoi(m[0][1]), Y: Atoi(m[0][2])}
		p1 := Point{X: Atoi(m[1][1]), Y: Atoi(m[1][2])}
		res = append(res, Line{P0: p0, P1: p1})
	}

	file.Close()

	return res
}

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

var re = regexp.MustCompile(`(\d+),(\d+)`)

type Point struct {
	N int
	X int
	Y int
}

type Line struct {
	P0 Point
	P1 Point
}

func main() {
	var lines []Line = ReadData()
	board := GetBoard(lines)

	for _, l := range lines {
		TraceLine(l, &board)
	}

	fmt.Println(GetOverlaps(&board))
}

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return int(n)
}

func GetBoard(a []Line) [][]Point {
	size := 0

	for _, l := range a {
		size = Max(size, l.P0.X, l.P0.Y, l.P1.X, l.P1.Y)
	}

	board := [][]Point{}

	for x := 0; x <= size; x++ {
		row := []Point{}
		for y := 0; y <= size; y++ {
			row = append(row, Point{X: x, Y: y})
		}
		board = append(board, row)
	}

	return board
}

func GetOverlaps(board *[][]Point) int {
	n := 0
	for _, row := range *board {
		for _, p := range row {
			if p.N >= 2 {
				n++
			}
		}
	}
	return n
}

func Max(a ...int) int {
	r := a[0]
	for i := 1; i < len(a); i++ {
		if r < a[i] {
			r = a[i]
		}
	}
	return r
}

func Sign(a, b int) int {
	if a > b {
		return -1
	} else if a < b {
		return 1
	}
	return 0
}

func TraceLine(l Line, board *[][]Point) {
	dx, sx := Abs(l.P1.X-l.P0.X), Sign(l.P0.X, l.P1.X)
	dy, sy := Abs(l.P1.Y-l.P0.Y), Sign(l.P0.Y, l.P1.Y)

	for i := 0; i <= Max(dx, dy); i++ {
		x := l.P0.X + i*sx
		y := l.P0.Y + i*sy
		(*board)[y][x].N += 1
	}
}
