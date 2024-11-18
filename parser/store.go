package parser

type StoreDelegate interface {
	Insert(address string, transaction Transaction)
	Get(address string) []Transaction
}

var delegate StoreDelegate = nil

func SetStoreDelegate(sd StoreDelegate) {
	if delegate != nil {
		panic("store delegate has been already initialized")
	}
	delegate = sd
}

func checkInitialized() {
	if delegate == nil {
		panic("store delegate must be initialized")
	}
}

func storeInsert(address string, transaction Transaction) {
	checkInitialized()
	delegate.Insert(address, transaction)
}

func storeGet(address string) []Transaction {
	checkInitialized()
	return delegate.Get(address)
}
