package blockchain

import "bytes"

// Transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// Locks signs the output
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Encode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
    out.PubKeyHash = pubKeyHash
}

func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
    return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func NewTXOutput(value int, address string) *TXOutput {
    txo := &TXOutput{value, nil}
    txo.Lock([]byte(address))
    return txo
}
