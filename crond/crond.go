package crond

import (
	"fmt"

	"github.com/abulo/layout/initial"

	"github.com/abulo/ratel/v3/core/task"
	"github.com/abulo/ratel/v3/core/task/driver/redis"
)

func CronJob() func() {
	redisHandler := initial.Core.Store.LoadRedis("redis")
	driverHandler := redis.NewDriver(redisHandler)
	cron := task.NewTask("WorkerService", driverHandler, task.WithLazyPick(true), task.WithSeconds())
	// 解析全局列队
	_ = cron.AddFunc("CommonQueue", task.JobLocal, "*/2 * * * * *", CommonQueue)
	cron.Start()
	return func() { cron.Stop() }
}

func CommonQueue() {
	fmt.Println("CommonQueue")
	return
}
