package go_chain

import "time"

//Block is block
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

//New create new block
func (block *Block) New(preBlock Block, BPM int) Block {

	block.BPM = BPM
	block.Timestamp = time.Now().String()
	block.Index = preBlock.Index + 1
	block.PrevHash = preBlock.Hash
	block.Hash = CalculateHash(*block)

	return *block

}

//IsValid verify block is valid or not
func (block *Block) IsValid(preblock Block) bool {

	if block.PrevHash != preblock.Hash {
		return false
	}
	if block.Index != preblock.Index+1 {
		return false
	}
	if CalculateHash(*block) != block.Hash {
		return false
	}

	return true

}
