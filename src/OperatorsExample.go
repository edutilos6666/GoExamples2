package main

import (
	"fmt"
)

type OperatorsExample struct  {

}

func (ops *OperatorsExample) test_ArithmeticOps() {
	var x , y = 10 , 3
	var resAdd = x + y
	var resSub  = x - y
	var resMul = x * y
	var resDiv = x / y
	var resMod = x % y
	x++
	var resInc = x
	x--
	var resDec = x

	fmt.Printf("<<Arithmetic Operators>>\n")
	fmt.Println("x = ", x, "y = ", y)
	fmt.Printf("x + y = %d\n", resAdd)
	fmt.Printf("x - y = %d\n", resSub)
	fmt.Printf("x * y = %d\n", resMul)
	fmt.Printf("x / y = %d\n", resDiv)
	fmt.Printf("x %% y = %d\n", resMod)
	fmt.Printf("x++ = %d\n", resInc)
	fmt.Printf("x-- = %d\n", resDec)
}



func (ops *OperatorsExample) test_RelationalOps() {
	var x , y = 10 , 3
	var eq = x == y
	var ne = x != y
	var gt = x > y
	var ge = x >= y
	var lt = x < y
	var le = x <= y
	fmt.Println("<<Relational Operators>>")
	fmt.Println("x = ", x, "y = ", y)
	fmt.Printf("x == y = %t\n", eq)
	fmt.Printf("x != y = %t\n", ne)
	fmt.Printf("x > y = %t\n", gt)
	fmt.Printf("x >= y %t\n", ge)
	fmt.Printf("x < y = %t\n", lt)
	fmt.Printf("x <= y = %t\n", le)
}


func (ops *OperatorsExample) test_LogicalOps() {
	var x , y = true , false
	var and = x && y
	var or = x || y
	var not_x = !x
	var not_y = !y
	fmt.Println("<<Logical Operators>>")
	fmt.Println("x = ", y , "y = ", y)
	fmt.Printf("x && y = %t\n", and)
	fmt.Printf("x || y = %t\n", or)
	fmt.Printf("!x = %t\n", not_x)
	fmt.Printf("!y = %t\n", not_y)
}


func (ops *OperatorsExample) test_BitwiseOps() {
	var x , y = 10 , 23
	var and = x & y
	var or = x | y
	var xor = x ^ y
	var leftShift = x << 2
	var rightShift = x >> 2
	fmt.Println("<<Bitwise Operators>>")
	fmt.Println("x = ", x , "y = ", y)
	fmt.Printf("x & y = %d\n", and)
	fmt.Printf("x | y = %d\n", or)
	fmt.Printf("x ^ y = %d\n", xor)
	fmt.Printf("x << 2 = %d\n", leftShift)
	fmt.Printf("x >> 2 = %d\n", rightShift)
}


