package main

import (
	"context"
	"fmt"
	"gitee.com/phper95/pkg/cache"
	"gitee.com/phper95/pkg/db"
	"gitee.com/phper95/pkg/es"
	"gitee.com/phper95/pkg/logger"
	"gitee.com/phper95/pkg/shutdown"
	"gitee.com/phper95/pkg/timeutil"
	"gitee.com/phper95/pkg/trace"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
	"net/http"
	"shop-search-api/config"
	"shop-search-api/internal/server/api"
	"time"
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
	logger.Warn("mysqlCfg", zap.Any("", mysqlCfg))
	err := db.InitMysqlClient(db.DefaultClient, mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.DBName)
	if err != nil {
		logger.Error("mysql init error", zap.Error(err))
		panic("initMysqlClient error")
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
		panic("initRedisClient error")
	}
}

func initESClient() {
	ESCfg := config.Cfg.Elasticsearch
	err := es.InitClientWithOptions(es.DefaultClient, ESCfg.Host,
		ESCfg.User,
		ESCfg.Password,
		es.WithScheme("https"))
	if err != nil {
		logger.Error("InitClientWithOptions error", zap.Error(err), zap.String("client", es.DefaultClient))
		panic(err)
	}
}

func initMongoClient() {
	// TO DO ...
}
func main() {
	router := api.InitRouter()
	listenAddr := fmt.Sprintf(":%d", config.Cfg.App.HttpPort)
	logger.Warn("start http server listening %s", zap.String("listenAddr", listenAddr))
	server := &http.Server{
		Addr:           listenAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Error("http server start error", zap.Error(err))
		}
	}()

	//优雅关闭
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				logger.Error("http server shutdown err", zap.Error(err))
			}

		},
	)
}
