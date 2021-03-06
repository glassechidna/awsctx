// Code generated by internal/generate/main.go. DO NOT EDIT.

package firehosectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
	"github.com/glassechidna/awsctx"
)

type Firehose interface {
	CreateDeliveryStreamWithContext(ctx context.Context, input *firehose.CreateDeliveryStreamInput, opts ...request.Option) (*firehose.CreateDeliveryStreamOutput, error)
	DeleteDeliveryStreamWithContext(ctx context.Context, input *firehose.DeleteDeliveryStreamInput, opts ...request.Option) (*firehose.DeleteDeliveryStreamOutput, error)
	DescribeDeliveryStreamWithContext(ctx context.Context, input *firehose.DescribeDeliveryStreamInput, opts ...request.Option) (*firehose.DescribeDeliveryStreamOutput, error)
	ListDeliveryStreamsWithContext(ctx context.Context, input *firehose.ListDeliveryStreamsInput, opts ...request.Option) (*firehose.ListDeliveryStreamsOutput, error)
	ListTagsForDeliveryStreamWithContext(ctx context.Context, input *firehose.ListTagsForDeliveryStreamInput, opts ...request.Option) (*firehose.ListTagsForDeliveryStreamOutput, error)
	PutRecordWithContext(ctx context.Context, input *firehose.PutRecordInput, opts ...request.Option) (*firehose.PutRecordOutput, error)
	PutRecordBatchWithContext(ctx context.Context, input *firehose.PutRecordBatchInput, opts ...request.Option) (*firehose.PutRecordBatchOutput, error)
	StartDeliveryStreamEncryptionWithContext(ctx context.Context, input *firehose.StartDeliveryStreamEncryptionInput, opts ...request.Option) (*firehose.StartDeliveryStreamEncryptionOutput, error)
	StopDeliveryStreamEncryptionWithContext(ctx context.Context, input *firehose.StopDeliveryStreamEncryptionInput, opts ...request.Option) (*firehose.StopDeliveryStreamEncryptionOutput, error)
	TagDeliveryStreamWithContext(ctx context.Context, input *firehose.TagDeliveryStreamInput, opts ...request.Option) (*firehose.TagDeliveryStreamOutput, error)
	UntagDeliveryStreamWithContext(ctx context.Context, input *firehose.UntagDeliveryStreamInput, opts ...request.Option) (*firehose.UntagDeliveryStreamOutput, error)
	UpdateDestinationWithContext(ctx context.Context, input *firehose.UpdateDestinationInput, opts ...request.Option) (*firehose.UpdateDestinationOutput, error)
}

type Client struct {
	firehoseiface.FirehoseAPI
	Contexter awsctx.Contexter
}

func New(base firehoseiface.FirehoseAPI, ctxer awsctx.Contexter) Firehose {
	return &Client{
		FirehoseAPI: base,
		Contexter: ctxer,
	}
}

var _ Firehose = (*firehose.Firehose)(nil)
var _ Firehose = (*Client)(nil)

func (c *Client) CreateDeliveryStreamWithContext(ctx context.Context, input *firehose.CreateDeliveryStreamInput, opts ...request.Option) (*firehose.CreateDeliveryStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "CreateDeliveryStream",
		Input:   input,
		Output:  (*firehose.CreateDeliveryStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.CreateDeliveryStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.CreateDeliveryStreamOutput), req.Error
}

func (c *Client) DeleteDeliveryStreamWithContext(ctx context.Context, input *firehose.DeleteDeliveryStreamInput, opts ...request.Option) (*firehose.DeleteDeliveryStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "DeleteDeliveryStream",
		Input:   input,
		Output:  (*firehose.DeleteDeliveryStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.DeleteDeliveryStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.DeleteDeliveryStreamOutput), req.Error
}

func (c *Client) DescribeDeliveryStreamWithContext(ctx context.Context, input *firehose.DescribeDeliveryStreamInput, opts ...request.Option) (*firehose.DescribeDeliveryStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "DescribeDeliveryStream",
		Input:   input,
		Output:  (*firehose.DescribeDeliveryStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.DescribeDeliveryStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.DescribeDeliveryStreamOutput), req.Error
}

func (c *Client) ListDeliveryStreamsWithContext(ctx context.Context, input *firehose.ListDeliveryStreamsInput, opts ...request.Option) (*firehose.ListDeliveryStreamsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "ListDeliveryStreams",
		Input:   input,
		Output:  (*firehose.ListDeliveryStreamsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.ListDeliveryStreamsWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.ListDeliveryStreamsOutput), req.Error
}

func (c *Client) ListTagsForDeliveryStreamWithContext(ctx context.Context, input *firehose.ListTagsForDeliveryStreamInput, opts ...request.Option) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "ListTagsForDeliveryStream",
		Input:   input,
		Output:  (*firehose.ListTagsForDeliveryStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.ListTagsForDeliveryStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.ListTagsForDeliveryStreamOutput), req.Error
}

func (c *Client) PutRecordWithContext(ctx context.Context, input *firehose.PutRecordInput, opts ...request.Option) (*firehose.PutRecordOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "PutRecord",
		Input:   input,
		Output:  (*firehose.PutRecordOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.PutRecordWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.PutRecordOutput), req.Error
}

func (c *Client) PutRecordBatchWithContext(ctx context.Context, input *firehose.PutRecordBatchInput, opts ...request.Option) (*firehose.PutRecordBatchOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "PutRecordBatch",
		Input:   input,
		Output:  (*firehose.PutRecordBatchOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.PutRecordBatchWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.PutRecordBatchOutput), req.Error
}

func (c *Client) StartDeliveryStreamEncryptionWithContext(ctx context.Context, input *firehose.StartDeliveryStreamEncryptionInput, opts ...request.Option) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "StartDeliveryStreamEncryption",
		Input:   input,
		Output:  (*firehose.StartDeliveryStreamEncryptionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.StartDeliveryStreamEncryptionWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.StartDeliveryStreamEncryptionOutput), req.Error
}

func (c *Client) StopDeliveryStreamEncryptionWithContext(ctx context.Context, input *firehose.StopDeliveryStreamEncryptionInput, opts ...request.Option) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "StopDeliveryStreamEncryption",
		Input:   input,
		Output:  (*firehose.StopDeliveryStreamEncryptionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.StopDeliveryStreamEncryptionWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.StopDeliveryStreamEncryptionOutput), req.Error
}

func (c *Client) TagDeliveryStreamWithContext(ctx context.Context, input *firehose.TagDeliveryStreamInput, opts ...request.Option) (*firehose.TagDeliveryStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "TagDeliveryStream",
		Input:   input,
		Output:  (*firehose.TagDeliveryStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.TagDeliveryStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.TagDeliveryStreamOutput), req.Error
}

func (c *Client) UntagDeliveryStreamWithContext(ctx context.Context, input *firehose.UntagDeliveryStreamInput, opts ...request.Option) (*firehose.UntagDeliveryStreamOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "UntagDeliveryStream",
		Input:   input,
		Output:  (*firehose.UntagDeliveryStreamOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.UntagDeliveryStreamWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.UntagDeliveryStreamOutput), req.Error
}

func (c *Client) UpdateDestinationWithContext(ctx context.Context, input *firehose.UpdateDestinationInput, opts ...request.Option) (*firehose.UpdateDestinationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "firehose",
		Action:  "UpdateDestination",
		Input:   input,
		Output:  (*firehose.UpdateDestinationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.FirehoseAPI.UpdateDestinationWithContext(ctx, input, opts...)
	})

	return req.Output.(*firehose.UpdateDestinationOutput), req.Error
}
