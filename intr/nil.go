package intr

import (
	"fmt"
)

func TMP() {
	a := pishih(1)
	if a == nil {
		fmt.Println("invalid arg")
	}

}

func pishih(arg int) GorbeSan {
	switch arg {
	case 1:
		return Cat{}
	case 2:
		return Tiger{}
	case 3:
		return Lion{}
	default:
		return nil
	}
}
