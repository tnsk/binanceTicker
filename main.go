package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Stream struct {
	Stream string `json:"stream"`
	Data   Data   `json:"data"`
}

type Data struct {
	CoinType string `json:"s"`
	Price    string `json:"p"`
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com/stream", nil)
	if err != nil {
		log.Fatalln("[!] Not Connected:", err)
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte("{\"method\":\"SUBSCRIBE\",\"params\":[\"!miniTicker@arr@3000ms\",\"usdttry@aggTrade\",\"btcusdt@aggTrade\",\"ltcusdt@aggTrade\",\"btcusdt@aggTrad\",\"usdttry@depth\",\"usdttry@kline_1d\"],\"id\":1}"))
	if err != nil {
		log.Fatalln("[!] Message Not Send:", err)
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatalln("[!] Message Not Read:", err)
		}
		data := Stream{}
		err = json.Unmarshal(message, &data)
		if err != nil {
			continue
		}

		if data.Stream == "ltcusdt@aggTrade" {
			log.Printf("LTC : %+v \n", data)
		}

		if data.Stream == "usdttry@aggTrade" {
			log.Printf("USDT : %+v \n", data)
		}

		if data.Stream == "btcusdt@aggTrade" {
			log.Printf("BTC : %+v \n", data)
		}
	}
}
