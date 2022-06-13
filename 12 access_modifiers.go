package main

import (
	"fmt"
	pk "1 Foundamentals/my_package"
)

/*
Summary:
	Capital case: Public
	Lower case: Private
*/

func access_modifiers() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	pk.PrintMessage("Hola Platzi")

	pk.printMessage1("Hola")
}
