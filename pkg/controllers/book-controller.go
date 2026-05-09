package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {

	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Failed to marshal books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	bookDetails, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}

	err := utils.ParseBody(r, createBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if createBook.Name == "" {
		http.Error(w, "Book name is required", http.StatusBadRequest)
		return
	}

	b, err := createBook.CreateBook()
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = models.DeleteBook(ID)
	if err != nil {
		http.Error(w, "Failed to delete book: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Book deleted successfully"}`))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	existingBook, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var updateData models.Book
	err = utils.ParseBody(r, &updateData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if updateData.Name != "" {
		existingBook.Name = updateData.Name
	}
	if updateData.Author != "" {
		existingBook.Author = updateData.Author
	}
	if updateData.Publication != "" {
		existingBook.Publication = updateData.Publication
	}

	updatedBook, err := models.UpdateBook(existingBook)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
