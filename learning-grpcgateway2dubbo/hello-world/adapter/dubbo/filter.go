package dubbo

import (
    "context"
    "dubbo.apache.org/dubbo-go/v3/common/extension"
    "dubbo.apache.org/dubbo-go/v3/filter"
    "dubbo.apache.org/dubbo-go/v3/protocol"
    inv "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
    "fmt"
)

func init() {
    extension.SetFilter("bos", newFilter)
}

type Filter struct {
    name string
}

func newFilter() filter.Filter {
    f := &Filter{
        name: "bos",
    }
    return f
}

func (f Filter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
    rpcInvocation := invocation.(*inv.RPCInvocation)
    fmt.Println("rpcInvocation: {}", rpcInvocation)
    return invoker.Invoke(ctx, rpcInvocation)
}

func (f Filter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
    return result
}




