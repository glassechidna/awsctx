// Code generated by internal/generate/main.go. DO NOT EDIT.

package kinesisvideoctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"github.com/aws/aws-sdk-go/service/kinesisvideo/kinesisvideoiface"
	"github.com/glassechidna/awsctx"
)

type KinesisVideo interface {
	CreateSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.CreateSignalingChannelInput, opts ...request.Option) (*kinesisvideo.CreateSignalingChannelOutput, error)
	CreateStreamWithContext(ctx context.Context, input *kinesisvideo.CreateStreamInput, opts ...request.Option) (*kinesisvideo.CreateStreamOutput, error)
	DeleteSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.DeleteSignalingChannelInput, opts ...request.Option) (*kinesisvideo.DeleteSignalingChannelOutput, error)
	DeleteStreamWithContext(ctx context.Context, input *kinesisvideo.DeleteStreamInput, opts ...request.Option) (*kinesisvideo.DeleteStreamOutput, error)
	DescribeImageGenerationConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeImageGenerationConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeImageGenerationConfigurationOutput, error)
	DescribeNotificationConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeNotificationConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeNotificationConfigurationOutput, error)
	DescribeSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.DescribeSignalingChannelInput, opts ...request.Option) (*kinesisvideo.DescribeSignalingChannelOutput, error)
	DescribeStreamWithContext(ctx context.Context, input *kinesisvideo.DescribeStreamInput, opts ...request.Option) (*kinesisvideo.DescribeStreamOutput, error)
	GetDataEndpointWithContext(ctx context.Context, input *kinesisvideo.GetDataEndpointInput, opts ...request.Option) (*kinesisvideo.GetDataEndpointOutput, error)
	GetSignalingChannelEndpointWithContext(ctx context.Context, input *kinesisvideo.GetSignalingChannelEndpointInput, opts ...request.Option) (*kinesisvideo.GetSignalingChannelEndpointOutput, error)
	ListSignalingChannelsWithContext(ctx context.Context, input *kinesisvideo.ListSignalingChannelsInput, opts ...request.Option) (*kinesisvideo.ListSignalingChannelsOutput, error)
	ListSignalingChannelsPagesWithContext(ctx context.Context, input *kinesisvideo.ListSignalingChannelsInput, cb func(*kinesisvideo.ListSignalingChannelsOutput, bool) bool, opts ...request.Option) error
	ListStreamsWithContext(ctx context.Context, input *kinesisvideo.ListStreamsInput, opts ...request.Option) (*kinesisvideo.ListStreamsOutput, error)
	ListStreamsPagesWithContext(ctx context.Context, input *kinesisvideo.ListStreamsInput, cb func(*kinesisvideo.ListStreamsOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *kinesisvideo.ListTagsForResourceInput, opts ...request.Option) (*kinesisvideo.ListTagsForResourceOutput, error)
	ListTagsForStreamWithContext(ctx context.Context, input *kinesisvideo.ListTagsForStreamInput, opts ...request.Option) (*kinesisvideo.ListTagsForStreamOutput, error)
	TagResourceWithContext(ctx context.Context, input *kinesisvideo.TagResourceInput, opts ...request.Option) (*kinesisvideo.TagResourceOutput, error)
	TagStreamWithContext(ctx context.Context, input *kinesisvideo.TagStreamInput, opts ...request.Option) (*kinesisvideo.TagStreamOutput, error)
	UntagResourceWithContext(ctx context.Context, input *kinesisvideo.UntagResourceInput, opts ...request.Option) (*kinesisvideo.UntagResourceOutput, error)
	UntagStreamWithContext(ctx context.Context, input *kinesisvideo.UntagStreamInput, opts ...request.Option) (*kinesisvideo.UntagStreamOutput, error)
	UpdateDataRetentionWithContext(ctx context.Context, input *kinesisvideo.UpdateDataRetentionInput, opts ...request.Option) (*kinesisvideo.UpdateDataRetentionOutput, error)
	UpdateImageGenerationConfigurationWithContext(ctx context.Context, input *kinesisvideo.UpdateImageGenerationConfigurationInput, opts ...request.Option) (*kinesisvideo.UpdateImageGenerationConfigurationOutput, error)
	UpdateNotificationConfigurationWithContext(ctx context.Context, input *kinesisvideo.UpdateNotificationConfigurationInput, opts ...request.Option) (*kinesisvideo.UpdateNotificationConfigurationOutput, error)
	UpdateSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.UpdateSignalingChannelInput, opts ...request.Option) (*kinesisvideo.UpdateSignalingChannelOutput, error)
	UpdateStreamWithContext(ctx context.Context, input *kinesisvideo.UpdateStreamInput, opts ...request.Option) (*kinesisvideo.UpdateStreamOutput, error)
}

