package main

import "fmt"

func main(){
	var f float64
	fmt.Println("Ingresa el lado del cuadrado")
	fmt.Scanf("%f", &f)
	output:=f*f
	fmt.Println("El area del cuadrado es: ",output)
}