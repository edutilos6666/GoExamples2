package main

import (
	"fmt"
	"container/list"
	"math"
)


type StructExample struct  {

}

type Person struct {
	id int64
	name string
	age int32
	wage float64
	active bool
}

func (runner *Person) setId(id int64) {
	runner.id = id
}

func (runner *Person) getId() int64 {
	return runner.id
}

func (runner *Person) setName(name string) {
	runner.name = name
}

func (runner *Person) getName() string {
	return runner.name
}

func (runner *Person) setAge(age int32) {
	runner.age = age
}

func (runner *Person) getAge() int32 {
	return runner.age
}

func (runner *Person) setWage(wage float64) {
	runner.wage = wage
}

func (runner *Person) getWage() float64 {
	return runner.wage
}

func (runner *Person) setActive(active bool) {
	runner.active = active
}

func (runner *Person) isActive() bool {
	return runner.active
}

func (runner *Person) toString() string {
	return fmt.Sprintf("Person(%d,%s,%d,%.2f,%t)", runner.id, runner.name,
	            runner.age, runner.wage, runner.active)
}





func (runner *StructExample) testPerson() {
	var p1, p2, p3 Person
	p1 = Person{1, "foo", 10 , 100.0, true}
	p2 = Person {2, "bar", 20, 200.0, false}
	p3 = Person {3, "bim", 30, 300.0, true}
	fmt.Println("<<People>>")
	fmt.Println("p1 = ", p1.toString())
	fmt.Println("p2 = ", p2.toString())
	fmt.Println("p3 = ", p3.toString())
	fmt.Println()
	p1.setId(1)
	p1.setName("new foo")
	p1.setAge(66)
	p1.setWage(666.666)
	p1.setActive(false)
	fmt.Println("<<p1 after update>>")
	fmt.Println("id = ", p1.getId())
	fmt.Println("name = ", p1.name)
	fmt.Println("age = ", p1.age)
	fmt.Println("wage = ", p1.getWage())
	p1.active = true
	fmt.Println("active = ", p1.active )
}



type PersonDAO struct {
	container map[int64]Person
}

func (runner *PersonDAO) save(person Person) {
	runner.container[person.id] = person
}

func (runner *PersonDAO) update(id int64, newP Person)  {
	delete(runner.container , id)
	runner.container[id] = newP
}

func (runner *PersonDAO) remove(id int64) {
	delete(runner.container, id)
}

func (runner *PersonDAO) findById(id int64) Person {
	return runner.container[id]
}

func (runner *PersonDAO) findAll() *list.List {
	var res  = list.New()
	for _, v := range(runner.container) {
		res.PushBack(v)
	}
	return res
}




func (runner *StructExample) testPersonDAO()  {
	var dao = PersonDAO{make(map[int64]Person)}
	var people = []Person {
		Person{1, "foo", 10, 100.0, true},
		Person{2, "bar", 20, 200.0, false},
		Person{3, "bim", 30, 300.0, true},
	}

	for _, el := range(people) {
		dao.save(el)
	}


	var all = dao.findAll()
	fmt.Println("<<all people>>")
	for el := all.Front(); el != nil ; el = el.Next() {
		value, _ := el.Value.(Person)
		fmt.Println(value.toString())
	}


}



type ComplexNumber struct  {
	re float64
	im float64
}


func (self *ComplexNumber) add(other ComplexNumber) ComplexNumber {
	var res = ComplexNumber{}
	res.re = self.re + other.re
	res.im = self.im + other.im
	return res
}

func (self *ComplexNumber) sub(other ComplexNumber) ComplexNumber {
	var res = ComplexNumber{}
	res.re = self.re - other.re
	res.im = self.im - other.im
	return res
}

func (self *ComplexNumber) mul(other ComplexNumber) ComplexNumber {
	var res = ComplexNumber{}
	res.re = self.re*other.re - self.im*other.im
	res.im = self.re*other.im + self.im*other.re
	return res
}

func (self *ComplexNumber) mulByFactor(factor float64) ComplexNumber {
	var res = ComplexNumber{}
	res.re = self.re*factor
	res.im = self.im*factor
	return res
}

func (self *ComplexNumber) divByFactor(factor float64) ComplexNumber {
	var res = ComplexNumber{}
	res.re = self.re / factor
	res.im = self.im / factor
	return res
}

func (self *ComplexNumber) div(other ComplexNumber) ComplexNumber {
	var factor = math.Pow(other.re , 2) + math.Pow(other.im , 2)
	var temp = ComplexNumber{other.re , -other.im}
	var res = self.mul(temp)
	res = res.divByFactor(factor)
	return res
}

func (self *ComplexNumber) toString() string {
	return fmt.Sprintf("%.2f + i*(%.2f)", self.re , self.im)
}



func (self *StructExample) test_ComplexNumber() {
	var cn1 , cn2 ComplexNumber
	cn1 = ComplexNumber{3, 3}
	cn2 = ComplexNumber{2, 2}
	var cnAdd = cn1.add(cn2)
	var cnSub = cn1.sub(cn2)
	var cnMul = cn1.mul(cn2)
	var cnDiv = cn1.div(cn2)
	fmt.Println("<<ComplexNumber Example>>")
	fmt.Println("cn1 = ", cn1.toString())
	fmt.Println("cn2 = ", cn2.toString())
	fmt.Println("cn1 + cn2 = ", cnAdd.toString())
	fmt.Println("cn1 - cn2 = ", cnSub.toString())
	fmt.Println("cn1 * cn2 = ", cnMul.toString())
	fmt.Println("cn1 / cn2 = ", cnDiv.toString())
}


