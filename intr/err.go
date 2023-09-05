package intr

import (
	"errors"
	"fmt"
	"strconv"
)

func Err() {
	a := "af10"

	b, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("invalid number , ", err)
		return

	}
	fmt.Println(b)
}

func start() {
	err := ErrorPune(2)
	if err != nil {
		if errors.As(err, &A{}) {
			fmt.Println("error type A :", err.Error())
		} else if errors.As(err, &d{}) {
			fmt.Println("error type d :", err.Error())
		} else if errors.As(err, &y{}) {
			fmt.Println("error type y :", err.Error())
		} else {
			fmt.Println("unknown error : ", err.Error())
		}
	}

}

func ErrorPune(arg uint) error {
	switch arg {
	case 1:
		a := A{}
		return a
	case 2:
		d := d{}
		return d
	default:
		y := y{}
		return y
	}
}

type A struct{}

func (a A) Error() string {
	b := "file address was invalid"
	return b
}

type d struct{}

func (b d) Error() string {
	q := "error b "
	return q
}

type y struct{}

func (y y) Error() string {
	p := "error y "
	return p
}
