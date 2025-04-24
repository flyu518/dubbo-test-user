package filter

import (
	"context"
	userLogger "user/pkg/logger"

	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
)

func init() {
	extension.SetFilter("myCustomFilter", NewMyClientFilter)
}

func NewMyClientFilter() filter.Filter {
	return &MyClientFilter{}
}

type MyClientFilter struct {
}

func (f *MyClientFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	userLogger.InjectTraceToGlobal(ctx)
	return invoker.Invoke(ctx, invocation)
}
func (f *MyClientFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	return result
}
