package main

import (
	"fmt"
	"reflect"
)

/*
	Anonymous function
*/

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked by the system")
	} else {
		fmt.Println("Welcome", name)
	}
}

/*
	function dapat dijadikan parameter di function lain
	formatnya adalah sebagai berikut

	func namaFunction (parameter tipeData, namaFunctionParameter (tipeData), tipeData) {}

	function sebagai parameter juga bisa diubah menjadi Type Declaration sebagai contoh dibawah helloWithFilter
*/

type filteredData func(string) string

func helloWithFilter(name string, filter filteredData) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)

}

func nameFilter(name string) string {
	if name == "Anjing" || name == "Babi" {
		return "..."
	} else {
		return name
	}

}

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

func variadic(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func variadicExp(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func functionAsValue(name string) string {
	return "Goodbye " + name
}

// normal loop factorial
func loopFactorial(val int) int {
	result := 1

	for i := val; i > 0; i-- {
		result *= i
	}
	return result
}

// vs using recrusive function
func factorialRecrusive(val int) int {
	if val == 1 {
		return 1
	} else {
		return val * factorialRecrusive(val-1)
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

	/*
		dalam variadic function dapat mengirim
		parameter kosong atau lebih dari 1 parameter,
		tipe variadic dalam sebuah function harus di deklarasi paling belakang dalam
		sebuah function, yang biasanya ditandai dengan ... (titik tiga didepan tipe data)
		contoh : func a (name string, age ...int) {} titik tiga setelah variable "i"
		menunjukan kalau i merupakan sebuah varargs atau variable argumens yang dapat diisi lebih dari 1 value
		contoh pemanggilan func a adalah a("budi", 10,11,12,13)
	*/

	slice := []int{10, 20, 11, 15, 30}
	total = variadic(slice...)
	fmt.Println(total)

	variadicExp(1, "exp", true, 10.5, []string{"tes", "obj"},
		map[string]int{"apple": 23, "orange": 10})

	goodBye := functionAsValue

	fmt.Println(goodBye("Kiryu"))

	// function sebagai parameter
	helloWithFilter("Kiryu", nameFilter)
	helloWithFilter("Anjing", nameFilter)
	helloWithFilter("Babi", nameFilter)

	blacklist := func(name string) bool {
		return name == "Kazama"
	}

	registerUser("Kazama", blacklist)
	registerUser("Kiryu", blacklist)

	loop := loopFactorial
	loopRecrusive := factorialRecrusive
	fmt.Println(loop(10))
	fmt.Println(loopRecrusive(10))

}
