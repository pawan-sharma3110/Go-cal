package helper

import "fmt"

var Num1 int
var Num2 int
var Operator string

func UserInputs() {
	fmt.Println("Enter fisrt number")
	fmt.Scan(&Num1)
	fmt.Println("Enter second number")
	fmt.Scan(&Num2)
	fmt.Println("Enter operator")
	fmt.Scan(&Operator)
}
func UserCondition() {
	switch Operator {
	case "+":
		fmt.Println(Num1 + Num2)
	case "-":

		fmt.Println(Num1 - Num2)
	case "*":
		fmt.Println(Num1 * Num2)
	case "/":
		fmt.Println(Num1 / Num2)
	default:
		fmt.Println("Invalid operator")
	}
}
