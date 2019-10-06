package main 

import ."fmt"
// 
func prova(x float64, a...int)float64{
	s:=0 
	for _,y:= range a {
		s+=y
	}
	return float64(s)/x
}

func main() {
	a:= []int{3,4,5} 
	Println(prova(2.0 ,5,6,7))
}
