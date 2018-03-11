package websocket

import (
	"sync"
	"os"
	"os/signal"
	"net/url"
	"log"
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
	"github.com/concertos/module/common"
	"github.com/concertos/module/entity"
)

const WS_SERVER_ADDR = "localhost:8081"

func (ws *WebSocket) handleInstallContainerEvent(wsm *common.WebSocketMessage) {
	var cInfo = new(entity.ContainerInfo)
	if err := json.Unmarshal(wsm.Content, cInfo); err != nil {
		log.Println("Erro handle msg : ", err)
		wsm.Content = []byte(err.Error())
	} else {

		wsm.Content = []byte("Operator success")
	}
	res, _ := json.Marshal(&wsm)
	ws.Send <- res
}

func (ws *WebSocket) Start() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: WS_SERVER_ADDR, Path: "/ws"}
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
			log.Println(string(message))

			var wsm = new(common.WebSocketMessage)
			json.Unmarshal(message, wsm)

			switch wsm.MessageType {
			case common.P_WS_INSTALL_CONTAINER:
				ws.handleInstallContainerEvent(wsm)
			default:
				res, _ := json.Marshal(&common.WebSocketMessage{
					MessageType: common.P_WS_INSTALL_CONTAINER,
					Content:     []byte("Unknown message type"),
				})
				ws.Send <- res
			}
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
			log.Println("System interrupt")
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
