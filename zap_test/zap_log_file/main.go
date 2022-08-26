package main

import (
	"go.uber.org/zap"
	"time"
)

func NewLogger() (*zap.Logger, error){
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = [] string {
		"./mytest.log",
		"stdout",
	}
	
	return cfg.Build()
}

func main()  {

	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}

	su := logger.Sugar()
	url := "https://imoc.com"
	defer su.Sync()
	su.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	su.Infof("Failed to fetch URL: %s", url)
}