package core

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormZapLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
	level         logger.LogLevel
}

func NewGormZapLog(log *zap.Logger, level logger.LogLevel, slow time.Duration) GormZapLogger {
	return GormZapLogger{
		ZapLogger:     log,
		SlowThreshold: slow * time.Millisecond,
		level:         level,
	}
}

// LogMode 实现 gormlogger.Interface 的 LogMode 方法
func (l GormZapLogger) LogMode(level logger.LogLevel) logger.Interface {
	return GormZapLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
		level:         level,
	}
}

// Info 实现 gormlogger.Interface 的 Info 方法
func (l GormZapLogger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Debugf(str, args...)
}

// Warn 实现 gormlogger.Interface 的 Warn 方法
func (l GormZapLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Warnf(str, args...)
}

// Error 实现 gormlogger.Interface 的 Error 方法
func (l GormZapLogger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Errorf(str, args...)
}

// Trace 实现 gormlogger.Interface 的 Trace 方法
func (l GormZapLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	if l.level == logger.Silent {
		return
	}
	// 获取运行时间
	elapsed := time.Since(begin)

	// 获取 SQL 请求和返回条数
	sql, rows := fc()
	logFields := []zap.Field{zap.String("sql", sql)}

	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Error("DB ErrRecordNotFound", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, zap.Error(err))
			l.logger().Error("DB Error", logFields...)
		}
		return
	}

	if l.level == logger.Error {
		return
	}

	logFields = append(logFields, zap.String("time", fmt.Sprintf("%v", elapsed)), zap.Int64("rows", rows))
	// 通用字段
	// logFields := []zap.Field{
	// 	zap.String("sql", sql),
	// 	zap.String("time", fmt.Sprintf("%v", elapsed)),
	// 	zap.Int64("rows", rows),
	// }

	// 慢查询日志
	if l.SlowThreshold > 0 && elapsed > l.SlowThreshold {
		l.logger().Warn("DB Slow Log", logFields...)
		return
	}

	if l.level == logger.Warn {
		return
	}

	// 记录所有 SQL 请求
	l.logger().Info("sql trace", logFields...)
}

// logger 内用的辅助方法，确保 Zap 内置信息 Caller 的准确性（如 paginator/paginator.go:148）
func (l GormZapLogger) logger() *zap.Logger {

	// 跳过 gorm 内置的调用
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)

	// 减去一次封装，以及一次在 logger 初始化里添加 zap.AddCallerSkip(1)
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			// 返回一个附带跳过行号的新的 zap logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}
