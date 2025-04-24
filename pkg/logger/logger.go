package logger

import (
	"context"
	"user/pkg/global"

	"github.com/dubbogo/gost/log/logger"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var baseLogger *zap.Logger

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	baseLogger = l
	logger.SetLogger(Sugar())
	global.Log = func() logger.Logger {
		return logger.GetLogger()
	}
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
	logger.SetLogger(loggerWithTrace.Sugar())
	global.Log = func() logger.Logger {
		return logger.GetLogger()
	}
}

// WithContext 返回带 traceId 的日志对象
func WithContext(ctx context.Context) *zap.SugaredLogger {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		return baseLogger.Sugar()
	}
	return baseLogger.Sugar().With(
		"traceId", spanCtx.TraceID().String(),
		"spanId", spanCtx.SpanID().String(),
	)
}

func Sugar() *zap.SugaredLogger {
	return baseLogger.Sugar()
}

func Info(args ...interface{}) {
	Sugar().Info(args...)
}

func Infof(format string, args ...interface{}) {
	Sugar().Infof(format, args...)
}

func Error(args ...interface{}) {
	Sugar().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	Sugar().Errorf(format, args...)
}

func Debug(args ...interface{}) {
	Sugar().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	Sugar().Debugf(format, args...)
}

func Warn(args ...interface{}) {
	Sugar().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	Sugar().Warnf(format, args...)
}
