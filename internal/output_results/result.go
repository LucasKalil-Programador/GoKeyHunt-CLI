package output_results

import (
	"btcgo/internal/domain"
	"btcgo/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"sort"
)

// Result

type Result struct {
	WalletIndex int    `json:"Wallet"`
	Key         string `json:"Key"`
	Wif         string `json:"Wif"`
}

func NewResult(key *big.Int, wallets domain.Wallets) *Result {
	address := utils.CreatePublicHash160(key)
	walletIndex := utils.Find(wallets.Addresses, address) + 1
	return &Result{WalletIndex: walletIndex, Key: fmt.Sprintf("%064x", key), Wif: utils.GenerateWif(key)}
}

func (r *Result) String() string {
	return fmt.Sprintf("WalletIndex: %d, Key: %s, Wif: %s", r.WalletIndex, r.Key, r.Wif)
}

// ResultArray

type ResultArray struct {
	Resuts []Result `json:"Wallets found"`
}

func NewEmptyResultArray() *ResultArray {
	return &ResultArray{}
}

// Define a type that implements sort.Interface for sorting by WalletIndex
type ByWalletIndex []Result

func (a ByWalletIndex) Len() int           { return len(a) }
func (a ByWalletIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWalletIndex) Less(i, j int) bool { return a[i].WalletIndex < a[j].WalletIndex }

// Function to sort an array of Result by WalletIndex field
func SortByWalletIndex(resuts []Result) []Result {
	sort.Sort(ByWalletIndex(resuts))
	return resuts
}

func (rArray *ResultArray) AppendIfNotExist(newResult Result) bool {
	index := sort.Search(len(rArray.Resuts), func(i int) bool {
		return rArray.Resuts[i].WalletIndex >= newResult.WalletIndex
	})

	if len(rArray.Resuts) <= index {
		rArray.Resuts = append(rArray.Resuts, newResult)
		return true
	}

	if rArray.Resuts[index] != newResult {
		rArray.Resuts = append(rArray.Resuts, Result{}) // add an empty interval to extend the slice
		copy(rArray.Resuts[index+1:], rArray.Resuts[index:])
		rArray.Resuts[index] = newResult
		return true
	}
	return false
}

func (rArray *ResultArray) Save(jsonPath string) bool {
	jsonData, err := json.MarshalIndent(rArray, "", "	")
	if err != nil {
		log.Println("Error on Marshal function:", err)
		return false
	}

	err = os.WriteFile(jsonPath, jsonData, 0644)
	if err != nil {
		log.Println("Error on write json file:", err)
		return false
	}
	return true
}

func Read(filePath string) (*ResultArray, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var resultsArray ResultArray
	if err := json.Unmarshal(bytes, &resultsArray); err != nil {
		return nil, err
	}

	resultsArray.Resuts = SortByWalletIndex(resultsArray.Resuts)
	return &resultsArray, nil
}

func ReadOrNew(filePath string) *ResultArray {
	IntervalArr, err := Read(filePath)
	if err != nil {
		return NewEmptyResultArray()
	}
	return IntervalArr
}
