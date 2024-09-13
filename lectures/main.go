package main

import (
	"fmt"
)

func main() {
	Serek := Lecturer{
		Employee: Employee{
			User: User{
				ID:        1,
				FirstName: "Azamat",
				LastName:  "Serek",
				Email:     "a.serek@kbtu.kz",
			},
			Position:   "Senior Lecturer",
			Department: "FIT",
			Hours:      40,
		},
	}

	Kakim := Student{
		User: User{
			ID:        2,
			FirstName: "Kakim",
			LastName:  "Nyssanov",
			Email:     "fkadjskh@kbtu.kz",
		},
		Address: "Almaty, Kazakhstan",
		Courses: make([]Course, 0),
	}

	Serek.Teach()
	Kakim.Study()
}

// func remove(slice []int, s int) []int {
// 	return append(slice[:s], slice[s+1:]...)
// }

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

type Student struct {
	User
	Address string
	Courses []Course
}

func (s Student) Study() {
	fmt.Println("Studying...")
}

type Study interface {
	Study()
}

type Employee struct {
	User
	Position   string
	Department string
	Hours      int
}

type Teacher interface {
	Teach()
}

type Lecturer struct {
	Employee
	Degree string
}

func (l Lecturer) Teach() {
	fmt.Println("Teaching...")
}

type Course struct {
	ID          int
	Name        string
	Department  string
	CreditHours int
}
