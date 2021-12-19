package main

import "fmt"

func main() {

	name := "Ujang"

	if name == "Dadang" {
		fmt.Println(false)
	} else if name == "Ujang" {
		fmt.Println(true)
	} else {
		fmt.Println("Siapa Anda?")
	}

	if length := len(name); length >= 5 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
