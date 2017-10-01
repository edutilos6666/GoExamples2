package main

import (
	"fmt"
)

type SliceExample struct {

}


func (self *SliceExample) example1()  {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d\n", len(numbers))
	fmt.Printf("cap = %d\n", cap(numbers))
	fmt.Printf("values = %v\n", numbers)
	//nil slice
	var numbers2 []int
	fmt.Println("<<nil slice>>")
	fmt.Printf("len = %d\n", len(numbers2))
	fmt.Printf("cap = %d\n", cap(numbers2))
	fmt.Printf("values = %v\n", numbers)
	if(numbers2 == nil) {
		fmt.Println("numbers2 is nil")
	} else {
		fmt.Println("numbers2 is not nil")
	}

	//append
	fmt.Println("<<append example>>")
	numbers2 = append(numbers2, 10)
	fmt.Printf("numbers2 = %v\n", numbers2)
	numbers2 = append(numbers2, 20, 30, 40, 50)
	fmt.Printf("numbers2 = %v\n", numbers2)
	//copy
	var numbers3 = make([]int , cap(numbers2), cap(numbers2)*2)
	copy(numbers3, numbers2)
	fmt.Printf("numbers3 = %v\n", numbers3)
}