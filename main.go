package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Id    int
	Title string
	Auth  string
}

var Books []Book = []Book{
	{
		Id:    1,
		Title: "The one",
		Auth:  "Wally Robert",
	},
	{
		Id:    2,
		Title: "The only",
		Auth:  "Wally Robert",
	},
	{
		Id:    3,
		Title: "The book",
		Auth:  "Wally Robert",
	},
}

func mainRoute(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World")
}

func listBook(rw http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(rw)
	encoder.Encode(Books)

}

func configServer() {
	http.HandleFunc("/", mainRoute)
	http.HandleFunc("/books", listBook)

	fmt.Println("Servidor esta rodando na porta 8000")
	http.ListenAndServe(":8000", nil) // DefaultServerMux
}

func main() {
	configServer()
}
