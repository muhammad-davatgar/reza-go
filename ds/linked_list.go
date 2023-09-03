package ds

import (
	"fmt"
)

func ls() {
	h1 := Person{Name: "reza"}

	fmt.Println(h1)

	fmt.Println(*h1.Next)

	h1.Next = new(Person)

	h1.Next.Name = "behzad"

	fmt.Println(h1.Next.Name)

}

type Person struct {
	Name string
	Next *Person
}
