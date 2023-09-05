package intr

import "fmt"

func main() {
	c := Cat{}
	Pet(c)

	t := Tiger{}

	l := Lion{}

	Pet(t)
	Pet(l)

}

type Cat struct{}

func (Cat) Meaw() {
	fmt.Println("Meaaaaaw")
}

type Tiger struct{}

func (Tiger) Meaw() {
	fmt.Println("Meaaaaaw")
}

type Lion struct{}

func (Lion) Meaw() {
	fmt.Println("roar")
}

func Pet(pishih GorbeSan) {
	pishih.Meaw()
}

type GorbeSan interface {
	Meaw()
}

type Dog struct{}

func (Dog) Bark() {
	fmt.Println("bark")
}
