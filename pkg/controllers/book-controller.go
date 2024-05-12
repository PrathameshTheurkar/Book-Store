package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PrathameshTheurkar/go-bookstore/pkg/models"
	"github.com/PrathameshTheurkar/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	AllBooks := models.GetAllBooks()
	res, err := json.Marshal(AllBooks)
	if err != nil {
		fmt.Println("Error during Json Parsing")
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stringId := params["id"]
	id, err := strconv.ParseInt(stringId, 0, 0)
	if err != nil {
		fmt.Println("Error while converting string to int")
		panic(err)
	}

	book, db := models.GetBookById(id)
	fmt.Println(db)

	res, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error while parsing to json")
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}

	utils.ParseBody(r, book)
	CreatedBook := book.CreateBook()
	res, err := json.Marshal(CreatedBook)
	if err != nil {
		fmt.Println("Error while parsing json")
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stringId := params["id"]
	id, err := strconv.ParseInt(stringId, 0, 0)
	if err != nil {
		fmt.Println("Error while converting string to int")
		panic(err)
	}

	book := models.DeleteBook(id)

	res, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error while parsing to json")
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	params := mux.Vars(r)
	stringId := params["id"]
	id, err := strconv.ParseInt(stringId, 0, 0)
	if err != nil {
		fmt.Println("Error while converting string to int")
		panic(err)
	}

	book, db := models.GetBookById(id)

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Save(&book)

	res, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Error whiling parsing to json")
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
