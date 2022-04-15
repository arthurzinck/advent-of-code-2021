package main

import (
	"fmt"
	"testing"
)

func TestCheckWinVertical(t *testing.T) {

	var testcase [][]int
	var param []int = []int{
		100, 95, 37, 52, 68,
		100, 1, 73, 96, 63,
		100, 49, 9, 42, 97,
		100, 81, 20, 11, 46,
		100, 24, 2, 34, 18}

	var param1 []int = []int{
		72, 100, 37, 52, 68,
		80, 100, 73, 96, 63,
		16, 100, 9, 42, 97,
		25, 100, 20, 11, 46,
		31, 100, 2, 34, 18}

	var param2 []int = []int{
		72, 95, 100, 52, 68,
		80, 1, 100, 96, 63,
		16, 49, 100, 42, 97,
		25, 81, 100, 11, 46,
		31, 24, 100, 34, 18}

	var param3 []int = []int{
		72, 95, 37, 100, 68,
		80, 1, 73, 100, 63,
		16, 49, 9, 100, 97,
		25, 81, 20, 100, 46,
		31, 24, 2, 100, 18}

	var param4 []int = []int{
		72, 95, 37, 52, 100,
		80, 1, 73, 96, 100,
		16, 49, 9, 42, 100,
		25, 81, 20, 11, 100,
		31, 24, 2, 34, 100}

	testcase = append(testcase, param, param1, param2, param3, param4)
	for _, test := range testcase {
		got := checkWinVertical(test)
		if got != true {
			t.Errorf("checkWinVertical = %t; want True", got)
		}
	}

}

func TestCheckWinHoriz(t *testing.T) {
	var testcase [][]int
	var param []int = []int{
		100, 100, 100, 100, 100,
		80, 1, 73, 96, 63,
		16, 49, 9, 42, 97,
		25, 81, 20, 11, 46,
		31, 24, 2, 34, 18}

	var param1 []int = []int{
		72, 95, 37, 52, 68,
		100, 100, 100, 100, 100,
		16, 49, 9, 42, 97,
		25, 81, 20, 11, 46,
		31, 24, 2, 34, 18}

	var param2 []int = []int{
		72, 95, 37, 52, 68,
		80, 1, 73, 96, 63,
		100, 100, 100, 100, 100,
		25, 81, 20, 11, 46,
		31, 24, 2, 34, 18}
	var param3 []int = []int{
		72, 95, 37, 52, 68,
		80, 1, 73, 96, 63,
		16, 49, 9, 42, 97,
		100, 100, 100, 100, 100,
		31, 24, 2, 34, 18}

	var param4 []int = []int{
		72, 95, 37, 52, 68,
		80, 1, 73, 96, 63,
		16, 49, 9, 42, 97,
		25, 81, 20, 11, 46,
		100, 100, 100, 100, 100}

	testcase = append(testcase, param, param1, param2, param3, param4)
	for _, test := range testcase {
		fmt.Println(test)
		got := checkWinHoriz(test)
		if got != true {
			t.Errorf("checkWinHoriz = %t; want True", got)
		}
	}
}
