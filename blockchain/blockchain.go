package main

import (
        "crypto/sha256"
        "encoding/json"
        "fmt"
        "strconv"
        "strings"
        "time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int

}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data) //converts block to json
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow) // concat of previous hash, current block's data, timestamp and Proof Of Work (pow)
	blockHash := sha256.Sum256([]byte(blockData)) //hashes the earlier concatenation with the SHA256 algorithm
	return fmt.Sprintf("%x", blockHash) // Returns the base 16 hash as a string
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
			b.pow++
			b.hash = b.calculateHash()
	}
}