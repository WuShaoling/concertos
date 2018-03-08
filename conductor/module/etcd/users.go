package etcd

//
//import (
//	"github.com/concertos/common"
//	"encoding/json"
//)
//
//// Add a user info to etcd
//func (c *Conductor) SetUser(user *common.UserInfo) error {
//	key := common.ETCD_PREFIX_USERS_INFO + user.Id
//	value, _ := json.Marshal(*user)
//	_, err := c.MyEtcdClent.Put(key, string(value), nil)
//	return err
//}
//
//// Return all users stored in etcd
//func (c *Conductor) GetAllUser() ([]common.UserInfo, error) {
//	resp, err := c.MyEtcdClent.Get(common.ETCD_PREFIX_USERS_INFO, nil)
//	var users []common.UserInfo
//	for _, v := range resp.Kvs {
//		var user common.UserInfo
//		json.Unmarshal([]byte(v.Value), &user)
//		users = append(users, user)
//	}
//	return users, err
//}
//
//// Query a user according to id
//func (c *Conductor) GetUser(id string) (*common.UserInfo, error) {
//	resp, err := c.MyEtcdClent.Get(common.ETCD_PREFIX_USERS_INFO+id, nil)
//	if nil != err {
//		return nil, err
//	}
//	var user common.UserInfo
//	for _, v := range resp.Kvs {
//		json.Unmarshal([]byte(v.Value), &user)
//	}
//	return &user, nil
//}
//
//func (c *Conductor) DeleteUser(id string) error {
//	return nil
//}
//
//func (c *Conductor) UpdateUser(user *common.UserInfo) error {
//	return nil
//}
