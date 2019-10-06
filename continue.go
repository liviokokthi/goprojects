// l'istruzione "continue" : "Continue" interrompe il ciclo piu intero e non l'intero corpo

package main

import (
	. "fmt"
)

func main() {

	var s, sum int

	Println("Inserire una sequenza di interi")

	for {

		Scan(&s)

		if s == 0 {

			break
		}

		if s%2 == 1 {

			continue
		}

		sum += s
	}

	Println("La somma totale dei numeri pari è:", sum)

}
