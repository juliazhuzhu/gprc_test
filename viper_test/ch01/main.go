package main

import (
	"fmt"
	"github.com/spf13/viper"
)
type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int `mapstructure:"port"`
}

func main(){
	v := viper.New()
	v.SetConfigFile("viper_test/ch01/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	servConfig := ServerConfig{}
	if err := v.Unmarshal(&servConfig); err != nil {
		panic(err)
	}
	fmt.Println(servConfig)
	fmt.Println(v.Get("name"))

}
