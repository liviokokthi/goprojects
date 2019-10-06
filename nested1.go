/* Cicli annidati MOD1: 

     *****
     ****
     ***
     **
     *      
                         */

package main 

import (
          ."fmt"
                   )

func main (){

	var i, n, j int
    
    Print("Inserire il numero di asterischi:_")
    Scan(&n)

    for i=0; i<n; i++{
    	for j=0; j<n-i;j++{

    		Print("*")
    	}
    	Println()
    }
}