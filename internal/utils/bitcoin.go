package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"golang.org/x/crypto/ripemd160"
)

// GenerateWif generates a Wallet Import Format (WIF) string from a given private key.
//
// This function takes a private key as a big.Int, converts it to a hexadecimal string,
// and processes it to generate a WIF string. The process includes prefixing, suffixing,
// computing checksums, and base58 encoding.
//
// Parameters:
// - privKeyInt: A pointer to a big.Int representing the private key.
//
// Returns:
// - string: The WIF string representation of the private key.
func GenerateWif(privKeyInt *big.Int) string {
	privKeyHex := fmt.Sprintf("%064x", privKeyInt)

	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	extendedKey := append([]byte{byte(0x80)}, privKeyBytes...)
	extendedKey = append(extendedKey, byte(0x01))

	firstSHA := sha256.Sum256(extendedKey)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:4]

	finalKey := append(extendedKey, checksum...)

	wif := Encode(finalKey)

	return wif
}

// CreatePublicHash160 generates a Hash160 from a given private key.
//
// This function takes a private key as a big.Int, derives the corresponding
// public key in compressed format, and computes the Hash160 (RIPEMD-160 after SHA-256).
//
// Parameters:
// - privKeyInt: A pointer to a big.Int representing the private key.
//
// Returns:
// - []byte: The Hash160 of the public key.
func CreatePublicHash160(privKeyInt *big.Int) []byte {
	privKeyBytes := privKeyInt.Bytes()

	privKey := secp256k1.PrivKeyFromBytes(privKeyBytes)

	compressedPubKey := privKey.PubKey().SerializeCompressed()

	pubKeyHash := hash160(compressedPubKey)

	return pubKeyHash
}

// hash160 computes the Hash160 of a given byte slice.
//
// This function takes a byte slice, computes its SHA-256 hash, and then computes
// the RIPEMD-160 hash of the SHA-256 hash.
//
// Parameters:
// - b: A byte slice to hash.
//
// Returns:
// - []byte: The Hash160 of the input byte slice.
func hash160(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sha256Hash := h.Sum(nil)

	r := ripemd160.New()
	r.Write(sha256Hash)
	return r.Sum(nil)
}
