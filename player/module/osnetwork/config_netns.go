package osnetwork

import (
	"fmt"
	"io/ioutil"
	"github.com/concertos/player/util"
	"os"
)

func ReadDir(path string) []string {
	dir_list, e := ioutil.ReadDir(path)
	if e != nil {
		fmt.Println("Read dir ", path, " error : ", e)
		return nil
	}
	files := make([]string, 0, 20)
	for _, v := range dir_list {
		files = append(files, v.Name())
	}
	return files
}

//check path is exist
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// create dir
func MakeDir(path string) error {
	return os.Mkdir(path, os.ModePerm);
}

func ConfigFiles() bool {
	// create docker netns dir if not exist
	dockerPath := "/var/run/docker/netns/"
	if !IsPathExist(dockerPath) {
		fmt.Println(dockerPath, " not exist", ", create path ", dockerPath)
		err := MakeDir(dockerPath)
		if err != nil {
			fmt.Println("!!! error : ", err, ", create dir "+dockerPath+" failed")
			return false
		}
	}

	// create system netns dir if not exist
	sysPath := "/var/run/netns/"
	if !IsPathExist(sysPath) {
		fmt.Println(sysPath+" not exist", ", create path ", sysPath)
		err := MakeDir(sysPath)
		if err != nil {
			fmt.Println("!!! error : ", err, ", create dir "+sysPath+" failed")
			return false
		}
	}

	files := ReadDir(dockerPath)
	for _, v := range files {
		util.ExecShell("ln -s " + dockerPath + v + " " + sysPath + v)
		fmt.Println(v)
	}
	return true
}
