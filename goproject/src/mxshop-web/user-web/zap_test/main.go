package main

import (
	//"time"

	//"net/url"

	"go.uber.org/zap"
)

func main() {
	//logger, _ := zap.NewProduction()//生产环境用法
	logger,_:=zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any；把内存数据同步硬盘
	//sugar := logger.Sugar()
	url := "https://imooc.com"
	logger.Info("failed to fetch URL",
		zap.String("url",url),
		zap.Int("nums",3),
	)
	// sugar.Infow("failed to fetch URL",
	// 	// Structured context as loosely typed key-value pairs.
	// 	"url", url,
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )
	// sugar.Infof("Failed to fetch URL: %s", url)
}