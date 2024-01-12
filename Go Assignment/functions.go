package main

import "fmt"

// Function declaration
func calculateBill(price, no int) int {
	totalPrice := price * no
	return totalPrice
}

// Function with multiple return values
func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

// Function with named return values
func rectPropsNamed(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // no explicit return value
}

// Blank identifier usage
func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	areaNamed, _ := rectPropsNamed(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f\n", areaNamed)
}
