package server

import (
	"context"
	"fmt"

	"github.com/abulo/layout/initial"
	"github.com/abulo/layout/internal/routes"
	"github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xhertz"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/spf13/cast"
)

func hertzPanicRecoveryHandler(ctx context.Context, newCtx *app.RequestContext, err interface{}, stack []byte) {
	logger.Logger.Error("startup")
	hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
	hlog.SystemLogger().Infof("Client: %s", newCtx.Request.Header.UserAgent())
	newCtx.AbortWithStatus(consts.StatusInternalServerError)
}

func (eng *Engine) NewHertzServer() error {
	var serverConf ServerConf
	if err := initial.Core.Config.BindStruct("Server.Api", &serverConf); err != nil {
		return err
	}
	//先获取这个服务是否是需要开启
	if !cast.ToBool(serverConf.Enable) {
		return fmt.Errorf("Server.Api is disabled")
	}
	client := xhertz.New()
	client.Host = cast.ToString(serverConf.Host)
	client.Port = cast.ToInt(serverConf.Port)
	client.Deployment = cast.ToString(serverConf.Deployment)
	client.EnableMetric = cast.ToBool(serverConf.EnableMetric)
	client.EnableTrace = cast.ToBool(serverConf.EnableTrace)
	client.EnableSlowQuery = cast.ToBool(serverConf.EnableSlowQuery)
	client.ServiceAddress = cast.ToString(serverConf.ServiceAddress)
	client.SlowQueryThresholdInMill = cast.ToInt64(serverConf.SlowQueryThresholdInMill)

	res := client.Build()
	res.Use(recovery.Recovery(recovery.WithRecoveryHandler(hertzPanicRecoveryHandler)))
	routes.InitRoute(res)
	return eng.Serve(res)
}
