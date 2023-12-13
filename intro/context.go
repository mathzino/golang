package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Membuat konteks dengan pembatalan setelah timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Pastikan untuk memanggil fungsi cancel untuk membersihkan sumber daya yang terkait dengan konteks

	// Memulai goroutine dengan konteks
	go performTask(ctx)

	// Menunggu untuk konteks dibatalkan atau selesai
	select {
	case <-ctx.Done():
		fmt.Println("Operasi dibatalkan atau timeout")
	}
}

func performTask(ctx context.Context) {
	// Menambahkan nilai ke konteks (opsional)
	ctxWithValue := context.WithValue(ctx, "key", "value")

	// Melakukan tugas (simulasi pekerjaan yang berlangsung)
	for i := 1; i <= 5; i++ {
		select {
		case <-time.After(time.Second):
			fmt.Println("Task in progress...",i)
		case <-ctx.Done():
			// Memeriksa apakah konteks telah dibatalkan
			fmt.Println("Task dibatalkan")
			return
		}
	}

	// Mengakses nilai dari konteks (jika diperlukan)
	if value := ctxWithValue.Value("key"); value != nil {
		fmt.Println("Value from context:", value)
	}

	fmt.Println("Task completed successfully")
}
