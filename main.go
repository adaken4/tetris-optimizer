package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var counter rune = 'A'

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		return
	}

	if !strings.HasSuffix(args[0], ".txt") {
		fmt.Println("ERROR")
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	var TetrominoString string
	var tetrominoes [][4][4]bool
	for scanner.Scan() {
		counter++
		TetrominoString += scanner.Text()
		if counter == 5 {
			if scanner.Text() != "" {
				fmt.Println("ERROR")
				return
			}
			validTetromino, err := BuildAndValidateTetromino(TetrominoString)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			tetrominoes = append(tetrominoes, validTetromino)
			counter = 0
			TetrominoString = ""
		} else if len(scanner.Text()) != 4 {
			fmt.Println("ERROR")
			return
		}
	}

	if counter != 0 && counter != 4 {
		fmt.Println("ERROR")
		return
	}
	if len(TetrominoString) != 0 {
		validTetromino, err := BuildAndValidateTetromino(TetrominoString)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tetrominoes = append(tetrominoes, validTetromino)
	}

	fmt.Println("Good File!")
	fmt.Println("Collected Tetrominoes:")
	for _, tetromino := range tetrominoes {
		printTetromino(tetromino)
	}
	// Further processing of tetrominoes to fit them in the smallest square possible
}

func BuildAndValidateTetromino(s string) ([4][4]bool, error) {
	if len(s) != 16 {
		return [4][4]bool{}, errors.New("")
	}
	var TetrominoGrid [4][4]bool
	pieces := 0
	for i, v := range s {
		row := i / 4
		col := i % 4
		if v == '.' {
			TetrominoGrid[row][col] = false
		} else if v == '#' {
			pieces++
			if pieces > 4 {
				return [4][4]bool{}, errors.New("tetrominoes have four pieces max")
			}
			TetrominoGrid[row][col] = true
		} else {
			return [4][4]bool{}, errors.New("bad tetromino format")
		}
	}
	if pieces != 4 {
		return [4][4]bool{}, errors.New("tetrominoes must have exactly four pieces")
	}
	if !ValidateConnections(TetrominoGrid) {
		return [4][4]bool{}, errors.New("tetrominoes must have at least six connecting sides")
	}

	return TetrominoGrid, nil
}

func ValidateConnections(grid [4][4]bool) bool {
	connections := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if grid[i][j] {
				for _, d := range directions {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < 4 && nj >= 0 && nj < 4 && grid[ni][nj] {
						connections++
					}
				}
			}
		}
	}

	return connections >= 6
}

func printTetromino(tetromino [4][4]bool) {

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if tetromino[i][j] {
				fmt.Print(string(counter))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	counter++
}
