package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Meenusiju/chart-view/internal/app"
	"github.com/Meenusiju/chart-view/internal/binance"
	"github.com/Meenusiju/chart-view/internal/websocket"
)

func TestMainSetup(t *testing.T) {

	symbols := []string{"BTCUSDT", "ETHUSDT", "PEPEUSDT"}

	// Create components
	binanceClient := binance.NewClient()
	hub := websocket.NewHub()

	// Create server
	server, err := app.NewServer(symbols, binanceClient, hub)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}

	// Check if server is not nil
	if server == nil {
		t.Fatal("Server should not be nil")
	}

	// Test WebSocket handler
	req, err := http.NewRequest("GET", "/ws", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hub.ServeWs)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

}
