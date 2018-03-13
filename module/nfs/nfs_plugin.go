package nfs

import (
	"sync"
	"github.com/concertos/module/common"
	"github.com/concertos/player/util"
	"os"
	"log"
)

func (nfs *NFSApi) UMount() {
	if _, err := util.ExecShell("umount " + common.GetNFSServerAddress() + ":" + common.NFS_MOUNT_PATH); nil != err {
		log.Println(err)
	}
}

func (nfs *NFSApi) Mount() {
	if _, err := os.Stat(common.NFS_MOUNT_PATH); nil != err {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(common.NFS_MOUNT_PATH, 0777); nil != err {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
	cmd := "sudo mount -t nfs " + common.GetNFSServerAddress() + ":" + common.NFS_MOUNT_PATH + " " + common.NFS_MOUNT_PATH
	if _, err := util.ExecShell(cmd); nil != err {
		log.Fatal(err)
	}
	log.Println("Mount to nfs server")
}

func (nfs *NFSApi) Mkdir(path string) error {
	return os.MkdirAll(path, 0777)
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
