package controller

import "net/http"

type Controller interface {
	CurrentBlock(w http.ResponseWriter, r *http.Request)
	GetAllTransactions(w http.ResponseWriter, r *http.Request)
	Subscribe(w http.ResponseWriter, r *http.Request)
}
