package main  
import (  
    "sort"  
    "fmt"  
)  
func main() {  
  
    intValue := []int{10, 20, 5, 8}  
    sort.Ints(intValue)  
    fmt.Println("Ints:   ", intValue)  
  
    floatValue := []float64{10.5, 20.5, 5.5, 8.5}  
    sort.Float64s(floatValue)  
    fmt.Println("floatValue:   ", floatValue)  
  
    stringValue := []string{"Raj", "Mohan", "Roy"}  
    sort.Strings(stringValue)  
    fmt.Println("Strings:", stringValue)  
  
    str := sort.Float64sAreSorted(floatValue)  
    fmt.Println("Sorted: ", str)  
}  