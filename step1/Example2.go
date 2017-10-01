package main
/*import (
	"fmt"
"math/rand"
)*/

/*
func main() {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}
	numbers3 := []int{7, 8, 9}
	sum := sumNumbers(numbers1, numbers2, numbers3)
	fmt.Println("sum = ", sum)

	r := rand.New(rand.NewSource(100))
	l1 := []float64{r.Float64(), r.Float64(), r.Float64()}
	l2 := []float64{r.Float64(), r.Float64(), r.Float64()}
	l3 := []float64{r.Float64(), r.Float64(), r.Float64()}

	fsum := sumFloats(l1, l2, l3)

	fmt.Println("l1 = ", l1)
	fmt.Println("l2 = ", l2)
	fmt.Println("l3 = ", l3)
	fmt.Println("fsum = ", fsum)

	sum = sumList(r.Int(), r.Int(), r.Int(), r.Int())
	fmt.Println("sumList = ", sum)
}
*/

func sumNumbers(numbers1 []int, numbers2 []int, numbers3 []int) int {
	var sum int = 0

	for _, v := range numbers1 {
		sum += v
	}

	for _, v := range numbers2 {
		sum += v
	}

	for _, v := range numbers3 {
		sum += v
	}

	return sum
}


func sumFloats(l1 []float64, l2 []float64, l3 []float64) float64 {
	var sum float64 = 0
	for _, v := range l1 {
		sum += v
	}

	for _, v := range l2 {
		sum += v
	}

	for _, v := range l3 {
		sum += v
	}

	return sum
}


func sumList(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}