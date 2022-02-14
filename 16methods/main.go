package main

import "fmt"

func main() {
	fmt.Println("structs in golang")

	aryan := User{"Aryan", "ary@go.dev", true, 19}
	fmt.Println(aryan)
	fmt.Println(aryan.Email)
	aryan.GetStatus()
	aryan.Newmail()
	fmt.Println(aryan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("The Status is", u.Status)
}

func (u User) Newmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email is", u.Email)
}
