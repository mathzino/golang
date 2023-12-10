package main

import "fmt"

func main() {
    // 1. Inisialisasi Nilai Awal Array
    var numbers1 [5]int
    numbers1 = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array 1:", numbers1)

    // 2. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
    numbers2 := [...]int{1, 2, 3, 4, 5}
    fmt.Println("Array 2:", numbers2)

    // 3. Array Multidimensi
    var matrix [3][3]int
    matrix = [3][3]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println("Array 3:")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Print(matrix[i][j], " ")
        }
        fmt.Println()
    }
}
