package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/howkyle/stockfolio-server/analysis"
)

func main() {
	http.HandleFunc("/analyze", analysis.AnalysisHandler())
	fmt.Printf("listening")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
