// Code generated by internal/generate/main.go. DO NOT EDIT.

package mediatailorctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/mediatailor/mediatailoriface"
	"github.com/glassechidna/awsctx"
)

type MediaTailor interface {
	DeletePlaybackConfigurationWithContext(ctx context.Context, input *mediatailor.DeletePlaybackConfigurationInput, opts ...request.Option) (*mediatailor.DeletePlaybackConfigurationOutput, error)
	GetPlaybackConfigurationWithContext(ctx context.Context, input *mediatailor.GetPlaybackConfigurationInput, opts ...request.Option) (*mediatailor.GetPlaybackConfigurationOutput, error)
	ListPlaybackConfigurationsWithContext(ctx context.Context, input *mediatailor.ListPlaybackConfigurationsInput, opts ...request.Option) (*mediatailor.ListPlaybackConfigurationsOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *mediatailor.ListTagsForResourceInput, opts ...request.Option) (*mediatailor.ListTagsForResourceOutput, error)
	PutPlaybackConfigurationWithContext(ctx context.Context, input *mediatailor.PutPlaybackConfigurationInput, opts ...request.Option) (*mediatailor.PutPlaybackConfigurationOutput, error)
	TagResourceWithContext(ctx context.Context, input *mediatailor.TagResourceInput, opts ...request.Option) (*mediatailor.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *mediatailor.UntagResourceInput, opts ...request.Option) (*mediatailor.UntagResourceOutput, error)
}

type Client struct {
	mediatailoriface.MediaTailorAPI
	Contexter awsctx.Contexter
}

func New(base mediatailoriface.MediaTailorAPI, ctxer awsctx.Contexter) MediaTailor {
	return &Client{
		MediaTailorAPI: base,
		Contexter: ctxer,
	}
}

var _ MediaTailor = (*mediatailor.MediaTailor)(nil)
var _ MediaTailor = (*Client)(nil)

func (c *Client) DeletePlaybackConfigurationWithContext(ctx context.Context, input *mediatailor.DeletePlaybackConfigurationInput, opts ...request.Option) (*mediatailor.DeletePlaybackConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mediatailor",
		Action:  "DeletePlaybackConfigurationWithContext",
		Input:   input,
		Output:  (*mediatailor.DeletePlaybackConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MediaTailorAPI.DeletePlaybackConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*mediatailor.DeletePlaybackConfigurationOutput), req.Error
}

func (c *Client) GetPlaybackConfigurationWithContext(ctx context.Context, input *mediatailor.GetPlaybackConfigurationInput, opts ...request.Option) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mediatailor",
		Action:  "GetPlaybackConfigurationWithContext",
		Input:   input,
		Output:  (*mediatailor.GetPlaybackConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MediaTailorAPI.GetPlaybackConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*mediatailor.GetPlaybackConfigurationOutput), req.Error
}

func (c *Client) ListPlaybackConfigurationsWithContext(ctx context.Context, input *mediatailor.ListPlaybackConfigurationsInput, opts ...request.Option) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mediatailor",
		Action:  "ListPlaybackConfigurationsWithContext",
		Input:   input,
		Output:  (*mediatailor.ListPlaybackConfigurationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MediaTailorAPI.ListPlaybackConfigurationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*mediatailor.ListPlaybackConfigurationsOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *mediatailor.ListTagsForResourceInput, opts ...request.Option) (*mediatailor.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mediatailor",
		Action:  "ListTagsForResourceWithContext",
		Input:   input,
		Output:  (*mediatailor.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MediaTailorAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*mediatailor.ListTagsForResourceOutput), req.Error
}

func (c *Client) PutPlaybackConfigurationWithContext(ctx context.Context, input *mediatailor.PutPlaybackConfigurationInput, opts ...request.Option) (*mediatailor.PutPlaybackConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mediatailor",
		Action:  "PutPlaybackConfigurationWithContext",
		Input:   input,
		Output:  (*mediatailor.PutPlaybackConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MediaTailorAPI.PutPlaybackConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*mediatailor.PutPlaybackConfigurationOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *mediatailor.TagResourceInput, opts ...request.Option) (*mediatailor.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mediatailor",
		Action:  "TagResourceWithContext",
		Input:   input,
		Output:  (*mediatailor.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MediaTailorAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*mediatailor.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *mediatailor.UntagResourceInput, opts ...request.Option) (*mediatailor.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "mediatailor",
		Action:  "UntagResourceWithContext",
		Input:   input,
		Output:  (*mediatailor.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MediaTailorAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*mediatailor.UntagResourceOutput), req.Error
}
