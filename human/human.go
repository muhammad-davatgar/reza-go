package human

import "fmt"

type Human struct {
	Name           string
	Age            int
	nationalNumber string
}

func (h *Human) SayName(message string) {
	fmt.Println("my Name is ", (*h).Name)
	fmt.Println(message)
}

func (h Human) SayAge() {
	fmt.Println("i am ", h.Age, " years old")
}

func (Human) SayRace() {

}

func (h *Human) SetNationalNumber(number string) {
	h.SayAge()
	h.nationalNumber = number
}

func (h *Human) getNationalNumber() string {
	return h.nationalNumber
}
