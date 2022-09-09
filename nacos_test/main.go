package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"hexmeet.com/grpctest/nacos_test/config"
	"time"
)


func main(){

	sc := []constant.ServerConfig{
		{
			IpAddr: "172.20.0.204",
			Port: 8848,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         "43be1373-70db-4890-b2b0-6ea409c3b55d", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// Create config client for dynamic configuration
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "DEV"})

	if err != nil {
		panic(err)
	}

	fmt.Println(content)
	serCfg := config.ServerConfig{}
	_ = json.Unmarshal([]byte(content), &serCfg)
	fmt.Println(serCfg)
	configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "DEV",
		OnChange: func (namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	time.Sleep(3600*time.Second)
}
