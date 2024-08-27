package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func AnimalTest() {
	var a Animal
	d := Dog{"Buddy"}
	a = d
	fmt.Println(a.Speak())

}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (n NokiaPhone) call() {
	fmt.Println("Calling from NokiaPhone")
}

type IPhone struct {
}

func (i IPhone) call() {

	fmt.Println("Calling from IPhone")
}

func PhonesTest() {
	phone := new(NokiaPhone)
	phone.call()
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

type Circle1 struct {
	radius float64
}

func (r Circle1) area() float64 {
	return 3.14 * r.radius * r.radius
}

func ShapeTest() {
	// r := new(Shape)
	s := Rectangle{10, 20}
	fmt.Println(s.area())

	c := Circle1{10}
	fmt.Println(c.area())

}

func main() {
	AnimalTest()
	PhonesTest()
	ShapeTest()
}
