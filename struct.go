package main

import "fmt"

type person struct {
	name string
	age  int
	phn  int64
}

func main() {
	id := person{name: "kirthi", age: 21, phn: 12345658}
	fmt.Println(id)
}
