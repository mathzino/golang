package main

import (
	"fmt"
)

func main() {
	// using var
	var text = "Hello World"
    fmt.Println(text)
        
    text = "Hello World diubah"
    fmt.Println(text)

	// using :=
	text1 := "ini teks 1"
    text1 = "ini teks 1 diubah"
	fmt.Println(text1)


	//using cont
	const judul = "Ini Judul"
    fmt.Println(judul)
    // jika kode di bawah ini di uncomment maka akan error
    // judul = "judul di ubah"
}