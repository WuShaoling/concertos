// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"log"
	"net/http"
	"sync"
)

func (ws *WebSocket) run() {
	for {
		select {
		case client := <-ws.register:
			ws.Clients[client] = true
		case client := <-ws.unregister:
			if _, ok := ws.Clients[client]; ok {
				delete(ws.Clients, client)
				close(client.send)
			}
		case message := <-ws.broadcast:
			log.Println(message)
			for client := range ws.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(ws.Clients, client)
				}
			}
		case clientid := <-ws.WriteTo:
			message := <-ws.WriteTo
			for client := range ws.Clients {
				if client.Id == string(clientid) {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(ws.Clients, client)
					}
					break
				}
			}
		}
	}
}

func (ws *WebSocket) Start() {
	log.Println("Start websocket server, listen on localhost:8081")
	go ws.run()
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		serveWs(ws, &writer, request)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

var webSocket *WebSocket
var once sync.Once

type WebSocket struct {
	Clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	WriteTo    chan []byte
}

func GetWebSocket() *WebSocket {
	once.Do(func() {
		webSocket = &WebSocket{
			broadcast:  make(chan []byte),
			WriteTo:    make(chan []byte),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			Clients:    make(map[*Client]bool),
		}
	})
	return webSocket
}
