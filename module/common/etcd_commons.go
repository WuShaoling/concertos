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
	"github.com/concertos/module/entity"
)

type MyEtcdClient struct {
}

func (e *MyEtcdClient) ConvertToUserInfo(resp *clientv3.GetResponse) *[]entity.UserInfo {
	var arr []entity.UserInfo
	for _, v := range resp.Kvs {
		var info entity.UserInfo
		json.Unmarshal([]byte(v.Value), &info)
		arr = append(arr, info)
	}
	return &arr
}

func (e *MyEtcdClient) ConvertToPlayerInfo(resp *clientv3.GetResponse) *[]entity.PlayerInfo {
	var arr []entity.PlayerInfo
	for _, v := range resp.Kvs {
		var info entity.PlayerInfo
		json.Unmarshal([]byte(v.Value), &info)
		arr = append(arr, info)
	}
	return &arr
}

func (e *MyEtcdClient) ConvertToContainerInfo(resp *clientv3.GetResponse) *[]entity.ContainerInfo {
	var arr []entity.ContainerInfo
	for _, v := range resp.Kvs {
		var info entity.ContainerInfo
		json.Unmarshal([]byte(v.Value), &info)
		arr = append(arr, info)
	}
	return &arr
}

func (e *MyEtcdClient) Put(key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := ectdClientV3.Put(ctx, key, val, opts...)
	cancel()
	if nil != err {
		log.Println("Error Put : ", err)
	}
	return resp
}

func (e *MyEtcdClient) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := ectdClientV3.Get(ctx, key, opts ...)
	cancel()
	if nil != err {
		log.Println("Error Get : ", err)
	}
	return resp
}

func (e *MyEtcdClient) Delete(key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_TIMEOUT)
	resp, err := ectdClientV3.Delete(ctx, key, opts ...)
	cancel()
	if nil != err {
		log.Println("Error Delete : ", err)
	}
	return resp
}

func (e *MyEtcdClient) CheckExist(id string) (bool) {
	if e.Get(id).Count > 0 {
		return true
	}
	return false
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
