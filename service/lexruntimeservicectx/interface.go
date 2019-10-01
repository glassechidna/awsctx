// Code generated by internal/generate/main.go. DO NOT EDIT.

package lexruntimeservicectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice/lexruntimeserviceiface"
	"github.com/glassechidna/awsctx"
)

type LexRuntimeService interface {
	DeleteSessionWithContext(ctx context.Context, input *lexruntimeservice.DeleteSessionInput, opts ...request.Option) (*lexruntimeservice.DeleteSessionOutput, error)
	GetSessionWithContext(ctx context.Context, input *lexruntimeservice.GetSessionInput, opts ...request.Option) (*lexruntimeservice.GetSessionOutput, error)
	PostContentWithContext(ctx context.Context, input *lexruntimeservice.PostContentInput, opts ...request.Option) (*lexruntimeservice.PostContentOutput, error)
	PostTextWithContext(ctx context.Context, input *lexruntimeservice.PostTextInput, opts ...request.Option) (*lexruntimeservice.PostTextOutput, error)
	PutSessionWithContext(ctx context.Context, input *lexruntimeservice.PutSessionInput, opts ...request.Option) (*lexruntimeservice.PutSessionOutput, error)
}

type Client struct {
	lexruntimeserviceiface.LexRuntimeServiceAPI
	Contexter awsctx.Contexter
}

func New(base lexruntimeserviceiface.LexRuntimeServiceAPI, ctxer awsctx.Contexter) LexRuntimeService {
	return &Client{
		LexRuntimeServiceAPI: base,
		Contexter: ctxer,
	}
}

var _ LexRuntimeService = (*lexruntimeservice.LexRuntimeService)(nil)
var _ LexRuntimeService = (*Client)(nil)

func (c *Client) DeleteSessionWithContext(ctx context.Context, input *lexruntimeservice.DeleteSessionInput, opts ...request.Option) (*lexruntimeservice.DeleteSessionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "lexruntimeservice",
		Action:  "DeleteSessionWithContext",
		Input:   input,
		Output:  (*lexruntimeservice.DeleteSessionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.LexRuntimeServiceAPI.DeleteSessionWithContext(ctx, input, opts...)
	})

	return req.Output.(*lexruntimeservice.DeleteSessionOutput), req.Error
}

func (c *Client) GetSessionWithContext(ctx context.Context, input *lexruntimeservice.GetSessionInput, opts ...request.Option) (*lexruntimeservice.GetSessionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "lexruntimeservice",
		Action:  "GetSessionWithContext",
		Input:   input,
		Output:  (*lexruntimeservice.GetSessionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.LexRuntimeServiceAPI.GetSessionWithContext(ctx, input, opts...)
	})

	return req.Output.(*lexruntimeservice.GetSessionOutput), req.Error
}

func (c *Client) PostContentWithContext(ctx context.Context, input *lexruntimeservice.PostContentInput, opts ...request.Option) (*lexruntimeservice.PostContentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "lexruntimeservice",
		Action:  "PostContentWithContext",
		Input:   input,
		Output:  (*lexruntimeservice.PostContentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.LexRuntimeServiceAPI.PostContentWithContext(ctx, input, opts...)
	})

	return req.Output.(*lexruntimeservice.PostContentOutput), req.Error
}

func (c *Client) PostTextWithContext(ctx context.Context, input *lexruntimeservice.PostTextInput, opts ...request.Option) (*lexruntimeservice.PostTextOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "lexruntimeservice",
		Action:  "PostTextWithContext",
		Input:   input,
		Output:  (*lexruntimeservice.PostTextOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.LexRuntimeServiceAPI.PostTextWithContext(ctx, input, opts...)
	})

	return req.Output.(*lexruntimeservice.PostTextOutput), req.Error
}

func (c *Client) PutSessionWithContext(ctx context.Context, input *lexruntimeservice.PutSessionInput, opts ...request.Option) (*lexruntimeservice.PutSessionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "lexruntimeservice",
		Action:  "PutSessionWithContext",
		Input:   input,
		Output:  (*lexruntimeservice.PutSessionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.LexRuntimeServiceAPI.PutSessionWithContext(ctx, input, opts...)
	})

	return req.Output.(*lexruntimeservice.PutSessionOutput), req.Error
}
