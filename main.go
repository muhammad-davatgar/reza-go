package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	z := 10

	j := z + 5
	PrintNum(0)
}

func PrintNum(num int) {
	fmt.Println(num)
	if num != 5 {
		PrintNum(num + 1)
	}
}
