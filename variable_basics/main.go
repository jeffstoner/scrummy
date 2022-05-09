package main

import "fmt"

func main() {
	var age int = 48

	fmt.Println("Age: ", age)

	// multi-variable declaration
	name, height := "Jeff", 73

	fmt.Printf("My name is %s and I am %d inches tall\n", name, height)
}