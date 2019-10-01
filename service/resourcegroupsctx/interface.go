// Code generated by internal/generate/main.go. DO NOT EDIT.

package resourcegroupsctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"github.com/aws/aws-sdk-go/service/resourcegroups/resourcegroupsiface"
	"github.com/glassechidna/awsctx"
)

type ResourceGroups interface {
	CreateGroupWithContext(ctx context.Context, input *resourcegroups.CreateGroupInput, opts ...request.Option) (*resourcegroups.CreateGroupOutput, error)
	DeleteGroupWithContext(ctx context.Context, input *resourcegroups.DeleteGroupInput, opts ...request.Option) (*resourcegroups.DeleteGroupOutput, error)
	GetGroupWithContext(ctx context.Context, input *resourcegroups.GetGroupInput, opts ...request.Option) (*resourcegroups.GetGroupOutput, error)
	GetGroupQueryWithContext(ctx context.Context, input *resourcegroups.GetGroupQueryInput, opts ...request.Option) (*resourcegroups.GetGroupQueryOutput, error)
	GetTagsWithContext(ctx context.Context, input *resourcegroups.GetTagsInput, opts ...request.Option) (*resourcegroups.GetTagsOutput, error)
	ListGroupResourcesWithContext(ctx context.Context, input *resourcegroups.ListGroupResourcesInput, opts ...request.Option) (*resourcegroups.ListGroupResourcesOutput, error)
	ListGroupsWithContext(ctx context.Context, input *resourcegroups.ListGroupsInput, opts ...request.Option) (*resourcegroups.ListGroupsOutput, error)
	SearchResourcesWithContext(ctx context.Context, input *resourcegroups.SearchResourcesInput, opts ...request.Option) (*resourcegroups.SearchResourcesOutput, error)
	TagWithContext(ctx context.Context, input *resourcegroups.TagInput, opts ...request.Option) (*resourcegroups.TagOutput, error)
	UntagWithContext(ctx context.Context, input *resourcegroups.UntagInput, opts ...request.Option) (*resourcegroups.UntagOutput, error)
	UpdateGroupWithContext(ctx context.Context, input *resourcegroups.UpdateGroupInput, opts ...request.Option) (*resourcegroups.UpdateGroupOutput, error)
	UpdateGroupQueryWithContext(ctx context.Context, input *resourcegroups.UpdateGroupQueryInput, opts ...request.Option) (*resourcegroups.UpdateGroupQueryOutput, error)
}

type Client struct {
	resourcegroupsiface.ResourceGroupsAPI
	Contexter awsctx.Contexter
}

func New(base resourcegroupsiface.ResourceGroupsAPI, ctxer awsctx.Contexter) ResourceGroups {
	return &Client{
		ResourceGroupsAPI: base,
		Contexter: ctxer,
	}
}

var _ ResourceGroups = (*resourcegroups.ResourceGroups)(nil)
var _ ResourceGroups = (*Client)(nil)

func (c *Client) CreateGroupWithContext(ctx context.Context, input *resourcegroups.CreateGroupInput, opts ...request.Option) (*resourcegroups.CreateGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "CreateGroupWithContext",
		Input:   input,
		Output:  (*resourcegroups.CreateGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.CreateGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.CreateGroupOutput), req.Error
}

func (c *Client) DeleteGroupWithContext(ctx context.Context, input *resourcegroups.DeleteGroupInput, opts ...request.Option) (*resourcegroups.DeleteGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "DeleteGroupWithContext",
		Input:   input,
		Output:  (*resourcegroups.DeleteGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.DeleteGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.DeleteGroupOutput), req.Error
}

func (c *Client) GetGroupWithContext(ctx context.Context, input *resourcegroups.GetGroupInput, opts ...request.Option) (*resourcegroups.GetGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "GetGroupWithContext",
		Input:   input,
		Output:  (*resourcegroups.GetGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.GetGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.GetGroupOutput), req.Error
}

func (c *Client) GetGroupQueryWithContext(ctx context.Context, input *resourcegroups.GetGroupQueryInput, opts ...request.Option) (*resourcegroups.GetGroupQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "GetGroupQueryWithContext",
		Input:   input,
		Output:  (*resourcegroups.GetGroupQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.GetGroupQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.GetGroupQueryOutput), req.Error
}

func (c *Client) GetTagsWithContext(ctx context.Context, input *resourcegroups.GetTagsInput, opts ...request.Option) (*resourcegroups.GetTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "GetTagsWithContext",
		Input:   input,
		Output:  (*resourcegroups.GetTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.GetTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.GetTagsOutput), req.Error
}

func (c *Client) ListGroupResourcesWithContext(ctx context.Context, input *resourcegroups.ListGroupResourcesInput, opts ...request.Option) (*resourcegroups.ListGroupResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "ListGroupResourcesWithContext",
		Input:   input,
		Output:  (*resourcegroups.ListGroupResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.ListGroupResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.ListGroupResourcesOutput), req.Error
}

func (c *Client) ListGroupsWithContext(ctx context.Context, input *resourcegroups.ListGroupsInput, opts ...request.Option) (*resourcegroups.ListGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "ListGroupsWithContext",
		Input:   input,
		Output:  (*resourcegroups.ListGroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.ListGroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.ListGroupsOutput), req.Error
}

func (c *Client) SearchResourcesWithContext(ctx context.Context, input *resourcegroups.SearchResourcesInput, opts ...request.Option) (*resourcegroups.SearchResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "SearchResourcesWithContext",
		Input:   input,
		Output:  (*resourcegroups.SearchResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.SearchResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.SearchResourcesOutput), req.Error
}

func (c *Client) TagWithContext(ctx context.Context, input *resourcegroups.TagInput, opts ...request.Option) (*resourcegroups.TagOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "TagWithContext",
		Input:   input,
		Output:  (*resourcegroups.TagOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.TagWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.TagOutput), req.Error
}

func (c *Client) UntagWithContext(ctx context.Context, input *resourcegroups.UntagInput, opts ...request.Option) (*resourcegroups.UntagOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "UntagWithContext",
		Input:   input,
		Output:  (*resourcegroups.UntagOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.UntagWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.UntagOutput), req.Error
}

func (c *Client) UpdateGroupWithContext(ctx context.Context, input *resourcegroups.UpdateGroupInput, opts ...request.Option) (*resourcegroups.UpdateGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "UpdateGroupWithContext",
		Input:   input,
		Output:  (*resourcegroups.UpdateGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.UpdateGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.UpdateGroupOutput), req.Error
}

func (c *Client) UpdateGroupQueryWithContext(ctx context.Context, input *resourcegroups.UpdateGroupQueryInput, opts ...request.Option) (*resourcegroups.UpdateGroupQueryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "resourcegroups",
		Action:  "UpdateGroupQueryWithContext",
		Input:   input,
		Output:  (*resourcegroups.UpdateGroupQueryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ResourceGroupsAPI.UpdateGroupQueryWithContext(ctx, input, opts...)
	})

	return req.Output.(*resourcegroups.UpdateGroupQueryOutput), req.Error
}
