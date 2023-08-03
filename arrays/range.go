package arrays

import "fmt"

func Range() {
	var myArray [10]int

	for i := range myArray {
		fmt.Printf("i : %d \n", i)
		myArray[i] = i + i*2
	}

	fmt.Println(myArray)

	for k, v := range myArray {
		fmt.Println("k : ", k)
		fmt.Println("v : ", v)
	}
}
