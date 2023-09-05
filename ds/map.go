package ds

import (
	"fmt"
)

func hm() {
	var hm map[string]int

	hm = map[string]int{}

	hm1 := map[string]int{}

	hm2 := make(map[string]int, 10)

	hm["reza"] = 10

	fmt.Println(hm["reza"])

	hm1["behzad"] = 20
	hm1["rostam"] = 30

	fmt.Println(hm1)
}

func ok() {
	a := make(map[string]int)

	a["reza"] = 16

	age, ok := a["reza"]
	if !ok {
		fmt.Println("invalid key")
	}
	fmt.Println(ok)
	fmt.Println(age)
}
