package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	d, r := ioutil.ReadFile("ak.txt")
	if r != nil {
		fmt.Println(r)
	}
	f, err := os.Open("ak.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	k, err := reader.Peek(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(k))

}
