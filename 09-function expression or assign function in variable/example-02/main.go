package main

import "fmt"

func sum() {
	add(2, 4)
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	sum()
	add(2, 3)
	add := func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}
	add(4, 5)
}

func init() {
	fmt.Println("I will be called first")
}
