package intr

import "fmt"

func main() {
	c := Cat{Name: "reza"}

	Pet(c)
}

type Cat struct {
	Name string
}

func (Cat) Meaw() {
	fmt.Println("cat")
}

type Tiger struct {
	Color string
}

func (Tiger)
 Meaw() {
	fmt.Println("tiger")
}

type GorbeSan interface {
	Meaw()
}

func Pet(arg GorbeSan) {
	arg.Meaw()

	j, ok := arg.(Cat)

	fmt.Println(j, ok)
}
