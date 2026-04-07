package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go 🚀")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(os.Environ())
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}