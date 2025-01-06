package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// InitLogger 初始化日志管理器，输出到控制台
func InitLogger(level zapcore.Level) {
	once.Do(func() {
		// 配置日志编码器
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.TimeKey = "time"
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		encoderConfig.LevelKey = "level"
		encoderConfig.MessageKey = "message"
		encoderConfig.CallerKey = "caller"
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

		// 创建日志核心，输出到控制台
		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),                                // 控制台编码器
			zapcore.Lock(zapcore.AddSync(zapcore.AddSync(zapcore.Lock(os.Stdout)))), // 输出到标准输出
			level, // 日志级别
		)

		// 创建全局 logger
		logger = zap.New(core, zap.AddCaller())
	})
}

// GetLogger 获取全局 logger
func GetLogger() *zap.Logger {
	if logger == nil {
		panic("Logger is not initialized. Call InitLogger first.")
	}
	return logger
}
