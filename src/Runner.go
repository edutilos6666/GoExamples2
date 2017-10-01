package main

import (
	"fmt"
	"errors"
)

func main() {
	//test_SimpleMath()
	//test_OperatorsExample()
	//test_DecisionMakingExample()
	//test_LoopExample()
	//test_FunctionExample()
	//test_StringExample()
	//test_ArrayExample()
	//test_PointerExample()
	//test_StructExample()
	//test_InterfaceExample()
	//test_SliceExample()
	//test_RangeExample()
	//test_Map()
	test_ErrorHandling()
}



func test_ErrorHandling()  {
	var safeDiv = func(a, b int64) (float64, error) {
		if b == 0 {
			return 0, errors.New("Division By Zero Error")
		}
		return float64(a) / float64(b), nil
	}

	var a , b int64 =  10 , 0
	var res, err = safeDiv(a, b)
	if err == nil {
		fmt.Printf("%d / %d = %.2f\n", a, b, res)
	} else {
		fmt.Printf("error message = %s\n", err)
	}

	a, b = 10 , 3
	res, err = safeDiv(a, b)
	if err == nil {
		fmt.Printf("%d / %d = %.2f\n", a, b, res)
	} else {
		fmt.Printf("error message = %s\n", err)
	}


}

func test_Map() {
	var names  map[string]int32
	names = make(map[string]int32)
	names["foo"] = 10
	names["bar"] = 20
	names["bim"] = 30
	fmt.Println("<<printing names>>")
	for k, v := range(names) {
		fmt.Printf("%s => %d\n", k, v)
	}

	//indexing
	value , ok := names["foo"]
	if ok {
		fmt.Printf("foo => %d\n", value)
	} else {
		fmt.Printf("foo is not key\n")
	}

	// delete
	delete(names, "foo")
	fmt.Println("<<printing names after delete>>")
	for k, v := range(names) {
		fmt.Printf("%s => %d\n", k, v)
	}
	fmt.Println()

}

func test_RangeExample() {
	var runner = RangeExample{}
	runner.example1()
}

func test_SliceExample() {
	var runner = SliceExample{}
	runner.example1()
}

func test_InterfaceExample() {
	var runner = InterfaceExample{}
	runner.example1()
}

func test_StructExample() {
	var runner = StructExample{}
	//runner.testPerson()
	//runner.testPersonDAO()
	runner.test_ComplexNumber()
}

func test_PointerExample() {
	var runner = PointerExample{}
	//runner.example1()
	//runner.example2()
	//runner.example3()
	runner.example4()
}


func test_ArrayExample() {
	var runner = ArrayExample{}
	//runner.example1()
	//runner.example2()
	runner.example3()
}


func test_StringExample() {
	var runner = StringExample{}
	runner.example1()
}

func test_FunctionExample() {
	var runner = FunctionExample{}
	//runner.example1()
	//runner.example2()
	//runner.example3()
	runner.test_Recursion()
}

func test_LoopExample() {
	var runner = LoopExample{}
	//runner.example1()
	//runner.example2()
	runner.example3()
}


func test_DecisionMakingExample()  {
	var runner = DecisionMakingExample{}
	//runner.test_If()
	//runner.test_Switch()
	runner.test_Select()
}


func test_OperatorsExample() {
	var runner = OperatorsExample {}
	//runner.test_ArithmeticOps()
	//runner.test_RelationalOps()
	//runner.test_LogicalOps()
	runner.test_BitwiseOps()
}


func test_SimpleMath()  {
	var sm = SimpleMath { 10.0 , 3.0 }
	var resAdd = sm.add()
	var resSub = sm.sub()
	var resMul = sm.mul()
	var resDiv = sm.div()
	fmt.Printf("<<SimpleMath Results>>\n")
	fmt.Println("x = ", sm.x, "y = ", sm.y)
	fmt.Printf("x + y = %.2f\n", resAdd)
	fmt.Printf("x - y = %.2f\n", resSub)
	fmt.Printf("x * y = %.2f\n", resMul)
	fmt.Printf("x / y = %.2f\n", resDiv)
}
