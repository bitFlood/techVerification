package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"log"
	"time"
)

// GenesisBlock is origin of chain
var GenesisBlock = Block{
	Index:      0,
	Hash:       CreateSHA3Hash([]byte("The new waves coming")),
	PHash:      CreateSHA3Hash([]byte("")),
	Timestamp:  time.Now().Unix(),
	TxRootHash: CreateSHA3Hash([]byte("Transaction")),

	Data: new(BlockData),
}

// Block is main resource of blockchain
type Block struct {
	// Header
	Index      int64
	Hash       []byte
	PHash      []byte
	Timestamp  int64
	TxRootHash []byte

	// Body
	Data *BlockData
}

// BlockData is body of Block
type BlockData struct {
	TxCount int64
	Tx      []Transaction
}

// CreateHash - create block's hash
func (block *Block) CreateHash() []byte {
	data, err := json.Marshal(block)
	if err != nil {
		log.Println(err)
		return nil
	}

	return CreateSHA3Hash(data)
}

// CreateNewBlock - create new block based on previous block
func CreateNewBlock(data BlockData) *Block {
	block := GetLastBlock()

	var nb = &Block{
		Index:     block.Index + 1,
		PHash:     block.Hash,
		Timestamp: time.Now().Unix(),
		Data:      &data,
	}

	nb.Hash = nb.CreateHash()

	return nb
}

// AppendByteArray appends byte array
func AppendByteArray(bef []byte, aft []byte) []byte {
	result := make([]byte, len(bef)+len(aft))

	for i := 0; i < len(bef); i++ {
		result[i] = bef[i]
	}

	for i := 0; i < len(aft); i++ {
		result[i+len(bef)] = aft[i]
	}

	return result
}

// GetTxRootHash makes root hash of transactions
func (block *Block) GetTxRootHash() []byte {
	//initialize
	length := len(block.Data.Tx)
	txBase := make([][]byte, length)

	//change string to byte
	for pos, obj := range block.Data.Tx {
		b64.StdEncoding.Encode(txBase[pos], []byte(obj.TxScript))
	}

	txLast := make([][]byte, length)
	copy(txLast, txBase)

	//make hashs
	for {
		txStart := make([][]byte, length)
		copy(txStart, txLast)
		length = length/2 + length%2
		txLast := make([][]byte, length)

		for i := 0; i < length; i++ {
			if i*2+1 >= len(txStart) {
				txLast[i] = CreateSHA3Hash(txStart[i*2])
				break
			}
			txLast[i] = CreateSHA3Hash(AppendByteArray(txStart[i*2], txStart[i*2+1]))
		}

		if length <= 1 {
			return txLast[0]
		}
	}
}
