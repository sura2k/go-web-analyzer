package main

import (
	"log"
	"net/http"

	"github.com/sura2k/go-web-analyzer/controllers/view"
)

func main() {
	//Handler Mappings
	http.HandleFunc("/", view.AnalyzerViewHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
