package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 4000")
	port := "4000"

	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":"+port, nil)
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World"}`))
}
