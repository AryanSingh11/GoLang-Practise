package main

import "fmt"

func main() {
	fmt.Println("This is Defer in Golang")
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()

}

func myDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Print(i)
	}
}
