package blockchain

import (
	"github.com/dgraph-io/badger" // This is our database import
)

const dbPath = "./tmp/blocks"

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}
