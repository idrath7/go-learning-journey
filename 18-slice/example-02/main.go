//slice with the help of slice literal

package main

import "fmt"

func main() {
	s := []int{1, 2, 5} //slice literal,array er shomoi just size likhtam,ekhane likhbo na
	fmt.Println("Slice:", s, "Length:", len(s), "Capacity:", cap(s))
}

//slide literal diye slice banaile length r capacity same hoi
