package main

import "fmt"

func a() {
	i := 0
	fmt.Println("First:", i)
	defer fmt.Println("Second:", i)
	i = i + 1
	fmt.Println("Third:", i)
	return
}

func main() {
	a()
}
