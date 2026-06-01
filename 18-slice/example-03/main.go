//slice with the help of make function with len

package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 5
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
