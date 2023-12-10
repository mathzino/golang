package main

import (
	"fmt"
	"net/http"
)

// Fungi Log yang berguna sebagai middleware
func log(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Ini dari middleware Log....\n")
    fmt.Println(r.URL.Path)
    next.ServeHTTP(w, r)
  })
}

// Fungsi GetMovie untuk mengampilkan text string di browser
func GetMovie(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Ini dari function GetMovie()"))
}

func main() {
  // konfigurasi server
  server := &http.Server{
    Addr: ":8080",
  }

  // routing
  http.Handle("/", log(http.HandlerFunc(GetMovie)))

  // jalankan server
  fmt.Println("server running at http://localhost:8080")
  server.ListenAndServe()
}