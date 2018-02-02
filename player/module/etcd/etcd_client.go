package etcd

import (
	"github.com/coreos/etcd/client"
	"github.com/shortid"
	"github.com/shirou/gopsutil/mem"
	"time"
	"sync"
	"log"
	"runtime"
	"os"
	"github.com/concertos/player/util"
	"encoding/json"
	"context"
	"github.com/concertos/config"
)

const ETCDENDPOINT = "http://127.0.0.1:2379"

var etcdClient *Client
var once sync.Once

type Client struct {
	KeysAPI client.KeysAPI
	Info    ClientInfo
}

type ClientInfo struct {
	Id       string
	Ips      []string
	Hostname string
	Memory   uint64
	Cpu      int
}

func (c *Client) HeartBeat() {
	for {
		key := "/players/" + c.Info.Hostname
		value, _ := json.Marshal(&c.Info)

		_, err := c.KeysAPI.Set(context.Background(), key, string(value), &client.SetOptions{
			TTL: time.Second * config.TTL,
		})
		if err != nil {
			log.Println("Error update EtcdClientInfo:", err)
		}
		//log.Fatal(key)
		time.Sleep(time.Second * config.HEARTBEAT)
	}
}

func NewEtcdClient() *Client {
	cfg := client.Config{
		Endpoints:               []string{ETCDENDPOINT},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}

	hostname, _ := os.Hostname()
	memory, _ := mem.VirtualMemory()
	player := &Client{
		Info: ClientInfo{
			Id:       shortid.MustGenerate(),
			Ips:      util.GetIps(),
			Hostname: hostname,
			Memory:   memory.Total,
			Cpu:      runtime.NumCPU(),
		},
		KeysAPI: client.NewKeysAPI(etcdClient),
	}
	return player
}

func GetEtcdClient() *Client {
	once.Do(func() {
		etcdClient = NewEtcdClient()
	})
	return etcdClient
}
