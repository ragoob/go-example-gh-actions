package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf("%d", Sum(5, 10))
		w.Write([]byte(s))
	})

	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte("Machine successfully shutdown"))
	})
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Sum(a int, b int) int {
	return a + b
}
