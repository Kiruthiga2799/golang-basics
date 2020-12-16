package main

import "fmt"

func calc(x, y, z int) (int, int, int) {
	var out1 = x + y + z
	var out2 = x - y - z
	var out3 = x + y - z
	return out1, out2, out3
}
func main() {
	n1 := 3
	n2 := 5
	n3 := 7
	result1, result2, result3 := calc(n1, n2, n3)
	fmt.Println(result1, result2, result3)
}
