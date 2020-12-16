package main

import "fmt"

func main() {
	subjects := []string{"english", "tamil", "maths", "science", "social"}
	for i, s := range subjects {
		fmt.Println(i, s)
	}
}
