package filter

import (
	"context"

	"user/pkg/config"
	pkgLogger "user/pkg/logger"

	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
)

func init() {
	extension.SetFilter(config.LogTraceFilterKey, NewLogTraceFilter)
}

func NewLogTraceFilter() filter.Filter {
	return &LogTraceFilter{}
}

type LogTraceFilter struct {
}

func (f *LogTraceFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	// 使用context-aware方式注入trace信息，获取增强的context
	ctxWithLogger := pkgLogger.InjectTrace(ctx)
	// 使用带有logger的context继续调用链
	return invoker.Invoke(ctxWithLogger, invocation)
}
func (f *LogTraceFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	return result
}
