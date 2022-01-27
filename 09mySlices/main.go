package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("coding on Slices")
	var fruitList = []string{"apple", "mango", "banana", "peach"}
	fmt.Printf("type of Slice %T", fruitList)
	fmt.Println(fruitList)

	fruitList = append(fruitList, "guava", "tomato")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highscore := make([]int, 3)

	highscore[0] = 345
	highscore[1] = 789
	highscore[2] = 909

	fmt.Println(highscore)

	highscore = append(highscore, 708, 412, 674)
	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)

	fmt.Println(sort.IntsAreSorted(highscore))

}
