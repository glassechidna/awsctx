// Code generated by internal/generate/main.go. DO NOT EDIT.

package iot1clickprojectsctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go/service/iot1clickprojects/iot1clickprojectsiface"
	"github.com/glassechidna/awsctx"
)

type IoT1ClickProjects interface {
	AssociateDeviceWithPlacementWithContext(ctx context.Context, input *iot1clickprojects.AssociateDeviceWithPlacementInput, opts ...request.Option) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error)
	CreatePlacementWithContext(ctx context.Context, input *iot1clickprojects.CreatePlacementInput, opts ...request.Option) (*iot1clickprojects.CreatePlacementOutput, error)
	CreateProjectWithContext(ctx context.Context, input *iot1clickprojects.CreateProjectInput, opts ...request.Option) (*iot1clickprojects.CreateProjectOutput, error)
	DeletePlacementWithContext(ctx context.Context, input *iot1clickprojects.DeletePlacementInput, opts ...request.Option) (*iot1clickprojects.DeletePlacementOutput, error)
	DeleteProjectWithContext(ctx context.Context, input *iot1clickprojects.DeleteProjectInput, opts ...request.Option) (*iot1clickprojects.DeleteProjectOutput, error)
	DescribePlacementWithContext(ctx context.Context, input *iot1clickprojects.DescribePlacementInput, opts ...request.Option) (*iot1clickprojects.DescribePlacementOutput, error)
	DescribeProjectWithContext(ctx context.Context, input *iot1clickprojects.DescribeProjectInput, opts ...request.Option) (*iot1clickprojects.DescribeProjectOutput, error)
	DisassociateDeviceFromPlacementWithContext(ctx context.Context, input *iot1clickprojects.DisassociateDeviceFromPlacementInput, opts ...request.Option) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error)
	GetDevicesInPlacementWithContext(ctx context.Context, input *iot1clickprojects.GetDevicesInPlacementInput, opts ...request.Option) (*iot1clickprojects.GetDevicesInPlacementOutput, error)
	ListPlacementsWithContext(ctx context.Context, input *iot1clickprojects.ListPlacementsInput, opts ...request.Option) (*iot1clickprojects.ListPlacementsOutput, error)
	ListProjectsWithContext(ctx context.Context, input *iot1clickprojects.ListProjectsInput, opts ...request.Option) (*iot1clickprojects.ListProjectsOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *iot1clickprojects.ListTagsForResourceInput, opts ...request.Option) (*iot1clickprojects.ListTagsForResourceOutput, error)
	TagResourceWithContext(ctx context.Context, input *iot1clickprojects.TagResourceInput, opts ...request.Option) (*iot1clickprojects.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *iot1clickprojects.UntagResourceInput, opts ...request.Option) (*iot1clickprojects.UntagResourceOutput, error)
	UpdatePlacementWithContext(ctx context.Context, input *iot1clickprojects.UpdatePlacementInput, opts ...request.Option) (*iot1clickprojects.UpdatePlacementOutput, error)
	UpdateProjectWithContext(ctx context.Context, input *iot1clickprojects.UpdateProjectInput, opts ...request.Option) (*iot1clickprojects.UpdateProjectOutput, error)
}

type Client struct {
	iot1clickprojectsiface.IoT1ClickProjectsAPI
	Contexter awsctx.Contexter
}

func New(base iot1clickprojectsiface.IoT1ClickProjectsAPI, ctxer awsctx.Contexter) IoT1ClickProjects {
	return &Client{
		IoT1ClickProjectsAPI: base,
		Contexter: ctxer,
	}
}

var _ IoT1ClickProjects = (*iot1clickprojects.IoT1ClickProjects)(nil)
var _ IoT1ClickProjects = (*Client)(nil)

func (c *Client) AssociateDeviceWithPlacementWithContext(ctx context.Context, input *iot1clickprojects.AssociateDeviceWithPlacementInput, opts ...request.Option) (*iot1clickprojects.AssociateDeviceWithPlacementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "AssociateDeviceWithPlacement",
		Input:   input,
		Output:  (*iot1clickprojects.AssociateDeviceWithPlacementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.AssociateDeviceWithPlacementWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.AssociateDeviceWithPlacementOutput), req.Error
}

func (c *Client) CreatePlacementWithContext(ctx context.Context, input *iot1clickprojects.CreatePlacementInput, opts ...request.Option) (*iot1clickprojects.CreatePlacementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "CreatePlacement",
		Input:   input,
		Output:  (*iot1clickprojects.CreatePlacementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.CreatePlacementWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.CreatePlacementOutput), req.Error
}

