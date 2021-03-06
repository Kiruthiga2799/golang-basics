package main

import (
	"bytes"
	"encoding/json"
	"log"
)

type user struct {
	username string
	password string
	gmail    string
}
type userDb struct {
	users []user
	Type  string
}

func main() {
	users := []user{
		{username: "kirthi", password: "abc", gmail: "kirthi@gmail.com"},
		{username: "deva", password: "def", gmail: "deva@gmail.com"},
		{username: "clera", password: "ghi", gmail: "clera@gmail.com"},
	}
	Db := userDb{users: users, Type: "simple"}
	var buf = new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(Db)
	f, err := os.create("user.Db.json")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.close()
	io.copy(f, buf)
}
