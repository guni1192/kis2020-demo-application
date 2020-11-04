package main

import (
	"log"
	"net/http"

	"github.com/guni1192/kis2020-demo-application/pkg/handler"
)

func main() {
	http.HandleFunc("/", handler.TopPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
