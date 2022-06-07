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

func getReturnMultipleVal() (string, string, int8) {
	return "Kiryu", "Kazuma", 22
}

func getFullName() (firstName string, lastName string, age int8) {
	firstName = "Kiryu"
	lastName = "Kazuma"
	age = 22
	return
}

func variadic(numbers ...int) (total int) {
	total = 0

	for _, number := range numbers {
		total += number
	}
	return total
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

	// pengembalian multiple function dengan semua value ditangkap
	first, last, age := getReturnMultipleVal()
	fmt.Println(first, last, age)

	/*
		semua return value dalam function yg
		mengembalikan multiple value harus di tangkap semua
		apa bila hanya ingin mengebalikan sebagian harus di tandai dengan "_"
	*/

	// pengembalian function dengan sebagian value ditangkap

	depan, _, umur := getReturnMultipleVal()
	fmt.Println(depan, umur)

	a, b, _ := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(c)

	total := variadic(10, 11, 10, 40)
	fmt.Println(total)
}
