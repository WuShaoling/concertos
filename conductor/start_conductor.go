package conductor

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"time"
	"github.com/concertos/conductor/module/etcd"
	"github.com/concertos/conductor/module/api"
	"flag"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		time.Sleep(1 * time.Second)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func playerWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		time.Sleep(1 * time.Second)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func StartConductor() {
	c := etcd.GetConductor()
	go c.Watch()

	api.Start()

	var addr = flag.String("addr", "0.0.0.0:8080", "service address")
	flag.Parse()
	http.HandleFunc("/consh", ws)
	http.HandleFunc("/player", playerWs)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
