package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	//var digitone int = 8
	//var digittwo float64 = 6.89

	//fmt.Println("sum is ", digittwo+float64(digitone))

	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5))

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
