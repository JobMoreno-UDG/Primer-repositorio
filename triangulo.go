package main

import "fmt"

func main() {
	var base float64
	var alto float64
	fmt.Print("Ingresa el alto del triangulo ")
	fmt.Scan(&alto)

	fmt.Print("Ingresa la base del triangulo ")
	fmt.Scan(&base)

	output := (base * alto) / 2
	fmt.Println("El area del triangulo es: ", output)
}
