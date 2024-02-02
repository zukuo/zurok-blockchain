package blockchain

import "bytes"

// Transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// Locks signs the output
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
    out.PubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
}

func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
    return bytes.Equal(out.PubKeyHash, pubKeyHash)
}

func NewTXOutput(value int, address string) *TXOutput {
    txo := &TXOutput{value, nil}
    txo.Lock([]byte(address))
    return txo
}
