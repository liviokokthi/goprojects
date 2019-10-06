// Contare il numero di parole in una stringa - v3

/* Package unicode reference manual: http://golang.org/pkg/unicode */

    package main

    import ( 
              "fmt"
              "unicode" 
                              )
    func main(){  // Aumentiamo il contatore se ciò che precede una lettera nel range "r" non è una lettera
       
	s:= "Ciao mamma   guarda come mi      diverto ancora"
	c:= 0
	precedente:= ' '

         for _, r:= range s{

         	if !unicode.IsLetter(precedente) && unicode.IsLetter(r){

         		c++
         	}

         	precedente = r
         }

  	fmt.Println ("Il numero di parole nel testo è",c)	

 }