func (c *Client) CreateProjectWithContext(ctx context.Context, input *iot1clickprojects.CreateProjectInput, opts ...request.Option) (*iot1clickprojects.CreateProjectOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "CreateProject",
		Input:   input,
		Output:  (*iot1clickprojects.CreateProjectOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.CreateProjectWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.CreateProjectOutput), req.Error
}

func (c *Client) DeletePlacementWithContext(ctx context.Context, input *iot1clickprojects.DeletePlacementInput, opts ...request.Option) (*iot1clickprojects.DeletePlacementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "DeletePlacement",
		Input:   input,
		Output:  (*iot1clickprojects.DeletePlacementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.DeletePlacementWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.DeletePlacementOutput), req.Error
}

func (c *Client) DeleteProjectWithContext(ctx context.Context, input *iot1clickprojects.DeleteProjectInput, opts ...request.Option) (*iot1clickprojects.DeleteProjectOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "DeleteProject",
		Input:   input,
		Output:  (*iot1clickprojects.DeleteProjectOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.DeleteProjectWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.DeleteProjectOutput), req.Error
}

func (c *Client) DescribePlacementWithContext(ctx context.Context, input *iot1clickprojects.DescribePlacementInput, opts ...request.Option) (*iot1clickprojects.DescribePlacementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "DescribePlacement",
		Input:   input,
		Output:  (*iot1clickprojects.DescribePlacementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.DescribePlacementWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.DescribePlacementOutput), req.Error
}

func (c *Client) DescribeProjectWithContext(ctx context.Context, input *iot1clickprojects.DescribeProjectInput, opts ...request.Option) (*iot1clickprojects.DescribeProjectOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "DescribeProject",
		Input:   input,
		Output:  (*iot1clickprojects.DescribeProjectOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.DescribeProjectWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.DescribeProjectOutput), req.Error
}

func (c *Client) DisassociateDeviceFromPlacementWithContext(ctx context.Context, input *iot1clickprojects.DisassociateDeviceFromPlacementInput, opts ...request.Option) (*iot1clickprojects.DisassociateDeviceFromPlacementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "DisassociateDeviceFromPlacement",
		Input:   input,
		Output:  (*iot1clickprojects.DisassociateDeviceFromPlacementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.DisassociateDeviceFromPlacementWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.DisassociateDeviceFromPlacementOutput), req.Error
}

func (c *Client) GetDevicesInPlacementWithContext(ctx context.Context, input *iot1clickprojects.GetDevicesInPlacementInput, opts ...request.Option) (*iot1clickprojects.GetDevicesInPlacementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "GetDevicesInPlacement",
		Input:   input,
		Output:  (*iot1clickprojects.GetDevicesInPlacementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.GetDevicesInPlacementWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.GetDevicesInPlacementOutput), req.Error
}

func (c *Client) ListPlacementsWithContext(ctx context.Context, input *iot1clickprojects.ListPlacementsInput, opts ...request.Option) (*iot1clickprojects.ListPlacementsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "ListPlacements",
		Input:   input,
		Output:  (*iot1clickprojects.ListPlacementsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.ListPlacementsWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.ListPlacementsOutput), req.Error
}

func (c *Client) ListProjectsWithContext(ctx context.Context, input *iot1clickprojects.ListProjectsInput, opts ...request.Option) (*iot1clickprojects.ListProjectsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "ListProjects",
		Input:   input,
		Output:  (*iot1clickprojects.ListProjectsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.ListProjectsWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.ListProjectsOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *iot1clickprojects.ListTagsForResourceInput, opts ...request.Option) (*iot1clickprojects.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*iot1clickprojects.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.ListTagsForResourceOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *iot1clickprojects.TagResourceInput, opts ...request.Option) (*iot1clickprojects.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "TagResource",
		Input:   input,
		Output:  (*iot1clickprojects.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *iot1clickprojects.UntagResourceInput, opts ...request.Option) (*iot1clickprojects.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*iot1clickprojects.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.UntagResourceOutput), req.Error
}

func (c *Client) UpdatePlacementWithContext(ctx context.Context, input *iot1clickprojects.UpdatePlacementInput, opts ...request.Option) (*iot1clickprojects.UpdatePlacementOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "UpdatePlacement",
		Input:   input,
		Output:  (*iot1clickprojects.UpdatePlacementOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.UpdatePlacementWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.UpdatePlacementOutput), req.Error
}

func (c *Client) UpdateProjectWithContext(ctx context.Context, input *iot1clickprojects.UpdateProjectInput, opts ...request.Option) (*iot1clickprojects.UpdateProjectOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "iot1clickprojects",
		Action:  "UpdateProject",
		Input:   input,
		Output:  (*iot1clickprojects.UpdateProjectOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.IoT1ClickProjectsAPI.UpdateProjectWithContext(ctx, input, opts...)
	})

	return req.Output.(*iot1clickprojects.UpdateProjectOutput), req.Error
}
