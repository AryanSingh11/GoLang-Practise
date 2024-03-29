package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Golang time study")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2022, time.January, 22, 00, 55, 00, 00, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
