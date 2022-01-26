package main

import "fmt"

func main() {
	fmt.Println("arrays in Golang")

	var fruitsList [5]string

	fruitsList[0] = "apple"
	fruitsList[1] = "peach"
	fruitsList[3] = "watermelon"
	fmt.Println(fruitsList)
	fmt.Println(len(fruitsList))

}
