package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// Constants definition
	const pi float64 = 3.14
	const pi2 = 3.1416
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)
	fmt.Printf("El valor es %v, y el tipo es %T", base, base)

	// Zero values (default values assigned by Go)
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)
}
