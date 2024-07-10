package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	for scanner.Scan() {
		counter++
		if counter == 5 {
			if scanner.Text() != "" {
				fmt.Println("ERROR")
				return
			}
			counter = 0
		} else if len(scanner.Text()) != 4 {
			fmt.Println("ERROR")
			return
		}
	}

	if counter != 0 && counter != 4 {
		fmt.Println("ERROR")
		return
	}

	fmt.Println("Good File!")

}
