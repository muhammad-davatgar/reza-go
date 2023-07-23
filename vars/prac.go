package vars

import "fmt"

func Tmp() {
	var name string

	var age int8
	var score float32
	var sex string

	age = 21
	score = 18.0
	sex = "male"

	name = "gotcha"
	fmt.Println("name: " + name)
	fmt.Println("age: ", age)
	fmt.Println("score: ", score)
	fmt.Println("sex: ", sex)

	age = age + 11
	age += 11
	score = score / 4
	score = 23 + score*12

	IsMale := true

	fmt.Println(IsMale)

}
