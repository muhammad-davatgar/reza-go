package concurrency

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go Print(10)
	go Print(20)

	wg.Wait()
}

func Print(arg int) {
	for i := 0; i < 50; i++ {
		fmt.Println(arg)
	}
	wg.Done()
}

// func main() { // thread main

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	fmt.Println("10")
// 	fmt.Println("20")

// 	go func() { // thread 10
// 		defer wg.Done()
// 		for i := 0; i < 50; i++ {
// 			fmt.Println(10)
// 		}
// 	}()

// 	go func() { // thread 20
// 		defer wg.Done()
// 		for i := 0; i < 50; i++ {
// 			fmt.Println(20)
// 		}
// 	}()

// 	wg.Wait()

// }
