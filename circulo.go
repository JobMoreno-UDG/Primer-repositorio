package main

import "fmt"

func main() {
	const Pi = 3.1415
	var f float64
	fmt.Println("Ingresa valor del radio")
	fmt.Scanf("%f", &f)
	output := Pi * (f * f)
	fmt.Println("El area del circulo es es: ", output)
}
