package main

import "fmt"

func main() {
	//var ptr *int
	//fmt.Println("code on pointers")
	//fmt.Println("value of", ptr)
	//fmt.Printf("type of %T", ptr)

	myNumber := 46
	var ptr = &myNumber

	fmt.Println("actual  value of pointer", ptr)
	fmt.Println("actual  value of pointer", *ptr)

	*ptr = *ptr * 2
	fmt.Println(*ptr)
}
