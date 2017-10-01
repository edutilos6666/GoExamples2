package main

import (
	"fmt"
)

func main() {
	//test_SimpleMath()
	//test_OperatorsExample()
	//test_DecisionMakingExample()
	//test_LoopExample()
	test_FunctionExample()
}



func test_FunctionExample() {
	var runner = FunctionExample{}
	//runner.example1()
	//runner.example2()
	runner.example3()
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
