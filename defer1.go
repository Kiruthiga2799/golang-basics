package main

import "fmt"

func main() {
	defer print("hi")
	defer print("hello")
	defer print("welcome")
	print("home")
	defer print("students")
}
func print(p string) {
	fmt.Println(p)
}
