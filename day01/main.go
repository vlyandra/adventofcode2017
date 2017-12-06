package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

// Solves the problem below:
//The captcha requires you to review a sequence of digits (your puzzle input) and find the sum of all digits that match the next digit in the list. The list is circular, so the digit after the last digit is the first digit in the list.
//
//For example:
//
//1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second digit and the third digit (2) matches the fourth digit.
//1111 produces 4 because each digit (all 1) matches the next.
//1234 produces 0 because no digit matches the next.
//91212129 produces 9 because the only digit that matches the next one is the last digit, 9.

func main() {
	entered := false
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	for !entered {
		var err error
		fmt.Println("Enter captcha:")
		/*
		input, err := reader.ReadString('\n')
		fmt.Println("Input was %s", input)
		if err != nil {
			fmt.Println("Error reading input. Try again.")
			continue
		}

		var sep []string
		if len(input) == 0 {
			fmt.Println("Captcha must be an integer length > 0. Try again.")
			continue
		} else {
			sep = strings.Split(input, "")
		}
		*/
		sum := 0
		prev := 0
		first := true
		firstint := 0
		current := 0

		for scanner.Scan() {
			ele := scanner.Text()
			fmt.Println("Ele: " + ele)
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
				//for i := 0; i < len(sep); i++ {
				//	if i == len(sep)-1 {
				//		next, err = strconv.Atoi(sep[0])
				//	} else {
				//		next, err = strconv.Atoi(sep[i+1])
				//	}
				//
				//	if err != nil {
				//		fmt.Println("Captcha must be an integer. Try again.")
				//		break
				//	}

				if current == prev {
					sum += current
				}

				prev = current
				fmt.Println("Current sum:", sum)
				//}
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