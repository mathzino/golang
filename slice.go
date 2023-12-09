package main

import "fmt"

func main() {
    // 1. Inisialisasi Slice
    slice1 := []int{1, 2, 3, 4, 5}
    fmt.Println("1. Slice 1:", slice1)

    // 2. Hubungan Slice Dengan Array & Operasi Slice
    array := [5]int{1, 2, 3, 4, 5}
    slice2 := array[1:4]
    fmt.Println("2. Slice 2 (dari array):", slice2)

    // 3. Slice Merupakan Tipe Data Reference
    slice3 := slice1
    slice3[0] = 10
    fmt.Println("3. Slice 3 (referensi):", slice3)
    fmt.Println("   Slice 1 setelah perubahan di Slice 3:", slice1)

    // 4. Fungsi len()
    fmt.Println("4. Panjang Slice 1:", len(slice1))

    // 5. Fungsi cap()
    fmt.Println("5. Kapasitas Slice 1:", cap(slice1))

    // 6. Fungsi append()
    slice1 = append(slice1, 6, 7, 8)
    fmt.Println("6. Slice 1 setelah append:", slice1)

    // 7. Fungsi copy()
    slice4 := make([]int, len(slice1))
    copy(slice4, slice1)
    fmt.Println("7. Slice 4 (copy dari Slice 1):", slice4)

    // 8. Pengaksesan Elemen Slice Dengan 3 Indeks
    fmt.Println("8. Elemen ke-2 dari Slice 1:", slice1[1])
    
    // 9. Alokasi Elemen Slice Menggunakan Keyword make
    slice5 := make([]int, 5, 10)
    fmt.Println("9. Slice 5 (make):", slice5)
}
