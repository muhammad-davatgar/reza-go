package generics

import (
	"fmt"
)

func Function() {
	a := WeddingCard{}

	printCard[WeddingCard](a)

}

func printCard[C Card](card C) {
	fmt.Println("card name:", card.Name())
}

type Card interface {
	Name() string
}

type WeddingCard struct {
}

func (WeddingCard) Name() string {
	return "wedding card"
}
