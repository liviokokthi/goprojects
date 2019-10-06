// Indovina se un numero è PRIMO

package main

import . "fmt"

func main() {

	var n, d int

	Print("Inserire un numero intero_")
	Scan(&n)

	primo := true

	for d = 2; d <= n && primo; d++ {

		if n%d == 0 {

			primo = false

			// Stampa illustrativa dei passi del ciclo
			Println(d, "divide", n)

		} else {
			Println(d, "NON divide", n)
			primo = true
		}
	}

	if primo {
		Println(n, "It's a good price for a chair")
	} else {
		Println(n, "NON è un numero PRIMO")
	}
}
