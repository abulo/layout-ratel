package service

import (
	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 系统访问记录
	// logger.RegisterLoginLogServiceServer(server.Server, &logger.SrvLoginLogServiceServer{
	// 	Server: server,
	// })

}
