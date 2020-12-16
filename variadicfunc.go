package main

import (
	"fmt"
	"strings"
)

func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

func main() {

	fmt.Println(joinstr("kir", "thi"))
	fmt.Println(joinstr("kir", "thi", "ga"))
	fmt.Println(joinstr("k", "i", "r", "t", "h", "i"))

}
