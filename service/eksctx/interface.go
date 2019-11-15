// Code generated by internal/generate/main.go. DO NOT EDIT.

package eksctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/glassechidna/awsctx"
)

type EKS interface {
	CreateClusterWithContext(ctx context.Context, input *eks.CreateClusterInput, opts ...request.Option) (*eks.CreateClusterOutput, error)
	CreateNodegroupWithContext(ctx context.Context, input *eks.CreateNodegroupInput, opts ...request.Option) (*eks.CreateNodegroupOutput, error)
	DeleteClusterWithContext(ctx context.Context, input *eks.DeleteClusterInput, opts ...request.Option) (*eks.DeleteClusterOutput, error)
	DeleteNodegroupWithContext(ctx context.Context, input *eks.DeleteNodegroupInput, opts ...request.Option) (*eks.DeleteNodegroupOutput, error)
	DescribeClusterWithContext(ctx context.Context, input *eks.DescribeClusterInput, opts ...request.Option) (*eks.DescribeClusterOutput, error)
	DescribeNodegroupWithContext(ctx context.Context, input *eks.DescribeNodegroupInput, opts ...request.Option) (*eks.DescribeNodegroupOutput, error)
	DescribeUpdateWithContext(ctx context.Context, input *eks.DescribeUpdateInput, opts ...request.Option) (*eks.DescribeUpdateOutput, error)
	ListClustersWithContext(ctx context.Context, input *eks.ListClustersInput, opts ...request.Option) (*eks.ListClustersOutput, error)
	ListClustersPagesWithContext(ctx context.Context, input *eks.ListClustersInput, cb func(*eks.ListClustersOutput, bool) bool, opts ...request.Option) error
	ListNodegroupsWithContext(ctx context.Context, input *eks.ListNodegroupsInput, opts ...request.Option) (*eks.ListNodegroupsOutput, error)
	ListNodegroupsPagesWithContext(ctx context.Context, input *eks.ListNodegroupsInput, cb func(*eks.ListNodegroupsOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *eks.ListTagsForResourceInput, opts ...request.Option) (*eks.ListTagsForResourceOutput, error)
	ListUpdatesWithContext(ctx context.Context, input *eks.ListUpdatesInput, opts ...request.Option) (*eks.ListUpdatesOutput, error)
	ListUpdatesPagesWithContext(ctx context.Context, input *eks.ListUpdatesInput, cb func(*eks.ListUpdatesOutput, bool) bool, opts ...request.Option) error
	TagResourceWithContext(ctx context.Context, input *eks.TagResourceInput, opts ...request.Option) (*eks.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *eks.UntagResourceInput, opts ...request.Option) (*eks.UntagResourceOutput, error)
	UpdateClusterConfigWithContext(ctx context.Context, input *eks.UpdateClusterConfigInput, opts ...request.Option) (*eks.UpdateClusterConfigOutput, error)
	UpdateClusterVersionWithContext(ctx context.Context, input *eks.UpdateClusterVersionInput, opts ...request.Option) (*eks.UpdateClusterVersionOutput, error)
	UpdateNodegroupConfigWithContext(ctx context.Context, input *eks.UpdateNodegroupConfigInput, opts ...request.Option) (*eks.UpdateNodegroupConfigOutput, error)
	UpdateNodegroupVersionWithContext(ctx context.Context, input *eks.UpdateNodegroupVersionInput, opts ...request.Option) (*eks.UpdateNodegroupVersionOutput, error)
}

type Client struct {
	eksiface.EKSAPI
	Contexter awsctx.Contexter
}

func New(base eksiface.EKSAPI, ctxer awsctx.Contexter) EKS {
	return &Client{
		EKSAPI: base,
		Contexter: ctxer,
	}
}

var _ EKS = (*eks.EKS)(nil)
var _ EKS = (*Client)(nil)

func (c *Client) CreateClusterWithContext(ctx context.Context, input *eks.CreateClusterInput, opts ...request.Option) (*eks.CreateClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "CreateCluster",
		Input:   input,
		Output:  (*eks.CreateClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.CreateClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.CreateClusterOutput), req.Error
}

func (c *Client) CreateNodegroupWithContext(ctx context.Context, input *eks.CreateNodegroupInput, opts ...request.Option) (*eks.CreateNodegroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "CreateNodegroup",
		Input:   input,
		Output:  (*eks.CreateNodegroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.CreateNodegroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.CreateNodegroupOutput), req.Error
}

func (c *Client) DeleteClusterWithContext(ctx context.Context, input *eks.DeleteClusterInput, opts ...request.Option) (*eks.DeleteClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "DeleteCluster",
		Input:   input,
		Output:  (*eks.DeleteClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.DeleteClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.DeleteClusterOutput), req.Error
}

func (c *Client) DeleteNodegroupWithContext(ctx context.Context, input *eks.DeleteNodegroupInput, opts ...request.Option) (*eks.DeleteNodegroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "DeleteNodegroup",
		Input:   input,
		Output:  (*eks.DeleteNodegroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.DeleteNodegroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.DeleteNodegroupOutput), req.Error
}

func (c *Client) DescribeClusterWithContext(ctx context.Context, input *eks.DescribeClusterInput, opts ...request.Option) (*eks.DescribeClusterOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "DescribeCluster",
		Input:   input,
		Output:  (*eks.DescribeClusterOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.DescribeClusterWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.DescribeClusterOutput), req.Error
}

func (c *Client) DescribeNodegroupWithContext(ctx context.Context, input *eks.DescribeNodegroupInput, opts ...request.Option) (*eks.DescribeNodegroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "DescribeNodegroup",
		Input:   input,
		Output:  (*eks.DescribeNodegroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.DescribeNodegroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.DescribeNodegroupOutput), req.Error
}

func (c *Client) DescribeUpdateWithContext(ctx context.Context, input *eks.DescribeUpdateInput, opts ...request.Option) (*eks.DescribeUpdateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "DescribeUpdate",
		Input:   input,
		Output:  (*eks.DescribeUpdateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.DescribeUpdateWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.DescribeUpdateOutput), req.Error
}

func (c *Client) ListClustersWithContext(ctx context.Context, input *eks.ListClustersInput, opts ...request.Option) (*eks.ListClustersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "ListClusters",
		Input:   input,
		Output:  (*eks.ListClustersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.ListClustersWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.ListClustersOutput), req.Error
}

func (c *Client) ListClustersPagesWithContext(ctx context.Context, input *eks.ListClustersInput, cb func(*eks.ListClustersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "ListClusters",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.EKSAPI.ListClustersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListNodegroupsWithContext(ctx context.Context, input *eks.ListNodegroupsInput, opts ...request.Option) (*eks.ListNodegroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "ListNodegroups",
		Input:   input,
		Output:  (*eks.ListNodegroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.ListNodegroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.ListNodegroupsOutput), req.Error
}

func (c *Client) ListNodegroupsPagesWithContext(ctx context.Context, input *eks.ListNodegroupsInput, cb func(*eks.ListNodegroupsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "ListNodegroups",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.EKSAPI.ListNodegroupsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *eks.ListTagsForResourceInput, opts ...request.Option) (*eks.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*eks.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListUpdatesWithContext(ctx context.Context, input *eks.ListUpdatesInput, opts ...request.Option) (*eks.ListUpdatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "ListUpdates",
		Input:   input,
		Output:  (*eks.ListUpdatesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.ListUpdatesWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.ListUpdatesOutput), req.Error
}

func (c *Client) ListUpdatesPagesWithContext(ctx context.Context, input *eks.ListUpdatesInput, cb func(*eks.ListUpdatesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "ListUpdates",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.EKSAPI.ListUpdatesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *eks.TagResourceInput, opts ...request.Option) (*eks.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "TagResource",
		Input:   input,
		Output:  (*eks.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *eks.UntagResourceInput, opts ...request.Option) (*eks.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*eks.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.UntagResourceOutput), req.Error
}

func (c *Client) UpdateClusterConfigWithContext(ctx context.Context, input *eks.UpdateClusterConfigInput, opts ...request.Option) (*eks.UpdateClusterConfigOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "UpdateClusterConfig",
		Input:   input,
		Output:  (*eks.UpdateClusterConfigOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.UpdateClusterConfigWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.UpdateClusterConfigOutput), req.Error
}

func (c *Client) UpdateClusterVersionWithContext(ctx context.Context, input *eks.UpdateClusterVersionInput, opts ...request.Option) (*eks.UpdateClusterVersionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "UpdateClusterVersion",
		Input:   input,
		Output:  (*eks.UpdateClusterVersionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.UpdateClusterVersionWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.UpdateClusterVersionOutput), req.Error
}

func (c *Client) UpdateNodegroupConfigWithContext(ctx context.Context, input *eks.UpdateNodegroupConfigInput, opts ...request.Option) (*eks.UpdateNodegroupConfigOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "UpdateNodegroupConfig",
		Input:   input,
		Output:  (*eks.UpdateNodegroupConfigOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.UpdateNodegroupConfigWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.UpdateNodegroupConfigOutput), req.Error
}

func (c *Client) UpdateNodegroupVersionWithContext(ctx context.Context, input *eks.UpdateNodegroupVersionInput, opts ...request.Option) (*eks.UpdateNodegroupVersionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eks",
		Action:  "UpdateNodegroupVersion",
		Input:   input,
		Output:  (*eks.UpdateNodegroupVersionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EKSAPI.UpdateNodegroupVersionWithContext(ctx, input, opts...)
	})

	return req.Output.(*eks.UpdateNodegroupVersionOutput), req.Error
}
