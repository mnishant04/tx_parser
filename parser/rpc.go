package parser

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type rpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type blockNumberResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      int    `json:"id"`
}

type blockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Number           string        `json:"number"`
		Hash             string        `json:"hash"`
		ParentHash       string        `json:"parentHash"`
		Sha3Uncles       string        `json:"sha3Uncles"`
		LogsBloom        string        `json:"logsBloom"`
		TransactionsRoot string        `json:"transactionsRoot"`
		StateRoot        string        `json:"stateRoot"`
		ReceiptsRoot     string        `json:"receiptsRoot"`
		Miner            string        `json:"miner"`
		Difficulty       string        `json:"difficulty"`
		TotalDifficulty  string        `json:"totalDifficulty"`
		ExtraData        string        `json:"extraData"`
		Size             string        `json:"size"`
		GasLimit         string        `json:"gasLimit"`
		GasUsed          string        `json:"gasUsed"`
		Timestamp        string        `json:"timestamp"`
		Transactions     []Transaction `json:"transactions"`
		Uncles           []interface{} `json:"uncles"`
		BaseFeePerGas    string        `json:"baseFeePerGas"`
		Nonce            string        `json:"nonce"`
		MixHash          string        `json:"mixHash"`
	} `json:"result"`
	Id int `json:"id"`
}

type rpcClient struct {
	url string
	seq int
}

func newRpcClient(url string) *rpcClient {
	return &rpcClient{url, 0}
}

func (client *rpcClient) doRequest(method string, params []interface{}) (*http.Response, error) {
	defer func() { client.seq++ }()
	req := rpcRequest{"2.0", method, params, client.seq}
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(client.url, "application/json", strings.NewReader(string(marshal)))
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func closeResponse(resp *http.Response) {
	func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
}

// GetRecentBlockNumber curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"seq":83}'
func (client *rpcClient) getRecentBlockNumber() (*blockNumberResponse, error) {
	resp, err := client.doRequest("eth_blockNumber", []interface{}{})

	if err != nil {
		return nil, err
	}

	defer closeResponse(resp)

	var ans = new(blockNumberResponse)
	if err := json.NewDecoder(resp.Body).Decode(ans); err != nil {
		return nil, err
	}
	return ans, nil
}

// GetBlockByNumber curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x1b4", true],"seq":1}'
func (client *rpcClient) getBlockByNumber(num string) (*blockResponse, error) {
	resp, err := client.doRequest("eth_getBlockByNumber", []interface{}{
		num, true,
	})

	if err != nil {
		return nil, err
	}

	defer closeResponse(resp)

	var ans = new(blockResponse)
	if err := json.NewDecoder(resp.Body).Decode(ans); err != nil {
		return nil, err
	}
	return ans, nil
}
