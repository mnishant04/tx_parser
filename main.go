package main

import (
	"context"
	"ethscanner/controller"
	_ "ethscanner/memstore"
	"ethscanner/parser"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const URL = "https://ethereum-rpc.publicnode.com"
const PORT = "8080"
	
func main() {
	p := parser.New(URL)
	controller := controller.NewEthHandler(p)
	handler := http.NewServeMux()
	handler.HandleFunc("GET /api/v1/currentblock", controller.CurrentBlock)
	handler.HandleFunc("GET /api/v1/getalltransactions", controller.GetAllTransactions)
	handler.HandleFunc("POST /api/v1/subscribe", controller.Subscribe)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", PORT),
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	go func() {
		log.Println("Starting server on port", PORT)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigchan
	log.Printf("Received signal: %s", sig)
	shutdownServer(server)
}

func shutdownServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server  Shutdown Failed: %+v", err)
	}
	log.Println("Server Exited  Gracefully")
}

