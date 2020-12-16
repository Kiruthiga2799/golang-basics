package main

import "fmt"

func divide(x, y int) int {
	var res = x / y
	return res
}
func main() {
	n1 := 6
	n2 := 3
	result := divide(n1, n2)
	fmt.Println(result)
}
