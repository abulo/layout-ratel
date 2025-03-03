package verify

import (
	"context"

	"github.com/abulo/layout/code"
	"github.com/abulo/layout/initial"
	"github.com/abulo/layout/internal/response"
	"github.com/abulo/layout/service/verify"
	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// Generate 生成验证码
func Generate(ctx context.Context, newCtx *app.RequestContext) {
	//判断这个服务能不能链接
	grpcClient, err := initial.Core.Client.LoadGrpc("grpc").Singleton()
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Grpc:验证码:Generate")
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.RPCError,
			"msg":  code.StatusText(code.RPCError),
		})
		return
	}
	//链接服务
	client := verify.NewVerifyServiceClient(grpcClient)
	request := &verify.GenerateRequest{}
	// 执行服务
	res, err := client.Generate(ctx, request)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"req": request,
			"err": err,
		}).Error("GrpcCall:验证码:Generate")
		fromError := status.Convert(err)
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": code.ConvertToHttp(fromError.Code()),
			"msg":  code.StatusText(code.ConvertToHttp(fromError.Code())),
		})
		return
	}
	if res.GetCode() != code.Success {
		response.JSON(newCtx, consts.StatusOK, utils.H{
			"code": res.GetCode(),
			"msg":  code.StatusText(res.GetCode()),
		})
		return
	}
	response.JSON(newCtx, consts.StatusOK, utils.H{
		"code": code.Success,
		"msg":  code.StatusText(code.Success),
		"data": utils.H{
			"verifyCodeId": res.GetData().GetVerifyCodeId(),
			"verifyImage":  res.GetData().GetVerifyImage(),
		},
	})
}
