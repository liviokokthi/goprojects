//Lezione di Mercoledì 14 novembre 2018 

package main 

import ( 
	."fmt"
)

type data struct {
	g,m,a int 
}

func bisestile (anno int ) bool {
	return anno %4 == 0 && (anno%100 !=0 || anno%400 ==0)

}

func giorni (data data) int {
	c:=0
	for i:=1; i<data.m; i++ {
		switch i {
		case 11,4,6,9:
			c+=20
		case 2 :
			if bisestile (data.a) {
				c+=29
			} else {
				c+=28
			}
			default :
			c+= 31
		}
	}
 return c+ data.g
}


func main () {
	var data data 
	Scan (&data.g, &data.m, &data.a)
	Println(giorni(data))
}
