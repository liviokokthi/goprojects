//19 novembre 2018

package main 

import (.
	"fmt"
)

func cambia(f []int,i int ,v int){
	Println(f)
	f[i]=v
}

func main() {
	f:= make([]int,5,10)
	//[0 0 0 0 0 x x x x x]
	s:=f[3:5]
	//[0 0 x x x x x]
	Println(cap(s))
	s = append(s,3)
	//[0 0 3 x x x x]
	Println(cap(s))
	Println(cap(f))
	// [0 0 0 0 0 3 x x x x]

}