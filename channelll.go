package main

import "fmt"

func worker(done chan bool) {
	fmt.Print("working...")
	fmt.Println("done")
	done <- false
}
func main() {
	done := make(chan bool, 0)
	go worker(done)
	<-done
}
