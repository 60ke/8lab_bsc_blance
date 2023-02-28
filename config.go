package main

import (
	"github.com/spf13/viper"
)

var Conf = &Config{}

type Config struct {
	DingUrl  string `toml:"ding-url" comment:"钉钉机器人接口"`
	Port     int    `toml:"port" comment:"服务端口默认为8546"`
	LogName  string `toml:"log-name" comment:"日志路径,默认为bsc_balance.log"`
	LogLevel string `toml:"log-level" comment:"日志等级int类型默认为info; debug || info || warn || error"`
	Token    string `toml:"token" comment:"用于客户端token校验"`
}

func init() {
	viper.SetDefault("ding-prefix", "bsc节点")
	viper.SetDefault("ding-url", "https://oapi.dingtalk.com/robot/send?access_token=")
	viper.SetDefault("ding-token", "48679669f51e7aa89c4544b17bfc4c1e4a565b975d067cf3fa420b0ae0c255ef")
	viper.SetDefault("port", 8546)
	viper.SetDefault("log-name", "bsc_balance.log")
	viper.SetDefault("log-level", "info")
	viper.SetDefault("token", "3D3781351A3EE9E4")
	viper.SetDefault("docker-name", "trust_zkbsc")

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	LogInit(viper.GetString("log-level"), viper.GetString("log-name"))
	if err != nil {
		Logger.Sugar().Error(err.Error())
	}
	Logger.Sugar().Info("bsc balance start")
}
