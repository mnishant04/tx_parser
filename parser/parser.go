package parser

type Transaction struct {
	Type                 string       `json:"type"`
	BlockHash            string       `json:"blockHash"`
	BlockNumber          string       `json:"blockNumber"`
	From                 string       `json:"from"`
	Gas                  string       `json:"gas"`
	Hash                 string       `json:"hash"`
	Input                string       `json:"input"`
	Nonce                string       `json:"nonce"`
	To                   string       `json:"to"`
	TransactionIndex     string       `json:"transactionIndex"`
	Value                string       `json:"value"`
	V                    string       `json:"v"`
	R                    string       `json:"r"`
	S                    string       `json:"s"`
	GasPrice             string       `json:"gasPrice"`
	MaxFeePerGas         string       `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string       `json:"maxPriorityFeePerGas,omitempty"`
	ChainId              string       `json:"chainId,omitempty"`
	AccessList           []AccessList `json:"accessList,omitempty"`
}

type AccessList struct {
	Address     string   `json:"address"`
	StorageKeys []string `json:"storageKeys"`
}

type Parser interface {

	GetCurrentBlock() int

	Subscribe(address string) bool

	GetTransactions(address string) []Transaction
}

type ethParser struct {
	d *daemon
}

func New(url string) Parser {
	p := &ethParser{newDaemon(url)}
	go p.d.run()
	return p
}

func (parser *ethParser) GetCurrentBlock() int {
	return int(parser.d.lastParsedBlock())
}

func (parser *ethParser) Subscribe(address string) bool {
	return parser.d.subscribe(address)
}

func (parser *ethParser) GetTransactions(address string) []Transaction {
	return storeGet(address)
}
