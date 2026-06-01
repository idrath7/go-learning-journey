//value assign with the help of append function-->append kintu nijhe ekta slice banabe and sheta banai rakhbe heap eh.

package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1, 2, 3) //first eh empty chilo,erpor value append korlam-->value append korsi slice eh tai (s,1,2,3)
	fmt.Println(s)
}
