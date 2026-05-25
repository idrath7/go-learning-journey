package main

import "fmt"

func call() func(x int, y int) {
	return add
}

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	sum := call() //function expression
	sum(5, 7)
}
