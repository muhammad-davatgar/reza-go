package loops

import "fmt"

func CB() {
	for i := 0; i < 13; i++ {
		if i == 5 {
			continue
		}
		if i == 8 {
			break
		}

		fmt.Println(i)
	}
}

func Reza() {
	for i := 0; i < 20; i++ {
		if i == 13 {
			continue
		}
		if i == 16 {
			break
		}
		fmt.Println(i)
	}
}
