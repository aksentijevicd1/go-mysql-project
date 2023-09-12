package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/aksentijevicd1/go-mysql-project/pkg/models"
	"github.com/aksentijevicd1/go-mysql-project/pkg/utils"
	"github.com/gorilla/mux"
)

type Books struct {
	l *log.Logger
}

func NewBooks(l *log.Logger) *Books {
	return &Books{l}
}

func (b *Books) GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Error while marshaling", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (b *Books) GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["bookId"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		b.l.Println("Error while converting to int")
		return
	}

	book, _ := models.GetBookById(id)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error while marshaling", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (b *Books) CreateBook(w http.ResponseWriter, r *http.Request) {
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)
	book := newBook.CreateBook()
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (b *Books) UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	idString := vars["bookId"]
	id, err := strconv.Atoi(idString)

	if err != nil {

		http.Error(w, "Unable to parse to int", http.StatusBadRequest)
		return
	}

	oldBook, db := models.GetBookById(id)

	if updateBook.Name != "" {

		oldBook.Name = updateBook.Name

	}
	if updateBook.Author != "" {

		oldBook.Author = updateBook.Author

	}
	if updateBook.Publication != "" {

		oldBook.Publication = updateBook.Publication

	}

	db.Save(&oldBook)
	res, err := json.Marshal(oldBook)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (b *Books) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["bookId"]
	id, err := strconv.Atoi(idString)

	if err != nil {

		http.Error(w, "Unable to parse to int", http.StatusBadRequest)
		return
	}
	delBook := models.DeleteBook(id)
	res, err := json.Marshal(delBook)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
