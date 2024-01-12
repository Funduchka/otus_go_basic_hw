package storage

import (
	"fmt"
	"os"
)

func Compare(q interface{}) bool {
	var c string
	fmt.Println("Enter what you want to print:")
	_, err := fmt.Scanf("%c", &c)
	if err != nil {
		fmt.Println("Incorrect input, please try again", err)
		os.Exit(1)
	}
	switch c. { 													// разобраться со свичам в структурах!!!!!!
	case ID:
		return book1.ID() > book2.ID()
	case Year:
		return book1.Year() > book2.Year()
	case Size:
		return book1.Size() > book2.Size()
	case Rate:
		return book1.Rate() > book2.Rate()
	default:
		return false
	}
}
