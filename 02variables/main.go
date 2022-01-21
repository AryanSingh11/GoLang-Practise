package main

import "fmt"

const LoginToken string = "millionr"

func main() {

	var username string = "Aryan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 262.6783444545
	fmt.Println(smallFloat)
	fmt.Printf("the variable is of type: %T \n", smallFloat)

	var testVariable int
	fmt.Println(testVariable)
	fmt.Printf("the default type is %T \n", testVariable)

	numberofUsers := 3000.090
	fmt.Println(numberofUsers)
	fmt.Printf("variable is of type %T \n", numberofUsers)

	fmt.Println(LoginToken)
	fmt.Printf("type is %T \n", LoginToken)
}
