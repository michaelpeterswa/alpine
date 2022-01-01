package blockchain

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

var (
	once sync.Once
)

// AlpineBlockchain is a structure that conains a mutex lock and the blockchain itself.
type AlpineBlockchain struct {
	mu          sync.Mutex
	Circulation int64            `json:"-"`
	TxPool      *TransactionPool `json:"transaction-pool"`
	Chain       []Block          `json:"blockchain"`
}

// Singleton Pattern using sync.Once to ensure that only one AlpineBlockchain is created.
// Seeds the blockchain with a genesis block
//
// Returns: *AlpineBlockchain and error
func InitAlpine(coins int64, txp *TransactionPool) (*AlpineBlockchain, error) {
	var blockchain *AlpineBlockchain

	// Create the singleton blockchain
	once.Do(func() {
		blockchain = &AlpineBlockchain{Circulation: coins, TxPool: txp}
	})

	// Instantiate the genesis block
	genesis := Block{
		Index:        0,
		Timestamp:    time.Now(),
		Transactions: []Transaction{},
		Proof:        0,
		PreviousHash: "",
	}

	err := blockchain.AddBlock(&genesis)
	if err != nil {
		return nil, err
	}
	return blockchain, nil
}

// Gets the last block on *AlpineBlockchain ab, used to calculate previous hash
//
// Returns: *Block
func (ab *AlpineBlockchain) GetLastBlock() *Block {
	chainLen := len(ab.Chain)
	if chainLen == 0 {
		return nil
	}
	return &ab.Chain[chainLen-1]
}

// Add *Block b to *AlpineBlockchain ab
//
// Returns: error
func (ab *AlpineBlockchain) AddBlock(b *Block) error {
	ab.mu.Lock()
	defer ab.mu.Unlock()

	b.Index = int64(len(ab.Chain) + 1)
	prevHash, err := ab.GetLastBlock().Hash()
	if err != nil {
		return err
	}
	b.PreviousHash = prevHash
	b.Proof = 0

	ab.Chain = append(ab.Chain, *b)
	return nil
}

// Prints indented JSON representing the current blockchain
//
// Returns: error
func (ab *AlpineBlockchain) PrintBlockchain() error {
	res, err := json.MarshalIndent(ab, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
