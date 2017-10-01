package main

import "fmt"

type LoopExample struct {

}


func (runner *LoopExample) example1()  {
	fmt.Println("<<for loop example>>")
	for i := 0 ; i< 10 ; i++ {
		fmt.Println("number = ", i)
	}

	fmt.Println()

	i := 0
	for i < 10 {
		fmt.Println("number = ", i)
		i++
	}

	fmt.Println()

	numbers := [6]float64 {10.0, 20.0, 30.0, 40.0}
	for i , x := range numbers {
		fmt.Printf("numbers[%d] = %.2f\n", i , x)
	}

	fmt.Println()
}



//break
func (runner *LoopExample) example2() {
	fmt.Println("<<Loop with break stmt>>")
	for i:= 0; i< 10 ; i++ {
		if i == 5 {
			break
		}

		fmt.Println("number = ", i)
	}
}

//continue
func (runner *LoopExample) example3() {
	fmt.Println("<<Loop with continue stmt>>")
	for i := 0 ; i< 10 ; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("number = ", i)
	}
}