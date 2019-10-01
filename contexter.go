package awsctx

import "context"

type Contexter interface {
	WrapContext(ctx context.Context, request *AwsRequest, inner func(ctx context.Context))
}

type ContexterFunc func(ctx context.Context, request *AwsRequest, inner func(ctx context.Context))

func (cf ContexterFunc) WrapContext(ctx context.Context, request *AwsRequest, inner func(ctx context.Context)) {
	cf(ctx, request, inner)
}

var NoopContexter = ContexterFunc(func(ctx context.Context, request *AwsRequest, inner func(ctx context.Context)) {
	inner(ctx)
})
