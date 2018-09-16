package main

import "fmt"

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"'`
	ISBN string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

var books = map[string]Book {
	"0345392018": Book{Title:"Cloud Native Go", Author:"James Cameroon", ISBN:"0345392018"},
	"0345391802": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
}

func AllBook() []Book {
	out:= make([]Book, len(books))
	idx:=0
	for first,book:= range books {
		fmt.Print(first)
		out[idx] = book
		idx++
	}
	return out
}

func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]

	if exists {
		return book.ISBN, false
	}

	books[book.ISBN] = book
	return book.ISBN , true
}

func GetBook(isbn string) (Book, bool) {
	book, exists := books[isbn]
	return book, exists
}

func UpdateBook(isbn string, book Book) (Book, bool) {
	_, exists := books[isbn]

	if exists {
		books[isbn] = book
	}
	return book, exists
}

func DeleteBook(isbn string) {
	delete(books, isbn)
}