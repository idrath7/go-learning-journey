package main

import "fmt"

func main() {
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(5, 6)
}

func init() {
	fmt.Println("I will be called first")
}
