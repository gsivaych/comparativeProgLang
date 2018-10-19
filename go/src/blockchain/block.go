package blockchain

import (
	"strconv"
	"strings"
	"encoding/hex"
	"crypto/sha256"
)

type Block struct {
	Generation uint64
	Difficulty uint8
	Data       string
	PrevHash   []byte
	Hash       []byte
	Proof      uint64
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	initialBlock := Block{
		Generation: 0,
		Difficulty: difficulty,
		Data: "",
		PrevHash: make([]byte, 32),
	}
	return initialBlock
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	newBlock := Block{
		Generation: prev_block.Generation + 1,
		Difficulty: prev_block.Difficulty,
		Data: data,
		PrevHash: prev_block.Hash,
	}
	return newBlock
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	HashString := strings.Join([]string{
		hex.EncodeToString(blk.PrevHash),
		strconv.Itoa(int(blk.Generation)),
		strconv.Itoa(int(blk.Difficulty)),
		blk.Data,
		strconv.Itoa(int(blk.Proof)),
		}, ":")
    hash := sha256.New()
    hash.Write([]byte(HashString))
    return hash.Sum(nil)
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	hash := blk.CalcHash()
	difficulty := blk.Difficulty
	last_d_bytes := hash[len(hash)-int(difficulty):]
	count := 0;
	for i := 0 ; i < len(last_d_bytes) ; i++ {
		if last_d_bytes[i] == 0 {
			count++
		}
	}
	if count == int(difficulty){
		return true
	} else {
			return false
		}
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
