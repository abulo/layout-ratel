package initial

import (
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/stores/redis"
	"github.com/spf13/cast"
)

type (
	RedisConf struct {
		ClientType   *string          `json:"ClientType"`
		Addr         *string          `json:"Addr"`
		Hosts        *StringSlice     `json:"Hosts"`
		Addrs        *StringMapString `json:"Addrs"`
		MasterName   *string          `json:"MasterName"`
		Database     *int             `json:"Database"`
		Password     *string          `json:"Password"`
		PoolSize     *int             `json:"PoolSize"`
		EnableMetric *bool            `json:"EnableMetric"`
		EnableTrace  *bool            `json:"EnableTrace"`
		LinkNode     *string          `json:"LinkNode"`
	}
	ProxyRedisConf struct {
		Name     *string `json:"Name"`
		LinkNode *string `json:"LinkNode"`
	}
)

// InitRedis load redis && returns an redis instance.
func (initial *Initial) InitRedis() *Initial {

	var redisConfigs []RedisConf
	if err := initial.Config.BindStruct("Redis", &redisConfigs); err != nil {
		panic(err)
	}

	links := make(map[string]*redis.Client)
	for _, item := range redisConfigs {
		opts := make([]redis.Option, 0)
		if item.ClientType == nil {
			panic("redis ClientType is empty")
		}
		opts = append(opts, redis.WithClientType(cast.ToString(item.ClientType)))
		if item.Hosts != nil {
			opts = append(opts, redis.WithHosts(*item.Hosts))
		}
		if item.Password != nil {
			opts = append(opts, redis.WithPassword(cast.ToString(item.Password)))
		}
		if item.Database != nil {
			opts = append(opts, redis.WithDatabase(cast.ToInt(item.Database)))
		}
		if item.PoolSize != nil {
			opts = append(opts, redis.WithPoolSize(cast.ToInt(item.PoolSize)))
		}
		opts = append(opts, redis.WithEnableMetric(cast.ToBool(item.EnableMetric)))
		opts = append(opts, redis.WithEnableTrace(cast.ToBool(item.EnableTrace)))
		if item.Addr != nil {
			opts = append(opts, redis.WithAddr(cast.ToString(item.Addr)))
		}
		if item.Addrs != nil {
			opts = append(opts, redis.WithAddrs(*item.Addrs))
		}
		if item.MasterName != nil {
			opts = append(opts, redis.WithMasterName(cast.ToString(item.MasterName)))
		}
		conn, err := redis.NewRedisClient(opts...)
		if err != nil {
			panic(err)
		}
		node := cast.ToString(item.LinkNode)
		links[node] = conn
	}

	var proxyConfigs []ProxyRedisConf
	if err := initial.Config.BindStruct("ProxyRedis", &proxyConfigs); err != nil {
		panic(err)
	}
	for _, item := range proxyConfigs {
		proxyPool := proxy.NewRedis()
		if item.LinkNode == nil {
			panic("redis LinkNode is empty")
		}
		node := cast.ToString(item.LinkNode)
		proxyPool.Store(links[node])
		if item.Name == nil {
			panic("redis Name is empty")
		}
		initial.Store.StoreRedis(cast.ToString(item.Name), proxyPool)
	}
	return initial
}
