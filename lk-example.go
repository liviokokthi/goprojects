// Lezione 15 Ottobre 2018
// input output

package main

import (
	. "fmt"
	_ "math"
)

func main() {

	var x, y = 3, 1.5 // se non metto i due punti prima dell'uguale le variabili fa differenza di tipo , con il := c'è un'assegnazione multipla

	print(y, x)                      //non c'è nessun spazio e non va a capo , usare "\n"per andare a capo con print , anche se più comodo usare sempre println
	println(x, "più uno uguale ", 4) // print che va a capo dopo che hai stampato e che inserisce spazio tra un oggetto e l'altro

	//funzione Scan

	var h, m int

	print("inserisci ore e minuti:")
	Scan(&h, &m)
	println("mancano", (23-h)*60+60-m, "minuti a Mezzanotte")

	//emoji time

	var s, z string

	print("inserisci Emoji :  ")
	Scan(&s, &z)
	println("Ecco le Emoji e il carattere che hai inserito ", s, z)

	// Synctactic Sugar : x++ ------> x=x+1  stessa identica roba

}
