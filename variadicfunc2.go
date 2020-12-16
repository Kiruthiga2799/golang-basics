package main

import (
	"fmt"
	"strings"
)

func joinstr(element ...string) string {
	return strings.Join(element, "-")
}

func main() {

	element := []string{"kirthi", "btech", "IT"}
	fmt.Println(joinstr(element...))
}
