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
