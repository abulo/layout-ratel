package server

import (
	"github.com/abulo/ratel/v3/core/app"
	"github.com/abulo/ratel/v3/core/logger"
	"github.com/sirupsen/logrus"
)

type Engine struct {
	app.Application
}

// NewGinEngine Gin框架
func NewGinEngine() *Engine {
	eng := &Engine{}
	//加载计划任务
	// eng.Schedule(eng.CrontabWork())
	// 注册函数
	// eng.RegisterHooks(hooks.Stage_AfterLoadConfig, eng.BeforeInit)
	if err := eng.Startup(
		eng.NewGinServer,
	); err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panic("startup")
	}
	return eng
}

// NewHertzEngine Hertz框架
func NewHertzEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.NewHertzServer,
	); err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panic("startup")
	}
	return eng
}

// NewGrpcEngine Grpc框架
func NewGrpcEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.NewGrpcServer,
	); err != nil {
		logger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Panic("startup")
	}
	return eng
}
