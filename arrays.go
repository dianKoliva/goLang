package main

import "fmt"

func main() {
	// var ages [3]int= [3]int{20,25,30}
	var ages = [3]int{20, 25, 30}
	var names = [3]string{"pp", "kk", "cc"}
	fmt.Println(ages, len(ages))
	fmt.Println(names)

	//slices
	var scores = []int{100, 50, 60}

	//we can't ppend to array but to slice
	scores = append(scores, 45)
	// fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]
	// [2:]
	// [:3]
	fmt.Println(rangeOne)
}
