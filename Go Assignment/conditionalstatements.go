package main

import (
	"fmt"
)

func main() {
	num := 10
	if num%2 == 0 {
		fmt.Println("The number", num, "is even")
		//return // Return early if the number is even
	}
	fmt.Println("The number", num, "is odd")

	num = 99

	if num <= 50 {
		fmt.Println("The number", num, "is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("The number", num, "is between 51 and 100")
	} else {
		fmt.Println("The number", num, "is greater than 100")
	}

	

	if num := 10; num%2 == 0 {
		// num is scoped to this if block
		fmt.Println("The number", num, "is even")
		// num is not accessible outside this block
	} else {
		fmt.Println("The number", num, "is odd")
	}

	/*************Avoiding automatic semicolon insertion************/

	// The following code will result in a compilation error:
	// Uncommenting this block will show the error.
	// if num%2 == 0 {
	// 	fmt.Println("The number is even")
	// };  // Compiler will insert a semicolon, causing an error

	// Instead, use this correct syntax:
	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}
	// In Go, it's idiomatic to return early from the function if possible
	num = 10
	if num%2 == 0 {
		fmt.Println("The number", num, "is even")
		return
	}
	fmt.Println("The number", num, "is odd")
}