package main

import (
	"fmt"
	"math/rand"
)

type ArrayExample struct {

}


func (runner *ArrayExample) example1() {
	var numbers = []int {10, 20, 30, 40, 50}
	fmt.Println("<<numbers>>")
	for i, n := range(numbers) {
		fmt.Printf("%d => %d\n", i , n)
	}

	fmt.Println()
	var doubles = []float64{10, 20, 30, 40,50}
	fmt.Println("<<doubles>>")
	for i, n := range(doubles) {
		fmt.Printf("%d => %.2f\n", i, n)
	}

	fmt.Println()
	var names = []string {"foo", "bar", "bim", "pako"}
	fmt.Println("<<names>>")
	for i, n := range(names) {
		fmt.Printf("%d => %s\n", i, n)
	}
	fmt.Println()

	var randomNumbers [10]int
/*	fmt.Println("len randomNumbers = ", len(randomNumbers))
	fmt.Println("cap randomNumbers = ", cap(randomNumbers))
	fmt.Println("random int = ", rand.Int())*/
	for i := 0; i < len(randomNumbers); i++ {
		randomNumbers[i] = rand.Int()%1000
	}

	fmt.Println("<<random Int32 numbers>>")
	for i, el := range(randomNumbers) {
		fmt.Printf("%d => %d\n", i , el)
	}
	fmt.Println()

	//fmt.Println("random double = ", rand.Float64())
	var randomDoubleNumbers [10]float64
	for i := 0; i< len(randomDoubleNumbers); i++ {
		randomDoubleNumbers[i] = rand.Float64()
	}
	fmt.Println("<<random float64 numbers>>")
	for i, el := range(randomDoubleNumbers) {
		fmt.Printf("%d => %.4f\n", i, el)
	}
	fmt.Println()

	for i := 0 ; i< len(randomNumbers) ; i++ {
		randomNumbers[i] = rand.Intn(1000)
	}
	fmt.Println()
	fmt.Println("<<testing rand.Intn(n)>>")
	for i, el := range(randomNumbers) {
		fmt.Printf("%d => %d\n", i, el)
	}
	fmt.Println()
}



//multi-dimensional array
func (runner *ArrayExample) example2() {
    var matrix = [3][4]int {
		{runner.generateRandomNumber(), runner.generateRandomNumber(), runner.generateRandomNumber(), runner.generateRandomNumber()},
		{runner.generateRandomNumber(), runner.generateRandomNumber(), runner.generateRandomNumber(), runner.generateRandomNumber()},
		{runner.generateRandomNumber(), runner.generateRandomNumber(), runner.generateRandomNumber(), runner.generateRandomNumber()},
	}

	fmt.Println("<<multi-dimensional array>>")
	for i := 0 ; i < len(matrix); i++ {
		for j := 0 ; j< len(matrix[i]) ; j++ {
			fmt.Printf("%d; ", matrix[i][j])
		}
		fmt.Println()
	}



	var matrix2 [3][4]int
	for i := 0; i< len(matrix2); i++ {
		for j := 0; j < len(matrix2[i]); j++ {
			matrix2[i][j] = runner.generateRandomNumber()
		}
	}

	fmt.Println("<<multi-dimensional array 2>>")
	for i := 0; i < len(matrix2); i++ {
		for j := 0; j < len(matrix2[i]); j++ {
			fmt.Printf("%d ; ", matrix2[i][j])
		}
		fmt.Println()
	}
}

func (runner *ArrayExample) generateRandomNumber() int {
	return rand.Intn(1000)
}


//passing array as paramater
func (runner *ArrayExample) example3() {
	var callable  = func(intNumbers []int , doubleNumbers []float64, names []string) {
		fmt.Println("<<intNumbers>>")
		for _, el := range(intNumbers) {
			fmt.Printf("%d; ", el)
		}
		fmt.Println()
		fmt.Println("<<doubleNumbers>>")
		for _ , el := range(doubleNumbers) {
			fmt.Printf("%.3f; ", el)
		}
		fmt.Println()
		fmt.Println("<<names>>")
		for _, el := range(names) {
			fmt.Printf("%s; ", el)
		}
		fmt.Println()
	}

	intNumbers := []int {runner.generateRandomNumber(), runner.generateRandomNumber(),
	runner.generateRandomNumber(), runner.generateRandomNumber(), runner.generateRandomNumber()}
	doubleNumbers := []float64 {rand.Float64(), rand.Float64(), rand.Float64(), rand.Float64()}
	names := []string {"foo", "bar", "bim", "pako"}
	callable(intNumbers, doubleNumbers, names)



}

