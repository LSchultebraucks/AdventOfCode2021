package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(day int) []string {
	filename := fmt.Sprintf("calendar/day%d/input.txt", day)
	file, errRead := os.Open(filename)
	if errRead != nil {
		log.Fatalf("Failed to read file.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	errClose := file.Close()
	if errClose != nil {
		log.Fatalf("Failed to close file.")
	}
	return text
}

func ReadTestInputFile(day int) []string {
	filename := fmt.Sprintf("calendar/day%d/input_test.txt", day)
	file, errRead := os.Open(filename)
	if errRead != nil {
		log.Fatalf("Failed to read file.")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	errClose := file.Close()
	if errClose != nil {
		log.Fatalf("Failed to close file.")
	}
	return text
}
