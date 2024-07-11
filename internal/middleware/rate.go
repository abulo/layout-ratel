package middleware

import (
	"context"

	"github.com/abulo/layout/initial"
	"github.com/abulo/layout/internal/response"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Limiter() app.HandlerFunc {
	return func(ctx context.Context, newCtx *app.RequestContext) {
		if !initial.Core.Limiter.Allow() {
			response.JSON(newCtx, 429, utils.H{
				"code": "429",
				"msg":  "too many requests",
			})
			newCtx.Abort()
			return
		}
		newCtx.Next(ctx)
	}
}
