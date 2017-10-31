package main

import (
	"errors"
	"strings"
)

// Interprete makes transaction text really works
func Interprete(tx string) bool {

	return true
}

// IsValidTransaction validates transaction
func IsValidTransaction(tx Transaction) bool {
	strs := strings.Split(tx.TxScript, " ")

	if !(len(strs) == 4 && strs[0] == "SEND" && strs[2] == "TO") {
		return false
	}

	return true
}

// Transaction contains transaction script and hash information
type Transaction struct {
	TxHash   []byte
	TxScript string
}

// Parse parses query and returns operation object
func (tx *Transaction) Parse(query string) (*Operation, error) {
	strs := strings.Split(query, " ")

	if !(len(strs) == 4 && strs[0] == "SEND" && strs[2] == "TO") {
		return nil, errors.New("parseERR : wrong syntax")
	}

	return &Operation{
		from: strs[1],
		to:   strs[3],
	}, nil
}

// Operation contains work info about parsed transaction
type Operation struct {
	from string
	to   string
}

// Operate asdf
func (op *Operation) Operate() error {

	return nil
}
