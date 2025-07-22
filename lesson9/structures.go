package main

import "fmt"

type Book struct {
	Title string
	Author string
	Year int
}

func main() {
	book := Book{
		Title: "1984",
		Author: "George Orwell",
		Year: 1949,
	}

	book2 := Book{
		Title: "Martin Iden",
		Author: "Jack London",
		Year: 1909,
	}

	book3 := Book{
		Title: "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Year: 1925,
	}

	books := []Book{book, book2, book3}

	for _, book := range books {
		fmt.Println("Название: ", book.Title, "Автор: ", book.Author, "Год: ", book.Year)
	}
	
}