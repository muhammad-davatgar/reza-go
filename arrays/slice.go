package arrays

import "fmt"

func Slice() {
	arr := []int{1, 2, 3}

	fmt.Println(arr[0])
	fmt.Println(arr[0:3])

	a := []int{1, 2, 3, 4, 5}

	b := a[0:3]

	fmt.Println(a)
	fmt.Println(b)

	a[0] = 40

	fmt.Println(a)
	fmt.Println(b)
}
