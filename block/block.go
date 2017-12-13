package block

import (
	"regexp"
	"time"

	"github.com/nbonfils/cryptopepe/pepe"
	"github.com/nbonfils/cryptopepe/schmekles"
	"log"
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
	ID        uint32
	Timestamp time.Time
	Bits      uint32
	Nonce     uint32
	Hash      string
	PrevHash  string
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
func NewBlock() Block {
	// TODO
	return Block{}
}

// IsValid check if a block is valid, ie hash fullfil requirement
func (b Block) IsValid() bool {
	ret, err := regexp.MatchString("^0{5}", b.Header.Hash)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

// Save writes the block on the disk
func (b Block) Save() {
	// TODO
}
