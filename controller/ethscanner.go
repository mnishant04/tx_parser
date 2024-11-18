package controller

import (
	"ethscanner/common"
	"ethscanner/parser"
	"fmt"
	"log"
	"net/http"
)

type ethHandler struct {
	p parser.Parser
}

func NewEthHandler(p parser.Parser) Controller {
	return ethHandler{
		p: p,
	}
}

func (e ethHandler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	transactions := e.p.GetTransactions(address)

	if len(transactions) == 0 {
		log.Println("Transactions Not found")
		common.SendResponse(w, nil, "Transaction Not Found", http.StatusNotFound)
		return
	}
	common.SendResponse(w, transactions, "Success", http.StatusOK)

}

func (e ethHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	var subs SubscribeResponse
	if !e.p.Subscribe(address) {
		lmsg := fmt.Sprintf("Address %s is already subscribed", address)
		log.Println(lmsg)
		subs.SubscriptionStatus = lmsg
		common.SendResponse(w, subs, "Success", http.StatusOK)
	}
	lmsg := fmt.Sprintf("Address %s subscribed successfully", address)
	subs.SubscriptionStatus = lmsg
	common.SendResponse(w, subs, "Success", http.StatusOK)
}

func (e ethHandler) CurrentBlock(w http.ResponseWriter, r *http.Request) {
	currentBlock := e.p.GetCurrentBlock()
	log.Printf("Last processed block: %d", currentBlock)
	common.SendResponse(w, GetCurrentBlockResponse{currentBlock}, "Success", http.StatusOK)
}
