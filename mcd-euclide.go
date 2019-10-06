// M.C.D. con l'equazione di Euclide

package main
import . "fmt"
func main() {
    
   Println ("Inserire due numeri interi: _")
   var x, y int 
   Scan (&x, &y)

   num:=1

   for x%y != 0 {

	Println ("Nella", num,"° divisione x:",x, "| y:",y) 
	// Println illustrativo, per vedere i valori ad ogni ciclo

	   x, y = y, x%y
	   
	   num++

   }

   Println ("Il Massimo Comun Divisore è ", y)

   // (!) in GO i cicli "WHILE" e "DO-WHILE" non esistono...

}

