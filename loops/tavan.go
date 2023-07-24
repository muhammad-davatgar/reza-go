package loops

import "fmt"

func Power() {
	tavan := 4
	paye := 3

	natije := paye

	for i := 0; i < tavan; i++ {
		natije *= paye

	}

	fmt.Println("outside : ", natije)
}
