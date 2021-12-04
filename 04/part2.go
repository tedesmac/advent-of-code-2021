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

func ReadData() ([]int, [][][]Cell) {
	path := ParseArgs()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		PrintHelp()
		os.Exit(1)
	}

	boards := [][][]Cell{}
	b := [][]Cell{}
	numbers := []int{}
	s := bufio.NewScanner(file)
	i := 0
	for s.Scan() {
		t := s.Text()

		if i == 0 {
			for _, n := range strings.Split(t, ",") {
				nn, _ := strconv.Atoi(n)
				numbers = append(numbers, nn)
			}
			i++
			continue
		}

		if len(t) == 0 {
			if len(b) != 0 {
				boards = append(boards, b)
				b = [][]Cell{}
			}
			continue
		}

		row := []Cell{}
		for _, n := range strings.Split(t, " ") {
			if len(n) == 0 {
				continue
			}
			nn, _ := strconv.Atoi(n)
			row = append(row, Cell{N: nn})
		}
		b = append(b, row)
	}

	boards = append(boards, b)

	file.Close()

	return numbers, boards
}

/* -----------------------------------------------------------------------------
 * Solution
 * ---------------------------------------------------------------------------*/

type Cell struct {
	Check bool
	N     int
}

func main() {
	I := 0
	N := 0
	wins := []int{}
	var numbers []int
	var boards [][][]Cell

	numbers, boards = ReadData()

	for _, n := range numbers {
		finish := false
		for i, board := range boards {
			MarkCheck(&board, n)
			if CheckWin(board) && !AlreadyWon(&wins, i) {
				wins = append(wins, i)

				if len(wins) == len(boards) {
					I = i
					N = n
					finish = true
					break
				}
			}
		}

		if finish {
			break
		}
	}

	sum := GetBoardSum(boards[I])
	fmt.Println(sum * N)
}

func AlreadyWon(a *[]int, n int) bool {
	for _, nn := range *a {
		if n == nn {
			return true
		}
	}
	return false
}

func CheckWin(b [][]Cell) bool {
	column := true
	row := true

	for i, r := range b {
		for j, _ := range r {
			column = column && b[j][i].Check
			row = row && b[i][j].Check
		}

		if column {
			return true
		}
		if row {
			return true
		}

		column = true
		row = true
	}

	return false
}

func GetBoardSum(board [][]Cell) int {
	sum := 0

	for _, row := range board {
		for _, cell := range row {
			if !cell.Check {
				sum += cell.N
			}
		}
	}

	return sum
}

func MarkCheck(b *[][]Cell, n int) {
	for i, row := range *b {
		for j, cell := range row {
			if cell.N == n {
				(*b)[i][j].Check = true
			}
		}
	}
}

func PrintBoard(board [][]Cell) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%v ", cell)
		}
		fmt.Println("\n")
	}
	fmt.Println("\n")
}
