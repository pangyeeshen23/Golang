package main

import "fmt"

// PYS - error is because the name is the same for main
func main() {
	var ages [3]int = [3]int{1, 2, 3}
	name := [4]string{"yoshi", "mario", "peach", "browser"}

	fmt.Println(ages, len(ages))
	fmt.Println(name, len(name))
}
