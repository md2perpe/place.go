package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func Connect() (c *websocket.Conn, err error) {
	dialer := &websocket.Dialer{}
	dialer.EnableCompression = false
	place_url := "YOUR WSS URL HERE"
	c, _, err = dialer.Dial(place_url, nil)
	if err != nil {
		return c, err
	}
	return c, err
}

type Place struct {
	Type         string `json:"type"`
	PlacePayload `json:"payload"`
}
type PlacePayload struct {
	Y      int       `json:"y"`
	X      int       `json:"x"`
	Color  int       `json:"color"`
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
}

func main() {
	c, _ := Connect()
	for {
		_, packet, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			continue
		}
		
		var p Place
		err := json.Unmarshal(packet, &p)
		if err != nil {
			panic(err)
		}
		
		p.Date = time.Now()
		pStr, _ := json.Marshal(p)
		fmt.Println(string(pStr))
	}
}
