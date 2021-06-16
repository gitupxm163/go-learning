package book

import "fmt"

type Book struct {
	id	int
	title string
	subject string
	author  string
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
