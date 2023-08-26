package func


import "fmt"

func main() {
	Power := HigherFunction()

	fmt.Println(Power(10))
}

func HigherFunction() func(a int) int {
	return func(i int) int {
		return i * i
	}
}


func Second(f func()int) int{
	return f()
}