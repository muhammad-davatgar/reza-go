package arrays

import "fmt"

func Methods() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	fmt.Println("myArray : ", myArray)
	n := append(myArray[:], 40, 50, 60)
	fmt.Println("n : ", n)

	fmt.Println("len : ", len(myArray))

}
