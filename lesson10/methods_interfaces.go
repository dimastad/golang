package main
import "fmt"


type Book struct {
	Title string
	Author string
	Year int
}

type Printer interface {
    PrintInfo()
}

func (b Book) PrintInfo() {
    fmt.Println("Книга:", b.Title)
}

func PrintBook(p Printer) {
	p.PrintInfo()
}

func main() {
	book := Book{
		Title: "1984",
		Author: "George Orwell",
		Year: 1949,
	}

	PrintBook(book)
}
