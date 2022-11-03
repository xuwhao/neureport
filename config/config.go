package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	StudentID string `json:"StudentID"`
	Password  string `json:"password"`
	Info      struct {
		Token                           string `json:"_token"`
		Jibenxinxi_shifoubenrenshangbao string `json:"jibenxinxi_shifoubenrenshangbao"`
		Profile                         struct {
			Xuegonghao  string `json:"xuegonghao"`
			Xingming    string `json:"xingming"`
			Suoshubanji string `json:"suoshubanji"`
		} `json:"profile"`
		Jiankangxinxi_muqianshentizhuangkuang    string        `json:"jiankangxinxi_muqianshentizhuangkuang"`
		Xingchengxinxi_weizhishifouyoubianhua    string        `json:"xingchengxinxi_weizhishifouyoubianhua"`
		Cross_city                               string        `json:"cross_city"`
		Qitashixiang_qitaxuyaoshuomingdeshixiang string        `json:"qitashixiang_qitaxuyaoshuomingdeshixiang"`
		Credits                                  string        `json:"credits"`
		Bmap_position                            string        `json:"bmap_position"`
		Bmap_position_latitude                   string        `json:"bmap_position_latitude"`
		Bmap_position_longitude                  string        `json:"bmap_position_longitude"`
		Bmap_position_address                    string        `json:"bmap_position_address"`
		Bmap_position_status                     string        `json:"bmap_position_status"`
		ProvinceCode                             string        `json:"ProvinceCode"`
		CityCode                                 string        `json:"CityCode"`
		Travels                                  []interface{} `json:"travels"`
	} `json:"info"`
}

// type Config struct {
// 	StudentID string `json:"StudentID"`
// 	Password string `json:"password"`
// 	Info Info `json:"info"`
// }
// type Info struct {
// 	_token string
// 	JibenxinxiShifoubenrenshangbao string `json:"jibenxinxi_shifoubenrenshangbao"`
// 	Profile Profile `json:"profile"`
// 	JiankangxinxiMuqianshentizhuangkuang string `json:"jiankangxinxi_muqianshentizhuangkuang"`
// 	XingchengxinxiWeizhishifouyoubianhua string `json:"xingchengxinxi_weizhishifouyoubianhua"`
// 	CrossCity string `json:"cross_city"`
// 	QitashixiangQitaxuyaoshuomingdeshixiang string `json:"qitashixiang_qitaxuyaoshuomingdeshixiang"`
// }

// type Profile struct {
// 	Xuegonghao string `json:"xuegonghao"`
// 	Xingming string `json:"xingming"`
// 	Suoshubanji string `json:"suoshubanji"`
// }

func (config *Config) SetToken(token string) {
	config.Info.Token = token
}

func (config *Config) GetToken() string {
	return config.Info.Token
}

// func (info *Info) SetToken(token string) {
//     info._token = token
// }

// func (info *Info) GetToken() string{
//     return info._token
// }

var instance *Config

func GetInstance(fpath string, fname string) *Config {

	// fmt.Println(fpath, fname)
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
