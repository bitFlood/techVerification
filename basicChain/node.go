package main

import (
	"net/http"
)

// MinerHandler is create new block and append to chain
func MinerHandler(w http.ResponseWriter, r *http.Request) {
	txs := []Transaction{
		Transaction{
			TxScript: "Hello!",
			TxHash:   CreateSHA3Hash([]byte("Hello!")),
		},
		Transaction{
			TxScript: "World!",
			TxHash:   CreateSHA3Hash([]byte("World!")),
		},
		Transaction{
			TxScript: "High Five!",
			TxHash:   CreateSHA3Hash([]byte("High Five!")),
		},
		Transaction{
			TxScript: "Nope",
			TxHash:   CreateSHA3Hash([]byte("Nope")),
		},
		Transaction{
			TxScript: "Sad",
			TxHash:   CreateSHA3Hash([]byte("Sad")),
		},
		Transaction{
			TxScript: "Gonnichiwa",
			TxHash:   CreateSHA3Hash([]byte("Gonnichiwa")),
		},
		Transaction{
			TxScript: "Annyung!",
			TxHash:   CreateSHA3Hash([]byte("Annyung!")),
		},
	}

	txCount := int64(len(txs))

	CreateNewBlock(BlockData{
		TxCount: txCount,
		Tx:      txs,
	})
}

// Run - run the node
func Run() {
	http.HandleFunc("/mineblock", MinerHandler)
	http.HandleFunc("/")

	http.ListenAndServe("3000", nil)
}

func main() {
	Run()
}
