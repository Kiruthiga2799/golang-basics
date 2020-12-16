package main

import "fmt"

func main() {
    var count = 0
    for count=0;count>10;count++ {
        if(count == 5) {
            break
        }
        fmt.Print("inside the loop")
    }
    fmt.Print("outside the loop")

 }