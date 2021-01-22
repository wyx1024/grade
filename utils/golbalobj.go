package utils

import (
	"encoding/json"
	"go-growth/ziface"
	"io/ioutil"
)

type globalobj struct {
	//sever
	TcpServer ziface.IServer
	Name     string
	TcpPort  int32
	Host     string
	//version
	Version        string
	MaxConn        int32
	MaxPackageSize int32
}

var GlobalObj *globalobj

func (g *globalobj) Reload()  {

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
		Name:           "ZinxServerApp",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		Version:        "V4.0",
		MaxConn:        2,
		MaxPackageSize: 512,
	}
	GlobalObj.Reload()
}
