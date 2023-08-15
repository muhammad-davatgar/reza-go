package main

import "fmt"

func main() {
	n := 5

	arr := make([]string, n)

	for i := range arr {
		arr[i] = " "
	}

	mid := 5 / 2
	arr[mid] = "*"

	// arr2 := []string{"*", "*", "*", "*", "*"}
	// DecreaseStars(&arr2)
	// fmt.Println(arr2)

	for i := 0; i < n*2-1; i++ {
		for j := range arr {
			fmt.Print(arr[j])
		}
		if i < mid {
			IncreaseStars(&arr)
		} else {
			DecreaseStars(&arr)
		}
		fmt.Println()
	}
}

func IncreaseStars(arrPointer *[]string) {
	found := false

	for i := range *arrPointer {
		if (*arrPointer)[i] == "*" && !found {
			(*arrPointer)[i-1] = "*"
			found = true
		}

		if (*arrPointer)[i] == " " && found {
			(*arrPointer)[i] = "*"
			found = false
		}
	}
}

func DecreaseStars(arrPointer *[]string) {
	found := false

	for i := range *arrPointer {
		if (*arrPointer)[i] == "*" && !found {
			(*arrPointer)[i] = " "
			found = true

		}

		if ((*arrPointer)[i] == " " || i == (len(*arrPointer)-1)) && found {
			if i == (len(*arrPointer) - 1) {
				(*arrPointer)[i] = " "
			} else {
				(*arrPointer)[i-1] = " "
			}
		}
	}
}
