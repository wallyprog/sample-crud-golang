package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Auth  string `json:"auth"`
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

func bookRoutes(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listBook(rw, r)
	} else if r.Method == "POST" {
		createBook(rw, r)
	}
}

func listBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(rw)
	encoder.Encode(Books)

}

func createBook(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		// catch
	}
	var newBook Book
	json.Unmarshal(body, &newBook)
	newBook.Id = len(Books) + 1
	Books = append(Books, newBook)

	encoder := json.NewEncoder(rw)
	encoder.Encode(newBook)
}

func configServer() {
	http.HandleFunc("/books", bookRoutes)

	fmt.Println("Servidor esta rodando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", nil)) // DefaultServerMux
}

func main() {
	configServer()
}
