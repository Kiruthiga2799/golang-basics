package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		channel <- "kirthi"
		channel <- "btech"
	}()
	a := <-channel
	
	fmt.Println(a)
	fmt.Println(<-channel)
}
