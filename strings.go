package main

import "fmt"

func main() {
	age := 12
	name := "Diane"
	fmt.Print("hello, ")
	fmt.Print("world ")
	//println on new line
	fmt.Println("i am ")
	fmt.Println("i am ")

	///printf for formating strings
	fmt.Printf("my name is %v and y age is %v", name, age)

	//sprintf (save formatte strings)
	var str = fmt.Sprintf("age is %v and names is %v ", age, name)
	fmt.Println(str)
}
