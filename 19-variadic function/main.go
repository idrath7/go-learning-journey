package main

import "fmt"

func print(numbers ...int) { //jotho kushi number pass kora jai,no limit
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

}

func main() {
	print(5, 6, 7, 8, 9, 10)
}
