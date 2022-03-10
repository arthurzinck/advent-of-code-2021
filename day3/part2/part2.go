package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/echojc/aocutil"
)

var one []string
var zero []string
var result_o2 []string
var result_co2 []string

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func oxygen(table []string, index int) {
	one = nil
	zero = nil
	fmt.Printf("Length of table %d \n", len(table))
	if len(table) > 1 && index < 12 {
		count_ones := 0
		count_zeroes := 0
		fmt.Printf("Column number %d\n", index)

		var item string
		for x := 0; x < len(table); x++ {
			// fmt.Printf("Column number %d  Line Number %d\n", index, x)
			if index < 12 {
				item = table[x][index : index+1]
			} else {
				item = table[x][index-1:]
			}
			if item == string('1') {
				count_ones++
				fmt.Printf("%s\n", table[x])
				one = append(one, table[x])
			} else {
				count_zeroes++
				zero = append(zero, table[x])
			}
		}
		if count_ones >= count_zeroes {
			fmt.Printf("ones wins\n")
			one = removeDuplicateStr(one)
			oxygen(one, index+1)
		} else {
			fmt.Printf("zeroes wins\n")
			zero = removeDuplicateStr(zero)
			oxygen(zero, index+1)
		}
	} else {
		result_o2 = table
	}
}

func co2(table []string, index int) {
	one = nil
	zero = nil
	fmt.Printf("Length of table %d \n", len(table))
	if len(table) > 1 && index < 12 {
		count_ones := 0
		count_zeroes := 0
		fmt.Printf("Column number %d\n", index)

		var item string
		for x := 0; x < len(table); x++ {
			// fmt.Printf("Column number %d  Line Number %d\n", index, x)
			if index < 12 {
				item = table[x][index : index+1]
			} else {
				item = table[x][index-1:]
			}
			if item == string('1') {
				count_ones++
				// fmt.Printf("%s\n", table[x])
				one = append(one, table[x])
			} else {
				count_zeroes++
				zero = append(zero, table[x])
			}
		}
		if count_ones < count_zeroes {
			// fmt.Printf("ones wins\n")
			one = removeDuplicateStr(one)
			co2(one, index+1)
		} else {
			// fmt.Printf("zeroes wins\n")
			zero = removeDuplicateStr(zero)
			co2(zero, index+1)
		}
	} else {
		result_co2 = table
	}
}
func binaryToDecimal(binary_num string) int {
	num, err := strconv.ParseInt(binary_num, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(num)
}

func main() {
	i, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := i.Strings(2021, 3)
	if err != nil {
		log.Fatal(err)
	}

	oxygen(lines, 0)

	co2(lines, 0)

	fmt.Printf("Oxygen : %s\n", result_o2[0])

	fmt.Printf("Co2 : %s\n", result_co2)
	CO2_decimal := binaryToDecimal(result_co2[0])
	fmt.Printf("%d\n", CO2_decimal)

	O2_decimal := binaryToDecimal(result_o2[0])
	fmt.Printf("%d\n", O2_decimal)

	day3_part2 := O2_decimal * CO2_decimal
	fmt.Printf("%d\n", day3_part2)
}
