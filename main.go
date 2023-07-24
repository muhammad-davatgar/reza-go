package main

import "fmt"

func main() {

	s := 10
	d := 10
	g := 10

	if s > d {
		fmt.Print("10 bigger than 8")
	} else if s < d {
		fmt.Print("10 smaller than 8")
	} else if d > g {
		fmt.Print("8 bigger than 2")
	} else if d < g {
		fmt.Print("8 smaller than 2")
	} else if s < g {
		fmt.Print("10 bigger than 2")
	} else if s > g {
		fmt.Print("10 smaller than 2")
	} else {
		fmt.Print("all numbers are equal ")
	}
}
