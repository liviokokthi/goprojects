// Lezione 24 Ottobre 2018
// Istruzione Continue
// input output
// Comandi rapidi fondamentali : ctrl (+) # = doppia barra commento ; Altgr (+) Pause/MaiuscDX (+) Graffe = {}

package main

import (
	. "fmt"
	_ "math"
)

// Continue : interrompe il corpo intenrno ma non il ciclo

func main() {

	s := 0
	
	for {
		var x int
		Scan(&x)

		if x < 0 {
			Println("STOP - STOP - STOP - STOP - STOP")
			break
		}

		if x%2 == 1 {
			continue
		}
		s += x
	}

	Println("PRIMO BREAK ", x%2)
	Println(s)

}
