package utils

import (
	"GoKeyHunt/internal/domain"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

// LoadData loads ranges and wallets data from JSON files.
//
// This function retrieves the root directory of the application, and then loads
// the ranges and wallets data from JSON files located in the "data" directory.
// It logs a fatal error if either file fails to load.
//
// Returns:
// - *domain.Ranges: A pointer to a Ranges structure containing the loaded ranges data.
// - *domain.Wallets: A pointer to a Wallets structure containing the loaded wallets data.
func LoadData() (*domain.Ranges, *domain.Wallets) {
	rootDir := GetRootDir()

	ranges, err := LoadRanges(filepath.Join(rootDir, "data", "ranges.json"))
	if err != nil {
		log.Fatalf("Failed to load ranges: %v", err)
	}

	wallets, err := LoadWallets(filepath.Join(rootDir, "data", "wallets.json"))
	if err != nil {
		log.Fatalf("Failed to load wallets: %v", err)
	}

	return ranges, wallets
}

// LoadRanges loads ranges data from a specified JSON file.
//
// This function reads the JSON file, unmarshals its content into a Ranges structure,
// and returns a pointer to the Ranges structure or an error if the operation fails.
//
// Parameters:
// - filename: The path to the JSON file containing the ranges data.
//
// Returns:
// - *domain.Ranges: A pointer to a Ranges structure containing the loaded ranges data.
// - error: An error if there is an issue reading or unmarshalling the file.
func LoadRanges(filename string) (*domain.Ranges, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var ranges domain.Ranges
	if err := json.Unmarshal(bytes, &ranges); err != nil {
		return nil, err
	}

	return &ranges, nil
}

// LoadWallets loads wallets data from a specified JSON file.
//
// This function reads the JSON file, unmarshals its content into a temporary structure,
// decodes the Base58 addresses, and returns a pointer to the Wallets structure or an error
// if the operation fails.
//
// Parameters:
// - filename: The path to the JSON file containing the wallets data.
//
// Returns:
// - *domain.Wallets: A pointer to a Wallets structure containing the loaded wallets data.
// - error: An error if there is an issue reading or unmarshalling the file.
func LoadWallets(filename string) (*domain.Wallets, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	type WalletsTemp struct {
		Addresses []string `json:"wallets"`
	}

	var walletsTemp WalletsTemp
	if err := json.Unmarshal(bytes, &walletsTemp); err != nil {
		return nil, err
	}

	var wallets domain.Wallets
	for _, address := range walletsTemp.Addresses {
		wallets.Addresses = append(wallets.Addresses, Decode(address)[1:21])
	}

	return &wallets, nil
}
