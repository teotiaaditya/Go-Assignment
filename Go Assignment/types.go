package main

import (
	"fmt"
)

func main() {
	// Boolean type
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	// Integer types
	var aInt int = 89
	bInt := 95
	fmt.Println("value of a is", aInt, "and b is", bInt)


	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	// Complex types
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cAdd := c1 + c2
	cMul := c1 * c2
	fmt.Println("sum:", cAdd)
	fmt.Println("product:", cMul)

	// String type
	first := "Aditya"
	last := "Teotia"
	name := first + " " + last
	fmt.Println("My name is", name)

	// Type conversion
	i := 55
	j := 67.8
	sumConverted := i + int(j)
	fmt.Println(sumConverted)

	// Assignment with explicit type conversion
	iAssignment := 10
	var jAssignment float64 = float64(iAssignment)
	fmt.Println("jAssignment", jAssignment)
}
