package main

import "fmt"

func main() {
	// Contoh penggunaan if else
	num := 10
	if num > 0 {
		fmt.Println("Angka positif")
	} else if num == 0 {
		fmt.Println("Angka nol")
	} else {
		fmt.Println("Angka negatif")
	}

	// Contoh penggunaan temporary variable di dalam if else
	if tempVar := num * 2; tempVar > 15 {
		fmt.Println("Temporary variable lebih besar dari 15")
	} else {
		fmt.Println("Temporary variable kurang dari atau sama dengan 15")
	}

	// Contoh penggunaan switch case
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Hari Senin")
	case "Tuesday":
		fmt.Println("Hari Selasa")
	case "Wednesday":
		fmt.Println("Hari Rabu")
	case "Thursday":
		fmt.Println("Hari Kamis")
	case "Friday":
		fmt.Println("Hari Jumat")
	default:
		fmt.Println("Hari lainnya")
	}

	// Contoh penggunaan fallthrough
	switch num := 2; {
	case num%2 == 0:
		fmt.Println("Angka genap")
		fallthrough
	case num%2 != 0:
		fmt.Println("Angka ganjil")
	}
}
