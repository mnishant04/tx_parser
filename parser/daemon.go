package parser

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func hexToInt(s string) int64 {
	parsed, _ := strconv.ParseInt(strings.Replace(s, "0x", "", -1), 16, 32)
	return parsed
}

func intToHex(num int64) string {
	return fmt.Sprintf("0x%x", num)
}

type daemon struct {
	lock        sync.RWMutex
	latestBlock int64
	client      *rpcClient
	subscribers map[string]bool
}

func newDaemon(url string) *daemon {
	return &daemon{sync.RWMutex{}, -1, newRpcClient(url), make(map[string]bool)}
}

func (daemon *daemon) subscribe(address string) bool {
	daemon.lock.Lock()
	defer daemon.lock.Unlock()
	if daemon.subscribers[address] {
		return false
	}
	daemon.subscribers[address] = true
	return true
}

func (daemon *daemon) lastParsedBlock() int64 {
	daemon.lock.RLock()
	defer daemon.lock.RUnlock()
	return daemon.latestBlock
}

func (daemon *daemon) run() {
	for {
		daemon.tick()
		time.Sleep(1 * time.Second)
	}
}

func (daemon *daemon) tick() {
	blockNumberResp, err := daemon.client.getRecentBlockNumber()
	if err != nil {
		panic(err)
	}
	blockNumber := hexToInt(blockNumberResp.Result)

	daemon.lock.Lock()
	defer daemon.lock.Unlock()

	if daemon.latestBlock == -1 {
		daemon.parseBlockByBlockNum(blockNumber)
	} else {
		for blockNum := daemon.latestBlock; blockNum <= blockNumber; blockNum++ {
			daemon.parseBlockByBlockNum(blockNum)
		}
	}

	daemon.latestBlock = blockNumber
}

func (daemon *daemon) parseBlockByBlockNum(block int64) {
	blockByNumberResp, err := daemon.client.getBlockByNumber(intToHex(block))
	if err != nil {
		panic(err)
	}

	for _, t := range blockByNumberResp.Result.Transactions {
		if daemon.subscribers[t.To] {
			storeInsert(t.To, t)
		}

		if daemon.subscribers[t.From] {
			storeInsert(t.From, t)
		}
	}
}
