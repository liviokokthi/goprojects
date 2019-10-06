/* Cicli annidati MOD2: 

      *
     **
    ***
   ****
  *****                     
                  */

  package main 

  import (
            ."fmt"
                    )

  func main(){

  	var i, j, k, n int

    Print("Inserire il numero di asterischi:_")
    Scan(&n)

  	    for i=0; i<n; i++ {

  	    	for j=n-i-1; j>=0; j-- {

  	    		Print(" ")

  	    	}
               
               for k=0; k<i+1; k++{

               	Print("*")
                
  	    	}

  	    	Println()
  	    }
  }