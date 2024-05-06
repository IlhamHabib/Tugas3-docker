package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo, Saya ilham habib sedang mengerjakan tugas 3")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server berjalan di http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
