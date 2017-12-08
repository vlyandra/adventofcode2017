package day02

import (
	"encoding/csv"
	"os"
	"io"
	"log"
	"fmt"
	"strconv"
	"bufio"
)

var r *csv.Reader
var csvFile *os.File

func init() {
	f, err := os.Open("src/day02/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	r = csv.NewReader(bufio.NewReader(f))
}

/*
PartOne solves the following problem:
For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.

For example, given the following spreadsheet:

5 1 9 5
7 5 3
2 4 6 8
The first row's largest and smallest values are 9 and 1, and their difference is 8.
The second row's largest and smallest values are 7 and 3, and their difference is 4.
The third row's difference is 6.
In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.
*/

func PartOne() {
	sum := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var min, max int
		for i, ele := range record {
			num, err := strconv.Atoi(ele)
			if err != nil {
				log.Fatal(err)
			}
			if i == 0 {
				min = num
				max = num
			} else if num > max {
				max = num
			} else if num < min {
				min = num
			}
		}
		sum += (max - min)
	}

	fmt.Println(sum)
}

/* PartTwo solves the following expansion to part 1:
It sounds like the goal is to find the only two numbers in each row where one evenly divides the other - that is, where the result of the division operation is a whole number. They would like you to find those numbers on each line, divide them, and add up each line's result.

For example, given the following spreadsheet:

5 9 2 8
9 4 7 3
3 8 6 5
In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
In the second row, the two numbers are 9 and 3; the result is 3.
In the third row, the result is 2.
In this example, the sum of the results would be 4 + 3 + 2 = 9.

What is the sum of each row's result in your puzzle input?
 */

func PartTwo() {
	sum := 0


}