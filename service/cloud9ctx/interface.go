// Code generated by internal/generate/main.go. DO NOT EDIT.

package cloud9ctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloud9"
	"github.com/aws/aws-sdk-go/service/cloud9/cloud9iface"
	"github.com/glassechidna/awsctx"
)

type Cloud9 interface {
	CreateEnvironmentEC2WithContext(ctx context.Context, input *cloud9.CreateEnvironmentEC2Input, opts ...request.Option) (*cloud9.CreateEnvironmentEC2Output, error)
	CreateEnvironmentMembershipWithContext(ctx context.Context, input *cloud9.CreateEnvironmentMembershipInput, opts ...request.Option) (*cloud9.CreateEnvironmentMembershipOutput, error)
	DeleteEnvironmentWithContext(ctx context.Context, input *cloud9.DeleteEnvironmentInput, opts ...request.Option) (*cloud9.DeleteEnvironmentOutput, error)
	DeleteEnvironmentMembershipWithContext(ctx context.Context, input *cloud9.DeleteEnvironmentMembershipInput, opts ...request.Option) (*cloud9.DeleteEnvironmentMembershipOutput, error)
	DescribeEnvironmentMembershipsWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentMembershipsInput, opts ...request.Option) (*cloud9.DescribeEnvironmentMembershipsOutput, error)
	DescribeEnvironmentMembershipsPagesWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentMembershipsInput, cb func(*cloud9.DescribeEnvironmentMembershipsOutput, bool) bool, opts ...request.Option) error
	DescribeEnvironmentStatusWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentStatusInput, opts ...request.Option) (*cloud9.DescribeEnvironmentStatusOutput, error)
	DescribeEnvironmentsWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentsInput, opts ...request.Option) (*cloud9.DescribeEnvironmentsOutput, error)
	ListEnvironmentsWithContext(ctx context.Context, input *cloud9.ListEnvironmentsInput, opts ...request.Option) (*cloud9.ListEnvironmentsOutput, error)
	ListEnvironmentsPagesWithContext(ctx context.Context, input *cloud9.ListEnvironmentsInput, cb func(*cloud9.ListEnvironmentsOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *cloud9.ListTagsForResourceInput, opts ...request.Option) (*cloud9.ListTagsForResourceOutput, error)
	TagResourceWithContext(ctx context.Context, input *cloud9.TagResourceInput, opts ...request.Option) (*cloud9.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *cloud9.UntagResourceInput, opts ...request.Option) (*cloud9.UntagResourceOutput, error)
	UpdateEnvironmentWithContext(ctx context.Context, input *cloud9.UpdateEnvironmentInput, opts ...request.Option) (*cloud9.UpdateEnvironmentOutput, error)
	UpdateEnvironmentMembershipWithContext(ctx context.Context, input *cloud9.UpdateEnvironmentMembershipInput, opts ...request.Option) (*cloud9.UpdateEnvironmentMembershipOutput, error)
}

type Client struct {
	cloud9iface.Cloud9API
	Contexter awsctx.Contexter
}

func New(base cloud9iface.Cloud9API, ctxer awsctx.Contexter) Cloud9 {
	return &Client{
		Cloud9API: base,
		Contexter: ctxer,
	}
}

var _ Cloud9 = (*cloud9.Cloud9)(nil)
var _ Cloud9 = (*Client)(nil)

func (c *Client) CreateEnvironmentEC2WithContext(ctx context.Context, input *cloud9.CreateEnvironmentEC2Input, opts ...request.Option) (*cloud9.CreateEnvironmentEC2Output, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "CreateEnvironmentEC2",
		Input:   input,
		Output:  (*cloud9.CreateEnvironmentEC2Output)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.CreateEnvironmentEC2WithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.CreateEnvironmentEC2Output), req.Error
}

func (c *Client) CreateEnvironmentMembershipWithContext(ctx context.Context, input *cloud9.CreateEnvironmentMembershipInput, opts ...request.Option) (*cloud9.CreateEnvironmentMembershipOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "CreateEnvironmentMembership",
		Input:   input,
		Output:  (*cloud9.CreateEnvironmentMembershipOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.CreateEnvironmentMembershipWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.CreateEnvironmentMembershipOutput), req.Error
}

