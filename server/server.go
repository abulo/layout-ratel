package server

import (
	"github.com/abulo/ratel/v3/core/app"
	"github.com/abulo/ratel/v3/core/logger"
	"github.com/sirupsen/logrus"
)

type Engine struct {
	app.Application
}

func NewHttpEngine() *Engine {
	eng := &Engine{}

	//加载计划任务
	// eng.Schedule(eng.CrontabWork())
	// 注册函数
	// eng.RegisterHooks(hooks.Stage_AfterLoadConfig, eng.BeforeInit)
	if err := eng.Startup(
		eng.NewHttpServer,
		// eng.ApiServer,
	); err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panic("startup")
	}
	return eng
}
