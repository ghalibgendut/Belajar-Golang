package main

import "fmt"

func helloFunction() {
	fmt.Println("Hello this is a Function before a continue operation!")
}

func helloFunc() {
	fmt.Println("Hello this is a Function after a continue operation!")
}

func main() {
	for i := 0; i < 5; i++ {
		helloFunction()

		if i < 3 {
			continue
		}

		helloFunc()
	}
}
