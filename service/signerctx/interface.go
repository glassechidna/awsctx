// Code generated by internal/generate/main.go. DO NOT EDIT.

package signerctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/signer"
	"github.com/aws/aws-sdk-go/service/signer/signeriface"
	"github.com/glassechidna/awsctx"
)

type Signer interface {
	CancelSigningProfileWithContext(ctx context.Context, input *signer.CancelSigningProfileInput, opts ...request.Option) (*signer.CancelSigningProfileOutput, error)
	DescribeSigningJobWithContext(ctx context.Context, input *signer.DescribeSigningJobInput, opts ...request.Option) (*signer.DescribeSigningJobOutput, error)
	GetSigningPlatformWithContext(ctx context.Context, input *signer.GetSigningPlatformInput, opts ...request.Option) (*signer.GetSigningPlatformOutput, error)
	GetSigningProfileWithContext(ctx context.Context, input *signer.GetSigningProfileInput, opts ...request.Option) (*signer.GetSigningProfileOutput, error)
	ListSigningJobsWithContext(ctx context.Context, input *signer.ListSigningJobsInput, opts ...request.Option) (*signer.ListSigningJobsOutput, error)
	ListSigningJobsPagesWithContext(ctx context.Context, input *signer.ListSigningJobsInput, cb func(*signer.ListSigningJobsOutput, bool) bool, opts ...request.Option) error
	ListSigningPlatformsWithContext(ctx context.Context, input *signer.ListSigningPlatformsInput, opts ...request.Option) (*signer.ListSigningPlatformsOutput, error)
	ListSigningPlatformsPagesWithContext(ctx context.Context, input *signer.ListSigningPlatformsInput, cb func(*signer.ListSigningPlatformsOutput, bool) bool, opts ...request.Option) error
	ListSigningProfilesWithContext(ctx context.Context, input *signer.ListSigningProfilesInput, opts ...request.Option) (*signer.ListSigningProfilesOutput, error)
	ListSigningProfilesPagesWithContext(ctx context.Context, input *signer.ListSigningProfilesInput, cb func(*signer.ListSigningProfilesOutput, bool) bool, opts ...request.Option) error
	PutSigningProfileWithContext(ctx context.Context, input *signer.PutSigningProfileInput, opts ...request.Option) (*signer.PutSigningProfileOutput, error)
	StartSigningJobWithContext(ctx context.Context, input *signer.StartSigningJobInput, opts ...request.Option) (*signer.StartSigningJobOutput, error)
}

type Client struct {
	signeriface.SignerAPI
	Contexter awsctx.Contexter
}

func New(base signeriface.SignerAPI, ctxer awsctx.Contexter) Signer {
	return &Client{
		SignerAPI: base,
		Contexter: ctxer,
	}
}

var _ Signer = (*signer.Signer)(nil)
var _ Signer = (*Client)(nil)

func (c *Client) CancelSigningProfileWithContext(ctx context.Context, input *signer.CancelSigningProfileInput, opts ...request.Option) (*signer.CancelSigningProfileOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "CancelSigningProfile",
		Input:   input,
		Output:  (*signer.CancelSigningProfileOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.CancelSigningProfileWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.CancelSigningProfileOutput), req.Error
}

func (c *Client) DescribeSigningJobWithContext(ctx context.Context, input *signer.DescribeSigningJobInput, opts ...request.Option) (*signer.DescribeSigningJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "DescribeSigningJob",
		Input:   input,
		Output:  (*signer.DescribeSigningJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.DescribeSigningJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.DescribeSigningJobOutput), req.Error
}

func (c *Client) GetSigningPlatformWithContext(ctx context.Context, input *signer.GetSigningPlatformInput, opts ...request.Option) (*signer.GetSigningPlatformOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "GetSigningPlatform",
		Input:   input,
		Output:  (*signer.GetSigningPlatformOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.GetSigningPlatformWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.GetSigningPlatformOutput), req.Error
}

func (c *Client) GetSigningProfileWithContext(ctx context.Context, input *signer.GetSigningProfileInput, opts ...request.Option) (*signer.GetSigningProfileOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "GetSigningProfile",
		Input:   input,
		Output:  (*signer.GetSigningProfileOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.GetSigningProfileWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.GetSigningProfileOutput), req.Error
}

func (c *Client) ListSigningJobsWithContext(ctx context.Context, input *signer.ListSigningJobsInput, opts ...request.Option) (*signer.ListSigningJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "ListSigningJobs",
		Input:   input,
		Output:  (*signer.ListSigningJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.ListSigningJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.ListSigningJobsOutput), req.Error
}

func (c *Client) ListSigningJobsPagesWithContext(ctx context.Context, input *signer.ListSigningJobsInput, cb func(*signer.ListSigningJobsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "ListSigningJobs",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SignerAPI.ListSigningJobsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListSigningPlatformsWithContext(ctx context.Context, input *signer.ListSigningPlatformsInput, opts ...request.Option) (*signer.ListSigningPlatformsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "ListSigningPlatforms",
		Input:   input,
		Output:  (*signer.ListSigningPlatformsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.ListSigningPlatformsWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.ListSigningPlatformsOutput), req.Error
}

func (c *Client) ListSigningPlatformsPagesWithContext(ctx context.Context, input *signer.ListSigningPlatformsInput, cb func(*signer.ListSigningPlatformsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "ListSigningPlatforms",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SignerAPI.ListSigningPlatformsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListSigningProfilesWithContext(ctx context.Context, input *signer.ListSigningProfilesInput, opts ...request.Option) (*signer.ListSigningProfilesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "ListSigningProfiles",
		Input:   input,
		Output:  (*signer.ListSigningProfilesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.ListSigningProfilesWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.ListSigningProfilesOutput), req.Error
}

func (c *Client) ListSigningProfilesPagesWithContext(ctx context.Context, input *signer.ListSigningProfilesInput, cb func(*signer.ListSigningProfilesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "ListSigningProfiles",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.SignerAPI.ListSigningProfilesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) PutSigningProfileWithContext(ctx context.Context, input *signer.PutSigningProfileInput, opts ...request.Option) (*signer.PutSigningProfileOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "PutSigningProfile",
		Input:   input,
		Output:  (*signer.PutSigningProfileOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.PutSigningProfileWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.PutSigningProfileOutput), req.Error
}

func (c *Client) StartSigningJobWithContext(ctx context.Context, input *signer.StartSigningJobInput, opts ...request.Option) (*signer.StartSigningJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "signer",
		Action:  "StartSigningJob",
		Input:   input,
		Output:  (*signer.StartSigningJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SignerAPI.StartSigningJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*signer.StartSigningJobOutput), req.Error
}
