package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() {

	// grett := "hello pretty"

	// fmt.Println(strings.Contains(grett, "hello"))
	// fmt.Println(strings.ReplaceAll(grett, "hello", "hi"))
	// fmt.Println(strings.Index(grett, "y"))

	ages := []int{20, 10, 40, 30, 80, 50, 70}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 10)
	fmt.Println(index)

	names := []string{"yoshi", "pp", "cc", "mario"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "cc"))
}
