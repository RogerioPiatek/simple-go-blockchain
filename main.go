package main

import (
	"crypto/sha256"
	"encoding/hex"
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

func main() {

}
