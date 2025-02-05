package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go Web Server is started!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
