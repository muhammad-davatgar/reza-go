package vars

import "fmt"

const (
	SPRING = iota
	SUMMER = iota
	FALL   = iota
	WINTER = iota
)

const (
	happy = 10
	sad
	natural
)

func Tmp2() {
	println("لطفا حالت خود را وارد کنید")
	var status string

	const (
		died   string = "died"
		live   string = "live"
		busy   string = "busy"
		unbusy string = "unbusy"
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
