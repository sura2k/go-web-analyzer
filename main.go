package main

import (
	"go-web-analyzer/controllers/view"
	"log"
	"net/http"
)

func main() {
	//Handler Mappings
	http.HandleFunc("/", view.AnalyzerHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
