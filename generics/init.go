package generics

import (
	"errors"
	"fmt"
)

func main() {
	err := ErrorPune("asddsa")

	if err != nil {
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println(err.Error())
		}
	}
}

func ErrorPune(name string) error {
	if name == "" {
		return ErrFileNotFound
	}

	return nil
}

var ErrFileNotFound = fmt.Errorf("file not found")
