package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
)

// Amount of reward
const subsidy = 10

// Transaction for currency transaction
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// Serialize a transaction
func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	HandleError(err)

	return encoded.Bytes()
}

// Hash a transaction
func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	return hash[:]
}

func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	if tx.IsCoinbase() {
		return
	}

	for _, vin := range tx.Vin {
		if prevTXs[hex.EncodeToString(vin.Txid)].ID == nil {
			log.Panic("ERROR: Previous transaction is incorrect")
		}
	}

	txCopy := tx.TrimmedCopy()

	for inID, vin := range txCopy.Vin {
		prevTx := prevTXs[hex.EncodeToString(vin.Txid)]
		txCopy.Vin[inID].Signature = nil
		txCopy.Vin[inID].PubKey = prevTx.Vout[vin.Vout].PubKeyHash
		txCopy.ID = txCopy.Hash()
		txCopy.Vin[inID].PubKey = nil

		r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.ID)
		HandleError(err)
		signature := append(r.Bytes(), s.Bytes()...)
		tx.Vin[inID].Signature = signature
	}

}

func (tx Transaction) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("--- Transaction %x:", tx.ID))

	for i, input := range tx.Vin {
		lines = append(lines, fmt.Sprintf("	   Input %d:", i))
		lines = append(lines, fmt.Sprintf("		TXID:	   %x", input.Txid))
		lines = append(lines, fmt.Sprintf("		Out:       %d", input.Vout))
		lines = append(lines, fmt.Sprintf("		Signature: %x", input.Signature))
		lines = append(lines, fmt.Sprintf("		PubKey:    %x", input.PubKey))
	}

	for i, output := range tx.Vout {
		lines = append(lines, fmt.Sprintf("	   Output %d:", i))
		lines = append(lines, fmt.Sprintf("		Value:  %d", output.Value))
		lines = append(lines, fmt.Sprintf("		Script: %x", output.PubKeyHash))
	}

	return strings.Join(lines, "\n")
}

func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	for _, vin := range tx.Vin {
		inputs = append(inputs, TXInput{vin.Txid, vin.Vout, nil, nil})
	}

	for _, vout := range tx.Vout {
		outputs = append(outputs, TXOutput{vout.Value, vout.PubKeyHash})
	}

	txCopy := Transaction{tx.ID, inputs, outputs}
	return txCopy
}

func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	for _, vin := range tx.Vin {
		if prevTXs[hex.EncodeToString(vin.Txid)].ID == nil {
			log.Panic("ERROR: Previous transaction is incorrect")
		}
	}

	txCopy := tx.TrimmedCopy()
	curve := elliptic.P256()

	for inID, vin := range tx.Vin {
		prevTx := prevTXs[hex.EncodeToString(vin.Txid)]
		txCopy.Vin[inID].Signature = nil
		txCopy.Vin[inID].PubKey = prevTx.Vout[vin.Vout].PubKeyHash
		txCopy.ID = txCopy.Hash()
		txCopy.Vin[inID].PubKey = nil

		r := big.Int{}
		s := big.Int{}
		sigLen := len(vin.Signature)
		r.SetBytes(vin.Signature[:(sigLen / 2)])
		s.SetBytes(vin.Signature[(sigLen / 2):])

		x := big.Int{}
		y := big.Int{}
		keyLen := len(vin.PubKey)
		x.SetBytes(vin.PubKey[:(keyLen / 2)])
		y.SetBytes(vin.PubKey[(keyLen / 2):])

		rawPubKey := ecdsa.PublicKey{curve, &x, &y}
		if ecdsa.Verify(&rawPubKey, txCopy.ID, &r, &s) == false {
			return false
		}
	}

	return true

}

// Checks if the transaction is coinbase
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		randData := make([]byte, 20)
		_, err := rand.Read(randData)
		HandleError(err)

		data = fmt.Sprintf("%x", randData)
	}

	txin := TXInput{[]byte{}, -1, nil, []byte(data)}
	txout := NewTXOutput(subsidy, to)
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{*txout}}
	tx.ID = tx.Hash()

	return &tx
}

func NewUTXOTransaction(from, to string, amount int, UTXOSet *UTXOSet) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	wallets, err := NewWallets()
	HandleError(err)
	wallet := wallets.GetWallet(from)
	pubKeyHash := HashPubKey(wallet.PublicKey)
	acc, validOutputs := UTXOSet.FindSpendableOutputs(pubKeyHash, amount)

	// If the account balance is insufficient, exit
	if acc < amount {
		log.Panic("Error: Not enough funds")
	}

	// Create inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		HandleError(err)

		for _, out := range outs {
			input := TXInput{txID, out, nil, wallet.PublicKey}
			inputs = append(inputs, input)
		}
	}

	// Create outputs
	outputs = append(outputs, *NewTXOutput(amount, to))

	// If there is change, create a change output
	if acc > amount {
		outputs = append(outputs, *NewTXOutput(acc-amount, from))
	}

	// Create transaction
	tx := Transaction{nil, inputs, outputs}
	tx.ID = tx.Hash()
	UTXOSet.Blockchain.SignTransaction(&tx, wallet.PrivateKey)

	return &tx
}
