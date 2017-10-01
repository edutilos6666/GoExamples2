package main

import (
	"fmt"
	"container/list"
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
