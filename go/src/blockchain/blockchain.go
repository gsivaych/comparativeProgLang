package blockchain

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	chain.Chain = append(chain.Chain, blk)
}

func (chain Blockchain) IsValid() bool {
	if chain.Chain[0].PrevHash != make([]byte, 32) {
		return false
	}
	if chain.Chain[0].Generation != uint64(0){
		return false
	}
	for i:= 0; i < len(chain.Chain);i++ {
		if chain.Chain[i].Difficulty != chain.Chain[0].Difficulty {
			return false
		}
		if chain.Chain[i+1].Generation != chain.Chain[i].Generation + 1 {
			return false
		}
		if chain.Chain[i+1].PrevHash != chain.Chain[i].Hash {
			return false
		}
		if chain.Chain[i].Hash != chain.Chain[i].CalcHash() {
			return false
		}
		if !chain.Chain[i].ValidHash() {
			return false
		}
	}
	return true
}
