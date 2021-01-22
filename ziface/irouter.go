package ziface

type IRouter interface {
	PreHeadler(IRequeset)
	Headler(IRequeset)
	PostHeadler(IRequeset)
}
