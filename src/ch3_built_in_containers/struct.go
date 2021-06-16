package main

import "fmt"

//define Book
type Book struct {
	id	int
	title string
	author string
	subject string
}
// define factory func to create instance of struct
func NewBook(id int, title, author, subject string) *Book{
	return &Book{
		id:      id,
		title:   title,
		author:  author,
		subject: subject,
	}
}

//define func of struct
func (book *Book) String() string {
	return 	fmt.Sprintf("id=%d, subject=%v, title=%v, author=%v\n",book.id,book.subject,book.title,book.author)
}

// get set
func (book *Book) GetTitle() string{
	return book.title
}

func (book *Book) SetTtile(title string) {
	book.title = title
}

// send struct as parameter
func printBook(book *Book)  {
	fmt.Printf("id=%d, subject=%v, title=%v, author=%v\n",book.id,book.subject,book.title,book.author)
	book.id = 2021
}

func main()  {

	//new book
	var book1 *Book
	book1 = new(Book)
	book1.author = "james"
	book1.id = 1001
	book1.subject = "golang in action"
	book1.title = "golang"

	fmt.Println(book1)

	// new book2
	book2 := Book{
		id:      1002,
		author:  "jordan",
		subject: "python in action",
		title:   "python",
	}


	fmt.Println(&book2)
	printBook(&book2)
	printBook(&book2)
	fmt.Println("=======")
	fmt.Println(book2.GetTitle())
	book2.SetTtile("python Go Go Go!!!")
	fmt.Println(book2.GetTitle())
	fmt.Println(book2.String())

	// use factory func to create instance of struct
	book3 := NewBook(1003, "Java", "gls", "Java in action")
	fmt.Println(book3.String())

}