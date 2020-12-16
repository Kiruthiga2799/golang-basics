package main

import (
	"fmt"
)

func main() {
	fmt.Println(Divide(10, 0))
	fmt.Println(Divide(20, 2))
	fmt.Println(Divide(9, 3))
}
func Divide(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	quotient := num1 / num2
	return quotient
}
