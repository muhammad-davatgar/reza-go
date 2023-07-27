package loops

import (
	"fmt"
	"time"
)

func T() {

	for {
		time.Sleep(time.Second * 1)
		fmt.Println("reza")
	}
}
