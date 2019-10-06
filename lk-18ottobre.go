// Lezione 15 Ottobre 2018
// input output

package main

import (
	. "fmt"
	_ "math"
)

func main() {

	/*VARIABILI BOOLEANE

	  x,y := true,false
	  z := x||y
	  if z {
	  }
	*/

	//ESERCIZIO ANNO BISESTILE
	var anno int
	Scan(&anno)
	if (anno%4 == 0) && !(anno%100 == 0 && anno%400 != 0) {
		Println("Bisestile")
	} else {
		Println("Bisestile")
	}
}
