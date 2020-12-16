package main

import "fmt"

func main() {
    defer printText("kirthi")
    defer printText("btech")
    printText("hi")
    defer printText("2020")
}
func printText(p string){
    fmt.Println(p)
}