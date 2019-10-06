//lezione 15 novembre 2018

package main

import (
	. "fmt"
)

func main () {
	var a []int
	//lo slice in questo modo ci risolve il dover restare all'interno della dimensione del vettore.
	for {
		var x int
		Scan(&x)
		if x == -1 {
			break
		}
		a = append (a,x)
		Println (a , len(a),cap(a))
	}
	s:=0.0

	for i := 0 ; i< len(a); i++ {
		s+= float64(a[i])
	}
	Println(s/float64(len(a)))
	
	/*
	primo esempio

	a = make([]int, 10 , 10)		//make viene usato per creare vari tipi di oggetti nuovi in go
	a= append(a,5)		//aggancia un ulteriore elemnto alla slice , aumentando dinamicamente anche la capacità
	Println (a , len(a),cap(a))

	*/
}