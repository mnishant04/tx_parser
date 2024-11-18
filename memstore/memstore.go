package memstore

import (
	"ethscanner/parser"
	"sync"
)

func init() {
	sd := &inMemoryStore{sync.RWMutex{}, make(map[string][]parser.Transaction)}
	parser.SetStoreDelegate(sd)
}

type inMemoryStore struct {
	lock sync.RWMutex
	data map[string][]parser.Transaction
}

func NewInMemoryStorage() Storage {
	return &inMemoryStore{
		lock: sync.RWMutex{},
		data: make(map[string][]parser.Transaction),
	}
}

func (s *inMemoryStore) Insert(address string, transaction parser.Transaction) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data[address] = append(s.data[address], transaction)
}

func (s *inMemoryStore) Get(address string) []parser.Transaction {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.data[address]
}
