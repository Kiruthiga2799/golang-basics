[package main

import "fmt"

func main() {
	defer print(1)
	defer print(2)
	defer print(3)
	print(0)
}
func print(p int) {
	fmt.Println(p)

}
