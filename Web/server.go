package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"math/rand"
)

	
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"ISBN"`
}

// Todo : Implement DB
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = rand.Intn(100000)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}