package main

import (
	"errors"
	"fmt"
)

type info interface {
	getInfo() (uint, string, error)
}

type User struct {
	Id   uint
	Name string
}

func (u User) getInfo() (uint, string, error) {
	if u.Id <= 0 || u.Name == "" {
		return 0, "", errors.New("id or name is empty")
	}
	return u.Id, u.Name, nil
}

func main() {
	user1 := User{
		Id:   0,
		Name: "Alice",
	}
	fmt.Println(user1.getInfo())
}
