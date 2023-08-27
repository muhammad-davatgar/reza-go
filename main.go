package main

import (
	"fmt"
)

type adad int

func main() {
	var nomre adad
	nomre = 10

	fmt.Println(nomre.Add(nomre))
}

func (a adad) Add(s adad) adad {
	return a + a
}
