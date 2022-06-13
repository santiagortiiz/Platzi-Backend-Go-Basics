package main

import "fmt"

type computer struct {
	ram   int
	brand string
	disk  int
}

func (myPC computer) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func stringers() {
	myPC := computer{ram: 16, brand: "msi", disk: 100}

	fmt.Println(myPC)
	//Alternative without overwrite String method: fmt.Printf("%+v", myPC)
}
