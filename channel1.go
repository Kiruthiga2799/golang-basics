package main

import "fmt"

func main() {
	channel := make(chan int)
	go func() {
		channel <- 123
		channel <- 456
	}()
	a := <-channel
	b := <-channel
	fmt.Println(a)
	fmt.Println(b)
}
