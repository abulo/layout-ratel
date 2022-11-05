package flash

import (
	"github.com/abulo/layout/initial"
	"github.com/abulo/ratel/v3/gin"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
)

func PutUrl(ctx *gin.Context, url string) {
	val, b := ctx.Get("CookiesID")
	if !b {
		return
	}
	session := initial.Core.Session(cast.ToString(val))
	session.Put(ctx.Request.Context(), util.NewReplacer(initial.Core.Config.String("cache.manager.url")), url)
}

func GetUrl(ctx *gin.Context) string {
	var res string
	val, b := ctx.Get("CookiesID")
	if !b {
		return res
	}
	session := initial.Core.Session(cast.ToString(val))
	backUrl := session.Get(ctx.Request.Context(), util.NewReplacer(initial.Core.Config.String("cache.manager.url")))
	return cast.ToString(backUrl)
}
