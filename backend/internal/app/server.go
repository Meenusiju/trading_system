package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Meenusiju/chart-view/internal/binance"
	"github.com/Meenusiju/chart-view/internal/websocket"
)

type Server struct {
	symbols       []string
	binanceClient *binance.Client
	hub           *websocket.Hub
}

func NewServer(symbols []string, binanceClient *binance.Client, hub *websocket.Hub) (*Server, error) {
	return &Server{
		symbols:       symbols,
		binanceClient: binanceClient,
		hub:           hub,
	}, nil
}

func (s *Server) Run() error {
	if err := s.binanceClient.Connect(s.symbols); err != nil {
		return err
	}

	tickChan := make(chan binance.TickData)
	candlestickChan := make(chan binance.Candlestick)

	go s.binanceClient.ReadTicks(tickChan)
	go s.binanceClient.AggregateTicks(tickChan, candlestickChan)
	go s.broadcastCandlesticks(candlestickChan)

	return http.ListenAndServe(":8081", nil)
}

func (s *Server) broadcastCandlesticks(candlestickChan <-chan binance.Candlestick) {
	for candlestick := range candlestickChan {
		data, err := json.Marshal(candlestick)
		if err != nil {
			log.Println("Error marshalling candlestick:", err)
			continue
		}
		s.hub.Broadcast(data)
	}
}
