// Code generated by internal/generate/main.go. DO NOT EDIT.

package transcribeservicectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/aws/aws-sdk-go/service/transcribeservice/transcribeserviceiface"
	"github.com/glassechidna/awsctx"
)

type TranscribeService interface {
	CreateVocabularyWithContext(ctx context.Context, input *transcribeservice.CreateVocabularyInput, opts ...request.Option) (*transcribeservice.CreateVocabularyOutput, error)
	DeleteTranscriptionJobWithContext(ctx context.Context, input *transcribeservice.DeleteTranscriptionJobInput, opts ...request.Option) (*transcribeservice.DeleteTranscriptionJobOutput, error)
	DeleteVocabularyWithContext(ctx context.Context, input *transcribeservice.DeleteVocabularyInput, opts ...request.Option) (*transcribeservice.DeleteVocabularyOutput, error)
	GetTranscriptionJobWithContext(ctx context.Context, input *transcribeservice.GetTranscriptionJobInput, opts ...request.Option) (*transcribeservice.GetTranscriptionJobOutput, error)
	GetVocabularyWithContext(ctx context.Context, input *transcribeservice.GetVocabularyInput, opts ...request.Option) (*transcribeservice.GetVocabularyOutput, error)
	ListTranscriptionJobsWithContext(ctx context.Context, input *transcribeservice.ListTranscriptionJobsInput, opts ...request.Option) (*transcribeservice.ListTranscriptionJobsOutput, error)
	ListVocabulariesWithContext(ctx context.Context, input *transcribeservice.ListVocabulariesInput, opts ...request.Option) (*transcribeservice.ListVocabulariesOutput, error)
	StartTranscriptionJobWithContext(ctx context.Context, input *transcribeservice.StartTranscriptionJobInput, opts ...request.Option) (*transcribeservice.StartTranscriptionJobOutput, error)
	UpdateVocabularyWithContext(ctx context.Context, input *transcribeservice.UpdateVocabularyInput, opts ...request.Option) (*transcribeservice.UpdateVocabularyOutput, error)
}

type Client struct {
	transcribeserviceiface.TranscribeServiceAPI
	Contexter awsctx.Contexter
}

func New(base transcribeserviceiface.TranscribeServiceAPI, ctxer awsctx.Contexter) TranscribeService {
	return &Client{
		TranscribeServiceAPI: base,
		Contexter: ctxer,
	}
}

var _ TranscribeService = (*transcribeservice.TranscribeService)(nil)
var _ TranscribeService = (*Client)(nil)

func (c *Client) CreateVocabularyWithContext(ctx context.Context, input *transcribeservice.CreateVocabularyInput, opts ...request.Option) (*transcribeservice.CreateVocabularyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "CreateVocabulary",
		Input:   input,
		Output:  (*transcribeservice.CreateVocabularyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.CreateVocabularyWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.CreateVocabularyOutput), req.Error
}

func (c *Client) DeleteTranscriptionJobWithContext(ctx context.Context, input *transcribeservice.DeleteTranscriptionJobInput, opts ...request.Option) (*transcribeservice.DeleteTranscriptionJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "DeleteTranscriptionJob",
		Input:   input,
		Output:  (*transcribeservice.DeleteTranscriptionJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.DeleteTranscriptionJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.DeleteTranscriptionJobOutput), req.Error
}

func (c *Client) DeleteVocabularyWithContext(ctx context.Context, input *transcribeservice.DeleteVocabularyInput, opts ...request.Option) (*transcribeservice.DeleteVocabularyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "DeleteVocabulary",
		Input:   input,
		Output:  (*transcribeservice.DeleteVocabularyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.DeleteVocabularyWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.DeleteVocabularyOutput), req.Error
}

func (c *Client) GetTranscriptionJobWithContext(ctx context.Context, input *transcribeservice.GetTranscriptionJobInput, opts ...request.Option) (*transcribeservice.GetTranscriptionJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "GetTranscriptionJob",
		Input:   input,
		Output:  (*transcribeservice.GetTranscriptionJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.GetTranscriptionJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.GetTranscriptionJobOutput), req.Error
}

func (c *Client) GetVocabularyWithContext(ctx context.Context, input *transcribeservice.GetVocabularyInput, opts ...request.Option) (*transcribeservice.GetVocabularyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "GetVocabulary",
		Input:   input,
		Output:  (*transcribeservice.GetVocabularyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.GetVocabularyWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.GetVocabularyOutput), req.Error
}

func (c *Client) ListTranscriptionJobsWithContext(ctx context.Context, input *transcribeservice.ListTranscriptionJobsInput, opts ...request.Option) (*transcribeservice.ListTranscriptionJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "ListTranscriptionJobs",
		Input:   input,
		Output:  (*transcribeservice.ListTranscriptionJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.ListTranscriptionJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.ListTranscriptionJobsOutput), req.Error
}

func (c *Client) ListVocabulariesWithContext(ctx context.Context, input *transcribeservice.ListVocabulariesInput, opts ...request.Option) (*transcribeservice.ListVocabulariesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "ListVocabularies",
		Input:   input,
		Output:  (*transcribeservice.ListVocabulariesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.ListVocabulariesWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.ListVocabulariesOutput), req.Error
}

func (c *Client) StartTranscriptionJobWithContext(ctx context.Context, input *transcribeservice.StartTranscriptionJobInput, opts ...request.Option) (*transcribeservice.StartTranscriptionJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "StartTranscriptionJob",
		Input:   input,
		Output:  (*transcribeservice.StartTranscriptionJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.StartTranscriptionJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.StartTranscriptionJobOutput), req.Error
}

func (c *Client) UpdateVocabularyWithContext(ctx context.Context, input *transcribeservice.UpdateVocabularyInput, opts ...request.Option) (*transcribeservice.UpdateVocabularyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transcribeservice",
		Action:  "UpdateVocabulary",
		Input:   input,
		Output:  (*transcribeservice.UpdateVocabularyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TranscribeServiceAPI.UpdateVocabularyWithContext(ctx, input, opts...)
	})

	return req.Output.(*transcribeservice.UpdateVocabularyOutput), req.Error
}
