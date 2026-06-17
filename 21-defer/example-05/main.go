package main

import "fmt"

func calculate() (result int) {
	fmt.Println("First", result)

	show := func() {
		result = result + 10
		fmt.Println("Defer", result)
	}

	defer show()

	result = 5

	p := func(a int) {
		fmt.Println("Ami", a)
	}

	defer p(result)

	defer fmt.Println(result)

	fmt.Println("Second", result)

	defer fmt.Println(10)

	return
}

func main() {
	a := calculate()
	fmt.Println("Main First", a)
}
