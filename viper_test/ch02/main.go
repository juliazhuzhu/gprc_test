package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}

type ServerConfig struct {
	Name string `mapstructure:"name"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
}

func GenEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func main(){

	debug := GenEnvInfo("MXSHOP_DEBUG")
	var configFileName string
	if debug {
		configFileName = "viper_test/ch02/config_debug.yaml"
	}else {
		configFileName = "viper_test/ch02/config_pro.yaml"
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	servConfig := ServerConfig{}
	if err := v.Unmarshal(&servConfig); err != nil {
		panic(err)
	}
	fmt.Println(servConfig)
	fmt.Println(v.Get("name"))
	v.WatchConfig()
	v.OnConfigChange(func (e fsnotify.Event){
		fmt.Println("config file changed", e.Name)
		_ = v.ReadInConfig()
		servConfig := ServerConfig{}
		if err := v.Unmarshal(&servConfig); err != nil {
			panic(err)
		}
		fmt.Println(servConfig)
		fmt.Println(v.Get("name"))
	})

	time.Sleep(time.Second*300)
}

