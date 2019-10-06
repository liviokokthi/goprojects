// Lezione 24 Ottobre 2018
// input output
// Comandi rapidi fondamentali : ctrl (+) # = doppia barra commento ; Altgr (+) Pause/MaiuscDX (+) Graffe = {}

package main

import (
	. "fmt"
	_ "math"
)

// Continue : interrompe il corpo intenrno ma non il ciclo

func main() {
	c := 0
	s := 0.0

	for i := 0; i < n; i++ {
		var x float64
		Scan(&x)
		if x == 0 {
			break
		}
		s += x
		c++
	}

	Println(s / float64(c))

}
