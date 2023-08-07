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
	"github.com/hugplus/go-walker/common/middleware"
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
	lock   sync.RWMutex
	engine http.Handler
	dbs    = make(map[string]*gorm.DB, 0)
)

func GetEngine() http.Handler {
	return engine
}

func SetEngine(aEngine http.Handler) {
	fmt.Println("init engine")
	engine = aEngine
}

func GetGinEngine() *gin.Engine {
	var r *gin.Engine
	lock.RLock()
	defer lock.RUnlock()
	if engine == nil {
		engine = gin.New()
	}
	switch engine.(type) {
	case *gin.Engine:
		r = engine.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}
	return r
}

func Init() {
	logInit()
	redisInit()
	dbInit()
}

func Run(appRs *[]func()) {
	if Cfg.Server.Mode == ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}

	initRouter()

	//初始化路由
	for _, f := range *appRs {
		f()
	}

	//服务启动参数
	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", Cfg.Server.Host, Cfg.Server.Port),
		Handler:        GetEngine(),
		ReadTimeout:    time.Duration(Cfg.Server.GetReadTimeout()),
		WriteTimeout:   time.Duration(Cfg.Server.GetWriteTimeout()),
		MaxHeaderBytes: 1 << 20,
	}

	//启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: ", err)
		}
	}()
	fmt.Printf("Server started Listen %s:%d \n", Cfg.Server.Host, Cfg.Server.Port)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
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

func initRouter() {
	//初始化gin
	r := GetGinEngine()
	middleware.InitMiddleware(r)
}
