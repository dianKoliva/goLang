package main

import "fmt"

func upName(x *string) {

	*x = "huhho"

}

func main() {
	name := "pp"
	memo := &name
	// fmt.Printf("memory : %v value :%v", memo, *memo)
	upName(memo)
	fmt.Println(name)
}
