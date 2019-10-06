// Trova la radice quadrata di un intero 

package main 
import . "fmt"

func main (){

	var x, r int
	
	Print("Inserire un intero:_") 
	Scan (&x)
	r=1

	for r*r < x {
		r++
	}

	/* Scrittura alternativa con ciclo "FOR" completo: 

	for r=1; r*r<x; r++ 

	*/

	Println ("La √ di",x, "è",r)

	// ctrl + c -> "Interrupt" di un ciclo infinito
}