package main

import "fmt"

func main() {
	var ctr = 0
	for ctr = 0; ctr < 10; ctr++ {
		if ctr == 5 {
			fmt.Println("5 is skipped")
			continue
		}
		fmt.Println(ctr)

	}
	fmt.Println("out off")
}
