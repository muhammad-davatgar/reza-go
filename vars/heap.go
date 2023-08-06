package vars

import "fmt"

func Heap() {
	name := new(string)

	var _ *string

	fmt.Scanln(name)

	fmt.Println(name)
	fmt.Println(*name)
}
