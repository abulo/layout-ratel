package service

import (
	"github.com/abulo/layout/service/verify"
	"github.com/abulo/ratel/v3/server/xgrpc"
)

// Registry 注册服务
func Registry(server *xgrpc.Server) {
	// 验证码服务->verify
	verify.RegisterVerifyServiceServer(server.Server, &verify.SrvVerifyServiceServer{
		Server: server,
	})
	// 注册服务
}
