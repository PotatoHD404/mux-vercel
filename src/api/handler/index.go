package function

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
	"net/http"
)

//var books []Book
var handler *APIHandler
var books []Book

type APIHandler struct {
	H  http.Handler
	Db *bun.DB
}

func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.H.ServeHTTP(w, r)
}

func NewAPIHandler() *APIHandler {
	res := &APIHandler{NewHttpHandler(), NewDB()}
	return res
}

func NewHttpHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/api/books", CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")
	return r
}

func NewDB() *bun.DB {
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One",
		Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two",
		Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	//dsn := os.Getenv("POSTGRESQL")
	//sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	//db := bun.NewDB(sqldb, pgdialect.New())

	//return db
	return nil
}

//var h

func InitProject() {
	//err := godotenv.Load("../../.env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if handler == nil {
		InitProject()
		handler = NewAPIHandler()
	}
	handler.ServeHTTP(w, r)
}

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct

// GetBooks Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//print(books)
	json.NewEncoder(w).Encode(books)
}

// GetBook Get single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			_ = json.NewEncoder(w).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Book{})
}

// CreateBook Add new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	_ = json.NewEncoder(w).Encode(book)
}

// UpdateBook Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			_ = json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// DeleteBook Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	_ = json.NewEncoder(w).Encode(books)
}

//Request sample
//{
//"isbn":"4545454",
//"title":"Book Three",
//"author":{"firstname":"Harry", "lastname":"White"}
//}
