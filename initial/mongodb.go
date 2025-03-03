package initial

import (
	"time"

	"github.com/abulo/ratel/v3/stores/mongodb"
	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/spf13/cast"
)

type (
	MongoDBConf struct {
		URI             *string `json:"URI"`
		MaxConnIdleTime *int    `json:"MaxConnIdleTime"`
		MaxPoolSize     *int    `json:"MaxPoolSize"`
		MinPoolSize     *int    `json:"MinPoolSize"`
		EnableMetric    *bool   `json:"EnableMetric"`
		EnableTrace     *bool   `json:"EnableTrace"`
		LinkNode        *string `json:"LinkNode"`
	}
	ProxyMongoDBConf struct {
		Name     *string `json:"Name"`
		LinkNode *string `json:"LinkNode"`
	}
)

// InitMongoDB load mongodb && returns an mongodb instance.
func (initial *Initial) InitMongoDB() *Initial {
	var mongodbConfigs []MongoDBConf
	if err := initial.Config.BindStruct("MongoDB", &mongodbConfigs); err != nil {
		panic(err)
	}
	links := make(map[string]*mongodb.MongoDB)
	for _, item := range mongodbConfigs {
		opts := make([]mongodb.Option, 0)
		if item.URI == nil {
			panic("mongodb uri is empty")
		}
		uri := cast.ToString(item.URI)
		if item.MaxConnIdleTime == nil {
			panic("mongodb MaxConnIdleTime is empty")
		}
		opts = append(opts, mongodb.WithMaxConnIdleTime(cast.ToDuration(item.MaxConnIdleTime)*time.Minute))
		if item.MaxPoolSize == nil {
			panic("mongodb MaxPoolSize is empty")
		}
		opts = append(opts, mongodb.WithMaxPoolSize(cast.ToUint64(item.MaxPoolSize)))
		if item.MinPoolSize == nil {
			panic("mongodb MinPoolSize is empty")
		}
		opts = append(opts, mongodb.WithMinPoolSize(cast.ToUint64(item.MinPoolSize)))
		client, err := mongodb.NewMongoDBClient(uri, opts...)
		if err != nil {
			panic(err)
		}
		client.SetEnableMetric(cast.ToBool(item.EnableMetric))
		client.SetEnableTrace(cast.ToBool(item.EnableTrace))
		if item.LinkNode == nil {
			panic("mongodb LinkNode is empty")
		}
		node := cast.ToString(item.LinkNode)
		links[node] = client
	}

	var proxyConfigs []ProxyMongoDBConf
	if err := initial.Config.BindStruct("ProxyMongoDB", &proxyConfigs); err != nil {
		panic(err)
	}
	for _, val := range proxyConfigs {
		proxyPool := proxy.NewMongoDB()
		if val.LinkNode == nil {
			panic("ProxyMongoDB LinkNode is empty")
		}
		linkNode := cast.ToString(val.LinkNode)
		proxyPool.Store(links[linkNode])
		if val.Name == nil {
			panic("ProxyMongoDB Name is empty")
		}
		node := cast.ToString(val.Name)
		initial.Store.StoreMongoDB(node, proxyPool)
	}
	return initial
}
