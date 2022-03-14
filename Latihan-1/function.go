package main

import "fmt"

func helloFunction() {
	fmt.Println("Hello this is a Function before a continue operation!")
}

func helloFunc() {
	fmt.Println("Hello this is a Function after a continue operation!")
}

func helloUser(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func returnVal(name string) string {
	if name == "" {
		return "Hello, may i know your name?"
	} else {
		return "Hello " + name + "!"
	}
}

func main() {
	for i := 0; i < 5; i++ {
		helloFunction()

		if i < 3 {
			continue
		}

		helloFunc()
	}

	firstName := "Ujang"
	lastName := "Jajang"
	helloUser(firstName, "Budiman")
	helloUser("Budi", lastName)

	result := returnVal("Dodi")
	fmt.Println(result)

	fmt.Println(returnVal(""))

}