type Client struct {
	kinesisvideoiface.KinesisVideoAPI
	Contexter awsctx.Contexter
}

func New(base kinesisvideoiface.KinesisVideoAPI, ctxer awsctx.Contexter) KinesisVideo {
	return &Client{
		KinesisVideoAPI: base,
		Contexter: ctxer,
	}
}

var _ KinesisVideo = (*kinesisvideo.KinesisVideo)(nil)
var _ KinesisVideo = (*Client)(nil)

func (c *Client) CreateSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.CreateSignalingChannelInput, opts ...request.Option) (*kinesisvideo.CreateSignalingChannelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "CreateSignalingChannel",
		Input:   input,
		Output:  (*kinesisvideo.CreateSignalingChannelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.CreateSignalingChannelWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.CreateSignalingChannelOutput), req.Error
}

func (c *Client) CreateStreamWithContext(ctx context.Context, input *kinesisvideo.CreateStreamInput, opts ...request.Option) (*kinesisvideo.CreateStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "CreateStream",
		Input:   input,
		Output:  (*kinesisvideo.CreateStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.CreateStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.CreateStreamOutput), req.Error
}

func (c *Client) DeleteSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.DeleteSignalingChannelInput, opts ...request.Option) (*kinesisvideo.DeleteSignalingChannelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "DeleteSignalingChannel",
		Input:   input,
		Output:  (*kinesisvideo.DeleteSignalingChannelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.DeleteSignalingChannelWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.DeleteSignalingChannelOutput), req.Error
}

func (c *Client) DeleteStreamWithContext(ctx context.Context, input *kinesisvideo.DeleteStreamInput, opts ...request.Option) (*kinesisvideo.DeleteStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "DeleteStream",
		Input:   input,
		Output:  (*kinesisvideo.DeleteStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.DeleteStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.DeleteStreamOutput), req.Error
}

func (c *Client) DescribeImageGenerationConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeImageGenerationConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeImageGenerationConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "DescribeImageGenerationConfiguration",
		Input:   input,
		Output:  (*kinesisvideo.DescribeImageGenerationConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.DescribeImageGenerationConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.DescribeImageGenerationConfigurationOutput), req.Error
}

func (c *Client) DescribeNotificationConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeNotificationConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeNotificationConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "DescribeNotificationConfiguration",
		Input:   input,
		Output:  (*kinesisvideo.DescribeNotificationConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.DescribeNotificationConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.DescribeNotificationConfigurationOutput), req.Error
}

func (c *Client) DescribeSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.DescribeSignalingChannelInput, opts ...request.Option) (*kinesisvideo.DescribeSignalingChannelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "DescribeSignalingChannel",
		Input:   input,
		Output:  (*kinesisvideo.DescribeSignalingChannelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.DescribeSignalingChannelWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.DescribeSignalingChannelOutput), req.Error
}

func (c *Client) DescribeStreamWithContext(ctx context.Context, input *kinesisvideo.DescribeStreamInput, opts ...request.Option) (*kinesisvideo.DescribeStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "DescribeStream",
		Input:   input,
		Output:  (*kinesisvideo.DescribeStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.DescribeStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.DescribeStreamOutput), req.Error
}

func (c *Client) GetDataEndpointWithContext(ctx context.Context, input *kinesisvideo.GetDataEndpointInput, opts ...request.Option) (*kinesisvideo.GetDataEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "GetDataEndpoint",
		Input:   input,
		Output:  (*kinesisvideo.GetDataEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.GetDataEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.GetDataEndpointOutput), req.Error
}

func (c *Client) GetSignalingChannelEndpointWithContext(ctx context.Context, input *kinesisvideo.GetSignalingChannelEndpointInput, opts ...request.Option) (*kinesisvideo.GetSignalingChannelEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "GetSignalingChannelEndpoint",
		Input:   input,
		Output:  (*kinesisvideo.GetSignalingChannelEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.GetSignalingChannelEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.GetSignalingChannelEndpointOutput), req.Error
}

func (c *Client) ListSignalingChannelsWithContext(ctx context.Context, input *kinesisvideo.ListSignalingChannelsInput, opts ...request.Option) (*kinesisvideo.ListSignalingChannelsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "ListSignalingChannels",
		Input:   input,
		Output:  (*kinesisvideo.ListSignalingChannelsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.ListSignalingChannelsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.ListSignalingChannelsOutput), req.Error
}

