package day01

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
}

// PartOne solves the problem below:
//The captcha requires you to review a sequence of digits (your puzzle input) and find the sum of all digits that match the next digit in the list. The list is circular, so the digit after the last digit is the first digit in the list.
//
//For example:
//
//1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
//1111 produces 4 because each digit (all 1) matches the next.
//1234 produces 0 because no digit matches the next.
//91212129 produces 9 because the only digit that matches the next one is the last digit, 9.

func PartOne() {
	entered := false

	scanner.Split(bufio.ScanRunes)

	for !entered {
		var err error
		fmt.Println("Enter captcha:")

		sum := 0
		prev := 0
		first := true
		firstint := 0
		current := 0

		for scanner.Scan() {
			ele := scanner.Text()
			if ele == "\n" {
				break
			} else {
				current, err = strconv.Atoi(scanner.Text())
				if err != nil {
					fmt.Println("Captcha must be an integer: Try again.")
					break
				}

				if first {
					firstint = current
					first = false
				}

				if current == prev {
					sum += current
				}

				prev = current
			}
		}

		if err != nil {
			continue
		}

		if current == firstint {
			sum += firstint
		}

		fmt.Println(sum)
		entered = true
	}
}

// PartTwo solves the following addition to part 1
//consider the digit halfway around the circular list. That is, if your list contains 10 items, only include a digit in your sum if the digit 10/2 = 5 steps forward matches it. Fortunately, your list has an even number of elements.
//
//For example:
//
//1212 produces 6: the list contains 4 items, and all four digits match the digit 2 items ahead.
//1221 produces 0, because every comparison is between a 1 and a 2.
//123425 produces 4, because both 2s match each other, but no other digit has a match.
//123123 produces 12.
//12131415 produces 4.

func PartTwo() {
	var capArray []int
	scanner.Split(bufio.ScanRunes)
	fmt.Println("Enter captcha:")

	for scanner.Scan() {
		ele := scanner.Text()
		if ele == "\n" {
			break
		} else {
			num, err := strconv.Atoi(ele)
			if err != nil {
				fmt.Println("Captcha must be an integer.")
				return
			}
			capArray = append(capArray, num)
		}
	}

	capLen := len(capArray)
	a := capArray
	ahead := capLen/2
	sum := 0
	for i, ele := range capArray {
		if ele == a[i + ahead] {
			sum += ele
		}

		a = append(a, ele)
	}
	fmt.Println(sum)
}