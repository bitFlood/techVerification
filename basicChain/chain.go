package main

import (
	"bytes"
)

// Chain is main chain of blocks
var Chain = []Block{GenesisBlock}

// GetLastBlock is get lastest block of blockchain
func GetLastBlock() Block {
	return Chain[len(Chain)-1]
}

// ValidateNewBlock is validate new block
func ValidateNewBlock(b Block) bool {
	vb := GetLastBlock()

	if vb.Index+1 != b.Index {
		return false
	} else if bytes.Compare(vb.Hash, b.PHash) != 0 {
		return false
	}

	return true
}
