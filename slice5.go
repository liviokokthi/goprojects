//19 novembre 2018

package main 

import (.
	"fmt"
)

func ce(x int, f[]int)bool {
	for _,y:= range f {
		if x==y {
			return true
		}
	}
	return false 
}

func trova(pagliaio,ago []int) []int {
	f:= make([]int,0)
	for _,x:= range pagliaio {
		if ce(x,ago) {
			f=append(f,x)
		}
	}
	return f 
}

func main () {
	Println(trova([]int {1,2,0,4,3,0,1,1,3}, int[] {0,1}))
}