package strct

import "fmt"

func Ananymos() {
	// var b int = 10
	// var i []int = []int{1, 2, 3}

	var s []struct {
		Name string
		Age  uint8
	} = []struct {
		Name string
		Age  uint8
	}{
		{Name: "reza", Age: 16},
		{Name: "behzad", Age: 32},
	}

	fmt.Println(s[0].Name)
}
