// Code generated by internal/generate/main.go. DO NOT EDIT.

package pollyctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/polly/pollyiface"
	"github.com/glassechidna/awsctx"
)

type Polly interface {
	DeleteLexiconWithContext(ctx context.Context, input *polly.DeleteLexiconInput, opts ...request.Option) (*polly.DeleteLexiconOutput, error)
	DescribeVoicesWithContext(ctx context.Context, input *polly.DescribeVoicesInput, opts ...request.Option) (*polly.DescribeVoicesOutput, error)
	GetLexiconWithContext(ctx context.Context, input *polly.GetLexiconInput, opts ...request.Option) (*polly.GetLexiconOutput, error)
	GetSpeechSynthesisTaskWithContext(ctx context.Context, input *polly.GetSpeechSynthesisTaskInput, opts ...request.Option) (*polly.GetSpeechSynthesisTaskOutput, error)
	ListLexiconsWithContext(ctx context.Context, input *polly.ListLexiconsInput, opts ...request.Option) (*polly.ListLexiconsOutput, error)
	ListSpeechSynthesisTasksWithContext(ctx context.Context, input *polly.ListSpeechSynthesisTasksInput, opts ...request.Option) (*polly.ListSpeechSynthesisTasksOutput, error)
	ListSpeechSynthesisTasksPagesWithContext(ctx context.Context, input *polly.ListSpeechSynthesisTasksInput, cb func(*polly.ListSpeechSynthesisTasksOutput, bool) bool, opts ...request.Option) error
	PutLexiconWithContext(ctx context.Context, input *polly.PutLexiconInput, opts ...request.Option) (*polly.PutLexiconOutput, error)
	StartSpeechSynthesisTaskWithContext(ctx context.Context, input *polly.StartSpeechSynthesisTaskInput, opts ...request.Option) (*polly.StartSpeechSynthesisTaskOutput, error)
	SynthesizeSpeechWithContext(ctx context.Context, input *polly.SynthesizeSpeechInput, opts ...request.Option) (*polly.SynthesizeSpeechOutput, error)
}

type Client struct {
	pollyiface.PollyAPI
	Contexter awsctx.Contexter
}

func New(base pollyiface.PollyAPI, ctxer awsctx.Contexter) Polly {
	return &Client{
		PollyAPI: base,
		Contexter: ctxer,
	}
}

var _ Polly = (*polly.Polly)(nil)
var _ Polly = (*Client)(nil)

func (c *Client) DeleteLexiconWithContext(ctx context.Context, input *polly.DeleteLexiconInput, opts ...request.Option) (*polly.DeleteLexiconOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "DeleteLexicon",
		Input:   input,
		Output:  (*polly.DeleteLexiconOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.DeleteLexiconWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.DeleteLexiconOutput), req.Error
}

func (c *Client) DescribeVoicesWithContext(ctx context.Context, input *polly.DescribeVoicesInput, opts ...request.Option) (*polly.DescribeVoicesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "DescribeVoices",
		Input:   input,
		Output:  (*polly.DescribeVoicesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.DescribeVoicesWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.DescribeVoicesOutput), req.Error
}

func (c *Client) GetLexiconWithContext(ctx context.Context, input *polly.GetLexiconInput, opts ...request.Option) (*polly.GetLexiconOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "GetLexicon",
		Input:   input,
		Output:  (*polly.GetLexiconOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.GetLexiconWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.GetLexiconOutput), req.Error
}

func (c *Client) GetSpeechSynthesisTaskWithContext(ctx context.Context, input *polly.GetSpeechSynthesisTaskInput, opts ...request.Option) (*polly.GetSpeechSynthesisTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "GetSpeechSynthesisTask",
		Input:   input,
		Output:  (*polly.GetSpeechSynthesisTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.GetSpeechSynthesisTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.GetSpeechSynthesisTaskOutput), req.Error
}

func (c *Client) ListLexiconsWithContext(ctx context.Context, input *polly.ListLexiconsInput, opts ...request.Option) (*polly.ListLexiconsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "ListLexicons",
		Input:   input,
		Output:  (*polly.ListLexiconsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.ListLexiconsWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.ListLexiconsOutput), req.Error
}

func (c *Client) ListSpeechSynthesisTasksWithContext(ctx context.Context, input *polly.ListSpeechSynthesisTasksInput, opts ...request.Option) (*polly.ListSpeechSynthesisTasksOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "ListSpeechSynthesisTasks",
		Input:   input,
		Output:  (*polly.ListSpeechSynthesisTasksOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.ListSpeechSynthesisTasksWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.ListSpeechSynthesisTasksOutput), req.Error
}

func (c *Client) ListSpeechSynthesisTasksPagesWithContext(ctx context.Context, input *polly.ListSpeechSynthesisTasksInput, cb func(*polly.ListSpeechSynthesisTasksOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "ListSpeechSynthesisTasks",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.PollyAPI.ListSpeechSynthesisTasksPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) PutLexiconWithContext(ctx context.Context, input *polly.PutLexiconInput, opts ...request.Option) (*polly.PutLexiconOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "PutLexicon",
		Input:   input,
		Output:  (*polly.PutLexiconOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.PutLexiconWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.PutLexiconOutput), req.Error
}

func (c *Client) StartSpeechSynthesisTaskWithContext(ctx context.Context, input *polly.StartSpeechSynthesisTaskInput, opts ...request.Option) (*polly.StartSpeechSynthesisTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "StartSpeechSynthesisTask",
		Input:   input,
		Output:  (*polly.StartSpeechSynthesisTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.StartSpeechSynthesisTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.StartSpeechSynthesisTaskOutput), req.Error
}

func (c *Client) SynthesizeSpeechWithContext(ctx context.Context, input *polly.SynthesizeSpeechInput, opts ...request.Option) (*polly.SynthesizeSpeechOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "polly",
		Action:  "SynthesizeSpeech",
		Input:   input,
		Output:  (*polly.SynthesizeSpeechOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PollyAPI.SynthesizeSpeechWithContext(ctx, input, opts...)
	})

	return req.Output.(*polly.SynthesizeSpeechOutput), req.Error
}
