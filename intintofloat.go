package main

import "fmt"

func divide(x,y int64) int64 {
    var res = x/y
    return res
}
func main() {
    var n1 int64 = 8
    var n2 int64 = 7
    var k float64 = float64(n1)
    var a float64 = float64(n2)
    var n float64 = k/a
    fmt.Println(n)
}