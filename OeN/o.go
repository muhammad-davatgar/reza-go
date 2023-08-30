package oen

import "fmt"

func main() {
	constant([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})

	N([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
}

func N(arr []int) {

	sum := 0
	for i := range arr {
		sum += arr[i]
	}

	fmt.Println(sum)
}

func constant(arr []int) {
	fmt.Println(arr)
}

func N2(arr []int) {
	sum := 0

	for range arr {
		for j := range arr {
			sum += arr[j]
		}
	}

	fmt.Println(sum)
}

func TwoN(arr []int) {
	sum := 0

	for i := range arr {
		sum += arr[i]
	}

	for i := range arr {
		sum += arr[i]
	}

	fmt.Println(sum)
}
