package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	_ "embed"

	"github.com/GaryBoone/GoStats/stats"
	"github.com/echojc/aocutil"
)

//go:embed 2021_5_ex1.txt
var i string

var map_vents_part1 [1000][1000]int
var map_vents_part2 [1000][1000]int

func main() {
	result1 := part1()
	result2 := part2()
	fmt.Println("Part 1 :", result1)
	fmt.Println("Part 2 :", result2)

	return
}

func getMaterials(test bool) [][]int {

	var lists [][]int
	var lines []string
	if !test {
		input, err := aocutil.NewInputFromFile("session_id")
		if err != nil {
			log.Fatal(err)
		}
		lines, err = input.Strings(2021, 5)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lines = strings.Split(i, "\n")
	}

	for _, line := range lines {
		var tmpstrarray []string
		var list []int
		tmpstrarray = strings.Split(strings.Replace(line, " -> ", ",", 1), ",")
		for _, tmpstr := range tmpstrarray {
			tmpint, err := strconv.Atoi(tmpstr)
			if err != nil {
				log.Fatal(err)
			}
			list = append(list, tmpint)
		}
		lists = append(lists, list)
	}
	return lists
}

func countIntersects(map_vents [1000][1000]int) int {
	var result int = 0
	for i := range map_vents {
		for j := range map_vents[i] {
			if map_vents[i][j] >= 2 {
				result++
			}
		}
	}
	return result
}

func getFormula(coords []int) (int, int) {
	var x []float64
	var y []float64

	for i := 0; i < len(coords)-1; i += 2 {
		x = append(x, float64(coords[i]))
		y = append(y, float64(coords[i+1]))
	}
	tmp0, tmp1, _, _, _, _ := stats.LinearRegression(x, y)
	if math.IsNaN(tmp1) || math.IsNaN(tmp0) {
		tmp0 = 0
		tmp1 = 0
	}
	var slope int = int(tmp0)
	var y_intercept int = int(tmp1)
	return slope, y_intercept
}

func setCoordonatesPart1(slope int, y_intercept int, x0 int, x1 int, y0 int, y1 int) {
	if x0 == x1 {
		if y0 < y1 {
			for y := y0; y <= y1; y++ {
				map_vents_part1[x0][slope*x0+y]++
			}
		} else if y0 > y1 {
			for y := y1; y <= y0; y++ {
				map_vents_part1[x0][slope*x0+y]++
			}
		} else {
			map_vents_part1[x0][y0]++
		}
	} else if y0 == y1 {
		if x0 < x1 {
			for x := x0; x <= x1; x++ {
				map_vents_part1[x][y0]++
			}
		} else if x0 > x1 {
			for x := x1; x <= x0; x++ {
				map_vents_part1[x][y0]++
			}
		} else {
			map_vents_part1[x0][y0]++
		}
	}

}

func setCoordonatesPart2(slope int, y_intercept int, x0 int, x1 int, y0 int, y1 int) {
	if x0 == x1 {
		if y0 < y1 {
			for y := y0; y <= y1; y++ {
				map_vents_part2[x0][slope*x0+y]++
			}
		} else if y0 > y1 {
			for y := y1; y <= y0; y++ {
				map_vents_part2[x0][slope*x0+y]++
			}
		} else if y0 == y1 {
			map_vents_part2[x0][y0]++
		} else {
			fmt.Println("Weird Case : x0=x1 && y0== y1")
		}

	} else if x0 > x1 {
		for x := x1; x <= x0; x++ {
			map_vents_part2[x][slope*x+y_intercept]++
		}
	} else if x0 < x1 {
		for x := x0; x <= x1; x++ {
			map_vents_part2[x][slope*x+y_intercept]++
		}
	} else {
		fmt.Println("uncaught case", x0, y0, x1, y1)
	}
}

func part1() int {
	var lists [][]int
	var result int

	lists = getMaterials(false)
	for _, list := range lists {
		slope, y_intercept := getFormula(list)
		setCoordonatesPart1(slope, y_intercept, list[0], list[2], list[1], list[3])
	}
	result = countIntersects(map_vents_part1)
	return result
}

func part2() int {
	var lists [][]int
	var result int

	lists = getMaterials(false)
	for _, list := range lists {
		slope, y_intercept := getFormula(list)
		setCoordonatesPart2(slope, y_intercept, list[0], list[2], list[1], list[3])
	}
	result = countIntersects(map_vents_part2)
	return result
}
