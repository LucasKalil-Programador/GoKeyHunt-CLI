package output_results

import (
	"GoKeyHunt/internal/domain"
	"GoKeyHunt/internal/utils"
	"fmt"
	"math/big"
	"sort"
)

// Result represents a result containing wallet index, key, and WIF (Wallet Import Format).
type Result struct {
	WalletIndex int    `json:"Wallet"` // The index of the wallet.
	Key         string `json:"Key"`    // The private key in hexadecimal format.
	Wif         string `json:"Wif"`    // The private key in Wallet Import Format.
}

// NewResult creates a new Result instance.
//
// This function takes a big.Int pointer representing the key and a domain.Wallets instance,
// then calculates the wallet index, key in hexadecimal format, and the WIF.
//
// Parameters:
// - key: A pointer to a big.Int representing the private key.
// - wallets: A domain.Wallets instance containing wallet addresses.
//
// Returns:
// - *Result: A pointer to the newly created Result instance.
func NewResult(key *big.Int, wallets domain.Wallets) *Result {
	address := utils.CreatePublicHash160(key)
	walletIndex := utils.Find(wallets.Addresses, address) + 1
	return &Result{WalletIndex: walletIndex, Key: fmt.Sprintf("%064x", key), Wif: utils.GenerateWif(key)}
}

// String returns the string representation of the Result.
//
// This method returns a formatted string containing the wallet index, key, and WIF.
//
// Returns:
// - string: A formatted string representation of the Result.
func (r *Result) String() string {
	return fmt.Sprintf("WalletIndex: %d, Key: %s, Wif: %s", r.WalletIndex, r.Key, r.Wif)
}

// ResultArray represents an array of Result instances.
type ResultArray struct {
	Resuts []Result `json:"Wallets found"` // A slice of Result instances.
}

// NewEmptyResultArray creates a new, empty ResultArray instance.
//
// This function initializes and returns an empty ResultArray.
//
// Returns:
// - *ResultArray: A pointer to the newly created, empty ResultArray instance.
func NewEmptyResultArray() *ResultArray {
	return &ResultArray{}
}

// ByWalletIndex is a type that implements sort.Interface for sorting by WalletIndex.
type ByWalletIndex []Result

// Len returns the number of elements in the collection.
//
// Returns:
// - int: The number of elements in the collection.
func (a ByWalletIndex) Len() int {
	return len(a)
}

// Swap swaps the elements with indexes i and j.
//
// Parameters:
// - i: The first index to swap.
// - j: The second index to swap.
func (a ByWalletIndex) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less reports whether the element with index i should sort before the element with index j.
//
// Parameters:
// - i: The first index to compare.
// - j: The second index to compare.
//
// Returns:
// - bool: True if the element with index i should sort before the element with index j.
func (a ByWalletIndex) Less(i, j int) bool {
	return a[i].WalletIndex < a[j].WalletIndex
}

// SortByWalletIndex sorts an array of Result by the WalletIndex field.
//
// This function takes a slice of Result and sorts it in place by WalletIndex.
//
// Parameters:
// - results: A slice of Result instances to be sorted.
//
// Returns:
// - []Result: The sorted slice of Result instances.
func SortByWalletIndex(results []Result) []Result {
	sort.Sort(ByWalletIndex(results))
	return results
}

// AppendIfNotExist appends a new Result to the ResultArray if it doesn't already exist.
//
// This method searches for the newResult in the ResultArray and appends it if it's not found.
//
// Parameters:
// - newResult: The Result instance to be appended.
//
// Returns:
// - bool: True if the newResult was appended, false otherwise.
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
