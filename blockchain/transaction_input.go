package blockchain

import "bytes"

// Transaction input
type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
    lockingHash := HashPubKey(in.PubKey)
    return bytes.Equal(lockingHash, pubKeyHash)
}
