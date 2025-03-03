package verify

import (
	"context"
	"time"

	"github.com/abulo/layout/code"
	"github.com/abulo/layout/initial"
	globalLogger "github.com/abulo/ratel/v3/core/logger"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/stores/redis"
	"github.com/abulo/ratel/v3/util"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

// SrvVerifyServiceServer 验证码服务
type SrvVerifyServiceServer struct {
	UnimplementedVerifyServiceServer
	Server *xgrpc.Server
}

// CaptchaDriver 驱动
type VerifyDriver struct {
	Driver base64Captcha.Driver
	Store  *redis.Client
}

// Generate  生成验证码
func (c *VerifyDriver) Generate(ctx context.Context) (id, b64s string, err error) {
	id, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, err := c.Driver.DrawCaptcha(content)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Captcha:验证码:Generate")
		return
	}
	verifyKey := util.NewReplacer(initial.Core.Config.String("Cache.Verify.Code"))
	_, err = c.Store.HSet(ctx, verifyKey, id, util.StrToLower(answer))
	c.Store.HExpire(ctx, verifyKey, time.Second*90, id)
	if err != nil {
		globalLogger.Logger.WithFields(logrus.Fields{
			"err": err,
		}).Error("Captcha:验证码:Generate")
		return
	}
	b64s = item.EncodeB64string()
	return
}

// Verify 验证
func (c *VerifyDriver) Verify(ctx context.Context, id, answer string, clear bool) bool {
	verifyKey := util.NewReplacer(initial.Core.Config.String("Cache.Verify.Code"))
	redisAnswer, err := c.Store.HGet(ctx, verifyKey, id)
	if err != nil {
		return false
	}
	match := util.StrToLower(redisAnswer) == util.StrToLower(answer)
	if clear {
		c.Store.HDel(ctx, redisAnswer, id)
	}
	return match
}

// NewCaptcha 创建验证码构造器
func (srv SrvVerifyServiceServer) NewCaptcha(store *redis.Client) *VerifyDriver {
	randCode := util.Rand(1, 4)
	var driver base64Captcha.Driver
	switch randCode {
	case 1:
		driverString := &base64Captcha.DriverString{
			Width:           360,
			Height:          120,
			Length:          5,
			NoiseCount:      0,
			ShowLineOptions: 14,
			Source:          "23456789qwertyuplkjhgfdsazxcvbnm",
			Fonts:           []string{"wqy-microhei.ttc"},
		}
		driver = driverString.ConvertFonts()
	case 2:
		driverMath := &base64Captcha.DriverMath{
			Width:           360,
			Height:          120,
			NoiseCount:      0,
			ShowLineOptions: 14,
			Fonts:           []string{"wqy-microhei.ttc"},
		}
		driver = driverMath.ConvertFonts()
	case 3:
		// driver = &base64Captcha.DriverDigit{
		// 	Height:   96,
		// 	Width:    360,
		// 	Length:   5,
		// 	DotCount: 99,
		// 	MaxSkew:  3,
		// }
		driverDigit := &base64Captcha.DriverString{
			Width:           360,
			Height:          120,
			Length:          5,
			NoiseCount:      0,
			ShowLineOptions: 14,
			Source:          "123456789",
			Fonts:           []string{"wqy-microhei.ttc"},
		}
		driver = driverDigit.ConvertFonts()
	case 4:
		driverChinese := &base64Captcha.DriverChinese{
			Width:           360,
			Height:          120,
			Length:          4,
			NoiseCount:      0,
			ShowLineOptions: 14,
			Source:          base64Captcha.TxtChineseCharaters,
			Fonts:           []string{"wqy-microhei.ttc"},
		}
		driver = driverChinese.ConvertFonts()
	}
	return &VerifyDriver{
		Driver: driver,
		Store:  store,
	}
}

// Generate 生成验证码
func (srv SrvVerifyServiceServer) Generate(ctx context.Context, request *GenerateRequest) (*GenerateResponse, error) {
	//这里需要去生成验证码
	redisHandler := initial.Core.Store.LoadRedis("redis")
	verifyDriver := srv.NewCaptcha(redisHandler)
	id, b64s, err := verifyDriver.Generate(ctx)
	if err != nil {
		return &GenerateResponse{}, status.Error(code.ConvertToGrpc(code.RedisError), err.Error())
	}
	return &GenerateResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: &GenerateObject{
			VerifyCodeId: id,
			VerifyImage:  b64s,
		},
	}, nil
}

// Verify 验证验证码
func (srv SrvVerifyServiceServer) Verify(ctx context.Context, request *VerifyRequest) (*VerifyResponse, error) {
	//这里需要去生成验证码
	redisHandler := initial.Core.Store.LoadRedis("redis")
	verifyDriver := srv.NewCaptcha(redisHandler)
	verifyCodeId := request.GetVerifyCodeId()
	verifyCode := request.GetVerifyCode()
	res := verifyDriver.Verify(ctx, verifyCodeId, verifyCode, true)
	return &VerifyResponse{
		Code: code.Success,
		Msg:  code.StatusText(code.Success),
		Data: &VerifyObject{
			Result: res,
		},
	}, nil
}
