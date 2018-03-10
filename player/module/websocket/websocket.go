package websocket

import (
	"sync"
	"flag"
	"os"
	"os/signal"
	"net/url"
	"time"
	"log"
	"github.com/gorilla/websocket"
)

func (ws *WebSocket) Start() {

	log.Println("Start websocket module")

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:8081", Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 2)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte("hello"))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("System interrupt, websocket module stop!")
			// To cleanly close a connection, a client should send a close
			// frame and wait for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			c.Close()
			return
		}
	}
}

type WebSocket struct {
}

var once sync.Once
var webSocket *WebSocket

func GetWetSocket() *WebSocket {
	once.Do(func() {
		webSocket = &WebSocket{}
	})
	return webSocket
}
