package main

import (
	"fmt"
	"strings"
)

func generateChessboard(size int) string {
	var board strings.Builder

	// Верхняя граница
	for i := 0; i < size; i++ {
		board.WriteString("|-----")
	}
	board.WriteString("|\n")

	// Основная часть доски
	for i := 0; i < size; i++ {
		board.WriteString("|") // Левая граница
		for j := 0; j < size; j++ {
			if (j+i)%2 == 0 {
				board.WriteString("  #  |")
			} else {
				board.WriteString("     |")
			}
		}
		board.WriteString("\n") // Нижняя граница строки
		for j := 0; j < size; j++ {
			board.WriteString("|-----")
		}
		board.WriteString("|\n")
	}

	return board.String()
}

func main() {
	var size int
	fmt.Print("Enter board size: ")
	fmt.Scan(&size)
	fmt.Print(generateChessboard(size))
}
