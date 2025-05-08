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
	pkgLogger.InjectTrace(ctx)
	return invoker.Invoke(ctx, invocation)
}
func (f *LogTraceFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	return result
}
