package routes

import (
	"context"

	"github.com/abulo/ratel/v3/server/xhertz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func InitRoute(handle *xhertz.Server) {
	handle.GET("/get", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "get")
	})
}
