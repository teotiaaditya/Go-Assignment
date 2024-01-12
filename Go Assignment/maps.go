package main

import (
	"fmt"
)

func main() {
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
        "mike": 9000,
	}
	employee := "jamie"
    salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)
}
