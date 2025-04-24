package logger

import (
	"context"

	"github.com/dubbogo/gost/log/logger"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var baseLogger *zap.Logger

type Logger struct {
	baseLogger *zap.Logger
}

func init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	baseLogger = l
	logger.SetLogger(&Logger{baseLogger: l})
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
