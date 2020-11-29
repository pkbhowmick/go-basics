package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
)

	
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"ISBN"`
}

// Todo : Implement DB
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error in encoding json object"}`))
		return
	}
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error in decoding json object"}`))
		return
	}
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	err = json.NewEncoder(w).Encode(book)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error in encoding json object"}`))
		return
	}
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	param := mux.Vars(r)
	var book Book
	for index, item := range books {
		if param["id"] == item.ID {
			book = books[index]
			err := json.NewEncoder(w).Encode(book)
			if err!=nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message":"Error in encoding json object"}`))
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message":"Book not exist!"}`))
}

func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	param := mux.Vars(r)
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error in decoding json object"}`))
		return
	}
	for index, item := range books {
		if param["id"] == item.ID {
			newBook.ID = item.ID
			books = append(books[:index],books[index+1:]...)
			books = append(books,newBook)
			err := json.NewEncoder(w).Encode(newBook)
			if err!=nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message":"Error in encoding json object"}`))
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message":"Book not exist!"}`))
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	param := mux.Vars(r)
	var book Book
	for index, item := range books {
		if param["id"] == item.ID {
			book = books[index]
			books = append(books[:index],books[index+1:]...)
			err := json.NewEncoder(w).Encode(book)
			if err!= nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"message":"Error in encoding json object"}`))
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message":"Book not exist!"}`))
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}