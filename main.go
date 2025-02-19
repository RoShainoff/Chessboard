package main

import (
	"fmt"
	"strings"
)

func generateChessboard(size uint32) string {
	var board strings.Builder

	// Верхняя граница
	for i := uint32(0); i < size; i++ {
		board.WriteString("|-----")
	}
	board.WriteString("|\n")

	// Основная часть доски
	for i := uint32(0); i < size; i++ {
		board.WriteString("|") // Левая граница
		for j := uint32(0); j < size; j++ {
			if (j+i)%2 == 0 {
				board.WriteString("  #  |")
			} else {
				board.WriteString("     |")
			}
		}
		board.WriteString("\n") // Нижняя граница строки
		for j := uint32(0); j < size; j++ {
			board.WriteString("|-----")
		}
		board.WriteString("|\n")
	}

	return board.String()
}

func main() {
	size := uint32(8)
	fmt.Print(generateChessboard(size))
}
