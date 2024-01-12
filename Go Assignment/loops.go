package main

import "fmt"

// Simple for loop
func main() {
	// Printing numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Using break to terminate a loop
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("line after for loop")

	// Using continue to skip iterations
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Nested for loop
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Label and break
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}

	// Alternative format for while loop
	var i int = 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println()
}
