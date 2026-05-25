package main

import "fmt"

func add(a int, b int) { //parameter-->a,b
	c := a + b
	fmt.Println(c)
}

func main() {
	add(2, 5) //argument-->2,5
}
