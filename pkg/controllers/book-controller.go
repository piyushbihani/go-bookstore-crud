package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/piyushbihani/go-bookstore-crud/pkg/models"
	"github.com/piyushbihani/go-bookstore-crud/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		panic(err)
	}
	bookDetails, _ := models.GetBookById(Id)
	if bookDetails == nil {
		w.Write([]byte("Book not found"))
		w.WriteHeader(http.StatusNotFound)
	} else {
		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := models.Book{}
	utils.ParseBody(r, &CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		panic(err)
	}
	book := models.DeleteBook(Id)
	if book == nil {
		w.Write([]byte("Book not found"))
		w.WriteHeader(http.StatusNotFound)
	} else {
		res, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := models.Book{}
	utils.ParseBody(r, &UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		panic(err)
	}
	bookDetails, db := models.GetBookById(Id)
	if bookDetails == nil {
		w.Write([]byte("Book not found"))
		w.WriteHeader(http.StatusNotFound)
	} else {
		if UpdateBook.Name != "" {
			bookDetails.Name = UpdateBook.Name
		}
		if UpdateBook.Author != "" {
			bookDetails.Author = UpdateBook.Author
		}
		if UpdateBook.Publication != "" {
			bookDetails.Publication = UpdateBook.Publication
		}

		db.Save(&bookDetails)
		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content_type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
