package main

import "fmt"

func key_value() {

	/* alternative way to declare a map
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	*/
	// Declare a map
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
