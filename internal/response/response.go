package response

import (
	"encoding/json"

	"github.com/abulo/layout/code"
	"github.com/abulo/layout/initial"
	"github.com/abulo/layout/internal/encrypt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func JSON(ctx any, status int, data map[string]any) {
	if !initial.Core.Config.Bool("encrypt.Disable", true) {
		if val, ok := data["data"]; ok {
			stringByte, err := json.Marshal(val)
			if err != nil {
				data["code"] = code.ParamInvalid
				data["msg"] = code.StatusText(code.ParamInvalid)
				data["data"] = nil
			} else {
				data["data"] = encrypt.AesCBCEncrypt(cast.ToString(stringByte), initial.Core.Config.String("encrypt.SecretKey"), initial.Core.Config.String("encrypt.Iv"), encrypt.PKCS7)
			}
		}
	}
	// 判断 泛型 ctx 是 *app.RequestContext 还是 *gin.Context
	if newCtx, ok := ctx.(*app.RequestContext); ok {
		newCtx.JSON(status, data)
	}
	if newCtx, ok := ctx.(*gin.Context); ok {
		newCtx.JSON(status, data)
	}
}
