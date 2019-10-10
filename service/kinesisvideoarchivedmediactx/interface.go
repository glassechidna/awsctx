// Code generated by internal/generate/main.go. DO NOT EDIT.

package kinesisvideoarchivedmediactx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia/kinesisvideoarchivedmediaiface"
	"github.com/glassechidna/awsctx"
)

type KinesisVideoArchivedMedia interface {
	GetDASHStreamingSessionURLWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error)
	GetHLSStreamingSessionURLWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error)
	GetMediaForFragmentListWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error)
	ListFragmentsWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput, opts ...request.Option) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error)
}

type Client struct {
	kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI
	Contexter awsctx.Contexter
}

func New(base kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI, ctxer awsctx.Contexter) KinesisVideoArchivedMedia {
	return &Client{
		KinesisVideoArchivedMediaAPI: base,
		Contexter: ctxer,
	}
}

var _ KinesisVideoArchivedMedia = (*kinesisvideoarchivedmedia.KinesisVideoArchivedMedia)(nil)
var _ KinesisVideoArchivedMedia = (*Client)(nil)

func (c *Client) GetDASHStreamingSessionURLWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideoarchivedmedia",
		Action:  "GetDASHStreamingSessionURL",
		Input:   input,
		Output:  (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoArchivedMediaAPI.GetDASHStreamingSessionURLWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput), req.Error
}

func (c *Client) GetHLSStreamingSessionURLWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideoarchivedmedia",
		Action:  "GetHLSStreamingSessionURL",
		Input:   input,
		Output:  (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoArchivedMediaAPI.GetHLSStreamingSessionURLWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput), req.Error
}

func (c *Client) GetMediaForFragmentListWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.GetMediaForFragmentListInput, opts ...request.Option) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideoarchivedmedia",
		Action:  "GetMediaForFragmentList",
		Input:   input,
		Output:  (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoArchivedMediaAPI.GetMediaForFragmentListWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput), req.Error
}

func (c *Client) ListFragmentsWithContext(ctx context.Context, input *kinesisvideoarchivedmedia.ListFragmentsInput, opts ...request.Option) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "kinesisvideoarchivedmedia",
		Action:  "ListFragments",
		Input:   input,
		Output:  (*kinesisvideoarchivedmedia.ListFragmentsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.KinesisVideoArchivedMediaAPI.ListFragmentsWithContext(ctx, input, opts...)
	})

	return req.Output.(*kinesisvideoarchivedmedia.ListFragmentsOutput), req.Error
}
