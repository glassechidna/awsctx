// Code generated by internal/generate/main.go. DO NOT EDIT.

package efsctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
	"github.com/glassechidna/awsctx"
)

type EFS interface {
	CreateFileSystemWithContext(ctx context.Context, input *efs.CreateFileSystemInput, opts ...request.Option) (*efs.FileSystemDescription, error)
	CreateMountTargetWithContext(ctx context.Context, input *efs.CreateMountTargetInput, opts ...request.Option) (*efs.MountTargetDescription, error)
	CreateTagsWithContext(ctx context.Context, input *efs.CreateTagsInput, opts ...request.Option) (*efs.CreateTagsOutput, error)
	DeleteFileSystemWithContext(ctx context.Context, input *efs.DeleteFileSystemInput, opts ...request.Option) (*efs.DeleteFileSystemOutput, error)
	DeleteMountTargetWithContext(ctx context.Context, input *efs.DeleteMountTargetInput, opts ...request.Option) (*efs.DeleteMountTargetOutput, error)
	DeleteTagsWithContext(ctx context.Context, input *efs.DeleteTagsInput, opts ...request.Option) (*efs.DeleteTagsOutput, error)
	DescribeFileSystemsWithContext(ctx context.Context, input *efs.DescribeFileSystemsInput, opts ...request.Option) (*efs.DescribeFileSystemsOutput, error)
	DescribeLifecycleConfigurationWithContext(ctx context.Context, input *efs.DescribeLifecycleConfigurationInput, opts ...request.Option) (*efs.DescribeLifecycleConfigurationOutput, error)
	DescribeMountTargetSecurityGroupsWithContext(ctx context.Context, input *efs.DescribeMountTargetSecurityGroupsInput, opts ...request.Option) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
	DescribeMountTargetsWithContext(ctx context.Context, input *efs.DescribeMountTargetsInput, opts ...request.Option) (*efs.DescribeMountTargetsOutput, error)
	DescribeTagsWithContext(ctx context.Context, input *efs.DescribeTagsInput, opts ...request.Option) (*efs.DescribeTagsOutput, error)
	ModifyMountTargetSecurityGroupsWithContext(ctx context.Context, input *efs.ModifyMountTargetSecurityGroupsInput, opts ...request.Option) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
	PutLifecycleConfigurationWithContext(ctx context.Context, input *efs.PutLifecycleConfigurationInput, opts ...request.Option) (*efs.PutLifecycleConfigurationOutput, error)
	UpdateFileSystemWithContext(ctx context.Context, input *efs.UpdateFileSystemInput, opts ...request.Option) (*efs.UpdateFileSystemOutput, error)
}

type Client struct {
	efsiface.EFSAPI
	Contexter awsctx.Contexter
}

func New(base efsiface.EFSAPI, ctxer awsctx.Contexter) EFS {
	return &Client{
		EFSAPI: base,
		Contexter: ctxer,
	}
}

var _ EFS = (*efs.EFS)(nil)
var _ EFS = (*Client)(nil)

func (c *Client) CreateFileSystemWithContext(ctx context.Context, input *efs.CreateFileSystemInput, opts ...request.Option) (*efs.FileSystemDescription, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "CreateFileSystemWithContext",
		Input:   input,
		Output:  (*efs.FileSystemDescription)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.CreateFileSystemWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.FileSystemDescription), req.Error
}

func (c *Client) CreateMountTargetWithContext(ctx context.Context, input *efs.CreateMountTargetInput, opts ...request.Option) (*efs.MountTargetDescription, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "CreateMountTargetWithContext",
		Input:   input,
		Output:  (*efs.MountTargetDescription)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.CreateMountTargetWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.MountTargetDescription), req.Error
}

func (c *Client) CreateTagsWithContext(ctx context.Context, input *efs.CreateTagsInput, opts ...request.Option) (*efs.CreateTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "CreateTagsWithContext",
		Input:   input,
		Output:  (*efs.CreateTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.CreateTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.CreateTagsOutput), req.Error
}

