package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"time"

	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hugplus/go-walker/common/utils"
	"github.com/hugplus/go-walker/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

var (
	Cfg    config.AppCfg
	Log    *zap.Logger
	Redis  *redis.Client
	dbs    map[string]*gorm.DB
	lock   sync.RWMutex
	engine http.Handler
	//AppRouters []func() // app路由
)

func GetEngine() http.Handler {
	return engine
}

func SetEngine(aEngine http.Handler) {
	engine = aEngine
}

func Init() {
	logInit()
	redisInit()
	dbInit()
}

func Run() {
	if Cfg.Server.Mode == ModeProd.String() {
		gin.SetMode(Cfg.Server.Mode)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", Cfg.Server.Host, Cfg.Server.Port),
		Handler: GetEngine(),
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen: ", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Printf("%s Shutdown Server ... \r\n", time.Now())

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func logInit() {
	//初始化日志
	Log = zapInit()
	zap.ReplaceGlobals(Log)
}

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func zapInit() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(Cfg.Logger.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", Cfg.Logger.Director)
		_ = os.Mkdir(Cfg.Logger.Director, os.ModePerm)
	}

	cores := Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if Cfg.Logger.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func redisInit() {
	redisCfg := Cfg.Cache
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		Log.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		Log.Info("redis connect ping response:", zap.String("pong", pong))
		Redis = rdb
	}
}
