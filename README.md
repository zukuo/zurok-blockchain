# The Zurok Blockchain

## Instructions for Running

### NODE 3000
Create wallet and blockchain
```bash
go run main.go createwallet
go run main.go createblockchain -address CENTRAL_NODE
cp blockchain_3000.db blockchain_genesis.db
```

### NODE 3001
Run 3 times:
```bash
go run main.go createwallet
```

### NODE 3000
```bash
go run main.go -from CENTRAL_NODE -to WALLET_1 -amount 10 -mine
go run main.go -from CENTRAL_NODE -to WALLET_2 -amount 10 -mine
go run main.go startnode
```

### NODE 3001
```bash
cp blockchain_genesis.db blockchain_3001.db
go run main.go startnode
```
Now all peers have been updated.

### NODE 3002
```bash
cp blockchain_genesis.db blockchain_3002.db
go run main.go createwallet
go run main.go startnode -miner MINER_WALLET
```

### NODE 3002
Should start mining a new block.

### NODE 3001
```bash
go run main.go startnode
```
Downloads new blocks
