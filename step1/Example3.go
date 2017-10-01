package main

import (
	"fmt"
)
/*

func main() {
	example7()
}
*/


//const keyword
func example7() {
	const id , name, age, wage, active = 1, "foo", 10, 100.0, false
	fmt.Println("<<const keyword>>")
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("age = %d\n", age)
	fmt.Printf("wage = %.2f\n", wage)
	fmt.Printf("active = %t\n", active)
	fmt.Printf("id type = %T\n", id)
	fmt.Printf("name type = %T\n", name)
	fmt.Printf("age type = %T\n", age)
	fmt.Printf("wage type = %T\n", wage)
	fmt.Printf("active type = %T\n", active)
}


//Dynamic multiple var declaration
func example6() {
	id , name , age, wage, active := 1, "foo", 10 , 100.0, true
	fmt.Printf("<<Multiple Variable Declaration>>\n")
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("age = %d\n", age)
	fmt.Printf("wage = %.2f\n", wage)
	fmt.Printf("active = %t\n", active)
	fmt.Printf("id type = %T\n", id)
	fmt.Printf("name type = %T\n", name)
	fmt.Printf("age type = %T\n", age)
	fmt.Printf("wage type = %T\n", wage)
	fmt.Printf("active type  = %T\n", active)
}




//multiple var declaration
func example5() {
	var id , name , age, wage, active = 1, "foo", 10 , 100.0, true
	fmt.Printf("<<Multiple Variable Declaration>>\n")
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("age = %d\n", age)
	fmt.Printf("wage = %.2f\n", wage)
	fmt.Printf("active = %t\n", active)
	fmt.Printf("id type = %T\n", id)
	fmt.Printf("name type = %T\n", name)
	fmt.Printf("age type = %T\n", age)
	fmt.Printf("wage type = %T\n", wage)
	fmt.Printf("active type  = %T\n", active)
}


//dynamic type declaration
func example4() {
	id := 1
	name := "foo"
	age := 10
	wage := 100.0
	active := false

	fmt.Printf("<<Dynamic Data Types>>\n")
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("age = %d\n", age)
	fmt.Printf("wage = %.2f\n", wage)
	fmt.Printf("active = %t\n", active)
	fmt.Printf("id type = %T\n", id)
	fmt.Printf("name type = %T\n", name)
	fmt.Printf("age type = %T\n", age)
	fmt.Printf("wage type = %T\n", wage)
	fmt.Printf("active type = %T\n", active)
}

func example3() {
	var id = 1
	var name = "foo"
	var age = 10
	var wage = 100.0
	var active = true
	fmt.Printf("<<Data Types>>\n")
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("age = %d\n", age)
	fmt.Printf("wage = %.2f\n", wage)
	fmt.Printf("active = %t\n", active)
	fmt.Printf("id type = %T\n", id)
	fmt.Printf("name type = %T\n", name)
	fmt.Printf("age type = %T\n", age)
	fmt.Printf("wage type = %T\n", wage)
	fmt.Printf("active type = %T\n", active)
}


func example2() {
	var id int64
	var name string
	var age int32
	var wage float64
	var active bool
	id = 2
	name = "bar"
	age = 20
	wage = 200.0
	active = false
	fmt.Printf("<<Data Types>>\n")
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("age = %d\n", age)
	fmt.Printf("wage = %.2f\n", wage)
	fmt.Printf("active = %t\n", active)
	fmt.Printf("id type = %T\n", id)
	fmt.Printf("name type = %T\n", name)
	fmt.Printf("age type = %T\n", age)
	fmt.Printf("wage type = %T\n", wage)
	fmt.Printf("active type = %T\n", active)
}


func example1() {
	var id int64
	var name string
	var age int32
	var wage float64
	var active bool

	id = 1
	name = "foo"
	age = 10
	wage = 100.0
	active = true

	fmt.Println("<<Variables>>")
	fmt.Println("id = ", id)
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	fmt.Println("wage = ", wage)
	fmt.Println("active = ", active)
}
