package ziface
type IServer interface {
	Start()
	Stop()
	//运行服务
	Server()
}