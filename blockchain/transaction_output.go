package blockchain

import (
	"bytes"
	"encoding/gob"
	"github.com/zukuo/zurok-blockchain/util"
	"github.com/zukuo/zurok-blockchain/wallet"
)

// Transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// Locks signs the output
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
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

type TXOutputs struct {
	Outputs []TXOutput
}

func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	util.HandleError(err)
	return buff.Bytes()
}

func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs
	enc := gob.NewDecoder(bytes.NewReader(data))
	err := enc.Decode(&outputs)
	util.HandleError(err)
	return outputs
}
