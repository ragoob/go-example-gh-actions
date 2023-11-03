package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf("%d", Sum(5, 10))
		w.Write([]byte(s))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func Sum(a int, b int) int {
	return a + b
}
