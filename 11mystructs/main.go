package main

import "fmt"

func main() {
	fmt.Println("structs in golang")

	aryan := User{"Aryan", "ary@go.dev", true, 19}
	fmt.Println(aryan)
	fmt.Println(aryan.Email)

}


type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}