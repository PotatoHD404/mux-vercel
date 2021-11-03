package main

import (
	function "hackaton-it-code-2.0/handler"
	"log"
	"net/http"
)

func main() {
	handler := function.NewAPIHandler()
	log.Fatal(http.ListenAndServe(":3000", handler.H))
}
