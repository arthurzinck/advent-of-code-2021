package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	_ "embed"

	"sort"

	"github.com/echojc/aocutil"
)

//go:embed 2021_7_example.txt
var i string

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
		lines, err = input.Strings(2021, 7)
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
	var med int

	list = getMaterials(false)
	med = median(list)
	fuel := fuelconsomation(list, med)
	result = fuel

	return result
}

func part2() int {
	var result int
	var list []int
	var med int
	var prev_fuel int
	var fuel int

	list = getMaterials(false)
	med = median(list)
	for i := med - 1000; i < med+1000; i++ {
		fuel = fuelconsomation2(list, i)
		if fuel < prev_fuel || prev_fuel == 0 {
			// fmt.Println(prev_fuel, fuel)
			prev_fuel = fuel
		}
	}
	result = prev_fuel
	return result
}
func avg(list []int) int {
	var sum int
	for j := range list {
		sum = list[j]
	}
	return sum / len(list)
}
func median(list []int) int {
	sort.Ints(list)
	index := len(list) / 2
	return list[index]
}

func fuelconsomation(list []int, avg int) int {
	var fuel int
	for i := range list {
		fuel = fuel + (int(math.Abs(float64(list[i] - avg))))
	}

	return fuel
}

func fuelconsomation2(list []int, position int) int {
	var fuel int
	var count int

	for i := range list {
		count = 0
		diff := int(math.Abs(float64(list[i] - position)))
		fuel = fuel + (int(math.Abs(float64(list[i] - position))))
		for j := diff; j > 0; j-- {
			fuel = fuel + count
			count++
		}
	}

	return fuel
}
