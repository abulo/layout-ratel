package server

import (
	"fmt"
	"time"

	"github.com/abulo/layout/initial"
	"github.com/abulo/layout/service"
	"github.com/abulo/ratel/v3/server/xgrpc"
	"github.com/abulo/ratel/v3/server/xgrpc/recovery"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"
)

var grpcPanicRecoveryHandler recovery.RecoveryHandlerFunc

func (eng *Engine) NewGrpcServer() error {

	var serverConf ServerConf
	if err := initial.Core.Config.BindStruct("Server.Grpc", &serverConf); err != nil {
		return err
	}
	//先获取这个服务是否是需要开启
	if !cast.ToBool(serverConf.Enable) {
		return fmt.Errorf("Server.Grpc is disabled")
	}

	client := xgrpc.New()
	client.Host = cast.ToString(serverConf.Host)
	client.Port = cast.ToInt(serverConf.Port)
	client.Deployment = cast.ToString(serverConf.Deployment)
	client.EnableMetric = cast.ToBool(serverConf.EnableMetric)
	client.EnableTrace = cast.ToBool(serverConf.EnableTrace)
	client.ServiceAddress = cast.ToString(serverConf.ServiceAddress)
	client.SlowQueryThresholdInMill = cast.ToInt64(serverConf.SlowQueryThresholdInMill)

	client.WithServerOption(
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
				PermitWithoutStream: true,            // Allow pings even when there are no active streams
			},
		),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
				MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
				MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
				Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
				Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
			},
		),
	)

	//防止服务端panic
	grpcPanicRecoveryHandler = func(p any) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
	client.WithStreamInterceptor(recovery.StreamServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)))
	client.WithUnaryInterceptor(recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)))
	res := client.MustBuild()
	//注册服务
	service.Registry(res)
	return eng.Serve(res)
}
