package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type student struct {
		Name       string `json:"name"`
		Age        int
		Percentage int
		Id         int
		Subject    []string
	}
	std1 := &student{Name: "kirthi", Age: 21, Percentage: 99, Id: 123, Subject: []string{"eng", "tamil", "maths"}}
	file, _ := json.Marshal(std1)
	fmt.Println(string(file))

	file1 := `{"Name" :"Kirthi","Age" :21,"Id" :123,"Subject":["eng","tamil","maths"]}`
	std2 := &student{}
	json.Unmarshal([]byte(file1), std2)
	fmt.Println(std2.Name)

}
