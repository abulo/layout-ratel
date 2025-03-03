package server

import (
	"fmt"

	"github.com/abulo/layout/initial"
	"github.com/abulo/ratel/v3/server/xgin"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FormatDateTime(str interface{}) string {
	if util.Empty(str) {
		return ""
	}
	return util.Date("Y-m-d H:i:s", str.(primitive.DateTime).Time())
}

func (eng *Engine) NewGinServer() error {
	var serverConf ServerConf
	if err := initial.Core.Config.BindStruct("Server.Admin", &serverConf); err != nil {
		return err
	}
	//先获取这个服务是否是需要开启
	if !cast.ToBool(serverConf.Enable) {
		return fmt.Errorf("Server.Admin is disabled")
	}

	client := xgin.New()
	client.Host = cast.ToString(serverConf.Host)
	client.Port = cast.ToInt(serverConf.Port)
	client.Deployment = cast.ToString(serverConf.Deployment)
	client.EnableMetric = cast.ToBool(serverConf.EnableMetric)
	client.EnableTrace = cast.ToBool(serverConf.EnableTrace)
	client.EnableSlowQuery = cast.ToBool(serverConf.EnableSlowQuery)
	client.ServiceAddress = cast.ToString(serverConf.ServiceAddress)
	client.SlowQueryThresholdInMill = cast.ToInt64(serverConf.SlowQueryThresholdInMill)
	// if !initial.Core.Config.Bool("DisableDebug", true) {
	// 	client.Mode = gin.DebugMode
	// } else {
	// 	client.Mode = gin.ReleaseMode
	// }
	server := client.Build()

	server.SetTrustedProxies([]string{"0.0.0.0/0"})

	//辅助函数
	// server.InitFuncMap()
	// server.AddFuncMap("config", initial.Core.Config.String)
	// server.AddFuncMap("marshalHtml", util.MarshalHTML)
	// server.AddFuncMap("marshalJs", util.MarshalJS)
	// server.AddFuncMap("static", util.Static)
	// server.AddFuncMap("js", util.JS)
	// server.AddFuncMap("formatDate", util.FormatDate)
	// server.AddFuncMap("formatDateTime", util.FormatDateTime)
	// server.AddFuncMap("unixTimeFormatDate", FormatDateTime)
	// server.AddFuncMap("inArray", util.InArray)
	// server.AddFuncMap("multiArray", util.MultiArray)
	// server.AddFuncMap("inMultiArray", util.InMultiArray)
	// server.AddFuncMap("empty", util.Empty)
	// server.AddFuncMap("divide", util.Divide)
	// server.AddFuncMap("add", util.Add)
	// server.AddFuncMap("strReplace", util.StrReplace)

	// // 开发模式
	// if !initial.Core.Config.Bool("DisableDebug", true) {
	// 	//模板
	// 	t, err := loadGlobTemplate(initial.Core.Path + "/view")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	server.LoadHTMLFiles(t...)
	// 	// server.Use(gin.Logger())
	// } else {
	// 	//加载模板文件
	// 	t, err := loadTemplate(server.FuncMap)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	server.SetHTMLTemplate(t)
	// }
	// //静态文件目录
	// server.Static("/static", initial.Core.Path+"/static")

	// // 通过 go-bindata-assetfs 提供的函数将静态资源
	// staticFS := assetfs.AssetFS{
	// 	Asset:    asset.Asset,
	// 	AssetDir: asset.AssetDir,
	// 	// AssetInfo: asset.AssetInfo,
	// 	Prefix:   "resource", // 访问路由index.html => 指向文件 resource/index.html
	// 	Fallback: "index.html",
	// }
	// server.StaticFS("/resource", &staticFS) //配置静态资源文件路由
	//添加路由
	// pprof.Register(server.Engine)
	// backstage.Route(server.Engine)
	// mobile.Route(server.Engine)
	return eng.Serve(server)
}

// // 加载模板文件
// func loadTemplate(funcMap template.FuncMap) (*template.Template, error) {
// 	t := template.New("").Delims("{{", "}}").Funcs(funcMap)

// 	for _, name := range view.AssetNames() {
// 		if !strings.HasSuffix(name, ".html") {
// 			continue
// 		}
// 		asset, err := view.Asset(name)
// 		if err != nil {
// 			continue
// 		}
// 		name := strings.Replace(name, "view/", "", 1)
// 		t, err = t.New(name).Parse(string(asset))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return t, nil
// }

// func loadGlobTemplate(dir string) ([]string, error) {
// 	fileList := []string{}
// 	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
// 		if !f.IsDir() {
// 			fileList = append(fileList, filepath.FromSlash(path))
// 		}
// 		return nil
// 	})
// 	return fileList, err
// }
