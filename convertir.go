package main

import "fmt"

func main(){
	var f float64
	fmt.Print("Ingresa los Grados en Fahrenheit: ")
	fmt.Scanf("%f", &f)
	output:=(f - 32) * 5/9
	fmt.Println(f," Fahrenheit es equivalente a ",output," en grados Celcius")
}