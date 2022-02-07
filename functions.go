package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("morning %v \n", n)
}

//we can also pass functions to funcs

func cycle(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func area(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {

	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]

	}

	return initials[0], "_"

}

//different return types

func main() {
	// sayGreeting("Pretty")
	// cycle([]string{"cloud", "tifa", "mummy"}, sayGreeting)
	// a1 := area(10.5)
	// fmt.Println(a1)

	fn, sn := getInitials("tifa lockhart")
	fmt.Println(fn, sn)

}
