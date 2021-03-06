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
	DescribeReportCreationWithContext(ctx context.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput, opts ...request.Option) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error)
	GetComplianceSummaryWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput, opts ...request.Option) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error)
	GetComplianceSummaryPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput, cb func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) bool, opts ...request.Option) error
	GetResourcesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetResourcesInput, opts ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error)
	GetResourcesPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetResourcesInput, cb func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool, opts ...request.Option) error
	GetTagKeysWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagKeysInput, opts ...request.Option) (*resourcegroupstaggingapi.GetTagKeysOutput, error)
	GetTagKeysPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagKeysInput, cb func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool, opts ...request.Option) error
	GetTagValuesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagValuesInput, opts ...request.Option) (*resourcegroupstaggingapi.GetTagValuesOutput, error)
	GetTagValuesPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagValuesInput, cb func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool, opts ...request.Option) error
	StartReportCreationWithContext(ctx context.Context, input *resourcegroupstaggingapi.StartReportCreationInput, opts ...request.Option) (*resourcegroupstaggingapi.StartReportCreationOutput, error)
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

func (c *Client) DescribeReportCreationWithContext(ctx context.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput, opts ...request.Option) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "DescribeReportCreation",
		Input:   input,
		Output:  (*resourcegroupstaggingapi.DescribeReportCreationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsTaggingAPIAPI.DescribeReportCreationWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroupstaggingapi.DescribeReportCreationOutput), req.Error
}

func (c *Client) GetComplianceSummaryWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput, opts ...request.Option) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "GetComplianceSummary",
		Input:   input,
		Output:  (*resourcegroupstaggingapi.GetComplianceSummaryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsTaggingAPIAPI.GetComplianceSummaryWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroupstaggingapi.GetComplianceSummaryOutput), req.Error
}

func (c *Client) GetComplianceSummaryPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput, cb func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "GetComplianceSummary",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ResourceGroupsTaggingAPIAPI.GetComplianceSummaryPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

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

func (c *Client) GetResourcesPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetResourcesInput, cb func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "GetResources",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ResourceGroupsTaggingAPIAPI.GetResourcesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
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

func (c *Client) GetTagKeysPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagKeysInput, cb func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "GetTagKeys",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ResourceGroupsTaggingAPIAPI.GetTagKeysPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
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

func (c *Client) GetTagValuesPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagValuesInput, cb func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "GetTagValues",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ResourceGroupsTaggingAPIAPI.GetTagValuesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) StartReportCreationWithContext(ctx context.Context, input *resourcegroupstaggingapi.StartReportCreationInput, opts ...request.Option) (*resourcegroupstaggingapi.StartReportCreationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroupstaggingapi",
		Action:  "StartReportCreation",
		Input:   input,
		Output:  (*resourcegroupstaggingapi.StartReportCreationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsTaggingAPIAPI.StartReportCreationWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroupstaggingapi.StartReportCreationOutput), req.Error
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
