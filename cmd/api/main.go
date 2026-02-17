package main

import (
	"log"
	"net/http"

	"github.com/denisrodrigues-code/lenovo-scraper-api/controller"
)

func main() {
	http.HandleFunc("/lenovo", controller.LenovoHandler)

	log.Println("API rodando em :8080")
	http.ListenAndServe(":8080", nil)
}
