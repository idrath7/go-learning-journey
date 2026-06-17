package main

import "fmt"

func calculate() (result int) {
	fmt.Println("First", result)
	show := func() { //anonymous function
		result = result + 10
		fmt.Println("Defer", result)
	}
	defer show() //instant run hobena,return er egh kore run hobe
	result = 5
	fmt.Println("Second", result)
	return
}

func main() {
	calculate()
}
