package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteInOutput writes a message to the output file.
//
// This function retrieves the root directory of the application, constructs the
// path to the "output.txt" file, and appends the given message to this file.
//
// Parameters:
// - message: The message to be written to the output file.
func WriteInOutput(message string) {
	rootDir := GetRootDir()
	appendToFile(filepath.Join(rootDir, "output.txt"), message)
}

// appendToFile appends text to a file.
//
// This function opens the specified file in append mode, writes the provided text
// to it, and then closes the file. If the file does not exist, it creates it.
//
// Parameters:
// - filename: The path to the file where the text will be appended.
// - text: The text to append to the file.
//
// Returns:
// - error: An error if there is an issue opening or writing to the file.
func appendToFile(filename, text string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}

	return nil
}
