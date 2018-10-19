package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/hex"
)

// some useful tests of Blocks
func TestBlock(t *testing.T){
	b0 := Initial(uint8(2))
	assert.Equal(t, b0.Generation, uint64(0), "Generation not zero for Initial Block.")
	assert.Equal(t, b0.Difficulty, uint8(2), "Difficulty mismatch for Initial Block.")
	assert.Equal(t, b0.Data, "", "Data mismatch for Initial Block.")
	assert.Equal(t, b0.PrevHash, make([]byte, 32), "PrevHash mismatch for Initial Block.")

	assert.False(t, b0.ValidHash(), "Invalid Hash being told Valid.")
	b0.SetProof(uint64(242278))
	assert.Equal(t, b0.Proof, uint64(242278), "SetProof didn't work.")
	// Set proof did CalcHash
	Hash0 := "29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000"
	assert.Equal(t, hex.EncodeToString(b0.Hash), Hash0, "Inapropriate Hash Calculated.")
	assert.True(t, b0.ValidHash(), "Valid Hash being told Invalid.")

	b1 := b0.Next("message")
	assert.Equal(t, b1.Generation, b0.Generation+1, "Generation not 1 for Next Block.")
	assert.Equal(t, b1.Difficulty, b0.Difficulty, "Difficulty mismatch for Next Block.")
	assert.Equal(t, b1.Data, "message", "Data mismatch for Second Block.")
	assert.Equal(t, b1.PrevHash, b0.Hash, "PrevHash mismatch for Second Block.")
}

func assertPanic(t *testing.T, f func()) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic while adding invalid block.")
        }
    }()
    f()
}

func TestBlockchain(t *testing.T)  {
		b0 := Initial(uint8(2))
		b0.SetProof(uint64(242278))
		newChain := new(Blockchain)
		newChain.Add(b0)
		assert.Equal(t, len(newChain), 1, "Adding Initial block to newChain failed.")

		b1 := b0.Next("message")

		b1.SetProof(uint64(242278)) // same proof as initial won't result in valid hash
		assertPanic(t, newChain.Add(b1))

		b1.SetProof(uint64(75729))
		newChain.Add(b1)
		assert.Equal(t, len(newChain), 2, "Adding Second block to theChain failed.")

		assert.True(t, newChain.IsValid(), "Valid chain being told invalid.")
}

func TestMining(t *testing.T){
	b0 := Initial(uint8(2))
	b0.Mine(1)
	proof := uint64(242278)
	hash :=	"29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000"
	assert.Equal(t, b0.Proof, proof, "Mining fail for Initial Block : Proof mismatch")
	assert.Equal(t, hex.EncodeToString(b0.Hash), hash, "Mining fail for Initial Block : HAsh mismatch")


	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	proof := uint64(41401)
	hash :=	"d558f4b9a0b5df021a98066efa5779758cd02f721ebcd8112872265799580000"
	assert.Equal(t, b1.Proof, proof, "Mining fail for Second Block : Proof mismatch")
	assert.Equal(t, hex.EncodeToString(b1.Hash), hash, "Mining fail for Second Block : HAsh mismatch")

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	proof := uint64(195955)
	hash :=	"b611d6fc52964c1c89c717731807c5785ca6bf50d0922b0fc2fdf283757c0000"
	assert.Equal(t, b2.Proof, proof, "Mining fail for Third Block : Proof mismatch")
	assert.Equal(t, hex.EncodeToString(b2.Hash), hash, "Mining fail for Third Block : HAsh mismatch")
}
