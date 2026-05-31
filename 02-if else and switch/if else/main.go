package main

import "fmt"

func main() {
	age := 18

	if age > 18 {
		fmt.Println("You are eligible to be married")
	} else if age < 18 {
		fmt.Println("You are not eligible to be married, but you can love someone")
	} else if age == 18 {
		fmt.Println("You are just a teenager, not eligible to be married")
	}

}
