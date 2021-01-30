package znet

import (
	"fmt"
	"go-growth/utils"
	"go-growth/ziface"
	"strconv"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
	WorkerPoolSize int32
	TaskQueue []chan ziface.IRequeset
}

func NewMsgHandle() ziface.IMsgHandle  {
	return &MsgHandle{
		Apis: map[uint32]ziface.IRouter{},
		WorkerPoolSize: utils.GlobalObj.WorkerPoolSize,
		TaskQueue: make([]chan ziface.IRequeset, utils.GlobalObj.WorkerPoolSize),
	}
}

func (m MsgHandle) DoMsgHandle(requeset ziface.IRequeset) {
	handler, ok := m.Apis[requeset.GetMsgId()]
	if !ok {
		fmt.Println("api msgID =", requeset.GetMsgId(),"is not found need register")
		return
	}
	handler.PreHeadler(requeset)
	handler.Headler(requeset)
	handler.PostHeadler(requeset)
}

func (m MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgId];ok {
		panic("repeat api msgID ="+ strconv.Itoa(int(msgId)))
	}
	m.Apis[msgId]= router
	fmt.Println("add msgid",msgId, "router succ！！")
}

func (m MsgHandle) StartWorkerPool()  {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		m.TaskQueue[i] = make(chan ziface.IRequeset, utils.GlobalObj.MaxWorkerTaskLen)
		go m.startOneWorker(i, m.TaskQueue[i])
	}
}

func (m MsgHandle) startOneWorker(workerID int, work chan ziface.IRequeset)  {
	fmt.Println("[Start Worker] WorkerID=",workerID)
	for  {
		select {
		case req :=<- work:
			m.DoMsgHandle(req)
		}
	}
}

func (m MsgHandle) SendMsgToQueue(req ziface.IRequeset)  {
	workerID := req.GetConnection().GetConnId()%10
	fmt.Println("[TaskQueue] Add ConnID=", req.GetConnection().GetConnId(),",MsgID=",req.GetMsgId(), "To queue workerID", workerID)
	m.TaskQueue[workerID] <- req
}