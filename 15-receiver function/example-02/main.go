package main

import "fmt"

type user struct {
	name string
	age  int
}

func printUserDetails(usr user) {
	fmt.Println("Name:", usr.name)
	fmt.Println("Age:", usr.age)
}

func (usr user) printDetails() { //custom type thaklei shudo matro receiver function banano jai
	fmt.Println("Name:", usr.name)
	fmt.Println("Age:", usr.age)
}

func (usr user) call(a int) {
	fmt.Println(usr.name)
	fmt.Println(a)
}

func main() {
	var user1 user
	user1 = user{
		name: "Idrath",
		age:  23,
	}
	user1.printDetails()
	user1.call(10)

	user2 := user{
		name: "Habib",
		age:  30,
	}
	printUserDetails(user2)
}
