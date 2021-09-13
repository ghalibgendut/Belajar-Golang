package main

import "fmt"

func main() {
	const result = 10 + 10
	fmt.Println(result)

	var (
		a = 10
		b = 10
		c = 0
	)

	c = a * b
	fmt.Println(c)

	var d = 10
	d += 10
	fmt.Println(d)

	d++
	fmt.Println(d)

	var negative = -100
	var positive = +100 // tidak wajib ditambah '+' karena secara default angka sudah berupa bilangan positive

	fmt.Println(negative)
	fmt.Println(positive)
}
