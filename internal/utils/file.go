package utils

import (
	"log"
	"os"
	"path/filepath"
)

// GetRootDir returns the directory of the executable file.
//
// This function retrieves the path of the currently running executable and returns its directory.
// If there is an error in obtaining the executable path, the function logs a fatal error and terminates the program.
//
// Returns:
// - string: The directory of the executable file.
func GetRootDir() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Erro ao obter o caminho do executável: %v\n", err)
	}
	return filepath.Dir(exePath)
}
