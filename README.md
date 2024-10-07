# Trading system (Go lang + React)

This is a OHLC candlestick chart service for our trading system. The chart displays candlestick patterns for `BTCUSD` `ETHUSDT` `PEPEUSDT` in a `1minute` time frame. The api endpoint is from [binance websocket](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams#aggregate-trade-streams)

## Tech stack

Frontend: `React` `Typescript` `css` `Apex chart`

Backend: `Go lang` `web socket`

## Prerequisites

- Go 1.x or higher
- Docker (for containerization)
- Node js

## Run the app

## `docker-compose up --build`

## Abort/stop container

## `docker-compose down`

## Folder Structure

## Backend

- `cmd/bff`: Main application entry point
- `internal/app`: Core application logic separated as modules
- `Dockerfile`: Docker configuration for containerization
- `vendor/`: Vendored dependencies

## Frontend

- `src`: Application entry point
- `src/components`: connecting chart service using backend api

Attached screenshot of chart displaying BTCUSDT in 1m time frame

<img width="1679" alt="screen shot for btcusd" src="https://github.com/user-attachments/assets/85053e9f-fbc4-42c8-9c9f-84cb838ebee0">




  
