package common

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"strings"
	"strconv"
	"encoding/json"
	"log"
	"time"
	"sync"
)

type MyEtcdClient struct {
	client *clientv3.Client
}

type PlayerInfoStruct struct {
	playerInfos []*PlayerInfo
}

type ContainerInfoStruct struct {
	containerInfos []*ContainerInfo
}

type UserInfoStruct struct {
	userInfos []*UserInfo
}

func (e *MyEtcdClient) Convert(ret interface{}, resp *clientv3.GetResponse) {
	switch data := ret.(type) {
	case PlayerInfoStruct:
		for _, v := range resp.Kvs {
			var info PlayerInfo
			json.Unmarshal([]byte(v.Value), &info)
			data.playerInfos = append(data.playerInfos, &info)
		}
		break
	case ContainerInfoStruct:
		for _, v := range resp.Kvs {
			var info ContainerInfo
			json.Unmarshal([]byte(v.Value), &info)
			data.containerInfos = append(data.containerInfos, &info)
		}
		break
	case UserInfoStruct:
		for _, v := range resp.Kvs {
			var info UserInfo
			json.Unmarshal([]byte(v.Value), &info)
			data.userInfos = append(data.userInfos, &info)
		}
		break
	default:
		log.Println("Unknown type")
		break
	}
}

func (e *MyEtcdClient) Put(key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := e.client.Put(ctx, key, val)
	cancel()
	return resp, err
}

func (e *MyEtcdClient) Get(key string, convertType int8, opts ...clientv3.OpOption) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := e.client.Get(ctx, key, opts ...)
	cancel()

	var result interface{}
	switch convertType {
	case 1:
		result = new(PlayerInfo)
		e.Convert(&result, resp)
		break
	case 2:
		result = new(ContainerInfo)
		e.Convert(&result, resp)
		break
	case 3:
		result = new(UserInfo)
		e.Convert(&result, resp)
		break
	}

	return result, err
}

func (e *MyEtcdClient) Delete(key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := e.client.Delete(ctx, key, opts ...)
	cancel()
	return resp, err
}

func (e *MyEtcdClient) GetErrorType(err error) int {
	strs := strings.Split(err.Error(), ":")
	if len(strs) <= 0 {
		return -1
	}
	code, _ := strconv.Atoi(strings.TrimSpace(strs[0]))
	return code
}

// clientv3.Client only one
var once2 sync.Once
var clientV3 *clientv3.Client

func GetClientV3() (*clientv3.Client) {
	once2.Do(func() {
		c, err := clientv3.New(clientv3.Config{
			Endpoints:   GetEtcdPoints(),
			DialTimeout: 2 * time.Second,
		})
		if err != nil {
			log.Fatal("Error: new etcd client error:", err)
		}
		clientV3 = c
		log.Println("+++++++++++++++++++++++++++++++===")
	})

	return clientV3
}


// MyEtcdClient only one
var once sync.Once
var etcdClient *MyEtcdClient

func GetMyEtcdClient() *MyEtcdClient {
	once.Do(func() {
		c := &MyEtcdClient{
			client: GetClientV3(),
		}
		etcdClient = c
		log.Println("------------------------------")
	})
	return etcdClient
}

//
//type Entity interface {
//	Convert(resp *clientv3.PutResponse) []Entity
//}
//
//func (p *PlayerInfo) Convert(resp *clientv3.GetResponse) []*common.PlayerInfo {
//	var infos []*common.PlayerInfo
//	for _, v := range resp.Kvs {
//		var info common.PlayerInfo
//		json.Unmarshal([]byte(v.Value), &info)
//		infos = append(infos, &info)
//	}
//	return infos
//}
//
//func (p *UserInfo) Convert(resp *clientv3.GetResponse) []*common.UserInfo {
//	var infos []*common.UserInfo
//	for _, v := range resp.Kvs {
//		var info common.UserInfo
//		json.Unmarshal([]byte(v.Value), &info)
//		infos = append(infos, &info)
//	}
//	return infos
//}
//
//func (p *ContainerInfo) Convert(resp *clientv3.GetResponse) []*common.ContainerInfo {
//	var infos []*common.ContainerInfo
//	for _, v := range resp.Kvs {
//		var info common.ContainerInfo
//		json.Unmarshal([]byte(v.Value), &info)
//		infos = append(infos, &info)
//	}
//	return infos
//}
