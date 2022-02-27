package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in GoLang")

	content := "This file needs to go in learnCodeOnline.txt"

	file, err := os.Create("./mylcofile.txt")

	// if err!=nil{
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)

	// if err!=nil{
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println("length is:", length)

	defer file.Close()

	readfile("./mylcofile.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	// if err!=nil{
	// 	panic(err)
	// }
	checkNilErr(err)

	fmt.Println(databyte)
	fmt.Println(string(databyte))

}


func checkNilErr(err error){
	if err!=nil{
		panic(err)
	}
}