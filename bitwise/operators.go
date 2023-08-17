package bitwise

import "fmt"

func main() {
	a := 10
	a = a << 1
	fmt.Println(a)

	b := 20
	b = b >> 2
	fmt.Println(b)

	k := 10
	j := 9
	z := k & j
	fmt.Println(z)

	q := k | j
	fmt.Println(q)

	xor := k ^ j
	fmt.Println(xor)

}
