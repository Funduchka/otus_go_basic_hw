package storage

import (
	"fmt"
	"os"
)

func Printer() {
	var c string
	fmt.Println("Enter what you want to print:")
	_, err := fmt.Scanf("%c", &c)
	if err != nil {
		fmt.Println("Incorrect input, please try again")
		os.Exit(1)
	}
}
