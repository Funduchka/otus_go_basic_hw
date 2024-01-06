package main

import (
	"fmt"
	"os"

	"github.com/Funduchka/otus_go_basic_hw/hw03_chessboard/board"
)

func main() {
	var size int // размер поля
	fmt.Println("Enter the size of the board:")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println("Incorrect input, please try again")
		os.Exit(1)
	}

	board.DrawBoard(size)
}
