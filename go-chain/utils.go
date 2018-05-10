package go_chain

import (
	"crypto/sha256"
	"encoding/hex"
)

func CalculateHash(block Block) string {

	record := string(block.Index) + block.Timestamp + block.PrevHash + string(block.BPM)
	hash := sha256.New()
	hash.Write([]byte(record))
	recordHashed := hash.Sum(nil)

	return hex.EncodeToString(recordHashed)

}
