import React, { useState, useEffect } from "react";
import Chart from "./components/Chart";
import SymbolSelector from "./components/SymbolSelector";
import LatestPrice from "./components/LatestPrice";
import { Candlestick } from "./types";
import "./App.css";

const App: React.FC = () => {
  const [symbol, setSymbol] = useState<string>("BTCUSDT");
  const [candlesticks, setCandlesticks] = useState<Candlestick[]>([]);
  const [latestPrice, setLatestPrice] = useState<number | null>(null);
  const [priceColor, setPriceColor] = useState<string>("black");
  const [lastCandlestickTime, setLastCandlestickTime] = useState<string>("");

  useEffect(() => {
    setLatestPrice(0);
    setPriceColor("black");
    setCandlesticks([]);

    const ws = new WebSocket("ws://localhost:8081/ws");

    ws.onmessage = (event) => {
      const candlestick: Candlestick = JSON.parse(event.data);
      if (candlestick.Symbol === symbol) {
        setCandlesticks((prev) => [...prev, candlestick]);
        setLatestPrice(candlestick.Close);
        setPriceColor(candlestick.Close > candlestick.Open ? "green" : "red");
        setLastCandlestickTime(
          new Date(candlestick.Timestamp).toUTCString().split(" ")[4]
        );
      }
    };

    return () => {
      ws.close();
    };
  }, [symbol]);

  const handleSymbolChange = (selectedSymbol: string) => {
    setSymbol(selectedSymbol);
  };

  return (
    <div className="container">
      <h1 className="header">Trading Chart View</h1>
      <SymbolSelector symbol={symbol} onSymbolChange={handleSymbolChange} />
      <LatestPrice
        price={latestPrice}
        color={priceColor}
        chartTime={lastCandlestickTime}
      />
      <Chart symbol={symbol} candlesticks={candlesticks} />
    </div>
  );
};

export default App;
