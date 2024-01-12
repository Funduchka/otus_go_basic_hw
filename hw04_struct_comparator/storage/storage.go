package storage

import (
	"fmt"
)

type Book struct {
	ID int				//ISBN
	Title string
	Author string
	Year int
	Size int
	Rate float32
}

func setBookID(b *Book) int {
	b.ID = ID
}

func getBookID(b *Book) int {
	return b.ID
}

func setBookTitle(b *Book) string {
	b.Title = Title
}

func getBookTitle(b *Book) string {
	return b.Title
}

func setBookAuthor(b *Book) string {
	b.Author = Author
}

func getBookAuthor(b *Book) string {
	return b.Author
}

func setBookYear(b *Book) int {
	b.Year = Year
}

func getBookYear(b *Book) int {
	return b.Year
}

func setBookSize(b *Book) int {
	b.Size = Size
}

func getBookSize(b *Book) int {
	return b.Size
}

func setBookRate(b *Book) float32 {
	b.Rate = Rate
}

func getBookRate(b *Book) float32 {
	return b.Rate
}



//разобраться, почему не работает
b1 := Book{
	ID: 1,
	Title "Troubled Blood",
	Author "Robert Galbraith",
	Year 2020,
	Size 944,
	Rate 4.6,
}

b2 := Book{
	ID: 2,
	Title "The Great Gatsby",
	Author "Scott Fitzgerald",
	Year 1925,
	Size 110,
	Rate 4.6,
}

