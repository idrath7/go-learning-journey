//slice with the help of make function with length and capacity

package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	s[0] = 2
	s[2] = 10
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
