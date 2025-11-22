package main

import (
	"fmt"
	"log"
	"net/http"

	handler "ascii-art-web/internal/handlers"
)

func main() {
	// Отдаём статические файлы: CSS, изображения
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("templates/style"))))

	// Хендлеры
	http.HandleFunc("/", handler.PostHandler)
	http.HandleFunc("/ascii-art", handler.PostHandler)

	fmt.Println("Server running at http://localhost:7050")
	log.Fatal(http.ListenAndServe(":7050", nil))
}

