package main

import "fmt"

func main() {
	go myFunc("first")
	go myFunc("second")
	go myFunc("third")
	fmt.Println("Done")
}
func myFunc(a string) {
	for i := 0; i < 3; i++ {
		fmt.Println(a, ":", i)

	}
}
