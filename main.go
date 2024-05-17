package main

import (
	"fmt"

	"blockchain-go/blockchain"
)

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	// This will join our previous block's relevant info with the new blocks
// 	hash := sha256.Sum256(info)
// 	//This performs the actual hashing algorithm
// 	b.Hash = hash[:]
// 	//If this ^ doesn't make sense, you can look up slice defaults
// }

// ----- ---- --- -- - - - -- ---

func main() {

	chain := blockchain.InitBlockChain()

	chain.AddBlock("first block after genesis")
	chain.AddBlock("second block after genesis")
	chain.AddBlock("third block after genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}

}
