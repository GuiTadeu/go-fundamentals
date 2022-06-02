package main

import (
	"net/http"
	"github.com/GuiTadeu/meli-go/web/class01/exercises"
)

func main() {
	http.HandleFunc("/hello", exercises.HelloHandler)
	http.HandleFunc("/transactions", exercises.GetAllHandler)
	http.ListenAndServe(":8080", nil)
}