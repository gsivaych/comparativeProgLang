package main

import (
	"fmt"
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
	inblk := Block{
		Generation: 0,
		Difficulty: difficulty,
		Data: "Data_string",
		PrevHash: []byte{0x00},
	}
	return inblk
}

func main(){
	x := Initial(2)
	fmt.Println(x);
}
