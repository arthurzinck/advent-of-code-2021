package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "embed"

	"github.com/echojc/aocutil"
)

//go:embed 2021_6_example.txt
var i string
var results []int

func main() {
	result1 := part1()
	fmt.Println("Part 1 :", result1)

	result2 := part2()
	fmt.Println("Part 2 :", result2)

	return
}

func getMaterials(test bool) []int {

	var list []int
	var lines []string
	if !test {
		input, err := aocutil.NewInputFromFile("session_id")
		if err != nil {
			log.Fatal(err)
		}
		lines, err = input.Strings(2021, 6)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lines = strings.Split(i, "\n")
	}
	for _, line := range lines {
		var tmpstrarray []string
		tmpstrarray = strings.Split(line, ",")
		for _, tmpstr := range tmpstrarray {
			tmpint, err := strconv.Atoi(tmpstr)
			if err != nil {
				log.Fatal(err)
			}
			list = append(list, tmpint)
		}
	}
	return list
}

func part1() int {
	var list []int
	var result int

	list = getMaterials(false)
	result = simulation(list, 80)

	return result
}

func part2() int {
	var result int
	var list []int
	var fishgroup [9]int

	list = getMaterials(false)
	fishgroup = triagefish(list)
	result = otherSimulation(fishgroup, 256)

	return result
}

func simulation(initial []int, duration int) int {
	for y := duration; y > 0; y-- {
		for i, x := range initial {
			if x == 0 {
				initial[i] = 6
				initial = append(initial, 8)
			} else {
				initial[i]--
			}
		}
		results = append(results, len(initial))
	}

	return len(initial)
}
func triagefish(initial []int) [9]int {
	var fishgroup [9]int
	for j := range initial {
		if initial[j] == 0 {
			fishgroup[0]++
		} else if initial[j] == 1 {
			fishgroup[1]++
		} else if initial[j] == 2 {
			fishgroup[2]++
		} else if initial[j] == 3 {
			fishgroup[3]++
		} else if initial[j] == 4 {
			fishgroup[4]++
		} else if initial[j] == 5 {
			fishgroup[5]++
		} else if initial[j] == 6 {
			fishgroup[6]++
		} else if initial[j] == 7 {
			fishgroup[7]++
		} else if initial[j] == 8 {
			fishgroup[8]++
		}

	}

	return fishgroup
}
func otherSimulation(fishgroup [9]int, duration int) int {
	var result int
	var newfishgroup [9]int

	for y := duration; y > 0; y-- {
		for i := 0; i < 8; i++ {
			newfishgroup[i] = fishgroup[i+1]
		}
		newfishgroup[8] = fishgroup[0]
		newfishgroup[6] = newfishgroup[6] + fishgroup[0]
		fishgroup = newfishgroup
	}
	for x := range fishgroup {
		result = result + fishgroup[x]
	}
	return result
}
