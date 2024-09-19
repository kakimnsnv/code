package main

import (
	"fmt"

	"github.com/kakimnsnv/golang-kbtu/lectures/3/1/add"
	"github.com/kakimnsnv/golang-kbtu/lectures/3/2/subtract"
)

func main() {
	res := add.Add(5, 3)
	fmt.Println("Sum:", res)
	res = subtract.Subtract(5, 3)
	fmt.Println("Subtraction:", res)
	fmt.Println("This is package 3")
}
