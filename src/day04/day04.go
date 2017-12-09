package day04

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"sort"
)

/* PartOne solves the following problem:
A passphrase consists of a series of words (lowercase letters) separated by spaces.

To ensure security, a valid passphrase must contain no duplicate words.

For example:

aa bb cc dd ee is valid.
aa bb cc dd aa is not valid - the word aa appears more than once.
aa bb cc dd aaa is valid - aa and aaa count as different words.
The system's full passphrase list is available as your puzzle input. How many passphrases are valid?
 */
func PartOne() {
	f, err := os.Open("src/day04/input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	r := bufio.NewScanner(f)
	count := 0
	for r.Scan() {
		in := r.Text()
		arr := strings.Split(in, " ")

		sort.Strings(arr)

		repeat := false
		for i := 0; i < (len(arr)-1); i++ {
			if arr[i+1] == arr[i] {
				repeat = true
				break
			}
		}
		if !repeat {
			count++
		}
	}
	fmt.Println(count)
}

/* PartTwo solves the following problem:
Now, a valid passphrase must contain no two words that are anagrams of each other - that is, a passphrase is invalid if any word's letters can be rearranged to form any other word in the passphrase.

For example:

abcde fghij is a valid passphrase.
abcde xyz ecdab is not valid - the letters from the third word can be rearranged to form the first word.
a ab abc abd abf abj is a valid passphrase, because all letters need to be used when forming another word.
iiii oiii ooii oooi oooo is valid.
oiii ioii iioi iiio is not valid - any of these words can be rearranged to form any other word.
Under this new system policy, how many passphrases are valid?
 */

func PartTwo() {
	f, err := os.Open("src/day04/input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	r := bufio.NewScanner(f)
	count := 0

	for r.Scan() {
		in := r.Text()
		arr := strings.Split(in, " ")

		var splitarr []string

		for _, ele := range arr {
			split := strings.Split(ele, "")
			sort.Strings(split)
			splitarr = append(splitarr, strings.Join(split, ""))
		}

		sort.Strings(splitarr)
		repeat := false

		for i := 0; i < (len(splitarr)-1); i++ {
			if splitarr[i+1] == splitarr[i] {
				repeat = true
				break
			}
		}
		if !repeat {
			count++
		}
	}
	fmt.Println(count)
}