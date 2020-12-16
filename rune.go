package main

import (
	"fmt"
	"reflect"
)

func main() {
	rune := 'k'
	fmt.Printf("%d \n", rune)
	fmt.Println(reflect.TypeOf(rune))
}
