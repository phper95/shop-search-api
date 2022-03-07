package main

import (
	"fmt"
	"gitee.com/phper95/pkg/cache"
	"gitee.com/phper95/pkg/db"
	"gitee.com/phper95/pkg/logger"
	"gitee.com/phper95/pkg/timeutil"
	"gitee.com/phper95/pkg/trace"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"net/http"
	"shop-search-api/config"
	"shop-search-api/internal/server/api"
)

func init() {
	config.LoadConfig()
	InitLog()
	initMysqlClient()
	initRedisClient()
	initMongoClient()
	initESClient()
}
func InitLog() {
	// 初始化 logger
	logger.InitLogger(
		logger.WithDisableConsole(),
		logger.WithTimeLayout(timeutil.CSTLayout),
		logger.WithFileRotationP(config.Cfg.App.AppLogPath),
	)
	defer func() {
		logger.Sync()
	}()
}
func initMysqlClient() {
	mysqlCfg := config.Cfg.Mysql
	err := db.InitMysqlClient(config.DefaultMysqlClient, mysqlCfg.User,
		mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.DBName,
		db.WithMaxOpenConn(mysqlCfg.MaxOpenConn),
		db.WithMaxIdleConn(mysqlCfg.MaxOpenConn),
		db.WithConnMaxLifeSecond(mysqlCfg.ConnMaxLifeSecond))
	if err != nil {
		logger.Error("mysql init error", zap.Error(err))
	}
}
func initRedisClient() {
	redisCfg := config.Cfg.Redis
	opt := redis.Options{
		Addr:         redisCfg.Host,
		DB:           redisCfg.DB,
		MaxRetries:   redisCfg.MaxRetries,
		PoolSize:     redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConn,
	}
	redisTrace := trace.Cache{
		Name:                  "redis",
		SlowLoggerMillisecond: 500,
		Logger:                logger.GetLogger(),
		AlwaysTrace:           config.Cfg.App.RunMode == config.RunModeDev,
	}
	err := cache.InitRedis(config.DefaultRedisClient, &opt, &redisTrace)
	if err != nil {
		logger.Error("redis init error", zap.Error(err))
	}
}

func initESClient() {
	// TO DO ...
}

func initMongoClient() {
	// TO DO ...
}
func main() {
	router := api.InitRouter()
	listenAddr := fmt.Sprintf(":%d", config.Cfg.App.HttpPort)
	server := &http.Server{
		Addr:           listenAddr,
		Handler:        router,
		ReadTimeout:    config.Cfg.App.ReadTimeout,
		WriteTimeout:   config.Cfg.App.WriteTimeout,
		MaxHeaderBytes: 1 << 20, //2^20,1MB
	}
	logger.Warn("start http server listening %s", zap.String("listenAddr", listenAddr))
	server.ListenAndServe()
}
