package main

import "fmt"

func main() {
	// strings
	var nameOne string = "mario"
	var nameTwo = "detail"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)
	nameThree = "cannel"
	fmt.Println(nameOne, nameTwo, nameThree)
	nameFour := "yoshi"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 25

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 27.98
	var scoreTwo float64 = 88888888888888888888888.41421321321

	name := "ethan"

	fmt.Println(scoreOne, scoreTwo)
	fmt.Println("my age is", ageOne, "and my name is", ageThree)

	//Printf
	fmt.Printf("my age is %v and my name is %v \n", ageOne, name)
	fmt.Printf("my age is %q and my name is %q \n", nameOne, name)
	fmt.Printf("age is of type %T \n", ageOne)
	fmt.Printf("you scored %0.2f points! \n", 225.55)
	fmt.Printf("you scored %0.2f points! \n", 225.55)

	//Sprintf
	var str = fmt.Sprintf("my age is %v and my name is %v \n", ageOne, name)
	fmt.Println(str)
}
