package main

import "fmt"

func main() {
	const a int = 12 // typed int (duh)
	const b = 13.4  // typeless "float64"
	const c = 25 // typeless "int"

	fmt.Printf("a: %T, b: %T, c: %T\n", a, b, c)

	var fred int = a // since a is of type int and fred is of type int, this works as expected
	var barney = b // b is typeless and barney is typeless but compiler infers type and makes barney float64

	fmt.Printf("fred is of type %T while barney is of type %T\n", fred, barney)

	// the following is a compile-time error as compiler does NOT implicitly cast a (int) to a float64
	// var illegal float64 = a

	var wilma float64 = float64(a) // using explicit type casting to get desired result

	fmt.Printf("wilma is of type %T\n", wilma)

	// again, without an explicit cast, this is a compile-time error. Compiler infers b is float64 but won't
	// implicitly cast to an int
	// var illegal int = b

	// this is a fun one. c is a typeless int while dino is an explicit float64. Unlike above, you might expect
	// the compiler to infer c as an int and throw a compile-time error but silently casts to a float64
	var dino float64 = c

	fmt.Printf("dino is of type %T\n", dino)


}