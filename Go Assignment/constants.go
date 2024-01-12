package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaring a constant
	const a = 50
	fmt.Println(a)

	// Constant reassignment is not allowed
	// const a = 55 // allowed
	// a = 89       // reassignment not allowed

	// Constants with values known at compile time
	var b = math.Sqrt(4)   // allowed
	// const c = math.Sqrt(4) // not allowed
	fmt.Println(b)
	// String Constants, Typed and Untyped Constants
	const hello = "Hello World"
	var name = hello
	fmt.Printf("type %T value %v", name, name)

	// Typed Constant
	const typedHello string = "Hello World"

	// Type Conversion with Constants
	i := 55
	j := 67.8
	sumConverted := i + int(j)
	fmt.Println(sumConverted)

	// Boolean Constants
	/*const trueConst = true
	type myBool bool
	var defaultBool = trueConst    // allowed
	var customBool myBool = trueConst // allowed
	// defaultBool = customBool        // not allowed
*/
	// Numeric Constants
	const aNumeric = 5
	var intVar int = aNumeric
	var int32Var int32 = aNumeric
	var float64Var float64 = aNumeric
	var complex64Var complex64 = aNumeric
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// Numeric Expressions
	var bNumeric = 5.9 / 8
	fmt.Printf("b's type is %T and value is %v", bNumeric, bNumeric)
}
