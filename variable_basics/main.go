package main

import "fmt"

func main() {
	var age int = 48

	fmt.Println("Age: ", age)

	// multi-variable declaration
	name, height := "Jeff", 73

	fmt.Printf("My name is %s and I am %d inches tall\n", name, height)

	var (
		upperLeftX int
		upperLeftY int
		lowerRightX int
		lowerRightY int
	)

	fmt.Println("Bounding box coordinates: ", upperLeftX, upperLeftY, lowerRightX, lowerRightY)

	var index, counter int
	index, counter = 0, 1

	// don't do this
	_, _ = index, counter

	// Silent type casting
	sum := 5 + 2.3
	fmt.Println(sum)

}