package websocket

import (
	"log"
	"net/http"
	"sync"
	"github.com/concertos/module/common"
)

type HttpWaitMsg struct {
	Id      string
	Channel chan string
}

func (ws *WebSocket) run() {
	for {
		select {
		case httpWaitMsg := <-ws.AddWait:
			ws.HttpWait[httpWaitMsg.Id] = httpWaitMsg.Channel
		case client := <-ws.register:
			ws.Clients[client] = true
		case client := <-ws.unregister:
			if _, ok := ws.Clients[client]; ok {
				delete(ws.Clients, client)
				close(client.send)
			}
		case message := <-ws.broadcast:
			log.Println("broadcast", message)
			for client := range ws.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(ws.Clients, client)
				}
			}
		case clientid := <-ws.WriteTo:
			log.Println("ws.WriteTo id : ", clientid)
			message := <-ws.WriteTo
			log.Println("ws.WriteTo message : ", message)
			for client, _ := range ws.Clients {
				log.Println("client.Id: ", client.Id)
				if client.Id == string(clientid) {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(ws.Clients, client)
					}
					break
					log.Println("Find player id")
				}
			}
		}
	}
}

func (ws *WebSocket) Start() {
	log.Println("websocket server listen on: ", common.GetWSServerAddr())
	go ws.run()
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		serveWs(ws, &writer, request)
	})
	log.Fatal(http.ListenAndServe(common.GetWSServerAddr(), nil))
}

var webSocket *WebSocket
var once sync.Once

type WebSocket struct {
	Clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	WriteTo    chan []byte

	AddWait  chan HttpWaitMsg
	HttpWait map[string]chan string
}

func GetWebSocket() *WebSocket {
	once.Do(func() {
		webSocket = &WebSocket{
			AddWait:  make(chan HttpWaitMsg),
			HttpWait: make(map[string]chan string),

			broadcast:  make(chan []byte),
			WriteTo:    make(chan []byte),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			Clients:    make(map[*Client]bool),
		}
	})
	return webSocket
}
