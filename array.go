package main

import "fmt"

func main() {
	var x [4]int
	var i, j int
	for i = 0; i < 4; i++ {
		x[i] = i + 9
	}
	for j = 0; j < 4; j++ {
		fmt.Printf("Element[%d] = %d\n", j, x[j])
	}
}
