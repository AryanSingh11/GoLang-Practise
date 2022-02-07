package main

import "fmt"

func main() {
	fmt.Println("ifelse in golang")

	loginCount := 23
	var result string

	if loginCount < 15 {
		result = "regular user"
	} else if loginCount > 15 {
		result = "power user💪"
	} else {
		result = "exactly 15"
	}

	fmt.Println(result)

	num := 3
	if num < 10 {
		fmt.Println("value is less than 10")
	} else {
		fmt.Println("value greater than 10")
	}

	if err != nil {
		return nil, err
	}

}
