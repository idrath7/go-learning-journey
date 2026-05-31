package main

import "fmt"

func getnumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func main() {
	a := 30
	b := 20
	p, q := getnumbers(a, b)
	fmt.Println(p)
	fmt.Println(q)

}
