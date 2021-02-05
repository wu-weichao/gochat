package config

import (
	"github.com/spf13/viper"
	"sync"
)

// app config
type AppConfig struct {
	RunMode string `mapstructure:"run_mode"`
}

var App = &AppConfig{}

// tcp
type TcpConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

var Tcp = &TcpConfig{}

// websocket
type WebSocketConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

var WebSocket = &WebSocketConfig{}

var once sync.Once

func init() {
	once.Do(func() {
		viper.AddConfigPath("./config")
		viper.SetConfigType("yaml")
		// load app.ini
		viper.SetConfigName("app")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		// load connect.ini
		viper.SetConfigName("connect")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		// format config
		viper.UnmarshalKey("app", &App)
		viper.UnmarshalKey("tcp", &Tcp)
		viper.UnmarshalKey("websocket", &WebSocket)
	})
}
