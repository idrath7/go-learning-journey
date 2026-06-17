//All defer function forms closure must

package main

import "fmt"

func calculateNamed() (result int) {
	fmt.Println("First", result)

	show := func() {
		result = result + 10
		fmt.Println("Defer", result)
	}

	defer show()

	result = 5
	fmt.Println("Second", result)

	return
}

func calculateNormal() int {
	result := 0

	fmt.Println("First", result)

	show := func() {
		result = result + 10
		fmt.Println("Defer", result)
	}

	defer show()

	result = 5
	fmt.Println("Second", result)

	return result
}

func main() {
	a := calculateNamed()
	fmt.Println("Main First", a)

	b := calculateNormal()
	fmt.Println("Main Second", b)
}
