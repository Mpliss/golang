package main 

import "fmt"

func main () { 
	var x int  
	var y int 

	fmt.Print("Tolong masukan nilai x dan y:")
	fmt.Scan(&x, &y)

	var fx = (1 / ((3 * x * x) + 10 )) + ((10 * y) + 7)
	fmt.Println(fx) // hasil sudah dibulatkan
}