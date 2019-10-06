// Indovina il numero esatto

package main
import ."fmt"

func main(){

	x:=100 // Numero da indovinare
	indovinato :=false
	
	for !indovinato{

		// Alternativa: for indovinato:=false; !indovinato{}

		Print("Indovina...")

		var n int
		Scan (&n)

		if n<x{

			Println("Il numero",n,"è minore del numero segreto!")
		
		}else if n>x{
			Println("Il numero",n,"è maggiore del numero segreto")
		}else{
			indovinato=true
			Println ("*** Indovinato!!! ***")

		}
	}

}