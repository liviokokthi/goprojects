//lezione 15 novembre 2018

package main

import (
	. "fmt"
)

func aumenta(s[]int) {
	s=append(s,5)

}

func main(){

	a := []int{15,9,-1,2,5,7}
	a = aumenta(a)
	Println(a)

}
