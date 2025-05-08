package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// contextKey 是用于存储context值的key类型
type contextKey string

// loggerContextKey 是存储logger的context key
const loggerContextKey = contextKey("dubbo-logger")

// InjectServiceName 注入服务名
// 正常是在初始化的时候调用，可以修改本地和全局的 logger
func InjectServiceName(name string) {
	dubboLogger := GetDubboLogger()
	dubboLogger.ZapLogger = dubboLogger.ZapLogger.With(
		zap.String("service", name),
	)
	dubboLogger.Logger = dubboLogger.ZapLogger.Sugar()

	SetGlobalLogger(dubboLogger)
}

// InjectTrace 注入 traceId、spanId 等
// 使用context-aware的方式处理trace信息，避免并发问题
// 返回包含logger的新context
func InjectTrace(ctx context.Context) context.Context {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		return ctx
	}

	// 获取原始logger
	originalLogger := GetDubboLogger()

	// 创建一个新的logger，添加trace相关字段
	traceLogger := &DubboLogger{
		ZapLogger: originalLogger.ZapLogger.With(
			zap.String("traceId", spanCtx.TraceID().String()),
			zap.String("spanId", spanCtx.SpanID().String()),
		),
		dynamicLevel: originalLogger.dynamicLevel,
	}

	// 设置Sugar Logger
	traceLogger.Logger = traceLogger.ZapLogger.Sugar()

	// 将logger存储在context中
	return context.WithValue(ctx, loggerContextKey, traceLogger)
}

// FromContext 从context中获取logger
func FromContext(ctx context.Context) *DubboLogger {
	if ctx == nil {
		return GetDubboLogger()
	}

	if logger, ok := ctx.Value(loggerContextKey).(*DubboLogger); ok {
		return logger
	}

	return GetDubboLogger()
}
