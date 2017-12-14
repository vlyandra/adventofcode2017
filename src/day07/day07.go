package day07

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"regexp"
	"strconv"
)

var tree []Program

func init() {
	tree = make([]Program, 0)
}

/* PartOne solves the following problem:
One program at the bottom supports the entire tower. It's holding a large disc, and on the disc are balanced several more sub-towers. At the bottom of these sub-towers, standing on the bottom disc, are other programs, each holding their own disc, and so on. At the very tops of these sub-sub-sub-...-towers, many programs stand simply keeping the disc below them balanced but with no disc of their own.

You offer to help, but first you need to understand the structure of these towers. You ask each program to yell out their name, their weight, and (if they're holding a disc) the names of the programs immediately above them balancing on that disc. You write this information down (your puzzle input). Unfortunately, in their panic, they don't do this in an orderly fashion; by the time you're done, you're not sure which program gave which information.

For example, if your list is the following:

pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
...then you would be able to recreate the structure of the towers that looks like this:

                gyxo
              /
         ugml - ebii
       /      \
      |         jptl
      |
      |         pbga
     /        /
tknk --- padx - havc
     \        \
      |         qoyq
      |
      |         ktlj
       \      /
         fwft - cntj
              \
                xhth
In this example, tknk is at the bottom of the tower (the bottom program), and is holding up ugml, padx, and fwft. Those programs are, in turn, holding up other programs; in this example, none of those programs are holding up any other programs, and are all the tops of their own towers. (The actual tower balancing in front of you is much larger.)

Before you're ready to help them, you need to make sure your information is correct. What is the name of the bottom program?
 */

type Program struct {
	Name string
	Weight int
	Leaf bool
	Children map[string]*Program
}

func PartOne() {
	f, err := os.Open("src/day07/input.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		in := s.Text()

		name, weight, children := parse(in)

		childmap := make(map[string]*Program)

		for _, ele := range children {
			child := findOrCreateProgram(ele)
			childmap[ele] = child
		}
		leaf := false
		if len(children) == 0 {
			leaf = true
		}

		p := findOrCreateProgram(name)
		p.Weight = weight
		p.Children = childmap
		p.Leaf = leaf
	}

	root := findRoot(&tree[0])
	fmt.Println(root.Name)
}

func parse(in string) (string, int, []string) {
	s := bufio.NewScanner(strings.NewReader(in))

	s.Split(bufio.ScanWords)

	var arr []string
	num := regexp.MustCompile(`\((\d+)\)`)
	comma := regexp.MustCompile(`(\w+),`)
	for s.Scan() {
		text := s.Text()
		if text == "->" {
			continue
		} else if num.MatchString(text) {
			text = num.ReplaceAllString(text, "$1")
		} else if comma.MatchString(text) {
			text = comma.ReplaceAllString(text, "$1")
		}
		arr = append(arr, text)
	}

	var name string
	var weight int
	var children []string
	for i, ele := range arr {
		if i == 0 {
			name = ele
		} else if i == 1 {
			weight, _ = strconv.Atoi(ele)
		} else {
			children = append(children, ele)
		}
	}

	return name, weight, children
}

func findOrCreateProgram(name string) *Program {
	for i := 0; i < len(tree); i++ {
		if tree[i].Name == name {
			return &tree[i]
		}
	}
	p := Program{Name: name}
	tree = append(tree, p)
	return &tree[len(tree)-1]
}

func findRoot(node *Program) *Program {
	var root *Program
	for i := 0; i < len(tree); i++ {
		if tree[i].Children[node.Name] != nil {
			root = findRoot(&tree[i])
		}
	}
	if root != nil {
		return root
	} else {
		return node
	}
}

/* PartTwo solves the following addition to Part One:
For any program holding a disc, each program standing on that disc forms a sub-tower. Each of those sub-towers are supposed to be the same weight, or the disc itself isn't balanced. The weight of a tower is the sum of the weights of the programs in that tower.

In the example above, this means that for ugml's disc to be balanced, gyxo, ebii, and jptl must all have the same weight, and they do: 61.

However, for tknk to be balanced, each of the programs standing on its disc and all programs above it must each match. This means that the following sums must all be the same:

ugml + (gyxo + ebii + jptl) = 68 + (61 + 61 + 61) = 251
padx + (pbga + havc + qoyq) = 45 + (66 + 66 + 66) = 243
fwft + (ktlj + cntj + xhth) = 72 + (57 + 57 + 57) = 243
As you can see, tknk's disc is unbalanced: ugml's stack is heavier than the other two. Even though the nodes above ugml are balanced, ugml itself is too heavy: it needs to be 8 units lighter for its stack to weigh 243 and keep the towers balanced. If this change were made, its weight would be 60.

Given that exactly one program is the wrong weight, what would its weight need to be to balance the entire tower?
 */

func PartTwo() {

}