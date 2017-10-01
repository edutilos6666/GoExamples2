package main

import (
	"fmt"
	"math"
)


type InterfaceExample struct {

}

type Shape interface {
	perimeter() float64
	area() float64
	toString() string
}


type Triangle struct {
	a float64
	b float64
	c float64
}

func (self Triangle) perimeter() float64 {
	return self.a + self.b + self.c
}

func (self Triangle) area() float64 {
	var S = self.perimeter() / 2
	return math.Sqrt(S*(S-self.a)*(S-self.b)*(S-self.c))
}


func (self Triangle) toString() string {
	return fmt.Sprintf("Triangle(%.2f,%.2f,%.2f)", self.a, self.b, self.c)
}


type Rectangle struct {
	width float64
	height float64
}


func (self Rectangle) perimeter() float64 {
	return 2*(self.width + self.height)
}

func (self Rectangle) area() float64 {
	return self.width * self.height
}

func (self Rectangle) toString() string {
	return fmt.Sprintf("Rectangle(%.2f, %.2f)", self.width , self.height)
}

type Circle struct {
	r float64
}

func (self Circle) perimeter() float64 {
	return 2*math.Pi*self.r
}

func (self Circle) area() float64 {
	return math.Pi*math.Pow(self.r, 2)
}

func (self Circle) toString() string {
	return fmt.Sprintf("Circle(%.2f)", self.r)
}


func (self *InterfaceExample) getPerimeter(shape Shape) float64 {
	return shape.perimeter()
}

func (self *InterfaceExample) getArea(shape Shape) float64 {
	return shape.area()
}

func (self *InterfaceExample) example1() {
	var s1 , s2 , s3 Shape
	s1 = Triangle{8, 9, 10}
	s2 = Rectangle{10, 20}
	s3 = Circle{10}
	fmt.Printf("<<Infos about %s>>\n", s1.toString())
	fmt.Printf("perimeter = %.2f\n", self.getPerimeter(s1))
	fmt.Printf("area = %.2f\n", self.getArea(s1))
	fmt.Println()
	fmt.Printf("<<Infos about %s>>\n", s2.toString())
	fmt.Printf("perimeter = %.2f\n", self.getPerimeter(s2))
	fmt.Printf("area = %.2f\n", self.getArea(s2))
	fmt.Println()
	fmt.Printf("<<Infos about %s>>\n", s3.toString())
	fmt.Printf("perimeter = %.2f\n", self.getPerimeter(s3))
	fmt.Printf("area = %.2f\n", self.getArea(s3))
	fmt.Println()

}










