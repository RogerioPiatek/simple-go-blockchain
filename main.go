package main

// Data model to our block, containing the necessary fields like the PrevHash
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func main() {

}
