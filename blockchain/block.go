package blockchain

import (
	"crypto/sha256"
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

//=======
// Types
//=======

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

//================
// Header Methods
//================

// String produce a pretty print of the Header
func (h *Header) String() string {
	return fmt.Sprintf(
		"ID: %v\nDate: %v\nBits: %v\nNonce %v\nMerkle Root: %v\nPrevious Hash: %v\nHash: %v\n",
		h.ID, h.Timestamp, h.Bits, h.Nonce, h.MerkleRoot, h.PrevHash, h.Hash,
	)
}

// Sum256 produce a string that is the hash (sha256) of all Header fields
// except "Hash"
func (h *Header) Sum256() string {
	headerString := fmt.Sprintf("%v%v%v%v%v%v", h.ID, h.Timestamp, h.Bits,
		h.Nonce, h.MerkleRoot, h.PrevHash)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(headerString)))
}

//===============
// Block Methods
//===============

// NewBlock generates a new block
func NewBlock() *Block {
	// TODO implement
	return &Block{}
}

// IsValid check if a block is valid, ie hash fullfil requirement
func (b *Block) IsValid() bool {
	log.Printf("[DEBUG] Checking block %v validity", b.Header.ID)

	// Check if target is reached
	match, err := regexp.MatchString("^0{5}", b.Header.Hash)
	if err != nil {
		log.Fatal(err)
	}
	if !match {
		return false
	}

	// Check for the MerkleRoot
	// in this order: reward -> schmekles tr -> pepe tr
	var hashes []byte
	hashes = append(hashes, b.Data.Reward.Hash[:]...)
	for _, t := range *b.Data.SchTransactions {
		hashes = append(hashes, t.Hash[:]...)
	}
	for _, t := range *b.Data.PepeTransactions {
		hashes = append(hashes, t.Hash[:]...)
	}
	rootHash := fmt.Sprintf("%x", sha256.Sum256(hashes))
	if rootHash != b.Header.MerkleRoot {
		return false
	}

	// Check for the Nonce validity
	if b.Header.Hash != b.Header.Sum256() {
		return false
	}

	// Transactions must be valid
	for _, t := range *b.Data.SchTransactions {
		if !t.IsValid() {
			return false
		}
	}

	// Transactions must be valid
	for _, t := range *b.Data.PepeTransactions {
		if !t.IsValid() {
			return false
		}
	}

	return true
}

// Save writes the block on the disk
func (b *Block) Save() {
	log.Printf("[DEBUG] Saving block %v on disk", b.Header.ID)

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
