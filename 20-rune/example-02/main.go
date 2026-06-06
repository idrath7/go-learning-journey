package main

import "fmt"

func main() {
	var a int8 = -128
	var b int8 = 127

	var x uint8 = 255

	var j float32 = 10.23343
	var k float64 = 10.4445

	var flag bool = false

	fmt.Println(a, b, x, j, k, flag)

	r := '🤍'
	s := "My name is Idrath Hossan"

	fmt.Printf("%c\n", r)
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", x)
	fmt.Printf("%.2f\n", j)
	fmt.Printf("%.2f\n", k)
	fmt.Printf("%v\n", flag)
	fmt.Printf("%s\n", s)
	fmt.Printf("The type of s is %T\n", s)
}