func (c *Client) DeleteFileSystemWithContext(ctx context.Context, input *efs.DeleteFileSystemInput, opts ...request.Option) (*efs.DeleteFileSystemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DeleteFileSystemWithContext",
		Input:   input,
		Output:  (*efs.DeleteFileSystemOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DeleteFileSystemWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DeleteFileSystemOutput), req.Error
}

func (c *Client) DeleteMountTargetWithContext(ctx context.Context, input *efs.DeleteMountTargetInput, opts ...request.Option) (*efs.DeleteMountTargetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DeleteMountTargetWithContext",
		Input:   input,
		Output:  (*efs.DeleteMountTargetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DeleteMountTargetWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DeleteMountTargetOutput), req.Error
}

func (c *Client) DeleteTagsWithContext(ctx context.Context, input *efs.DeleteTagsInput, opts ...request.Option) (*efs.DeleteTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DeleteTagsWithContext",
		Input:   input,
		Output:  (*efs.DeleteTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DeleteTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DeleteTagsOutput), req.Error
}

func (c *Client) DescribeFileSystemsWithContext(ctx context.Context, input *efs.DescribeFileSystemsInput, opts ...request.Option) (*efs.DescribeFileSystemsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeFileSystemsWithContext",
		Input:   input,
		Output:  (*efs.DescribeFileSystemsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DescribeFileSystemsWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DescribeFileSystemsOutput), req.Error
}

func (c *Client) DescribeLifecycleConfigurationWithContext(ctx context.Context, input *efs.DescribeLifecycleConfigurationInput, opts ...request.Option) (*efs.DescribeLifecycleConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeLifecycleConfigurationWithContext",
		Input:   input,
		Output:  (*efs.DescribeLifecycleConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DescribeLifecycleConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DescribeLifecycleConfigurationOutput), req.Error
}

func (c *Client) DescribeMountTargetSecurityGroupsWithContext(ctx context.Context, input *efs.DescribeMountTargetSecurityGroupsInput, opts ...request.Option) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeMountTargetSecurityGroupsWithContext",
		Input:   input,
		Output:  (*efs.DescribeMountTargetSecurityGroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DescribeMountTargetSecurityGroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DescribeMountTargetSecurityGroupsOutput), req.Error
}

func (c *Client) DescribeMountTargetsWithContext(ctx context.Context, input *efs.DescribeMountTargetsInput, opts ...request.Option) (*efs.DescribeMountTargetsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeMountTargetsWithContext",
		Input:   input,
		Output:  (*efs.DescribeMountTargetsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DescribeMountTargetsWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DescribeMountTargetsOutput), req.Error
}

func (c *Client) DescribeTagsWithContext(ctx context.Context, input *efs.DescribeTagsInput, opts ...request.Option) (*efs.DescribeTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeTagsWithContext",
		Input:   input,
		Output:  (*efs.DescribeTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DescribeTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DescribeTagsOutput), req.Error
}

func (c *Client) ModifyMountTargetSecurityGroupsWithContext(ctx context.Context, input *efs.ModifyMountTargetSecurityGroupsInput, opts ...request.Option) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "ModifyMountTargetSecurityGroupsWithContext",
		Input:   input,
		Output:  (*efs.ModifyMountTargetSecurityGroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.ModifyMountTargetSecurityGroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.ModifyMountTargetSecurityGroupsOutput), req.Error
}

func (c *Client) PutLifecycleConfigurationWithContext(ctx context.Context, input *efs.PutLifecycleConfigurationInput, opts ...request.Option) (*efs.PutLifecycleConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "PutLifecycleConfigurationWithContext",
		Input:   input,
		Output:  (*efs.PutLifecycleConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.PutLifecycleConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.PutLifecycleConfigurationOutput), req.Error
}

func (c *Client) UpdateFileSystemWithContext(ctx context.Context, input *efs.UpdateFileSystemInput, opts ...request.Option) (*efs.UpdateFileSystemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "UpdateFileSystemWithContext",
		Input:   input,
		Output:  (*efs.UpdateFileSystemOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.UpdateFileSystemWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.UpdateFileSystemOutput), req.Error
}
