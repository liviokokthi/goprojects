// Esercizio del "numero primo" con l'istruzione "break" e valori Booleani

package main 

import (
           ."fmt"
                    )

func main( )
{

     var n int
     var primo bool
     
     Print ("Inserire un numero Intero: _")
     Scan(&n)

     primo = true

     for i:=2; i < n; i++ {

     	if n%i == 0{

     		primo = false

     		Println ("Il numero",n,"NON è un numero primo")

     		break
     	}

     }

     if primo {

     	Println ("Il numero",n,"è un numero primo")
     }
}