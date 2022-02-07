package main

import "fmt"

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }

	names := []string{"pp", "cc", "oo"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, value := range names {
		fmt.Printf("at index %v we have %v \n", index, value)

	}

	//when we don't want no index
	for _, value := range names {
		fmt.Printf(" we have %v \n", value)

	}
}
