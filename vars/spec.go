package vars

import "fmt"

func Tmp4() {
	var i int
	i = 10

	fmt.Printf("i : %d", i)
	fmt.Println()
	name := "reza"
	var age uint
	age = 15
	fmt.Printf("hello %s age : %v", name, age)
	fmt.Println()

	fmt.Printf("name : %T \t", name)
	fmt.Printf("age : %T \n", age)

	fmt.Println("blah \a")
}
