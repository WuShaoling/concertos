package main

import (
	"log"
	"os"
	"github.com/concertos/module/common"
)

func main() {

	if _, err := os.Stat(common.NFS_MOUNT_LOCAL_PATH); nil != err {
		if os.IsNotExist(err) { //if not exist, create it
			log.Println("not exist")
		} else {
			log.Println("err", err)
		}
	} else {
		log.Println("存在")
	}

}