func (c *Client) ListSignalingChannelsPagesWithContext(ctx context.Context, input *kinesisvideo.ListSignalingChannelsInput, cb func(*kinesisvideo.ListSignalingChannelsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "ListSignalingChannels",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.KinesisVideoAPI.ListSignalingChannelsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListStreamsWithContext(ctx context.Context, input *kinesisvideo.ListStreamsInput, opts ...request.Option) (*kinesisvideo.ListStreamsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "ListStreams",
		Input:   input,
		Output:  (*kinesisvideo.ListStreamsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.ListStreamsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.ListStreamsOutput), req.Error
}

func (c *Client) ListStreamsPagesWithContext(ctx context.Context, input *kinesisvideo.ListStreamsInput, cb func(*kinesisvideo.ListStreamsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "ListStreams",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.KinesisVideoAPI.ListStreamsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *kinesisvideo.ListTagsForResourceInput, opts ...request.Option) (*kinesisvideo.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*kinesisvideo.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTagsForStreamWithContext(ctx context.Context, input *kinesisvideo.ListTagsForStreamInput, opts ...request.Option) (*kinesisvideo.ListTagsForStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "ListTagsForStream",
		Input:   input,
		Output:  (*kinesisvideo.ListTagsForStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.ListTagsForStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.ListTagsForStreamOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *kinesisvideo.TagResourceInput, opts ...request.Option) (*kinesisvideo.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "TagResource",
		Input:   input,
		Output:  (*kinesisvideo.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.TagResourceOutput), req.Error
}

func (c *Client) TagStreamWithContext(ctx context.Context, input *kinesisvideo.TagStreamInput, opts ...request.Option) (*kinesisvideo.TagStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "TagStream",
		Input:   input,
		Output:  (*kinesisvideo.TagStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.TagStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.TagStreamOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *kinesisvideo.UntagResourceInput, opts ...request.Option) (*kinesisvideo.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*kinesisvideo.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.UntagResourceOutput), req.Error
}

func (c *Client) UntagStreamWithContext(ctx context.Context, input *kinesisvideo.UntagStreamInput, opts ...request.Option) (*kinesisvideo.UntagStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "UntagStream",
		Input:   input,
		Output:  (*kinesisvideo.UntagStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.UntagStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.UntagStreamOutput), req.Error
}

func (c *Client) UpdateDataRetentionWithContext(ctx context.Context, input *kinesisvideo.UpdateDataRetentionInput, opts ...request.Option) (*kinesisvideo.UpdateDataRetentionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "UpdateDataRetention",
		Input:   input,
		Output:  (*kinesisvideo.UpdateDataRetentionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.UpdateDataRetentionWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.UpdateDataRetentionOutput), req.Error
}

func (c *Client) UpdateImageGenerationConfigurationWithContext(ctx context.Context, input *kinesisvideo.UpdateImageGenerationConfigurationInput, opts ...request.Option) (*kinesisvideo.UpdateImageGenerationConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "UpdateImageGenerationConfiguration",
		Input:   input,
		Output:  (*kinesisvideo.UpdateImageGenerationConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.UpdateImageGenerationConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.UpdateImageGenerationConfigurationOutput), req.Error
}

func (c *Client) UpdateNotificationConfigurationWithContext(ctx context.Context, input *kinesisvideo.UpdateNotificationConfigurationInput, opts ...request.Option) (*kinesisvideo.UpdateNotificationConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "UpdateNotificationConfiguration",
		Input:   input,
		Output:  (*kinesisvideo.UpdateNotificationConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.UpdateNotificationConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.UpdateNotificationConfigurationOutput), req.Error
}

func (c *Client) UpdateSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.UpdateSignalingChannelInput, opts ...request.Option) (*kinesisvideo.UpdateSignalingChannelOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "UpdateSignalingChannel",
		Input:   input,
		Output:  (*kinesisvideo.UpdateSignalingChannelOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.UpdateSignalingChannelWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.UpdateSignalingChannelOutput), req.Error
}

func (c *Client) UpdateStreamWithContext(ctx context.Context, input *kinesisvideo.UpdateStreamInput, opts ...request.Option) (*kinesisvideo.UpdateStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideo",
		Action:  "UpdateStream",
		Input:   input,
		Output:  (*kinesisvideo.UpdateStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoAPI.UpdateStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideo.UpdateStreamOutput), req.Error
}
