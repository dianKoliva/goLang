package main

import "fmt"

func main() {
	name := "pp"
	memo := &name
	fmt.Printf("memory : %v value :%v", memo, *memo)
}
