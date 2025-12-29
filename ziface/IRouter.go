package ziface

type IRouter interface {
	// PreHandle 在请求处理前调用
	PreHandle(request IRequest)

	// Handle 在请求处理主逻辑调用
	Handle(request IRequest)

	// PostHandle 在请求处理后调用
	PostHandle(request IRequest)
}
