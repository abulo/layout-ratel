package initial

import (
	"fmt"
	"time"

	"github.com/abulo/ratel/v3/client"
	"github.com/abulo/ratel/v3/client/grpc"
	"github.com/abulo/ratel/v3/client/grpc/balancer"
	"github.com/abulo/ratel/v3/client/grpc/balancer/p2c"
	"github.com/abulo/ratel/v3/registry/etcdv3"
	"github.com/abulo/ratel/v3/util"
	"github.com/spf13/cast"
	"google.golang.org/grpc/keepalive"
)

func (initial *Initial) InitRegistry() *Initial {
	configs := initial.Config.Get("etcd")
	list := configs.(map[string]interface{})
	links := make(map[string]*etcdv3.Config)
	for node, nodeConfig := range list {
		config := etcdv3.New()
		res := nodeConfig.(map[string]interface{})
		if Name := cast.ToString(res["Name"]); Name != "" {
			config.SetName(Name)
			config.SetNode(Name)
		}
		if Endpoints := cast.ToStringSlice(res["Endpoints"]); len(Endpoints) > 0 {
			config.SetEndpoints(Endpoints)
		}
		if CertFile := cast.ToString(res["CertFile"]); CertFile != "" {
			config.SetCertFile(CertFile)
		}
		if KeyFile := cast.ToString(res["KeyFile"]); KeyFile != "" {
			config.SetKeyFile(KeyFile)
		}
		if CaCert := cast.ToString(res["CaCert"]); CaCert != "" {
			config.SetCaCert(CaCert)
		}
		if BasicAuth := cast.ToBool(res["BasicAuth"]); BasicAuth {
			config.SetBasicAuth(true)
		} else {
			config.SetBasicAuth(false)
		}
		if UserName := cast.ToString(res["UserName"]); UserName != "" {
			config.SetUserName(UserName)
		}
		if Password := cast.ToString(res["Password"]); Password != "" {
			config.SetPassword(Password)
		}
		if ConnectTimeout := cast.ToString(res["ConnectTimeout"]); ConnectTimeout != "" {
			config.SetConnectTimeout(util.Duration(ConnectTimeout))
		}
		if Secure := cast.ToBool(res["Secure"]); Secure {
			config.SetSecure(true)
		} else {
			config.SetSecure(false)
		}
		if AutoSyncInterval := cast.ToString(res["AutoSyncInterval"]); AutoSyncInterval != "" {
			config.SetAutoSyncInterval(util.Duration(AutoSyncInterval))
		}
		if Prefix := cast.ToString(res["Prefix"]); Prefix != "" {
			config.SetPrefix(Prefix)
		}
		links["etcd."+node] = config
	}
	proxyConfigs := initial.Config.Get("proxyetcd")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := client.NewEtcdConfig()
		if node := cast.ToString(val["Node"]); node != "" {
			proxyPool.Store(links[node])
		}
		if Name := cast.ToString(val["Name"]); Name != "" {
			initial.Client.StoreEtcd(Name, proxyPool)
		}
	}
	return initial
}

func (initial *Initial) InitGrpc() *Initial {
	configs := initial.Config.Get("grpc")
	list := configs.(map[string]interface{})
	links := make(map[string]*grpc.Config)
	for node, nodeConfig := range list {
		config := grpc.New()
		res := nodeConfig.(map[string]interface{})
		if Name := cast.ToString(res["Name"]); Name != "" {
			config.SetName(Name)
		}
		if BalancerName := cast.ToString(res["BalancerName"]); BalancerName != "" {
			config.SetBalancerName(BalancerName)
		}
		if Address := cast.ToString(res["Address"]); Address != "" {
			config.SetAddress(Address)
		}
		if Block := cast.ToBool(res["Block"]); Block {
			config.SetBlock(true)
		} else {
			config.SetBlock(false)
		}
		if DialTimeout := cast.ToString(res["DialTimeout"]); DialTimeout != "" {
			config.SetDialTimeout(util.Duration(DialTimeout))
		}
		if ReadTimeout := cast.ToString(res["ReadTimeout"]); ReadTimeout != "" {
			config.SetReadTimeout(util.Duration(ReadTimeout))
		}
		if Direct := cast.ToBool(res["Direct"]); Direct {
			config.SetDirect(true)
		} else {
			config.SetDirect(false)
		}
		if SlowThreshold := cast.ToString(res["SlowThreshold"]); SlowThreshold != "" {
			config.SetSlowThreshold(util.Duration(SlowThreshold))
		}
		if Debug := cast.ToBool(res["Debug"]); Debug {
			config.SetDebug(true)
		} else {
			config.SetDebug(false)
		}
		if DisableTraceInterceptor := cast.ToBool(res["DisableTraceInterceptor"]); DisableTraceInterceptor {
			config.SetDisableTraceInterceptor(true)
		} else {
			config.SetDisableTraceInterceptor(false)
		}
		if DisableAidInterceptor := cast.ToBool(res["DisableAidInterceptor"]); DisableAidInterceptor {
			config.SetDisableAidInterceptor(true)
		} else {
			config.SetDisableAidInterceptor(false)
		}
		if DisableTimeoutInterceptor := cast.ToBool(res["DisableTimeoutInterceptor"]); DisableTimeoutInterceptor {
			config.SetDisableTimeoutInterceptor(true)
		} else {
			config.SetDisableTimeoutInterceptor(false)
		}
		if DisableMetricInterceptor := cast.ToBool(res["DisableMetricInterceptor"]); DisableMetricInterceptor {
			config.SetDisableMetricInterceptor(true)
		} else {
			config.SetDisableMetricInterceptor(false)
		}
		if DisableAccessInterceptor := cast.ToBool(res["DisableAccessInterceptor"]); DisableAccessInterceptor {
			config.SetDisableAccessInterceptor(true)
		} else {
			config.SetDisableAccessInterceptor(false)
		}
		if AccessInterceptorLevel := cast.ToString(res["AccessInterceptorLevel"]); AccessInterceptorLevel != "" {
			config.SetAccessInterceptorLevel(AccessInterceptorLevel)
		}
		config.KeepAlive = &keepalive.ClientParameters{
			Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
			Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
			PermitWithoutStream: true,             // send pings even without active streams
		}
		if Etcd := cast.ToString(res["Etcd"]); Etcd != "" {
			config.SetEtcd(initial.Client.LoadEtcd(Etcd))
		}

		config.BalancerName = balancer.NameSmoothWeightRoundRobin
		config.BalancerName = p2c.Name
		links["grpc."+node] = config
	}

	proxyConfigs := initial.Config.Get("proxygrpc")
	proxyRes := proxyConfigs.([]map[string]interface{})
	for _, val := range proxyRes {
		proxyPool := client.NewGrpcConfig()
		if node := cast.ToString(val["Node"]); node != "" {
			proxyPool.Store(links[node])
		}
		if Name := cast.ToString(val["Name"]); Name != "" {
			fmt.Println(Name, "nnnnnnnn")
			initial.Client.StoreGrpc(Name, proxyPool)
		}
	}

	return initial
}
