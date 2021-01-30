package ziface

type IMsgHandle interface {
	DoMsgHandle(IRequeset)
	AddRouter(uint32,IRouter)
	StartWorkerPool()
 	SendMsgToQueue(IRequeset)
}
