package main

import "fmt"

func main() {
	var name string
	// var age uint
	var money int32
	// var tmp int64
	// var score float64
	// var dead bool

	name = "reza"
	// age = 20
	money = 200
	// score = 17.5
	// dead = false

	fmt.Print(money / 2)
	fmt.Println(name + " mir ebrahimi")
	fmt.Println(name)

	name = name + "mir ebrahimi"
	fmt.Println(name)

	money = money + 100
	fmt.Println(money)

	var product1 bool
	var product2 bool

	product1 = money != 400
	product2 = money == 300

	fmt.Println("product 1 : ", product1)
	fmt.Println("product 2 : ", product2)

}
