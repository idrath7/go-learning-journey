package main

import "fmt"

func main() {
	x := 20
	p := &x
	fmt.Println(p)
	*p = 30
	fmt.Println("Value at address p=", *p)

}
