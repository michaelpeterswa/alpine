package blockchain

import (
	"container/list"
	"sync"
)

type TransactionPool struct {
	mu     sync.Mutex
	txPool *list.List
}

type PendingTransactions struct {
	Ptx []Transaction `json:"pending-transactions"`
}

func NewTransactionPool() *TransactionPool {
	return &TransactionPool{txPool: list.New()}
}

func (tp *TransactionPool) Push(t Transaction) {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	_ = tp.txPool.PushBack(t)
}

func (tp *TransactionPool) Pop() Transaction {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	return tp.txPool.Remove(tp.txPool.Front()).(Transaction)
}

func (tp *TransactionPool) List() PendingTransactions {
	pending := PendingTransactions{}

	for e := tp.txPool.Front(); e != nil; e = e.Next() {
		pending.Ptx = append(pending.Ptx, e.Value.(Transaction))
	}

	return pending
}

func (tp *TransactionPool) Len() int {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	return tp.txPool.Len()
}
