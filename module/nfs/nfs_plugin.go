package nfs

import (
	"sync"
	"github.com/concertos/module/common"
	"github.com/concertos/module/util"
	"os"
	"log"
)

func (nfs *NFSApi) Start() {
	nfs.UMount()
	nfs.Mount()
}

func (nfs *NFSApi) UMount() {
	if _, err := util.ExecShell("umount " + common.NFS_MOUNT_LOCAL_PATH); nil != err {
		log.Println(err)
	}
}

func (nfs *NFSApi) Mount() {
	if _, err := os.Stat(common.NFS_MOUNT_LOCAL_PATH); nil != err {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(common.NFS_MOUNT_LOCAL_PATH, 0777); nil != err {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
	cmd := "sudo mount -v " + common.GetNFSMountRemoteAddr() + " " + common.NFS_MOUNT_LOCAL_PATH
	if _, err := util.ExecShell(cmd); nil != err {
		log.Fatal(err)
	}
	log.Println("mount ", common.NFS_MOUNT_LOCAL_PATH, " to nfs server ", common.GetNFSMountRemoteAddr())
}

func (nfs *NFSApi) CreateContainerRootPath(cname string) error {

	_, err := os.Stat(common.NFS_MOUNT_LOCAL_PATH + cname);
	if nil == err { //存在
		return nil
	}

	if os.IsNotExist(err) { //if not exist, create it
		if err := os.MkdirAll(common.NFS_MOUNT_LOCAL_PATH+cname, 0777); nil != err {
			log.Println("create path: ", cname, " error: ", err)
			return err
		} else {
			return nil
		}
	}

	return err
}

var NfsApi *NFSApi
var once sync.Once

type NFSApi struct {
}

func GetNFSApi() *NFSApi {

	once.Do(func() {
		NfsApi = new(NFSApi)
	})

	return NfsApi
}
