package routes

import (
	"github.com/abulo/layout/api/verify"
	"github.com/abulo/ratel/v3/server/xhertz"
)

func GlobalInitRoute(handle *xhertz.Server) {
	route := handle.Group("/v1")
	{
		// 验证码生成
		route.GET("/verify/generate", verify.Generate)
	}
}
