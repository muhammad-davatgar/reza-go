package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Println("inter studet count")
	fmt.Scanln(&a)

	i, _ := strconv.Atoi(a)

	Myarray := make([]string, i)

	for s := 0; s < i; s++ {
		fmt.Println("interr student", s+1, "name")
		fmt.Scanln(&Myarray[s])
	}
	fmt.Println(Myarray)

}
