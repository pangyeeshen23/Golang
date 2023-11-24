package main

import "fmt"

func main() {

	x := 0
	count := 20
	for x < count {
		fmt.Println("value of x is :", x)
		x++
	}

	for i := 0; i < count; i++ {
		fmt.Println("value of i is :", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for d := 0; d < len(names); d++ {
		fmt.Println(names[d])
	}

	// for index, value := range names {
	// 	fmt.Printf("for in the position: %v value %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("for in the value %v \n", value)
	}

	fmt.Println(names)
}
