package main
import "fmt"

func main() {

	fmt.Println("W3Adda - Go For Each with Strings")
	for i,s := range "W3Adda"{
		fmt.Printf("%U represents %c and it is at position %d\n", s, s, i)
	}
}