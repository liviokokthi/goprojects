//19 novembre 2018

package main 

import (
	."fmt"
)

func main() {
	s:= []int {1,2}
	s = append(s,s...)
	Println(s)
	s = append(s,s...)
	Println(s)
	s = append(s,s...)
	Println(s)
	s = append(s,s...)
	Println(s)
	s = append(s,s...)
	Println(s)
}