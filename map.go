package main

import "fmt"

func main() {
    // 1. Penggunaan Map
    var personMap map[string]int
    personMap = make(map[string]int)

    // 2. Inisialisasi Nilai Map
    personMap["Alice"] = 30
    personMap["Bob"] = 25
    personMap["Charlie"] = 35

    fmt.Println("1. Penggunaan Map:")
    fmt.Println("   Person Map:", personMap)

    // 3. Menghapus Item Map
    delete(personMap, "Bob")
    fmt.Println("2. Setelah Menghapus Item 'Bob':")
    fmt.Println("   Person Map:", personMap)

    // 4. Deteksi Keberadaan Item Dengan Key Tertentu
    age, exists := personMap["Alice"]
    if exists {
        fmt.Println("3. Umur Alice adalah", age)
    } else {
        fmt.Println("3. Alice tidak ditemukan di map.")
    }

    _, exists = personMap["Bob"]
    if exists {
        fmt.Println("   Bob ditemukan di map.")
    } else {
        fmt.Println("   4. Bob tidak ditemukan di map.")
    }
}
