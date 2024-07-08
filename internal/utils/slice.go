package utils

import "bytes"

// Contains checks if a slice of byte slices contains a specific byte slice.
// It returns true if the item is found in the slice, and false otherwise.
//
// Parameters:
// - slice: A slice of byte slices to search within.
// - item: The byte slice to search for.
//
// Returns:
// - bool: true if the item is found, false otherwise.
func Contains(slice [][]byte, item []byte) bool {
	for _, a := range slice {
		if bytes.Equal(a, item) {
			return true
		}
	}
	return false
}

// Find searches for a specific byte slice in a slice of byte slices and returns its index.
// It returns the index of the item if found, and -1 otherwise.
//
// Parameters:
// - slice: A slice of byte slices to search within.
// - item: The byte slice to search for.
//
// Returns:
// - int: The index of the item if found, -1 otherwise.
func Find(slice [][]byte, item []byte) int {
	for i, a := range slice {
		if bytes.Equal(a, item) {
			return i
		}
	}
	return -1
}
