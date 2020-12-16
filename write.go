package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	mes := []byte("hello")
	err := ioutil.WriteFile("kirthi.txt", mes, 0645)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("kirthi.txt")
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("kirthu")
	f.Sync()
}
