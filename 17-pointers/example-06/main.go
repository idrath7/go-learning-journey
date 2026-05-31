//with the help of struct

package main

import "fmt"

type idrath struct {
	name   string
	age    int
	salary float64
}

func main() {
	obj := idrath{
		name:   "Idrath",
		age:    23,
		salary: 0,
	}
	fmt.Println(obj)
}
