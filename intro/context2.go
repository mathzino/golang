package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Membuat konteks dengan pembatalan setelah 3 detik
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Menjalankan dua goroutine dengan konteks yang sama
	go performTask(ctx, "Task 1")
	go performTask(ctx, "Task 2")

	// Menunggu untuk konteks dibatalkan atau selesai
	select {
	case <-ctx.Done():
		fmt.Println("Operasi utama selesai atau dibatalkan")
	}
}

func performTask(ctx context.Context, taskName string) {
	// Menambahkan nilai ke konteks
	ctxWithValue := context.WithValue(ctx, "taskName", taskName)

	// Melakukan tugas (simulasi pekerjaan yang berlangsung)
	for i := 1; i <= 5; i++ {
		select {
		case <-time.After(time.Second):
			fmt.Printf("%s in progress...\n", taskName)
		case <-ctx.Done():
			// Memeriksa apakah konteks telah dibatalkan
			fmt.Printf("%s dibatalkan\n", taskName)
			return
		}
	}

	// Mengakses nilai dari konteks
	if value, ok := ctxWithValue.Value("taskName").(string); ok {
		fmt.Printf("%s completed successfully\n", value)
	}
}
