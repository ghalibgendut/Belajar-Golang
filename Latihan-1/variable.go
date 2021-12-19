package main

import "fmt"

func main() {
	var name string

	name = "Ghalib"
	fmt.Println(name)

	name = "Ghalib Sasmito"
	fmt.Println(name)

	var alias = "Ujenk"
	fmt.Println(alias)

	alias = "Dadang"
	fmt.Println(alias)

	var age int8 = 23
	fmt.Println(age)

	var (
		firstName = "Ghalib"
		lastName  = "Sasmito"
	)

	fmt.Println("Nama Lengkap : ", firstName+" "+lastName)

	const sebuahVariable = "a"
	fmt.Println(sebuahVariable)

	// sebuahVariable = "B"
	// fmt.Println(sebuahVariable)

}
