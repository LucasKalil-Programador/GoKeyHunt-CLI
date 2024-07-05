package utils

import "bytes"

func Contains(slice [][]byte, item []byte) bool {
	for _, a := range slice {
		if bytes.Equal(a, item) {
			return true
		}
	}
	return false
}

func Find(slice [][]byte, item []byte) int {
	for i, a := range slice {
		if bytes.Equal(a, item) {
			return i
		}
	}
	return -1
}
