// Code generated by internal/generate/main.go. DO NOT EDIT.

package resourcegroupstaggingapictx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi/resourcegroupstaggingapiiface"
	"github.com/glassechidna/awsctx"
)

type ResourceGroupsTaggingAPI interface {
	GetResourcesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetResourcesInput, opts ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error)
	GetTagKeysWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagKeysInput, opts ...request.Option) (*resourcegroupstaggingapi.GetTagKeysOutput, error)
	GetTagValuesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagValuesInput, opts ...request.Option) (*resourcegroupstaggingapi.GetTagValuesOutput, error)
	TagResourcesWithContext(ctx context.Context, input *resourcegroupstaggingapi.TagResourcesInput, opts ...request.Option) (*resourcegroupstaggingapi.TagResourcesOutput, error)
	UntagResourcesWithContext(ctx context.Context, input *resourcegroupstaggingapi.UntagResourcesInput, opts ...request.Option) (*resourcegroupstaggingapi.UntagResourcesOutput, error)
}

type Client struct {
	resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI
	Contexter awsctx.Contexter
}

func New(base resourcegroupstaggingapiiface.ResourceGroupsTaggingAPIAPI, ctxer awsctx.Contexter) ResourceGroupsTaggingAPI {
	return &Client{
		ResourceGroupsTaggingAPIAPI: base,
		Contexter: ctxer,
	}
}

var _ ResourceGroupsTaggingAPI = (*resourcegroupstaggingapi.ResourceGroupsTaggingAPI)(nil)
var _ ResourceGroupsTaggingAPI = (*Client)(nil)

func (c *Client) GetResourcesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetResourcesInput, opts ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "GetResources",
		Input:   input,
		Output:  (*resourcegroupstaggingapi.GetResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsTaggingAPIAPI.GetResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroupstaggingapi.GetResourcesOutput), req.Error
}

func (c *Client) GetTagKeysWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagKeysInput, opts ...request.Option) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "GetTagKeys",
		Input:   input,
		Output:  (*resourcegroupstaggingapi.GetTagKeysOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsTaggingAPIAPI.GetTagKeysWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroupstaggingapi.GetTagKeysOutput), req.Error
}

func (c *Client) GetTagValuesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagValuesInput, opts ...request.Option) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "GetTagValues",
		Input:   input,
		Output:  (*resourcegroupstaggingapi.GetTagValuesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsTaggingAPIAPI.GetTagValuesWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroupstaggingapi.GetTagValuesOutput), req.Error
}

func (c *Client) TagResourcesWithContext(ctx context.Context, input *resourcegroupstaggingapi.TagResourcesInput, opts ...request.Option) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "TagResources",
		Input:   input,
		Output:  (*resourcegroupstaggingapi.TagResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsTaggingAPIAPI.TagResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroupstaggingapi.TagResourcesOutput), req.Error
}

func (c *Client) UntagResourcesWithContext(ctx context.Context, input *resourcegroupstaggingapi.UntagResourcesInput, opts ...request.Option) (*resourcegroupstaggingapi.UntagResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "UntagResources",
		Input:   input,
		Output:  (*resourcegroupstaggingapi.UntagResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsTaggingAPIAPI.UntagResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroupstaggingapi.UntagResourcesOutput), req.Error
}
