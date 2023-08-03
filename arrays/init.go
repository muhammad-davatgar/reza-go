package arrays

import "fmt"

func Init() {
	var myArray [5]string

	for i := 0; i < 5; i++ {
		fmt.Println("input name for student ", i+1)
		fmt.Scanln(&myArray[i])
	}
	myArray[4] = "reza"
	fmt.Println(myArray[4])

}
