package main

import "fmt"

func print(numbers *[3]int) {
	fmt.Println(numbers) //Ami print(&arr) diye array-er address pathacchi.Function (numbers *[3]int) diye oi address ta pointer value hisebe receive kortese.
}

func main() {
	arr := [3]int{1, 2, 3}
	print(&arr) //pass by reference
}
