package blockchain

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/nbonfils/cryptopepe/pepe"
	"github.com/nbonfils/cryptopepe/schmekles"
)

//=========
// Structs
//=========

// Block is a basic structure for our blockchain
type Block struct {
	Header *Header
	Data   *Data
}

// Header is the header of the Block
type Header struct {
	ID         uint32
	Timestamp  time.Time
	Bits       uint32
	Nonce      uint32
	MerkleRoot string
	PrevHash   string
	Hash       string
}

// Data is the data shared throughout the blockchain
type Data struct {
	Reward           *schmekles.Transaction
	SchTransactions  *[]schmekles.Transaction
	PepeTransactions *[]pepe.Transaction
}

//=========
// Methods
//=========

// NewBlock generates a new block
func NewBlock() *Block {
	// TODO implement
	return &Block{}
}

// IsValid check if a block is valid, ie hash fullfil requirement
func (b *Block) IsValid() bool {
	// Check if target is reached
	ret, err := regexp.MatchString("^0{5}", b.Header.Hash)
	if err != nil {
		log.Fatal(err)
	}

	// TODO check for MerkleRoot and for Nonce

	// Transactions must be valid
	for _, t := range *b.Data.SchTransactions {
		ret = ret && t.IsValid()
	}

	// Transactions must be valid
	for _, t := range *b.Data.PepeTransactions {
		ret = ret && t.IsValid()
	}

	return ret
}

// Save writes the block on the disk
func (b *Block) Save() {
	// Create the directory if it does not already exists
	if _, err := os.Stat(ChainDir); os.IsNotExist(err) {
		err := os.Mkdir(ChainDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Pad the filename with 0 so it stays sorted in dir
	index := fmt.Sprintf("%09d", b.Header.ID)

	// Open/create the block file
	filename := filepath.Join(ChainDir, index)
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Save the block as a json
	block, err := json.Marshal(*b)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(block)
	if err != nil {
		log.Fatal(err)
	}
}
