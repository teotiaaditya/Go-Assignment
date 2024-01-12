package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare a variable without assigning a value
	var age int
	fmt.Println("My age is", age)
	age = 22
	fmt.Println("My age is", age)
	age = 54
	fmt.Println("My new age is", age)

	// Declare a variable with an initial value
	var initialAge int = 29
	fmt.Println("My age is", initialAge)

	var inferredAge = 29
	fmt.Println("My age is", inferredAge)

	// Declare multiple variables
	var width, height = 100, 50
	fmt.Println("Width is", width, "Height is", height)

	// Short hand declaratio
	count := 10
	fmt.Println("Count =", count)

	// Declare variables of different types in a single statement
	var (
		name = "Aditya Teotia"
		roll int = 850
	)
	fmt.Println("Name:", name, "Enrollment Number:", roll, "Height:")

	a, b := 145.8, 543.8
	minValue := math.Min(a, b)
	fmt.Println("Minimum value is", minValue)
}
