package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var bookTitle = "Go语言"
	fmt.Println(Books{bookTitle, "www.runoob.com", "go语言教程", 6495407})
	fmt.Println(Books{title: bookTitle, author: "www.runoob.com", subject: "Go语言教程", book_id: 6495407})
	fmt.Println(Books{title: bookTitle, author: "www.runoob.com"})

	var book1 Books
	var book2 Books

	book1.title = bookTitle
	book1.author = "www.runoob.com"
	book1.subject = "Go语言教程"
	book1.book_id = 6495407

	book2 = book1

	book2.book_id = 666

	fmt.Println(book1)
	fmt.Println(book2)

	printBook(&book1)
	printBook(&book2)
}

func printBook(book *Books) {
	fmt.Println("Print book!")
	fmt.Printf("Title is %s\n", book.title)
}
