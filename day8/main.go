package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	_ "embed"

	"github.com/echojc/aocutil"
)

//go:embed 2021_8_example.txt
var i string

func main() {
	result1 := part1()
	fmt.Println("Part 1 :", result1)

	result2 := part2()
	fmt.Println("Part 2 :", result2)

	return
}

func getMaterials(test bool) ([]string, []string) {
	var list0 []string
	var list1 []string
	var lines []string
	if !test {
		input, err := aocutil.NewInputFromFile("session_id")
		if err != nil {
			log.Fatal(err)
		}
		lines, err = input.Strings(2021, 8)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lines = strings.Split(i, "\n")
	}
	for _, line := range lines {
		tmpstrarray0 := strings.Split(line, "|")
		for _, tmpstr := range strings.Split(tmpstrarray0[0], " ") {
			list0 = append(list0, tmpstr)
		}
		for _, tmpstr := range strings.Split(tmpstrarray0[1], " ") {
			if tmpstr == "" {
				continue
			}
			list1 = append(list1, strings.TrimSpace(tmpstr))
		}
	}
	return list0, list1
}

func part1() int {
	var list []string
	var number int

	_, list = getMaterials(false)
	for i := range list {
		if len(list[i]) == 2 || len(list[i]) == 3 || len(list[i]) == 4 || len(list[i]) == 7 {
			number++
		}
	}
	return number
}

func part2() int {
	var list []string
	var output []string
	var output_value int
	var value int = 0

	list, output = getMaterials(false)
	for i, j := 0, 0; i < (len(list) - 10); i, j = i+11, j+4 {
		value = 0
		zero, one, two, three, four, five, six, seven, eight, nine := guessDigits(list[i : i+10])
		results := getNum(output[j:j+4], zero, one, two, three, four, five, six, seven, eight, nine)
		for c, item := range results {
			i, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal(err)
			}
			if c == 0 {
				i = i * 1000
			}
			if c == 1 {
				i = i * 100
			}
			if c == 2 {
				i = i * 10
			}
			value = value + i
		}
		output_value = output_value + value
	}
	return output_value
}

func getNum(output []string, zero string, one string, two string, three string, four string, five string, six string, seven string, eight string, nine string) []string {
	var results []string
	for _, v := range output {
		value := sortString(v)
		if value == zero {
			results = append(results, "0")
		} else if value == one {
			results = append(results, "1")
		} else if value == two {
			results = append(results, "2")
		} else if value == three {
			results = append(results, "3")
		} else if value == four {
			results = append(results, "4")
		} else if value == five {
			results = append(results, "5")
		} else if value == six {
			results = append(results, "6")
		} else if value == seven {
			results = append(results, "7")
		} else if value == eight {
			results = append(results, "8")
		} else if value == nine {
			results = append(results, "9")
		}
	}
	return results
}

func guessDigits(list []string) (zero string, one string, two string, three string, four string, five string, six string, seven string, eight string, nine string) {
	var len5 []string
	var len6 []string

	for i := range list {
		if len(list[i]) == 2 {
			one = list[i]
		} else if len(list[i]) == 3 {
			seven = list[i]
		} else if len(list[i]) == 4 {
			four = list[i]
		} else if len(list[i]) == 5 {
			len5 = append(len5, list[i])
		} else if len(list[i]) == 6 {
			len6 = append(len6, list[i])
		} else if len(list[i]) == 7 {
			eight = list[i]
		}
	}

	for i, str := range len6 {
		var count int = 0

		for _, item := range one {
			if strings.ContainsRune(str, item) {
				count++
			}
		}
		if count == 1 {
			six = str
			len6[i] = ""
			continue
		}
		count = 0
		for _, item := range four {
			if strings.ContainsRune(str, item) {
				count++
			}
		}
		if count == 4 {
			nine = str
			len6[i] = ""
		} else if count == 3 {
			zero = str
			len6[i] = ""
		}
	}

	for i, str := range len5 {
		var count int = 0

		for _, item := range one {
			if strings.ContainsRune(str, item) {
				count++
			}
		}
		if count == 2 {
			three = str
			len5[i] = ""
		} else if count == 1 {
			count = 0
			for _, item := range four {
				if strings.ContainsRune(str, item) {
					count++
				}
			}
			if count == 2 {
				two = str
				len5[i] = ""
			} else {
				five = str
				len5[i] = ""
			}
		}
	}
	return sortString(zero), sortString(one), sortString(two), sortString(three), sortString(four), sortString(five), sortString(six), sortString(seven), sortString(eight), sortString(nine)
}

func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}
