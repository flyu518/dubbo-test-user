package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

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
// 因为通过 context 动态设置，不要修改本地的 logger 的值，仅仅获取后设置到全局的 logger 里面
func InjectTrace(ctx context.Context) {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		return
	}

	dubboLogger := GetDubboLogger()
	dubboLogger.ZapLogger = dubboLogger.ZapLogger.With(
		zap.String("traceId", spanCtx.TraceID().String()),
		zap.String("spanId", spanCtx.SpanID().String()),
	)

	dubboLogger.Logger = dubboLogger.ZapLogger.Sugar()

	SetGlobalLogger(dubboLogger)
}
