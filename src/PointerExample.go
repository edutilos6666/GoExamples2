package main

import (
	"fmt"
	"math/rand"
)

type PointerExample struct  {

}


func (runner *PointerExample) example1() {
	var id int64 = 1
	var name = "foo"
	var age = 10
	var wage = 100.0
	var active = true

	var idPtr = &id
	var namePtr = &name
	var agePtr = &age
	var wagePtr = &wage
	var activePtr = &active

	fmt.Println("<<Pointer Example>>")
	fmt.Printf("Address of id = %x and %x\n", &id, idPtr)
	fmt.Printf("Address of name = %x and %x\n", &name, namePtr)
	fmt.Printf("Address of age = %x and %x\n", &age, agePtr)
	fmt.Printf("Address of wage = %x and %x\n", &wage, wagePtr)
	fmt.Printf("Address of active = %x and %x\n", &active , activePtr)
	fmt.Println()
	fmt.Printf("Value of id = %d and %d\n", id , *idPtr)
	fmt.Printf("Value of name = %s and %s\n", name , *namePtr)
	fmt.Printf("Value of age = %d and %d\n", age, *agePtr)
	fmt.Printf("Value of wage = %.2f and %.2f\n", wage, *wagePtr)
	fmt.Printf("Value of active = %t and %t\n", active , *activePtr)
	fmt.Println()
}


//pointer to array
func (runner *PointerExample) example2() {
	var numbers = []int {
		rand.Intn(1000), rand.Intn(1000), rand.Intn(1000),
		rand.Intn(1000), rand.Intn(1000), rand.Intn(1000),
	}
	//const size = len(numbers)   // that does not work

	var numbersPtr [6]*int  // only constant value for len allowed
	for i := 0 ; i< len(numbersPtr); i++  {
		numbersPtr[i] = &numbers[i]
	}

	fmt.Println("<<Iterating over array of pointers>>")
	for _, el := range(numbersPtr) {
		fmt.Printf("address = %x , value = %d\n", el , *el)
	}
	fmt.Println()
}


//pointer to pointer
func (runner *PointerExample) example3() {
	var id int64 = 1
	var name = "foo"
	var age = 10
	var wage = 100.0
	var active = true

	var idPtr = &id
	var namePtr = &name
	var agePtr = &age
	var wagePtr = &wage
	var activePtr = &active

	var idPtrPtr **int64 = &idPtr
	var namePtrPtr **string = &namePtr
	var agePtrPtr **int = &agePtr
	var wagePtrPtr **float64 = &wagePtr
	var activePtrPtr **bool = &activePtr

	fmt.Println("<<Pointer details>>")
	fmt.Printf("Address of id = %x , %x , %x\n", &id , idPtr, *idPtrPtr)
	fmt.Printf("Address of name = %x, %x, %x\n", &name, namePtr, *namePtrPtr)
	fmt.Printf("Address of age = %x, %x, %x\n", &age, agePtr, *agePtrPtr)
	fmt.Printf("Address of wage = %x, %x, %x\n", &wage, wagePtr, *wagePtrPtr)
	fmt.Printf("Address of active = %x, %x, %x\n", &active, activePtr, *activePtrPtr)
	fmt.Println()
	fmt.Printf("Value of id = %d, %d, %d\n", id, *idPtr, **idPtrPtr)
	fmt.Printf("Value of name = %s, %s, %s\n", name, *namePtr, **namePtrPtr)
	fmt.Printf("Value of age = %d, %d, %d\n", age, *agePtr, **agePtrPtr)
	fmt.Printf("Value of wage = %.2f, %.2f, %.2f\n", wage, *wagePtr, **wagePtrPtr)
	fmt.Printf("Value of active = %t, %t, %t\n", active, *activePtr, **activePtrPtr)
	fmt.Println()
}


//passing pointer to the function
func (runner *PointerExample) example4() {
	var callable = func (id *int64 , name *string , age *int32, wage *float64, active *bool) {
		fmt.Println("<<Printing Properties>>")
		fmt.Printf("Id: Address = %x , value = %d\n", id, *id)
		fmt.Printf("Name: Address = %x, value = %s\n", name , *name)
		fmt.Printf("Age: Address = %x , value = %d\n", age, *age)
		fmt.Printf("Wage: Address = %x, value = %.2f\n", wage, *wage)
		fmt.Printf("Active: Address = %x, value = %t\n", active, *active)
		fmt.Println()
	}

	var id int64 = 1
	var name string = "foo"
	var age int32 = 10
	var wage float64 = 100.0
	var active bool = true
	callable(&id, &name, &age, &wage, &active)


	var doubleSummaryStatistics = func(numbers [6]*float64) (float64, float64, float64, float64) {
		var sum , avg , min , max float64
		min = *numbers[0]
		max = *numbers[0]

		for _, el := range(numbers) {
			sum += (*el)
			if min > (*el) {
				min = (*el)
			}

			if max < (*el) {
				max = (*el)
			}
		}

		avg = sum / float64(len(numbers))
		return sum , avg, min , max
	}

	var numbers = []float64 {10, 20, 30, 40, 50, 60}
	var numbersPtr [6]*float64
	for i, _ := range(numbers) {
		numbersPtr[i] = &numbers[i]
	}


	fmt.Printf("numbersPtr = ")
	for _, el := range(numbersPtr) {
		fmt.Printf("%.2f ; ", *el)
	}
	fmt.Println()

	var sum , avg , min , max = doubleSummaryStatistics(numbersPtr)
	fmt.Println("<<Double Summary Statistics>>")
	fmt.Println("sum = ", sum)
	fmt.Println("avg = ", avg)
	fmt.Println("min = ", min)
	fmt.Println("max = ", max)

}

