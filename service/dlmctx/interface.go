// Code generated by internal/generate/main.go. DO NOT EDIT.

package dlmctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dlm"
	"github.com/aws/aws-sdk-go/service/dlm/dlmiface"
	"github.com/glassechidna/awsctx"
)

type DLM interface {
	CreateLifecyclePolicyWithContext(ctx context.Context, input *dlm.CreateLifecyclePolicyInput, opts ...request.Option) (*dlm.CreateLifecyclePolicyOutput, error)
	DeleteLifecyclePolicyWithContext(ctx context.Context, input *dlm.DeleteLifecyclePolicyInput, opts ...request.Option) (*dlm.DeleteLifecyclePolicyOutput, error)
	GetLifecyclePoliciesWithContext(ctx context.Context, input *dlm.GetLifecyclePoliciesInput, opts ...request.Option) (*dlm.GetLifecyclePoliciesOutput, error)
	GetLifecyclePolicyWithContext(ctx context.Context, input *dlm.GetLifecyclePolicyInput, opts ...request.Option) (*dlm.GetLifecyclePolicyOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *dlm.ListTagsForResourceInput, opts ...request.Option) (*dlm.ListTagsForResourceOutput, error)
	TagResourceWithContext(ctx context.Context, input *dlm.TagResourceInput, opts ...request.Option) (*dlm.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *dlm.UntagResourceInput, opts ...request.Option) (*dlm.UntagResourceOutput, error)
	UpdateLifecyclePolicyWithContext(ctx context.Context, input *dlm.UpdateLifecyclePolicyInput, opts ...request.Option) (*dlm.UpdateLifecyclePolicyOutput, error)
}

type Client struct {
	dlmiface.DLMAPI
	Contexter awsctx.Contexter
}

func New(base dlmiface.DLMAPI, ctxer awsctx.Contexter) DLM {
	return &Client{
		DLMAPI: base,
		Contexter: ctxer,
	}
}

var _ DLM = (*dlm.DLM)(nil)
var _ DLM = (*Client)(nil)

func (c *Client) CreateLifecyclePolicyWithContext(ctx context.Context, input *dlm.CreateLifecyclePolicyInput, opts ...request.Option) (*dlm.CreateLifecyclePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dlm",
		Action:  "CreateLifecyclePolicy",
		Input:   input,
		Output:  (*dlm.CreateLifecyclePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DLMAPI.CreateLifecyclePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*dlm.CreateLifecyclePolicyOutput), req.Error
}

func (c *Client) DeleteLifecyclePolicyWithContext(ctx context.Context, input *dlm.DeleteLifecyclePolicyInput, opts ...request.Option) (*dlm.DeleteLifecyclePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dlm",
		Action:  "DeleteLifecyclePolicy",
		Input:   input,
		Output:  (*dlm.DeleteLifecyclePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DLMAPI.DeleteLifecyclePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*dlm.DeleteLifecyclePolicyOutput), req.Error
}

func (c *Client) GetLifecyclePoliciesWithContext(ctx context.Context, input *dlm.GetLifecyclePoliciesInput, opts ...request.Option) (*dlm.GetLifecyclePoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dlm",
		Action:  "GetLifecyclePolicies",
		Input:   input,
		Output:  (*dlm.GetLifecyclePoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DLMAPI.GetLifecyclePoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*dlm.GetLifecyclePoliciesOutput), req.Error
}

func (c *Client) GetLifecyclePolicyWithContext(ctx context.Context, input *dlm.GetLifecyclePolicyInput, opts ...request.Option) (*dlm.GetLifecyclePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dlm",
		Action:  "GetLifecyclePolicy",
		Input:   input,
		Output:  (*dlm.GetLifecyclePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DLMAPI.GetLifecyclePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*dlm.GetLifecyclePolicyOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *dlm.ListTagsForResourceInput, opts ...request.Option) (*dlm.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dlm",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*dlm.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DLMAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*dlm.ListTagsForResourceOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *dlm.TagResourceInput, opts ...request.Option) (*dlm.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dlm",
		Action:  "TagResource",
		Input:   input,
		Output:  (*dlm.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DLMAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*dlm.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *dlm.UntagResourceInput, opts ...request.Option) (*dlm.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dlm",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*dlm.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DLMAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*dlm.UntagResourceOutput), req.Error
}

func (c *Client) UpdateLifecyclePolicyWithContext(ctx context.Context, input *dlm.UpdateLifecyclePolicyInput, opts ...request.Option) (*dlm.UpdateLifecyclePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "dlm",
		Action:  "UpdateLifecyclePolicy",
		Input:   input,
		Output:  (*dlm.UpdateLifecyclePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DLMAPI.UpdateLifecyclePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*dlm.UpdateLifecyclePolicyOutput), req.Error
}
