package initial

import (
	"time"

	"github.com/abulo/ratel/v3/stores/proxy"
	"github.com/abulo/ratel/v3/stores/sql"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type (
	SqlConf struct {
		// LinkNode = "sql.host1"

		Host         *string `json:"Host"`
		Port         *string `json:"Port"`
		Username     *string `json:"Username"`
		Password     *string `json:"Password"`
		Charset      *string `json:"Charset"`
		Database     *string `json:"Database"`
		ParseTime    *bool   `json:"ParseTime"`
		TimeZone     *string `json:"TimeZone"`
		SslMode      *bool   `json:"SslMode"`
		DialTimeOut  *string `json:"DialTimeOut"`
		ReadTimeOut  *string `json:"ReadTimeOut"`
		MaxIdleConns *int    `json:"MaxIdleConns"`
		MaxOpenConns *int    `json:"MaxOpenConns"`
		MaxLifetime  *int    `json:"MaxLifetime"`
		MaxIdleTime  *int    `json:"MaxIdleTime"`
		EnableMetric *bool   `json:"EnableMetric"`
		EnableTrace  *bool   `json:"EnableTrace"`
		EnableDebug  *bool   `json:"EnableDebug"`
		DriverName   *string `json:"DriverName"`
		LinkNode     *string `json:"LinkNode"`
	}
	ProxySqlConf struct {
		Name   *string      `json:"Name"`
		Master *StringSlice `json:"Master"`
		Slave  *StringSlice `json:"Slave"`
	}
)

// InitSql load mysql && returns an mysql instance.
func (initial *Initial) InitSql() *Initial {
	var sqlConfigs []SqlConf
	if err := initial.Config.BindStruct("Sql", &sqlConfigs); err != nil {
		panic(err)
	}
	links := make(map[string]*gorm.DB)
	for _, item := range sqlConfigs {
		opts := make([]sql.Option, 0)
		if Host := cast.ToString(item.Host); Host != "" {
			opts = append(opts, sql.WithHost(Host))
		}
		if Port := cast.ToString(item.Port); Port != "" {
			opts = append(opts, sql.WithPort(Port))
		}
		if Username := cast.ToString(item.Username); Username != "" {
			opts = append(opts, sql.WithUsername(Username))
		}
		if Password := cast.ToString(item.Password); Password != "" {
			opts = append(opts, sql.WithPassword(Password))
		}
		if Charset := cast.ToString(item.Charset); Charset != "" {
			opts = append(opts, sql.WithCharset(Charset))
		}
		if Database := cast.ToString(item.Database); Database != "" {
			opts = append(opts, sql.WithDatabase(Database))
		}
		if ParseTime := cast.ToBool(item.ParseTime); ParseTime {
			opts = append(opts, sql.WithParseTime(ParseTime))
		}
		if TimeZone := cast.ToString(item.TimeZone); TimeZone != "" {
			opts = append(opts, sql.WithTimeZone(TimeZone))
		}
		if SslMode := cast.ToBool(item.SslMode); SslMode {
			opts = append(opts, sql.WithSslMode(SslMode))
		}
		if DialTimeOut := cast.ToString(item.DialTimeOut); DialTimeOut != "" {
			opts = append(opts, sql.WithDialTimeOut(DialTimeOut))
		}
		if ReadTimeOut := cast.ToString(item.ReadTimeOut); ReadTimeOut != "" {
			opts = append(opts, sql.WithReadTimeOut(ReadTimeOut))
		}
		if MaxIdleConns := cast.ToInt(item.MaxIdleConns); MaxIdleConns > 0 {
			opts = append(opts, sql.WithMaxIdleConns(MaxIdleConns))
		}
		if MaxOpenConns := cast.ToInt(item.MaxOpenConns); MaxOpenConns > 0 {
			opts = append(opts, sql.WithMaxOpenConns(MaxOpenConns))
		}
		if MaxLifetime := cast.ToInt(item.MaxLifetime); MaxLifetime > 0 {
			opts = append(opts, sql.WithMaxLifetime(time.Duration(MaxLifetime)*time.Second))
		}
		if MaxIdleTime := cast.ToInt(item.MaxIdleTime); MaxIdleTime > 0 {
			opts = append(opts, sql.WithMaxIdleTime(time.Duration(MaxIdleTime)*time.Second))
		}
		opts = append(opts, sql.WithEnableMetric(cast.ToBool(item.EnableMetric)))
		opts = append(opts, sql.WithEnableTrace(cast.ToBool(item.EnableTrace)))
		if DriverName := cast.ToString(item.DriverName); DriverName != "" {
			opts = append(opts, sql.WithDriverName(DriverName))
		}
		if EnableDebug := cast.ToBool(item.EnableDebug); EnableDebug {
			opts = append(opts, sql.WithEnableDebug(EnableDebug))
		}
		if item.LinkNode == nil {
			panic("sql LinkNode is empty")
		}
		node := cast.ToString(item.LinkNode)
		client, err := sql.NewClient(opts...)
		if err != nil {
			panic(err)
		}
		link, err := client.SqlConn()
		if err != nil {
			panic(err)
		}
		links[node] = link
	}

	var proxyConfigs []ProxySqlConf
	if err := initial.Config.BindStruct("ProxySql", &proxyConfigs); err != nil {
		panic(err)
	}
	for _, item := range proxyConfigs {
		proxyPool := proxy.NewSQL()
		if item.Name == nil {
			panic("sql Name is empty")
		}
		if item.Master != nil {
			for _, v := range *item.Master {
				proxyPool.SetWrite(links[cast.ToString(v)])
			}
		}
		if item.Slave != nil {
			for _, v := range *item.Slave {
				proxyPool.SetRead(links[cast.ToString(v)])
			}
		}
		name := cast.ToString(item.Name)
		initial.Store.StoreSQL(name, proxyPool)
	}
	return initial
}
