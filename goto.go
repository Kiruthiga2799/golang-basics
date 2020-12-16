package main

import "fmt"

func main() {
	var age int
election:
	fmt.Print("what is your age : ")
	fmt.Scanln(&age)
	if age > 17 {
		fmt.Print("eligible to vote.\n")
		goto election
	} else {
		fmt.Print("not eligible to vote")
	}

}
