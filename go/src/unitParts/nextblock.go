package main

import (
	"fmt"
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
	// HashString := "0000000000000000000000000000000000000000000000000000000000000000:0:2::242278"
    hash := sha256.New()
    hash.Write([]byte(HashString))
    return hash.Sum(nil)
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	hash := blk.CalcHash()
	difficulty := blk.Difficulty
	last_d_bytes := hash[len(hash)-int(difficulty):]
	fmt.Println(last_d_bytes)
	count := 0;
	for i := 0 ; i < len(last_d_bytes) ; i++ {
		if last_d_bytes[i] == 0 {
			count++
		}
	}
	fmt.Println(count)
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

func main()  {
	fmt.Printf("Entry\n")
	x := Initial(2)
	x.SetProof(242278)
	fmt.Printf(hex.EncodeToString(x.CalcHash()))
	fmt.Println()
	flag := x.ValidHash()
	if flag == true {
		fmt.Printf("Yes !! a valid hash it is.\n")
	}else {
		fmt.Printf("No !! it's not a valid hash.\n")
	}
	fmt.Printf("Exit\n")
}
