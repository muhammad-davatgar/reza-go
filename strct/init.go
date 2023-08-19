package main

import "fmt"

func main() {
	// var stud Student

	stud := Student{
		FirstName:   "reza",
		LastName:    "mir-ebrahimi",
		CodeMelli:   110,
		Average:     20,
		FatherName:  "yassol",
		ClassName:   "ryazi",
		Major:       "nohom",
		ChairNumber: 13,
	}

	// var stud2 Student
	// stud2 = Student{}

	// var stud3 Student
	// stud3.FirstName = "hosen"
	// ...

	PrintKarname(&stud)

	var classStudent []Student

}

type Student struct {
	FirstName   string
	LastName    string
	CodeMelli   int
	Average     float32
	FatherName  string
	ClassName   string
	Major       string
	ChairNumber uint8
}

func PrintKarname(stud *Student) {
	fmt.Println(stud.FirstName)
}

type Class struct {
	Students []Student
}
