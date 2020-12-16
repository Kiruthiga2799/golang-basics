package main

import "fmt"

func main() {
	var age int
	i := 0
	for i = 0; i < 3; i++ {
		fmt.Print("what is your age : ")
		fmt.Scanln(&age)
		if age > 17 {
			fmt.Print("elligile to vote \n")
		} else {
			fmt.Print("not eligible \n")
		}
	}

}
