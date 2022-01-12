package main

import "fmt"

func main() {

	// Perulangan menggunakan break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //Break akan menghentikan perulangan secara keseluruhan
		}
		fmt.Println("Perulangan ke ", i)
	}

	// Perulangan meggunakan Continue
	for k := 0; k < 10; k++ {
		if k%2 == 0 {
			continue //Continue akan menghentikan perulangan, tapi dia akan melanjutkan post statement selanjutnya
		}

		fmt.Println("Perulangan ke ", k)

	}

}
