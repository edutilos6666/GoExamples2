package main

import (
	"fmt"
)

type RangeExample struct  {

}

func (self *RangeExample) example1() {
	var numbers = []int{10, 20, 30, 40, 50}
	fmt.Println("<<numbers>>")
	for i, n := range(numbers) {
		fmt.Printf("numbers[%d] = %d\n", i, n)
	}
	fmt.Println()
	var doubleNumbers = []float64{10, 20, 30, 40,50}
	fmt.Println("<<double numbers>>")
	for i, n := range(doubleNumbers) {
		fmt.Printf("doubleNumbers[%d] = %.2f\n", i, n)
	}
	fmt.Println()
	var names = []string{"foo", "bar", "bim"}
	fmt.Println("<<names>>")
	for i, name := range(names) {
		fmt.Printf("names[%d] = %s\n", i , name)
	}
	fmt.Println()
	var countries = map[string]string {"England":"Longdon", "Germany":"Berlin",
	                "Turkey":"Ankara", "Azerbaijan":"Baku"}
	fmt.Println("<<Countries & Capitals>>")
	for country, capital := range(countries) {
		fmt.Printf("%s => %s\n", country, capital)
	}
	fmt.Println()
}
