package server

import (
	"github.com/abulo/layout/initial"
	"github.com/abulo/layout/service"
	"github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/spf13/cast"
)

func (eng *Engine) NewGrpcServer() error {
	configAdmin := initial.Core.Config.Get("server.grpc")
	cfg := configAdmin.(map[string]interface{})
	//先获取这个服务是否是需要开启
	if disable := cast.ToBool(cfg["Disable"]); disable {
		logger.Logger.Error("server.grpc is disabled")
		return nil
	}
	client := xgrpc.New()
	client.Host = cast.ToString(cfg["Host"])
	client.Port = cast.ToInt(cfg["Port"])
	client.Deployment = cast.ToString(cfg["Deployment"])
	client.DisableMetric = cast.ToBool(cfg["DisableMetric"])
	client.DisableTrace = cast.ToBool(cfg["DisableTrace"])
	client.ServiceAddress = cast.ToString(cfg["ServiceAddress"])
	client.SlowQueryThresholdInMilli = cast.ToInt64(cfg["SlowQueryThresholdInMilli"])
	res := client.MustBuild()
	//注册服务
	service.Registry(res)
	return eng.Serve(res)
}
