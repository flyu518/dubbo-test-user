package logger

import (
	"context"
	"time"

	"github.com/dubbogo/gost/log/logger"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	baseLogger *zap.Logger
	logConfig  zap.Config
	callerSkip int = 1
)

type Logger struct {
	baseLogger *zap.Logger
}

func init() {
	// 创建自定义的生产环境配置
	logConfig = zap.NewProductionConfig()

	// 自定义时间格式编码器，这里使用的是 ISO8601 格式
	// 你可以替换为其他格式，例如 RFC3339 或自定义格式
	logConfig.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}

	// 使用自定义配置创建日志记录器并设置CallerSkip
	l, err := logConfig.Build(zap.AddCallerSkip(callerSkip))
	if err != nil {
		panic(err)
	}
	baseLogger = l
	logger.SetLogger(&Logger{baseLogger: l})
}

func GetLogger() *Logger {
	return &Logger{baseLogger: baseLogger}
}

// SetCallerSkip 设置日志调用者跳过层级
// 用于控制日志输出中显示的代码位置
// 值越大，跳过的调用栈层级越多
func SetCallerSkip(skip int) {
	if skip < 0 {
		skip = 0
	}

	// 保存新的callerSkip值
	callerSkip = skip

	// 重新构建日志记录器，使用AddCallerSkip选项
	newLogger, err := logConfig.Build(zap.AddCallerSkip(callerSkip))
	if err != nil {
		baseLogger.Error("无法更新CallerSkip设置", zap.Error(err))
		return
	}

	// 替换全局日志记录器
	_ = baseLogger.Sync()
	baseLogger = newLogger
	logger.SetLogger(&Logger{baseLogger: newLogger})
}

// GetCallerSkip 获取当前的CallerSkip设置
func GetCallerSkip() int {
	return callerSkip
}

// SetLoggerLevel 设置日志级别
// level 可以是以下值之一："debug", "info", "warn", "error", "dpanic", "panic", "fatal"
func SetLoggerLevel(level string) error {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		return err
	}

	// 更新配置中的日志级别
	logConfig.Level.SetLevel(zapLevel)

	// 使用更新后的配置创建新的日志记录器，保持CallerSkip设置
	newLogger, err := logConfig.Build(zap.AddCallerSkip(callerSkip))
	if err != nil {
		return err
	}

	// 替换全局日志记录器
	// 确保先关闭旧的日志记录器，以释放资源
	_ = baseLogger.Sync()
	baseLogger = newLogger
	logger.SetLogger(&Logger{baseLogger: newLogger})

	return nil
}

// GetLoggerLevel 获取日志级别
func GetLoggerLevel() string {
	return logConfig.Level.String()
}

// InjectTraceToGlobal 将 traceId 注入到全局 logger 中
func InjectTraceToGlobal(ctx context.Context) {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		return
	}
	loggerWithTrace := baseLogger.With(
		zap.String("traceId", spanCtx.TraceID().String()),
		zap.String("spanId", spanCtx.SpanID().String()),
	)
	logger.SetLogger(&Logger{baseLogger: loggerWithTrace})
}

func (l *Logger) Debug(args ...interface{}) {
	l.baseLogger.Sugar().Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.baseLogger.Sugar().Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.baseLogger.Sugar().Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.baseLogger.Sugar().Error(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.baseLogger.Sugar().Debugf(format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.baseLogger.Sugar().Infof(format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.baseLogger.Sugar().Warnf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.baseLogger.Sugar().Errorf(format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.baseLogger.Sugar().Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.baseLogger.Sugar().Fatalf(format, args...)
}
