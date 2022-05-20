package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age int = 48

	fmt.Println("Age: ", age)

	// multi-variable declaration
	name, height := "Jeff", 73

	fmt.Printf("My name is %s and I am %d inches tall\n", name, height)

	var (
		upperLeftX  int
		upperLeftY  int
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

	// Go does NOT error on overflow EXCEPT during compile
	var over1 uint8 = 255 // uint8 max value is 255
	over1++               // this is fine. it "wraps" around to  0
	fmt.Println(over1)

	// This is a compiler error
	// var over2 uint8 = 255 + 1

	// converting numbers
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	var a string = strconv.Itoa(i)
	var b, _ = strconv.Atoi(s2)
	var c string = fmt.Sprintf("%.2f", f)
	var d, _ = strconv.ParseFloat(s1, 64)

	fmt.Println(a, b, c, d)
}
