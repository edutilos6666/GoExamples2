package main

import (
	"fmt"
)

type DecisionMakingExample struct {

}

func (runner *DecisionMakingExample) test_If() {
	var age = 15
	if(age > 0 && age < 10)  {
		fmt.Println("You are a kid.")
	} else if(age >= 10 && age < 20)  {
		fmt.Println("You are a teenager.")
	} else if(age >= 20 && age < 50) {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are an elderly.")
	}

	var username ,password = "foo", "bar"
	if(username == "foo" && password == "bar") {
		fmt.Println("You are logged in as a user.")
	} else if(username == "admin" && password == "admin") {
		fmt.Println("You are logged in as an admin.")
	} else {
		fmt.Println("Login failed.")
	}
}


func (runner *DecisionMakingExample) test_Switch() {
	var marks = 90
	var grade string
	switch marks {
	case 100, 90: grade = "A"
	case 80: grade = "B"
	case 70: grade = "C"
	case 60: grade = "D"
	case 50: grade = "E"
	default: grade = "F"
	}

	switch {
	case grade  == "A" :
		fmt.Println("Excellent!")
	case grade == "B":
		fmt.Println("Very good!")
	case grade == "C" || grade == "D":
		fmt.Println("Good!")
	case grade == "E":
		fmt.Println("Satisfactory!")
	default:
		fmt.Println("Better try again!")
	}


	var x interface {}


	switch x.(type) {
	case nil:
		fmt.Println("type is nil.")
	case int8:
		fmt.Println("type is int8.")
	case int16:
		fmt.Println("type is int16.")
	case int32:
		fmt.Println("type is int32.")
	case int64:
		fmt.Println("type is int64.")
	case float32:
		fmt.Println("type is float32.")
	case float64:
		fmt.Println("type is float64.")
	case bool:
		fmt.Println("type is bool.")
	case string:
		fmt.Println("type is string.")
	default:
		fmt.Println("unknown type to the system.")
	}
}

func (runner *DecisionMakingExample) test_Select()  {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1 , " from c1")
	case c1 <- i1:
		fmt.Println("sent ", i1 , "to c1")
	case i2 = <-c2:
		fmt.Println("received ", i2 , " from c2")
	case c2 <- i2:
		fmt.Println("sent ", i2 , " to c2")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("received ", i3 , "from c3")
		} else {
			fmt.Println("c3 is closed")
		}

	default:
		fmt.Println("No communications")

	}
}


