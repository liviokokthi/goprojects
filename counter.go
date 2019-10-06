// Contare il numero di parole in una stringa 

/* Package unicode reference manual: http://golang.org/pkg/unicode */

package main

import (
            ."fmt"
                    )
func main(){
	
	s:="Ciao mamma   guarda come mi diverto"

	c:=1 // il contatore parte da "1" visto che la prima parola del range parte senza uno spazio davanti


	for _, r:= range s{

	    if r==' '{      // se nel range del testo, trovo uno spazio vuoto, allora conto una parola
 	       c++
	    }
	}

	Println ("Il numero di parole nel testo è",c)  // algoritmo migliorabile, visto che il programma non riesce a distinguere piu spazi di fila, calcolandoli come parole
}