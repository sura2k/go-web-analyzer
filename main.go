package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sura2k/go-web-analyzer/config"
	"github.com/sura2k/go-web-analyzer/controllers/view"
)

func main() {
	//Handler Mappings
	http.HandleFunc("/", view.AnalyzerViewHandler)

	serverPort := config.Config.Server.Port
	log.Println("Server started on port ", serverPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
}
