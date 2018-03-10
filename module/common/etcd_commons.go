package common

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"strings"
	"strconv"
	"log"
	"time"
	"sync"
	"encoding/json"
)

type MyEtcdClient struct {
}

func (e *MyEtcdClient) ConvertToUserInfo(resp *clientv3.GetResponse) *[]UserInfo {
	var arr []UserInfo
	for _, v := range resp.Kvs {
		var info UserInfo
		json.Unmarshal([]byte(v.Value), &info)
		arr = append(arr, info)
	}
	return &arr
}

func (e *MyEtcdClient) ConvertToPlayerInfo(resp *clientv3.GetResponse) *[]PlayerInfo {
	var arr []PlayerInfo
	for _, v := range resp.Kvs {
		var info PlayerInfo
		json.Unmarshal([]byte(v.Value), &info)
		arr = append(arr, info)
	}
	return &arr
}

func (e *MyEtcdClient) ConvertToContainerInfo(resp *clientv3.GetResponse) *[]ContainerInfo {
	var arr []ContainerInfo
	for _, v := range resp.Kvs {
		var info ContainerInfo
		json.Unmarshal([]byte(v.Value), &info)
		arr = append(arr, info)
	}
	return &arr
}

func (e *MyEtcdClient) Put(key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := ectdClientV3.Put(ctx, key, val)
	cancel()
	return resp, err
}

func (e *MyEtcdClient) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := ectdClientV3.Get(ctx, key, opts ...)
	cancel()
	return resp, err
}

func (e *MyEtcdClient) Delete(key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := ectdClientV3.Delete(ctx, key, opts ...)
	cancel()
	return resp, err
}

func (e *MyEtcdClient) CheckExist(id string) (bool, error) {
	if resp, err := e.Get(id); err != nil {
		return false, err
	} else {
		if len(resp.Kvs) > 0 {
			return true, nil
		}
	}
	return false, nil
}

func (e *MyEtcdClient) GetErrorType(err error) int {
	strs := strings.Split(err.Error(), ":")
	if len(strs) <= 0 {
		return -1
	}
	code, _ := strconv.Atoi(strings.TrimSpace(strs[0]))
	return code
}

var once2 sync.Once
var ectdClientV3 *clientv3.Client

func (e *MyEtcdClient) GetClientV3() (*clientv3.Client) {
	once2.Do(func() {
		var err error
		ectdClientV3, err = clientv3.New(clientv3.Config{
			Endpoints:   GetEtcdPoints(),
			DialTimeout: 2 * time.Second,
		})
		if err != nil {
			log.Fatal("Error: new common client error:", err)
		}
		//defer ectdClientV3.Close()
	})
	return ectdClientV3
}

var once sync.Once
var myEtcdClient *MyEtcdClient

func GetMyEtcdClient() *MyEtcdClient {
	once.Do(func() {
		myEtcdClient = &MyEtcdClient{}
		myEtcdClient.GetClientV3()
	})
	return myEtcdClient
}
