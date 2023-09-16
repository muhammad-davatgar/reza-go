// You can edit this code!
// Click here and start typing.
package generics

import (
	"fmt"
)

func init() {
	a := Stack[Tiger]{
		arr: []Tiger{Tiger{Name: "babr"}},
	}

	fmt.Println(a.arr)

}

type Stack[T GorbeSan] struct {
	arr []T
}

type GorbeSan interface {
	Meaw()
}

type Tiger struct {
	Name string
}

func (Tiger) Meaw() {
	fmt.Println("emeaw")
}
