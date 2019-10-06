
// Il numero primo con l'istruzione "break" senza variabili Booleane

package main 

import (
            ."fmt" 
                    )

func main (){

	var n, i int
     
    Print("Inserire un numero intero:_")
    Scan(&n) 

    // temp = n

	for i=2; i<n; i++ {

		if n%i == 0{
			break
		}
	} 	

	if i < n-1 {

		Println("Il numero",n,"NON è un numero primo")
	
	} else {

		Println("Il numero",n,"è un numero primo")
	}

}                    