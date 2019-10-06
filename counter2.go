// Contare il numero di parole in una stringa 

/* Package unicode reference manual: http://golang.org/pkg/unicode */

package main

import (
            ."fmt"
					)
					
func main(){
	
	s:="Ciao mamma   guarda come mi      diverto ancora"
	c:=0
	precedente:=' '

	for _,r:= range s{

		if precedente == ' ' && r != ' '{

			c++
		}

		precedente = r
	} 

  	Println ("Il numero di parole nel testo è",c)	
}