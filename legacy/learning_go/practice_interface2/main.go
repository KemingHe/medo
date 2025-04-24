package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// readFileToStdout reads the content of the given filename and writes it to stdout
func readFileToStdout(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	bytesWritten, err := io.Copy(os.Stdout, file)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	log.Printf("Successfully read %d bytes from %s\n", bytesWritten, filename)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing text file name in command line arguments, exiting...")
	}

	fileName := os.Args[1]
	if err := readFileToStdout(fileName); err != nil {
		log.Fatal(err)
	}
}
