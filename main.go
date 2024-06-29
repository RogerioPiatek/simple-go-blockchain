package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Data model to our block, containing the necessary fields like the PrevHash
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

// Calculates the Hash of our Block
func calculateHash(block Block) string {
	//mounting the record, basically the fields of the struct
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Needs to receive oldBlock so we get the PrevHash
func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()

	//auto increment the index based on previous block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	//BPM is receive as a parameter
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

// Validates if our block is valid, by comparing properties from both blocks (old and new Block)
func isBlockValid(newBlock Block, oldBlock Block) bool {
	//If they're not equal, our logic was broken somewhere
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	//all the tests passed
	return true
}

// Ensure that the most updated Node is used as the blockchain (most blocks)
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func main() {

}
