package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days{
	// 	fmt.Printf("index is and value is %v\n", day)
	// }

	rogvalue := 1

	for rogvalue < 10 {
		if rogvalue == 2 {
			goto mygithub
		}

		if rogvalue == 5 {
			rogvalue++
			continue
		}
		fmt.Println(rogvalue)
		rogvalue++

	}
mygithub:
	fmt.Println("get to github.com/AryanSingh11")

}
