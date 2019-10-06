//lezione 15 novembre 2018

package main

import (
	. "fmt"
)

func main () {
	a := []int {15,9,-1}
	s=0
	for _,x:= range a {
		s +=x
	}
	println(s)
}

// le stringhe sono slice immutabili di array di byte 
/*

// ESERCIZIO PER ESERCITARSI CON LO SLICE 

func main (){
	a:= []int {15,9,-1,2,5,7}
	s:= a[2:6]
	t:= s[1:3]
	t[1]=100
	Println(a)
}

// DISEGNO DEL PROG ALLA LAVAGNA 

	15	9	-1	2	5	5
a	^						^
s			^				^
t				^		^


// ESERCIZIO TRE

func aumenta(s[]int) {
	for i,_:= range s {
		s[i]++
	}
}

func main(){

	a:=[]int{15,9,-1,2,5,7}
	aumenta(a[0:3])
	Println(a)

}
