package main

import (
	"golang.org/x/crypto/sha3"
)

// CreateSHA3Hash is "title is content."
func CreateSHA3Hash(origin []byte) []byte {
	hasher := sha3.New512()
	hasher.Write(origin)
	return hasher.Sum(nil)
}
