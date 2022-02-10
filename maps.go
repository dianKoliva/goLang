package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":   45.44,
		"bugari": 500.00,
		"coffee": 3000.000,
	}

	fmt.Println(menu)
	fmt.Println(menu["bugari"])

	for k, v := range menu {
		fmt.Println(k, "-", v)

	}

	phonebook := map[int]string{
		89989989: "pretty",
		2345678:  "kellia",
		98765456: "koko",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[89989989])

}
