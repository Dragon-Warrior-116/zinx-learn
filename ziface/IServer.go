package ziface

type IServer interface {
	Start()
	Stop()
	//运行服务
	Server()

	// 添加路由
	AddRouter(router IRouter)
}
