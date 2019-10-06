// Numero di cifre in un INT

package main
import ."fmt"

func main(){

    var n, c, num int 
	
    Print("Inserire un intero:_")
	Scan (&n)	
	
	/* Prima Versione
	c=0;
	for i:=1; n/(foat64)i>0; i*=10{
		c++
		Println ("Il numero",n/i," contiene",c," cifre!")
	} */
	
	num = n

	for c=0; num>0; num/=10{
		c++
	}

	Println ("Il numero",n,"contiene",c,"cifre!")
}