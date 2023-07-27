package switch_command

import "fmt"

func Reza() {
	println("لطفا حالت خود را وارد کنید")
	var status string

	const (
		died   = "died"
		live   = "live"
		busy   = "busy"
		unbusy = "unbusy"
	)

	fmt.Scanln(&status)

	switch status {
	case died:
		fmt.Println("تو مرده هستی")

	case live:
		fmt.Println("تو زنده هستی")

	case busy:
		fmt.Println("تو مشغول هستی")

	case unbusy:
		fmt.Println("تو بی کار هستی")

	}

}
