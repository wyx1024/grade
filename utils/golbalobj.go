package utils

import (
	"encoding/json"
	"go-growth/ziface"
	"io/ioutil"
)

type globalobj struct {
	//sever
	TcpServer ziface.IServer
	Name      string
	TcpPort   int32
	Host      string
	//version
	Version          string
	MaxConn          int
	MaxPackageSize   int32
	WorkerPoolSize   int32
	MaxWorkerTaskLen int32
}

var GlobalObj *globalobj

func (g *globalobj) Reload() {

	data, err := ioutil.ReadFile("conf/app.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, GlobalObj)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObj = &globalobj{
		Name:             "ZinxServerApp",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		Version:          "V0.9",
		MaxConn:          2,
		MaxPackageSize:   512,
		MaxWorkerTaskLen: 1024,
		WorkerPoolSize:   10,
	}
	GlobalObj.Reload()
}
