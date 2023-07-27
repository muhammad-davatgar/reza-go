package switch_command

import "fmt"

const (
	FirstGrade  string = "first"
	SecondGrade string = "second"
	ThirdGrade  string = "third"
)

const (
	IPV4 = "v4"
	IPv6 = "v6"
)

func Tmp() {
	fmt.Println("enter your grade")

	var grade string
	fmt.Scan(&grade)

	// if grade == FirstGrade {
	// 	fmt.Println("you are first grade")
	// } else if grade == SecondGrade {
	// 	fmt.Println("you are second grade")
	// } else if grade == ThirdGrade {
	// 	fmt.Println("you are third grade")
	// } else {
	// 	fmt.Println("non registerd")
	// }

	switch grade {
	case FirstGrade:
		fmt.Println("you are first grade")
	case SecondGrade:
		fmt.Println("you are second grade")
	case ThirdGrade:
		fmt.Println("you are third grade")
	default:
		fmt.Println("non registerd")
	}
}
