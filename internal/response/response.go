package response

import (
	"encoding/json"

	"github.com/abulo/layout/code"
	"github.com/abulo/layout/initial"
	"github.com/abulo/layout/internal/encrypt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/spf13/cast"
)

func JSON(ctx *app.RequestContext, status int, data utils.H) {
	if !initial.Core.Config.Bool("Encrypt.Disable", true) {
		if val, ok := data["data"]; ok {
			stringByte, err := json.Marshal(val)
			if err != nil {
				data["code"] = code.ParamInvalid
				data["msg"] = code.StatusText(code.ParamInvalid)
				data["data"] = nil
			} else {
				data["data"] = encrypt.AesCBCEncrypt(cast.ToString(stringByte), initial.Core.Config.String("Encrypt.SecretKey"), initial.Core.Config.String("Encrypt.Iv"), encrypt.PKCS7)
			}
		}
	}
	ctx.JSON(status, data)
}
