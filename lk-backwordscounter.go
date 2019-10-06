// Contenitore di Codice d'esempio
// Comandi rapidi fondamentali : ctrl (+) # = doppia barra commento ; Altgr (+) Pause/MaiuscDX (+) Graffe = {}
// Continue : interrompe il corpo intenrno ma non il ciclo
//Scan(&variabile) : inserisci quello che vuoi da tastiera
package main

import . "fmt"

func main() {
	var r, i, t int
	var s string
	print("inserisci range: ")
	Scan(&t)
	print("inserisci un numero e ti stamperò tutti i numeri precedenti fino a 0: ")
	Scan(&r)
	if r <= t {
		for i = 0; i < r; i-- {
			println(i)
		}
		print("corretto?: ")
		Scan(&s)
		if s == "si" {
			println("Complimenti, sapevi che range moltiplicato per il tup numero fa:", t*r)
		} else if s == "no" {
			println("CI STANNOH TRACCIANDOH,STAKKAH STAKKAH")
		}
	} else {
		println("Numero Troppo alto")
	}
}
