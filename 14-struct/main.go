package main

import "fmt"

type user struct { //custom type
	name string //member variable or property
	age  int
}

func main() {
	var user1 user
	user1 = user{ //instantiate
		name: "Habib",
		age:  30,
	}
	fmt.Println("Name:", user1.name)
	fmt.Println("Age:", user1.age)

	user2 := user{ //Instance or object
		name: "Idrath",
		age:  23,
	}
	fmt.Println("Name:", user2.name)
	fmt.Println("Age:", user2.age)

}
