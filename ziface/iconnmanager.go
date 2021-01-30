package ziface

type IConnManager interface {
	Add(IConnection)
	Remove(IConnection)
	Get(uint322 uint32)(IConnection, error)
	GetLen()int
	ClearConn()
}
