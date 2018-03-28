package websocket

import (
	"sync"
	"os"
	"os/signal"
	"net/url"
	"log"
	"github.com/gorilla/websocket"
	"time"
	"github.com/concertos/module/common"
)

func (ws *WebSocket) Start() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: common.WS_SERVER_ADDR, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	ws.conn = c
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer ws.conn.Close()

	done := make(chan struct{})

	go func() {
		defer ws.conn.Close()
		defer close(done)
		for {
			_, message, err := ws.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			ws.HandleMsg(message)
		}
	}()

	for {
		select {
		case msg := <-ws.Send:
			log.Println("Write msg", string(msg))
			err := ws.conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("System interrupt, websocket disconnected")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			ws.conn.Close()
			return
		}
	}
}

var once sync.Once
var webSocket *WebSocket

type WebSocket struct {
	conn *websocket.Conn
	Send chan []byte
}

func GetWebSocket() *WebSocket {
	once.Do(func() {
		webSocket = &WebSocket{
			Send: make(chan []byte, 1024),
		}
	})
	return webSocket
}
