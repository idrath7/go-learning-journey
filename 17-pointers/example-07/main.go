package main

import "fmt"

type idrath struct {
	name   string
	age    int
	salary float64
}

func main() {
	obj := idrath{
		name:   "idrath",
		age:    23,
		salary: 0,
	}
	p := &obj
	fmt.Println(p.name)
}
