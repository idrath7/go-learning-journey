package main

import "fmt"

func main() {
	fmt.Println("Welcome to the application")

	var name string

	fmt.Println("Enter your name -")
	fmt.Scanln(&name)

	fmt.Println("-----", name)

}
