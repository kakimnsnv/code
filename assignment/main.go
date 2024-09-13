package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	as1()

	as2()

	as3()

	as4()

	as5()

	as6()

	as7()

}

func as1() {
	fmt.Println("Hello, World!")
}

func as2() {
	var i int = 1
	var f float64 = 0.0
	var s string = "sfs"
	var b bool = true

	t := 1

	fmt.Printf("%v %v %v %v %v\n", i, f, s, b, t)
}

func as3() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("The number is positive.")
	} else if num < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}

	sum := 0

	for i := 0; i <= 10; i++ {
		sum += i
	}

	fmt.Println("Sum:", sum)

	var day int
	fmt.Print("Enter a number (1-7) for the day of the week: ")
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input. Please enter a number between 1 and 7.")
	}
}

func as4() {
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	f, s := swap("hello", "world")
	fmt.Println("Swapped strings:", f, s)

	c, r := divide(10, 3)
	fmt.Println("Quotient:", c, "Remainder:", r)
}

//? as4

func add(a int, b int) int {
	return a + b
}

func swap(str1 string, str2 string) (string, string) {
	return str2, str1
}

func divide(a int, b int) (int, int) {
	c := a / b
	r := a % b
	return c, r
}

// ? as5
func as5() {
	p := Person{"John", 30}
	p.Greet()

	m := Manager{Employee{"Jane", 100}, "Engineering"}
	m.Work()
	fmt.Printf("Manager Name: %v, ID: %v, , Department: %v\n", m.name, m.id, m.deparment)
}

type Person struct {
	name string
	age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %v, and i'm %v years old\n", p.name, p.age)
}

type Employee struct {
	name string
	id   int
}

func (e Employee) Work() {
	fmt.Printf("Employee Name: %s, ID: %d\n", e.name, e.id)
}

type Manager struct {
	Employee
	deparment string
}

func as6() {
	c := Circle{radius: 5}
	r := Rectangle{width: 10, height: 5}

	PrintArea(c)
	PrintArea(r)
}

type Shape interface {
	Area() float64
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func as7() {
	p := Product{"Laptop", 500, 10}
	jsonString, err := p.toJson()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonString)

	p2, err := fromJson(jsonString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func (p Product) toJson() (string, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func fromJson(jsonString string) (Product, error) {
	var p Product
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}
