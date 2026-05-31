//eta just bujhar jnno,eta receiver function er code na

package main

import "fmt"

type user struct {
	name string //member variable or property
	age  int
}

func printUserDetails(usr user) {
	fmt.Println("Name:", usr.name)
	fmt.Println("Age:", usr.age)
}

func main() {
	var user1 user
	user1 = user{ //instantiate
		name: "Idrath",
		age:  23,
	}
	printUserDetails(user1)
	user2 := user{
		name: "habib",
		age:  30,
	}
	printUserDetails(user2)

}
