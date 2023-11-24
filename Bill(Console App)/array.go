package main

import "fmt"

// PYS - error is because the name is the same for main
func main() {
	var ages [3]int = [3]int{1, 2, 3}
	name := [4]string{"yoshi", "mario", "peach", "browser"}

	fmt.Println(ages, len(ages))
	fmt.Println(name, len(name))

	// slices
	var scores = []int{100, 50, 60}
	scores[2] = 55
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := name[1:3]
	rangeTwo := name[2:]
	rangeThree := name[:3]
	fmt.Println(rangeOne, rangeTwo)
	fmt.Println(rangeThree)
	rangeOne = append(rangeOne, "koppa")

	fmt.Println(rangeOne)
}
