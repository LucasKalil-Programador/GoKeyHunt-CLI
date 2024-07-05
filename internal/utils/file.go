package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetRootDir() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Erro ao obter o caminho do executável: %v\n", err)
	}
	return filepath.Dir(exePath)
}
