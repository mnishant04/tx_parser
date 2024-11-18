package memstore

import "ethscanner/parser"

type Storage interface {
	Insert(address string, transaction parser.Transaction)
	Get(address string) []parser.Transaction
}

func Insert(s Storage, address string, transaction parser.Transaction) {
	s.Insert(address, transaction)
}

func Get(s Storage, address string) {
	s.Get(address)
}
