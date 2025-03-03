package initial

import (
	"github.com/abulo/ratel/v3/config"
)

// GetEnvironment ...
func (initial *Initial) GetEnvironment(dir, key string) string {
	envConfig := config.New()
	if err := envConfig.LoadDir(dir); err != nil {
		panic(err)
	}
	envConfig.WatchConfig()
	return envConfig.String(key)
}

// InitConfig set app config toml files
func (initial *Initial) InitConfig(dir string) *Initial {
	Config := config.New()
	if err := Config.LoadDir(dir); err != nil {
		panic(err)
	}
	Config.WatchConfig()
	initial.Config = Config
	return initial
}
