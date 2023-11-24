package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hello there friends !"

	// fmt.Println(strings.Contains(greeting, "there"))
	// fmt.Println(strings.ReplaceAll(greeting, "there", "now"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "fr"))
	// fmt.Println(strings.Split(greeting, "e"))

	// fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 45, 21, 52, 23, 66, 78, 99}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 22)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser"))

}
