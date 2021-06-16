package book

import (
	"fmt"
	"reflect"
)

type Book struct {
	id	int			"this is book id"
	title string
	subject string
	author  string
}

type TechBook struct {
	cat string
	int //count
	Book
}

func NewTechBook(id, typeInt int, title, subject, author, cat string) *TechBook {
	book := NewBook(id, title, subject, author)
	techBook := new(TechBook)
	techBook.cat = cat
	techBook.int = typeInt
	techBook.Book = *book
	return techBook
}

func NewBook(id int, title, subject, author string) *Book {
	return &Book{id,title,subject,author}
}

func (book *Book) String() string {
	return fmt.Sprintf("id=%d, title=%v, author=%v，subject=%v\n",
		book.id,
		book.title,
		book.author,
		book.subject,
		)
}

func RefTag(book Book, idx int)  {
	bookType := reflect.TypeOf(book)
	idxfField := bookType.Field(idx)
	tag := idxfField.Tag
	fmt.Println(tag)
}