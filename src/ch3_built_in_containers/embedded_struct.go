package main

import (
	"fmt"
	"go-learning/src/ch3_built_in_containers/book"
)

func main()  {
	bk := book.NewBook(1001, "golang", "golang in action", "james")
	fmt.Println(bk.String())
	book.RefTag(*bk, 0)

	techBook := book.NewTechBook(6001, 10, "golang", "golang in action", "james", "tech")
	fmt.Println(techBook.String())
}
