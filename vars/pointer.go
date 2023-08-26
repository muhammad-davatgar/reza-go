package vars

import "fmt"

func Pr() {

	a := 10                      // var a int
	aPointer := &a               // var aPointer *int
	aPointerPointer := &aPointer // var aPointerPointer **
	fmt.Println(a)
	fmt.Println(aPointer)
	fmt.Println(*aPointer)
	fmt.Println(aPointerPointer)
	fmt.Println(*aPointerPointer)
	fmt.Println(**aPointerPointer)

	var b *int

}
