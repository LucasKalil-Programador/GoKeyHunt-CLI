package utils

import (
	"btcgo/internal/domain"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

func LoadData() (*domain.Ranges, *domain.Wallets) {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Erro ao obter o caminho do executável: %v\n", err)
	}
	rootDir := filepath.Dir(exePath)

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
