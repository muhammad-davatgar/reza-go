package func

import "fmt"

func main() {

	Power(2, 3)

	i, j := Reza(10)

	fmt.Println(i)
	fmt.Println(j)
}

func Power(intfirst, intsecond int) {
	natije := intfirst
	var i int
	for i = 1; i < intsecond; i++ {
		natije *= intfirst
	}

	fmt.Println(natije)
}

func Reza(x int) (int, int) {
	return x - 1, x + 1
}
