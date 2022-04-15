package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func main() {
	result1 := part1()
	result2 := part2()
	fmt.Println("Part 1 :", result1)
	fmt.Println("Part 2 :", result2)

	return
}

func getMaterials() ([]int, [][]int) {

	var listOfBoards [][]int
	var picked_number []int

	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	lines, err := i.Strings(2021, 4)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range strings.Split(lines[0], ",") {
		tmp, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		picked_number = append(picked_number, tmp)
	}

	for x := 2; x < len(lines); x += 6 {
		var board []int
		tmp_str := strings.Split(strings.Join(lines[x:x+5], " "), " ")
		for _, value := range tmp_str {
			if value == "" {
				continue
			}
			tmp, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			board = append(board, tmp)
		}
		if len(board) != 25 {
			fmt.Println("There is an error!")
		}
		listOfBoards = append(listOfBoards, board)
	}
	return picked_number, listOfBoards
}

func checkWinHoriz(board []int) bool {
	win := true
	for i := 0; i < len(board)-1; i += 5 {
		for _, cell := range board[i : i+5] {
			if cell != 100 {
				win = false
				break
			} else {
				win = true

			}
		}
		if win {
			return win
		}
	}
	return false
}

func checkWinVertical(board []int) bool {
	win := true
	for i := 0; i < 5; i++ {
		for j := 0; j < 25; j += 5 {
			if board[i%5+j] != 100 {
				win = false
				break
			} else {
				win = true
			}
		}
		if win {
			return win
		}
	}
	return false
}

func part1() int {
	var picks []int
	var boards [][]int
	var result int

	picks, boards = getMaterials()

	for _, pick := range picks {
		for _, board := range boards {
			for i, cell := range board {
				if cell == pick {
					board[i] = 100
					if checkWinHoriz(board) || checkWinVertical(board) {
						sum := 0
						for _, value := range board {
							if value == 100 {
								continue
							}
							sum += value
						}
						result = sum * pick
						return sum * pick
					}
				}
			}
		}
	}
	return result
}

func part2() int {
	var picks []int
	var boards [][]int
	var result int

	picks, boards = getMaterials()

	var hasWin = make([]bool, len(boards))

	for _, pick := range picks {
		for t, board := range boards {
			if hasWin[t] {
				continue
			}
			for i, cell := range board {
				if cell == pick {
					board[i] = 100
					if checkWinHoriz(board) || checkWinVertical(board) {
						hasWin[t] = true
						sum := 0
						for _, value := range board {
							if value == 100 {
								continue
							}
							sum += value
						}
						result = sum * pick
					}
				}
			}
		}
	}
	return result
}
