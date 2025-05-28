package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Hello from Go!")
	})

	port := "8080"
	fmt.Println("Server listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
