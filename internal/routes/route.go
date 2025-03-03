package routes

import (
	"github.com/abulo/layout/internal/middleware"
	"github.com/abulo/ratel/v3/server/xhertz"
)

func InitRoute(handle *xhertz.Server) {
	handle.Use(middleware.Limiter(), middleware.Request())
	GlobalInitRoute(handle)
}
