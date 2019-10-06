// Lezione 24 Ottobre 2018
// Istruzione Break
// input output
// Comandi rapidi fondamentali : ctrl (+) # = doppia barra commento ; Altgr (+) Pause/MaiuscDX (+) Graffe = {}

package main

import (
	. "fmt"
	_ "math"
)

/* Istruzione break programma base

func main() {

	s:=0
	for {
		var x int
		Scan(&x)
		if x<0 {
			break
		}
		s+=x
	}

Println(s)

}

*/

func main() {
	print("inserisci un numero e vedremo se è primo: ")
	var n, d int
	Scan(&n)

	for d = 2; d < n; d++ {
		if n%d == 0 {
			break
		}
	}

	//se dopo un ciclo la condizione di controllo è ancora vera vuol dire che il ciclo è finito

	if d < n {
		Println("Non è primo")
	} else {
		Println("è Primo,Yuppi")
	}

}