func (c *Client) DeleteEnvironmentWithContext(ctx context.Context, input *cloud9.DeleteEnvironmentInput, opts ...request.Option) (*cloud9.DeleteEnvironmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "DeleteEnvironment",
		Input:   input,
		Output:  (*cloud9.DeleteEnvironmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.DeleteEnvironmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.DeleteEnvironmentOutput), req.Error
}

func (c *Client) DeleteEnvironmentMembershipWithContext(ctx context.Context, input *cloud9.DeleteEnvironmentMembershipInput, opts ...request.Option) (*cloud9.DeleteEnvironmentMembershipOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "DeleteEnvironmentMembership",
		Input:   input,
		Output:  (*cloud9.DeleteEnvironmentMembershipOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.DeleteEnvironmentMembershipWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.DeleteEnvironmentMembershipOutput), req.Error
}

func (c *Client) DescribeEnvironmentMembershipsWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentMembershipsInput, opts ...request.Option) (*cloud9.DescribeEnvironmentMembershipsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "DescribeEnvironmentMemberships",
		Input:   input,
		Output:  (*cloud9.DescribeEnvironmentMembershipsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.DescribeEnvironmentMembershipsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.DescribeEnvironmentMembershipsOutput), req.Error
}

func (c *Client) DescribeEnvironmentMembershipsPagesWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentMembershipsInput, cb func(*cloud9.DescribeEnvironmentMembershipsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "DescribeEnvironmentMemberships",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.Cloud9API.DescribeEnvironmentMembershipsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeEnvironmentStatusWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentStatusInput, opts ...request.Option) (*cloud9.DescribeEnvironmentStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "DescribeEnvironmentStatus",
		Input:   input,
		Output:  (*cloud9.DescribeEnvironmentStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.DescribeEnvironmentStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.DescribeEnvironmentStatusOutput), req.Error
}

func (c *Client) DescribeEnvironmentsWithContext(ctx context.Context, input *cloud9.DescribeEnvironmentsInput, opts ...request.Option) (*cloud9.DescribeEnvironmentsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "DescribeEnvironments",
		Input:   input,
		Output:  (*cloud9.DescribeEnvironmentsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.DescribeEnvironmentsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.DescribeEnvironmentsOutput), req.Error
}

func (c *Client) ListEnvironmentsWithContext(ctx context.Context, input *cloud9.ListEnvironmentsInput, opts ...request.Option) (*cloud9.ListEnvironmentsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "ListEnvironments",
		Input:   input,
		Output:  (*cloud9.ListEnvironmentsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.ListEnvironmentsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.ListEnvironmentsOutput), req.Error
}

func (c *Client) ListEnvironmentsPagesWithContext(ctx context.Context, input *cloud9.ListEnvironmentsInput, cb func(*cloud9.ListEnvironmentsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "ListEnvironments",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.Cloud9API.ListEnvironmentsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *cloud9.ListTagsForResourceInput, opts ...request.Option) (*cloud9.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*cloud9.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.ListTagsForResourceOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *cloud9.TagResourceInput, opts ...request.Option) (*cloud9.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "TagResource",
		Input:   input,
		Output:  (*cloud9.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *cloud9.UntagResourceInput, opts ...request.Option) (*cloud9.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*cloud9.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.UntagResourceOutput), req.Error
}

func (c *Client) UpdateEnvironmentWithContext(ctx context.Context, input *cloud9.UpdateEnvironmentInput, opts ...request.Option) (*cloud9.UpdateEnvironmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "UpdateEnvironment",
		Input:   input,
		Output:  (*cloud9.UpdateEnvironmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.UpdateEnvironmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.UpdateEnvironmentOutput), req.Error
}

func (c *Client) UpdateEnvironmentMembershipWithContext(ctx context.Context, input *cloud9.UpdateEnvironmentMembershipInput, opts ...request.Option) (*cloud9.UpdateEnvironmentMembershipOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cloud9",
		Action:  "UpdateEnvironmentMembership",
		Input:   input,
		Output:  (*cloud9.UpdateEnvironmentMembershipOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Cloud9API.UpdateEnvironmentMembershipWithContext(ctx, input, opts...)
	})

	return req.Output.(*cloud9.UpdateEnvironmentMembershipOutput), req.Error
}
