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

func main() {

}
