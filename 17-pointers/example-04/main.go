package main

import "fmt"

func print(numbers [3]int) {
	fmt.Println(numbers)
}

func main() {
	arr := [3]int{1, 2, 3}
	print(arr)
}
