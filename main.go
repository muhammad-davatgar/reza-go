package main

import (
	"reza/human"
)

func main() {

}

func PK() { // packages
	h1 := human.Human{
		Name: "reza",
		Age:  16,
	}

	h1.SayRace()
	h1.SetNationalNumber("1712")
	// h1.getNationalNumber()
}
