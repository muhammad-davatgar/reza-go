package intr

import "fmt"

func stringer() {
	reza := Person{Name: "reza", Phone: "asdsa", Pet: "gorbe", Country: "iran"}
	fmt.Println(reza)
}

type Person struct {
	Name    string
	Phone   string
	Pet     string
	Country string
}

func (p Person) String() string {
	return p.Name + " " + p.Country
}
