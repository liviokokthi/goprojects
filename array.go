//lezione 15 novembre 2018

package main

import (
	. "fmt"
)

/* esempio esplicativo inziale 

func main()  {
	var a [5]int
	a[0] = 6
	a[1] = 6
	a[2] = 6
	Println (a) // se usassi Println (len(a)) otterei in stampa la lunghezza del vettore , che in questo caso è 5
	// a:= [10]int {2:5} mette direttamente nel vettore 2 il numero 5

}
*/

func main() {
	const dim = 10
	Println ("inserisci i valori nel tuo vettore di dimensione: ",dim) 
	a:=[dim] float64 {}

	for i := 0 ; i<dim ; i++ {
		Scan(&a[i])
	}
	s:= 0.0
	for i := 0 ; i<dim ;i++{
		s+=a[i]
	}
	Println (s/dim)

}