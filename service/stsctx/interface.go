// Code generated by internal/generate/main.go. DO NOT EDIT.

package stsctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/glassechidna/awsctx"
)

type STS interface {
	AssumeRoleWithContext(ctx context.Context, input *sts.AssumeRoleInput, opts ...request.Option) (*sts.AssumeRoleOutput, error)
	AssumeRoleWithSAMLWithContext(ctx context.Context, input *sts.AssumeRoleWithSAMLInput, opts ...request.Option) (*sts.AssumeRoleWithSAMLOutput, error)
	AssumeRoleWithWebIdentityWithContext(ctx context.Context, input *sts.AssumeRoleWithWebIdentityInput, opts ...request.Option) (*sts.AssumeRoleWithWebIdentityOutput, error)
	DecodeAuthorizationMessageWithContext(ctx context.Context, input *sts.DecodeAuthorizationMessageInput, opts ...request.Option) (*sts.DecodeAuthorizationMessageOutput, error)
	GetAccessKeyInfoWithContext(ctx context.Context, input *sts.GetAccessKeyInfoInput, opts ...request.Option) (*sts.GetAccessKeyInfoOutput, error)
	GetCallerIdentityWithContext(ctx context.Context, input *sts.GetCallerIdentityInput, opts ...request.Option) (*sts.GetCallerIdentityOutput, error)
	GetFederationTokenWithContext(ctx context.Context, input *sts.GetFederationTokenInput, opts ...request.Option) (*sts.GetFederationTokenOutput, error)
	GetSessionTokenWithContext(ctx context.Context, input *sts.GetSessionTokenInput, opts ...request.Option) (*sts.GetSessionTokenOutput, error)
}

type Client struct {
	stsiface.STSAPI
	Contexter awsctx.Contexter
}

func New(base stsiface.STSAPI, ctxer awsctx.Contexter) STS {
	return &Client{
		STSAPI: base,
		Contexter: ctxer,
	}
}

var _ STS = (*sts.STS)(nil)
var _ STS = (*Client)(nil)

func (c *Client) AssumeRoleWithContext(ctx context.Context, input *sts.AssumeRoleInput, opts ...request.Option) (*sts.AssumeRoleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sts",
		Action:  "AssumeRoleWithContext",
		Input:   input,
		Output:  (*sts.AssumeRoleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.STSAPI.AssumeRoleWithContext(ctx, input, opts...)
	})

	return req.Output.(*sts.AssumeRoleOutput), req.Error
}

func (c *Client) AssumeRoleWithSAMLWithContext(ctx context.Context, input *sts.AssumeRoleWithSAMLInput, opts ...request.Option) (*sts.AssumeRoleWithSAMLOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sts",
		Action:  "AssumeRoleWithSAMLWithContext",
		Input:   input,
		Output:  (*sts.AssumeRoleWithSAMLOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.STSAPI.AssumeRoleWithSAMLWithContext(ctx, input, opts...)
	})

	return req.Output.(*sts.AssumeRoleWithSAMLOutput), req.Error
}

func (c *Client) AssumeRoleWithWebIdentityWithContext(ctx context.Context, input *sts.AssumeRoleWithWebIdentityInput, opts ...request.Option) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sts",
		Action:  "AssumeRoleWithWebIdentityWithContext",
		Input:   input,
		Output:  (*sts.AssumeRoleWithWebIdentityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.STSAPI.AssumeRoleWithWebIdentityWithContext(ctx, input, opts...)
	})

	return req.Output.(*sts.AssumeRoleWithWebIdentityOutput), req.Error
}

func (c *Client) DecodeAuthorizationMessageWithContext(ctx context.Context, input *sts.DecodeAuthorizationMessageInput, opts ...request.Option) (*sts.DecodeAuthorizationMessageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sts",
		Action:  "DecodeAuthorizationMessageWithContext",
		Input:   input,
		Output:  (*sts.DecodeAuthorizationMessageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.STSAPI.DecodeAuthorizationMessageWithContext(ctx, input, opts...)
	})

	return req.Output.(*sts.DecodeAuthorizationMessageOutput), req.Error
}

func (c *Client) GetAccessKeyInfoWithContext(ctx context.Context, input *sts.GetAccessKeyInfoInput, opts ...request.Option) (*sts.GetAccessKeyInfoOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sts",
		Action:  "GetAccessKeyInfoWithContext",
		Input:   input,
		Output:  (*sts.GetAccessKeyInfoOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.STSAPI.GetAccessKeyInfoWithContext(ctx, input, opts...)
	})

	return req.Output.(*sts.GetAccessKeyInfoOutput), req.Error
}

func (c *Client) GetCallerIdentityWithContext(ctx context.Context, input *sts.GetCallerIdentityInput, opts ...request.Option) (*sts.GetCallerIdentityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sts",
		Action:  "GetCallerIdentityWithContext",
		Input:   input,
		Output:  (*sts.GetCallerIdentityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.STSAPI.GetCallerIdentityWithContext(ctx, input, opts...)
	})

	return req.Output.(*sts.GetCallerIdentityOutput), req.Error
}

func (c *Client) GetFederationTokenWithContext(ctx context.Context, input *sts.GetFederationTokenInput, opts ...request.Option) (*sts.GetFederationTokenOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sts",
		Action:  "GetFederationTokenWithContext",
		Input:   input,
		Output:  (*sts.GetFederationTokenOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.STSAPI.GetFederationTokenWithContext(ctx, input, opts...)
	})

	return req.Output.(*sts.GetFederationTokenOutput), req.Error
}

func (c *Client) GetSessionTokenWithContext(ctx context.Context, input *sts.GetSessionTokenInput, opts ...request.Option) (*sts.GetSessionTokenOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sts",
		Action:  "GetSessionTokenWithContext",
		Input:   input,
		Output:  (*sts.GetSessionTokenOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.STSAPI.GetSessionTokenWithContext(ctx, input, opts...)
	})

	return req.Output.(*sts.GetSessionTokenOutput), req.Error
}
