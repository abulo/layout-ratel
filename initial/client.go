package initial

import (
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

type (
	EtcdConf struct {
		Name             *string      `json:"Name"`
		Endpoints        *StringSlice `json:"Endpoints"`
		CertFile         *string      `json:"CertFile"`
		KeyFile          *string      `json:"KeyFile"`
		CaCert           *string      `json:"CaCert"`
		BasicAuth        *bool        `json:"BasicAuth"`
		UserName         *string      `json:"UserName"`
		Password         *string      `json:"Password"`
		ConnectTimeout   *string      `json:"ConnectTimeout"`
		Secure           *bool        `json:"Secure"`
		AutoSyncInterval *string      `json:"AutoSyncInterval"`
		Prefix           *string      `json:"Prefix"`
		EnableTrace      *bool        `json:"EnableTrace"`
		LinkNode         *string      `json:"LinkNode"`
	}
	ProxyEtcdConf struct {
		Name     *string `json:"Name"`
		LinkNode *string `json:"LinkNode"`
	}

	GrpcConf struct {
		Name                     *string `json:"Name"`
		BalancerName             *string `json:"BalancerName"`
		Address                  *string `json:"Address"`
		Block                    *bool   `json:"Block"`
		DialTimeout              *string `json:"DialTimeout"`
		ReadTimeout              *string `json:"ReadTimeout"`
		Direct                   *bool   `json:"Direct"`
		SlowThreshold            *string `json:"SlowThreshold"`
		EnableDebug              *bool   `json:"EnableDebug"`
		EnableTraceInterceptor   *bool   `json:"EnableTraceInterceptor"`
		EnableAidInterceptor     *bool   `json:"EnableAidInterceptor"`
		EnableTimeoutInterceptor *bool   `json:"EnableTimeoutInterceptor"`
		EnableMetricInterceptor  *bool   `json:"EnableMetricInterceptor"`
		EnableAccessInterceptor  *bool   `json:"EnableAccessInterceptor"`
		AccessInterceptorLevel   *string `json:"AccessInterceptorLevel"`
		Etcd                     *string `json:"Etcd"`
		LinkNode                 *string `json:"LinkNode"`
	}
	ProxyGrpcConf struct {
		Name     *string `json:"Name"`
		LinkNode *string `json:"LinkNode"`
	}
)

func (initial *Initial) InitRegistry() *Initial {
	var etcdConfs []EtcdConf
	if err := initial.Config.BindStruct("Etcd", &etcdConfs); err != nil {
		panic(err)
	}
	links := make(map[string]*etcdv3.Config)
	for _, item := range etcdConfs {
		config := etcdv3.New()
		if item.Name == nil {
			panic("etcd name is empty")
		}
		config.SetName(cast.ToString(item.Name))
		config.SetNode(cast.ToString(item.Name))
		if item.Endpoints == nil {
			panic("etcd endpoints is empty")
		}
		config.SetEndpoints(*item.Endpoints)
		if item.CertFile != nil {
			config.SetCertFile(cast.ToString(item.CertFile))
		}
		if item.KeyFile != nil {
			config.SetKeyFile(cast.ToString(item.KeyFile))
		}
		if item.CaCert != nil {
			config.SetCaCert(cast.ToString(item.CaCert))
		}
		config.SetBasicAuth(cast.ToBool(item.BasicAuth))
		if item.UserName != nil {
			config.SetUserName(cast.ToString(item.UserName))
		}
		if item.Password != nil {
			config.SetPassword(cast.ToString(item.Password))
		}
		if item.ConnectTimeout != nil {
			config.SetConnectTimeout(util.Duration(cast.ToString(item.ConnectTimeout)))
		}
		config.SetSecure(cast.ToBool(item.Secure))
		if item.AutoSyncInterval != nil {
			config.SetAutoSyncInterval(util.Duration(cast.ToString(item.AutoSyncInterval)))
		}
		if item.Prefix != nil {
			config.SetPrefix(cast.ToString(item.Prefix))
		}
		if item.EnableTrace != nil {
			config.SetEnableTrace(cast.ToBool(item.EnableTrace))
		}
		node := cast.ToString(item.LinkNode)
		links[node] = config
	}
	var proxyEtcdConfs []ProxyEtcdConf
	if err := initial.Config.BindStruct("ProxyEtcd", &proxyEtcdConfs); err != nil {
		panic(err)
	}
	for _, val := range proxyEtcdConfs {
		proxyPool := client.NewEtcdConfig()
		if val.LinkNode == nil {
			panic("ProxyEtcd LinkNode is empty")
		}
		linkNode := cast.ToString(val.LinkNode)
		proxyPool.Store(links[linkNode])
		if val.Name == nil {
			panic("ProxyEtcd Name is empty")
		}
		node := cast.ToString(val.Name)
		initial.Client.StoreEtcd(node, proxyPool)
	}
	return initial
}

func (initial *Initial) InitGrpc() *Initial {
	var grpcConfs []GrpcConf
	if err := initial.Config.BindStruct("Grpc", &grpcConfs); err != nil {
		panic(err)
	}
	links := make(map[string]*grpc.Config)
	for _, item := range grpcConfs {
		config := grpc.New()
		if item.Name != nil {
			config.SetName(cast.ToString(item.Name))
		}
		if item.BalancerName != nil {
			config.SetBalancerName(cast.ToString(item.BalancerName))
		}
		if item.Address != nil {
			config.SetAddress(cast.ToString(item.Address))
		}
		if item.Block == nil {
			panic("grpc block is empty")
		}
		config.SetBlock(cast.ToBool(item.Block))
		if item.DialTimeout != nil {
			config.SetDialTimeout(util.Duration(cast.ToString(item.DialTimeout)))
		}
		if item.ReadTimeout != nil {
			config.SetReadTimeout(util.Duration(cast.ToString(item.ReadTimeout)))
		}
		if item.Direct == nil {
			panic("grpc direct is empty")
		}
		config.SetDirect(cast.ToBool(item.Direct))

		if item.SlowThreshold != nil {
			config.SetSlowThreshold(util.Duration(cast.ToString(item.SlowThreshold)))
		}
		config.SetEnableDebug(cast.ToBool(item.EnableDebug))
		config.SetEnableTraceInterceptor(cast.ToBool(item.EnableTraceInterceptor))
		config.SetEnableAidInterceptor(cast.ToBool(item.EnableAidInterceptor))
		config.SetEnableTimeoutInterceptor(cast.ToBool(item.EnableTimeoutInterceptor))
		config.SetEnableMetricInterceptor(cast.ToBool(item.EnableMetricInterceptor))
		config.SetEnableAccessInterceptor(cast.ToBool(item.EnableAccessInterceptor))
		if item.AccessInterceptorLevel != nil {
			config.SetAccessInterceptorLevel(cast.ToString(item.AccessInterceptorLevel))
		}

		config.KeepAlive = &keepalive.ClientParameters{
			Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
			Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
			PermitWithoutStream: true,             // send pings even without active streams
		}
		if Etcd := cast.ToString(item.Etcd); Etcd != "" {
			config.SetEtcd(initial.Client.LoadEtcd(Etcd))
		}
		config.BalancerName = balancer.NameSmoothWeightRoundRobin
		config.BalancerName = p2c.Name
		node := cast.ToString(item.LinkNode)
		links[node] = config
	}

	var proxyGrpcConfs []ProxyGrpcConf
	if err := initial.Config.BindStruct("ProxyGrpc", &proxyGrpcConfs); err != nil {
		panic(err)
	}

	for _, val := range proxyGrpcConfs {
		proxyPool := client.NewGrpcConfig()
		if val.LinkNode == nil {
			panic("ProxyGrpc LinkNode is empty")
		}
		linkNode := cast.ToString(val.LinkNode)
		proxyPool.Store(links[linkNode])
		if val.Name == nil {
			panic("ProxyGrpc Name is empty")
		}
		name := cast.ToString(val.Name)
		initial.Client.StoreGrpc(name, proxyPool)
	}

	return initial
}
