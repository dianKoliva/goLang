package main

import (
	"fmt"
	"math"
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

//different return types

func main() {
	// sayGreeting("Pretty")
	// cycle([]string{"cloud", "tifa", "mummy"}, sayGreeting)
	// a1 := area(10.5)
	// fmt.Println(a1)

}
