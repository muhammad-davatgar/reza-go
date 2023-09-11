package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	a := Person{FName: "reza", LName: "mir", NationalCode: 10, FatherName: "rassol", Job: "student"}

	v, _ := json.Marshal(a)
	fmt.Println(string(v))
	_ = os.WriteFile("name.json", v, 0644)

	// s := fmt.Sprintf("%s,%s,%d,%s,%s", a.FName, a.LName, a.NationalCode, a.FatherName, a.Job)
	// d1 := []byte(s)
	// err := os.WriteFile("name.txt", d1, 0644)
	// check(err)

	// defer f.Close()

	// w.Flush()

}

type Person struct {
	FName        string
	LName        string
	NationalCode uint
	FatherName   string
	Job          string
	Pen          Pen
}

type Pen struct {
	Length uint8
	Color  string
}
