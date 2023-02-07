package server

import (
	"github.com/abulo/layout/initial"
	"github.com/abulo/layout/internal/routes"
	"github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xhertz"
	"github.com/spf13/cast"
)

func (eng *Engine) NewHertzServer() error {
	configAdmin := initial.Core.Config.Get("server.admin")
	cfg := configAdmin.(map[string]interface{})
	//先获取这个服务是否是需要开启
	if disable := cast.ToBool(cfg["Disable"]); disable {
		logger.Logger.Error("server.admin is disabled")
		return nil
	}
	client := xhertz.New()
	client.Host = cast.ToString(cfg["Host"])
	client.Port = cast.ToInt(cfg["Port"])
	client.Deployment = cast.ToString(cfg["Deployment"])
	client.DisableMetric = cast.ToBool(cfg["DisableMetric"])
	client.DisableTrace = cast.ToBool(cfg["DisableTrace"])
	client.DisableSlowQuery = cast.ToBool(cfg["DisableSlowQuery"])
	client.ServiceAddress = cast.ToString(cfg["ServiceAddress"])
	client.SlowQueryThresholdInMilli = cast.ToInt64(cfg["SlowQueryThresholdInMilli"])

	res := client.Build()
	routes.InitRoute(res)
	return eng.Serve(res)
}
