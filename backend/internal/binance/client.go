package binance

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

type TickData struct {
	Symbol    string
	Price     float64
	Timestamp int64
}

type Candlestick struct {
	Symbol    string  `json:"Symbol"`
	Open      float64 `json:"Open"`
	High      float64 `json:"High"`
	Low       float64 `json:"Low"`
	Close     float64 `json:"Close"`
	Timestamp int64   `json:"Timestamp"`
}

func NewClient() *Client {
	return &Client{}
}

const BINANCE_ENDPOINT = "wss://stream.binance.com:9443/ws/"

func (c *Client) Connect(symbols []string) error {
	streams := make([]string, len(symbols))
	for i, symbol := range symbols {
		streams[i] = strings.ToLower(symbol) + "@aggTrade"
	}

	url := BINANCE_ENDPOINT + strings.Join(streams, "/")

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	c.conn = conn
	return nil
}

func (c *Client) ReadTicks(tickChan chan<- TickData) {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading from WebSocket: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		var data struct {
			Symbol    string `json:"s"`
			Price     string `json:"p"`
			Timestamp int64  `json:"T"`
		}

		if err := json.Unmarshal(message, &data); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		price, err := strconv.ParseFloat(data.Price, 64)
		if err != nil {
			log.Printf("Error parsing price: %v", err)
			continue
		}

		tickChan <- TickData{
			Symbol:    data.Symbol,
			Price:     price,
			Timestamp: data.Timestamp,
		}
	}
}

func (c *Client) AggregateTicks(tickChan <-chan TickData, candlestickChan chan<- Candlestick) {
	candlesticks := make(map[string]*Candlestick)

	for tick := range tickChan {
		minute := tick.Timestamp / 60000 * 60000
		key := fmt.Sprintf("%s-%d", tick.Symbol, minute)

		if candlestick, exists := candlesticks[key]; exists {
			candlestick.High = math.Max(candlestick.High, tick.Price)
			candlestick.Low = math.Min(candlestick.Low, tick.Price)
			candlestick.Close = tick.Price
		} else {
			candlesticks[key] = &Candlestick{
				Symbol:    tick.Symbol,
				Open:      tick.Price,
				High:      tick.Price,
				Low:       tick.Price,
				Close:     tick.Price,
				Timestamp: minute,
			}
		}

		now := time.Now().Unix() * 1000
		for k, v := range candlesticks {
			if v.Timestamp+60000 <= now {
				candlestickChan <- *v
				delete(candlesticks, k)
			}
		}
	}
}
