package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var i int
	// i = -10

	// a := uint(i)
	// fmt.Println(a)
	// fmt.Printf("i : %T \n", i)
	// fmt.Printf("a : %T", a)

	// var j float32

	// j = 10.99

	// fmt.Println(int(j))

	// d := 10

	// s := fmt.Sprintf("a %d", d)

	// fmt.Printf("s : %T", s)

	s := "23"
	var a uint64

	i, _ := strconv.ParseInt(s, 10, 0)
	fmt.Println(i)
	fmt.Printf("i : %T", i)

	a = uint64(i)
	fmt.Printf("a : %T ", a)
}
