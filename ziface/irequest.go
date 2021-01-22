package ziface

type IRequeset interface {
	GetConnection() IConnection
	GetData() []byte
}
