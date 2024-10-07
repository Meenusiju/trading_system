package main

import (
	"log"
	"net/http"

	"github.com/Meenusiju/chart-view/internal/app"
	"github.com/Meenusiju/chart-view/internal/binance"
	"github.com/Meenusiju/chart-view/internal/websocket"
)

func main() {
	symbols := []string{"BTCUSDT", "ETHUSDT", "PEPEUSDT"}

	binanceClient := binance.NewClient()
	hub := websocket.NewHub()

	server, err := app.NewServer(symbols, binanceClient, hub)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	go hub.Run()

	http.HandleFunc("/ws", hub.ServeWs)

	log.Println("Server starting on :8081")
	if err := server.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
