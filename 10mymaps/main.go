package main

import (
	"fmt"
)


func main()  {
	fmt.Println("maps in golang")

   languages:= make(map[string]string)

   
   languages["rb"]="ruby"
   languages["py"]="python"
   languages["js"]="javascript"

   fmt.Println("language list", languages)
   fmt.Println("language single", languages["js"])

   delete(languages, "rb")
   fmt.Println("language list new", languages)

//loops in golang
   for key, value := range languages{
	   fmt.Println("for key %v value is %v\n", key, value )
   }
}