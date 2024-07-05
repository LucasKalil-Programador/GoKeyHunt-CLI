package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteInOutput(message string) {
	rootDir := GetRootDir()
	appendToFile(filepath.Join(rootDir, "output.txt"), message)
}

func appendToFile(filename, text string) error {
	// Abre o arquivo no modo de adição
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	// Escreve a string no final do arquivo
	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %v", err)
	}

	return nil
}
