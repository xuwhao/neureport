package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	StudentID string
	Password  string
	Info      map[string]string
}

var instance *Config

func GetInstance(fpath string, fname string) *Config {

	fmt.Println(fpath, fname)
	if instance == nil {
		config := viper.New()
		config.AddConfigPath(fpath)
		config.SetConfigName(fname)
		config.SetConfigType("json")
		if err := config.ReadInConfig(); err != nil {
			panic(err)
		}
		config.WatchConfig()
		config.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			if err := config.ReadInConfig(); err != nil {
				panic(err)
			}
		})
		// fmt.Println(config.GetString("appId"))
		// fmt.Println(config.GetString("secret"))
		// fmt.Println(config.GetString("host.address"))
		// fmt.Println(config.GetString("host.port"))

		//直接反序列化为Struct
		var configjson Config
		if err := config.Unmarshal(&configjson); err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		instance = &configjson
	}

	return instance

	// fmt.Println(configjson.Host)
}
