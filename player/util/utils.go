package util

import (
	"net"
	"log"
	"fmt"
	"os/exec"
	"bytes"
	"errors"
	"encoding/json"
)

func MyJsonMarshal(info interface{}) []byte {
	res, err := json.Marshal(info)
	if nil != err {
		log.Println("MyJsonMarshal: ", err)
	}
	return res
}

func ExecShell(s string) (string, error) {
	fmt.Println(s)
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if nil != err {
		return "", errors.New(string(out.Bytes()))
	}
	return string(out.Bytes()), nil
}

func GetIps() []string {
	ips := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal("Error, get ips failed, ", err.Error())
		return ips
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return ips
}
