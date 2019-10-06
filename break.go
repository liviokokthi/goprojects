// istruzione "break"

package main

import (
         ."fmt"
                )

func main(){
	
	var x, s int

	Println ("Inserire una sequenza di interi interrotta da uno \"0\":")
        
	for s=0; ; s+=x {

		Scan (&x)

	if x == 0{

		break
	}

}

	Println ("La somma di tutti i numeri è:",s)

}