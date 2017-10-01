package main

import "fmt"

type FunctionExample struct {

}


//simple functions
func (runner *FunctionExample) example1() {
	var x , y = 10 , 3
	var max = findMax(x,y)
	var min = findMin(x, y)
	fmt.Println("x = ", x, "y = ", y)
	fmt.Println("max = ", max)
	fmt.Println("min = ", min)
}

func findMax(x, y int) int {
   if x > y {
	   return x
   } else {
	   return y
   }
}

func findMin(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}



//function with multiple returns
func (runner *FunctionExample) example2() {
	var sum , avg, min , max = doubleSummaryStatistics(10.0, 20.0, 30.0, 40.0, 50.0)
	fmt.Println("<<Double Summary Statistics>>")
	fmt.Printf("sum = %.2f\n", sum)
	fmt.Printf("avg = %.2f\n", avg)
	fmt.Printf("min = %.2f\n", min)
	fmt.Printf("max = %.2f\n", max)

	var numbers = []float64{10.0, 20.0, 30.0, 40.0, 50.0}
	sum , avg , min, max = doubleSummaryStatisticsWithArray(numbers)
	fmt.Println()
	fmt.Println("<<Double Summary Statistics with Array>>")
	fmt.Printf("sum = %.2f\n", sum)
	fmt.Printf("avg = %.2f\n", avg)
	fmt.Printf("min = %.2f\n", min)
	fmt.Printf("max = %.2f\n", max)
}

func doubleSummaryStatistics(numbers ...float64) (float64, float64, float64, float64) {
	var sum , avg, min , max float64
	min = numbers[0]
	max = numbers[0]
	for _, el := range(numbers) {
		sum += el
		if min > el {
			min = el
		}
		if max < el {
			max = el
		}
	}

	avg = sum / float64(len(numbers))

	return sum, avg, min , max

}

func doubleSummaryStatisticsWithArray(numbers []float64) (float64, float64, float64, float64) {
	var sum , avg, min ,max float64
	min = numbers[0]
	max = numbers[0]

	for _, el := range(numbers) {
		sum += el
		if min > el {
			min = el
		}
		if max < el {
			max = el
		}
	}

	avg = sum / float64(len(numbers))
	return sum , avg, min , max
}


//call by reference
func (runner *FunctionExample) example3() {
	var x , y int = 10 , 3
	fmt.Println("<<Before Swap>>")
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
	swapByRef(&x, &y)
	fmt.Println("<<After Swap>>")
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}
func swapByRef(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}


//recursion
func (self *FunctionExample) factorial(number int64) int64 {
	if number <= 1 {
		return 1
	} else {
		return number*self.factorial(number-1)
	}
}

func (self *FunctionExample) fibonacci(index int64) int64 {
	if index == 0 || index == 1 {
		return index
	}
	return self.fibonacci(index-1) + self.fibonacci(index-2)
}

func (self *FunctionExample) test_Recursion() {
	var fact_10 = self.factorial(10)
	var fact_5 = self.factorial(5)
	fmt.Println("<<Recursion Examples>>")
	fmt.Printf("10! = %d\n", fact_10)
	fmt.Printf("5! = %d\n", fact_5)
	fmt.Println()
	var i int64
	for i = 0 ; i < 15; i++  {
		fmt.Printf("%d ; ", self.fibonacci(i))
	}
	fmt.Println()
}