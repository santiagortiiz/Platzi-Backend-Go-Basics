package main

import "fmt"

func fmt_class() {
	// Declaración de variables
	helloMessage := "Hello"
	worldMesssage := "world"

	// Println: Add new line at the of the string
	fmt.Println(helloMessage, worldMesssage)
	fmt.Println(helloMessage, worldMesssage)

	// Printf: Add a function to the string
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf: store string
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}
