package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Println("Hello, ", name)
}

func sayBye(name string) {
	fmt.Println("Goodbye, ", name)
}

func cycleNames(names []string, fun func(string)) {
	for _, value := range names {
		fun(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * (r * r)
}

func main() {
	sayGreeting("ethan")
	sayBye("ludwig")

	cycleNames([]string{"cloud", "tifa", "dragon"}, sayGreeting)
	cycleNames([]string{"cloud", "detail", "ambient"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15.5)

	fmt.Println("value :", a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}
