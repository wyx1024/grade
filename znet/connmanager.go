package znet

import (
	"errors"
	"fmt"
	"go-growth/ziface"
	"strconv"
	"sync"
)

type ConnManager struct {
	Connections map[uint32]ziface.IConnection
	connLook sync.RWMutex
}

func NewConnManager() ziface.IConnManager  {
	return &ConnManager{
		Connections: make(map[uint32]ziface.IConnection),
	}
}
func (connMgr ConnManager) Add(conn ziface.IConnection) {
	connMgr.connLook.Lock()
	defer connMgr.connLook.Unlock()

	connMgr.Connections[conn.GetConnId()] = conn
	fmt.Println("[ConnMgr] Add New ConnID=", conn.GetConnId(),"Now Len=", connMgr.GetLen())
}

func (connMgr ConnManager) Remove(conn ziface.IConnection) {
	connMgr.connLook.Lock()
	defer connMgr.connLook.Unlock()
	delete(connMgr.Connections, conn.GetConnId())
	fmt.Println("[ConnMgr] Remove ConnID=", conn.GetConnId(),"Now Len=", connMgr.GetLen())
}

func (connMgr ConnManager) Get(connID uint32) (ziface.IConnection, error){
	connMgr.connLook.RLock()
	defer connMgr.connLook.RUnlock()
	fmt.Println("[ConnMgr] Get ConnID=", connID)
	if conn, ok := connMgr.Connections[connID]; ok{
		return conn, nil
	}
	return nil, errors.New("[ConnMgr] NOT FOUND ConnID="+strconv.Itoa(int(connID)))
}

func (connMgr ConnManager) GetLen() int {
	return len(connMgr.Connections)
}

func (connMgr ConnManager) ClearConn() {
	connMgr.connLook.Lock()
	defer connMgr.connLook.Unlock()
	fmt.Println("[ConnMgr] Clear Conn")
	for connId, conn := range connMgr.Connections {
		conn.Stop()
		delete(connMgr.Connections,connId)
	}
}

