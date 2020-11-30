package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"encoding/base64"
	"os"
	"github.com/dgrijalva/jwt-go"
	"time"
)

	
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"ISBN"`
}

var SIGNATURE_KEY = []byte(os.Getenv("SIGNATURE_KEY"))

// Todo : Implement DB
var books []Book

func getToken(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	authHeader:= r.Header.Get("Authorization")
	authString := strings.Split(authHeader," ")
	if len(authString) < 2 || authString[0] != "Basic" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Invalid authorization header"}`))
		return
	}
	decodedString, err := base64.StdEncoding.DecodeString(authString[1])
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Failed to decode authrization header"}`))
		return
	}
	finalString := string(decodedString)
	userString := strings.Split(finalString,":")
	if len(userString) < 2 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Invalid authorization header"}`))
		return
	}
	username := userString[0]
	password := userString[1]

	if username!=os.Getenv("ADMIN_USER") || password!=os.Getenv("ADMIN_PASS") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Wrong username or password"}`))
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"createdAt": time.Now(),
	})
	tokenString, err := token.SignedString(SIGNATURE_KEY)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Failed to generate token"}`))
		return
	}
	w.Write([]byte(`{"token": `+tokenString+`}`))
}

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
	router.HandleFunc("/api/authenticate",getToken).Methods("POST")
	router.HandleFunc("/api/books",getBooks).Methods("GET")
	router.HandleFunc("/api/books",createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}