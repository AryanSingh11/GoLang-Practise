package main

import "fmt"

func main() {
	fmt.Println("from function one")

	result, myMessage := proAdder(12, 14, 2, 345, 7, 8)

	fmt.Println(result)
	fmt.Println(myMessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This is the second string output from proAdder function"
}

// func greeter() {
// 	fmt.Println("Namastey from Golang")
// }
