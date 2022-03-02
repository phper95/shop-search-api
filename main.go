package main

import (
	"fmt"
	"shop-search-api/config"
	"shop-search-api/internal/pkg/logger"
	"shop-search-api/internal/pkg/timeutil"
)

func init() {
	config.LoadConfig()
}
func InitLog() {
	// 初始化 access logger
	accessLogger, err := logger.NewLogger(
		logger.WithDisableConsole(),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileRotationP(config.Cfg.App.AccessLogPath),
	)
	if err != nil {
		panic(err)
	}

	// 初始化 app logger
	appLogger, err := logger.NewLogger(
		logger.WithDisableConsole(),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileRotationP(config.Cfg.App.AppLogPath),
	)
	defer func() {
		_ = accessLogger.Sync()
		_ = appLogger.Sync()
	}()
}
func main() {

	fmt.Println(config.Cfg)

}
