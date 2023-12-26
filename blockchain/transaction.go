package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// Amount of reward
const subsidy = 10

// Transaction for currency transaction
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// Transaction input
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

// Transaction output
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

func (tx *Transaction) SetID() {
    var encoded bytes.Buffer
    var hash [32]byte

    enc := gob.NewEncoder(&encoded)
    err := enc.Encode(tx)
    HandleError(err)
    hash = sha256.Sum256(encoded.Bytes())
    tx.ID = hash[:]
}

// Checks if the transaction is coinbase
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

// Check whether the address initiated the transaction
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

// Checks if the output can be unlocked with the provided data
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}

func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	// Find unspent outputs to spend
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	// If the account balance is insufficient, exit
	if acc < amount {
		log.Panic("Error: Not enough funds")
	}

	// Create inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		HandleError(err)

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Create outputs
	outputs = append(outputs, TXOutput{amount, to})

	// If there is change, create a change output
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from})
	}

	// Create transaction
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
