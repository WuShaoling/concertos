package util

import (
	"net"
	"log"
	"fmt"
	"os/exec"
	"bytes"
	"errors"
	"encoding/json"
	"os"
	"io/ioutil"
	"path/filepath"
	"bufio"
	"io"
)

func MyJsonMarshal(info interface{}) []byte {
	res, err := json.Marshal(info)
	if nil != err {
		log.Println("MyJsonMarshal: ", err)
	}
	return res
}

func ExecShell(s string) (string, error) {
	//fmt.Println(s)
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	result := string(out.Bytes())

	if nil != err {
		log.Println("exec error : ", result)
		return "", errors.New(result)
	}
	log.Println("exec result: ", result)
	return result, nil
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

func WriteFile(data, path string) {
	file, error := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0766);
	if error != nil {
		fmt.Println(error);
	}
	file.WriteString(data)
	file.Close();
}

func ReadFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("read file error: ", err)
		return nil, err
	}
	return b, err
}

func ReadLines(path string) ([]string, error) {
	files, _ := filepath.Glob(path)
	var lines []string
	for _, file := range files {
		f, err := os.OpenFile(file, os.O_RDONLY, 0666)
		if err != nil {
			log.Println("read file line error : ", err)
			f.Close()
			return nil, err
		}
		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			lines = append(lines, line)
		}
	}

	return lines, nil
}
