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
	CreateAccessPointWithContext(ctx context.Context, input *efs.CreateAccessPointInput, opts ...request.Option) (*efs.CreateAccessPointOutput, error)
	CreateFileSystemWithContext(ctx context.Context, input *efs.CreateFileSystemInput, opts ...request.Option) (*efs.FileSystemDescription, error)
	CreateMountTargetWithContext(ctx context.Context, input *efs.CreateMountTargetInput, opts ...request.Option) (*efs.MountTargetDescription, error)
	CreateTagsWithContext(ctx context.Context, input *efs.CreateTagsInput, opts ...request.Option) (*efs.CreateTagsOutput, error)
	DeleteAccessPointWithContext(ctx context.Context, input *efs.DeleteAccessPointInput, opts ...request.Option) (*efs.DeleteAccessPointOutput, error)
	DeleteFileSystemWithContext(ctx context.Context, input *efs.DeleteFileSystemInput, opts ...request.Option) (*efs.DeleteFileSystemOutput, error)
	DeleteFileSystemPolicyWithContext(ctx context.Context, input *efs.DeleteFileSystemPolicyInput, opts ...request.Option) (*efs.DeleteFileSystemPolicyOutput, error)
	DeleteMountTargetWithContext(ctx context.Context, input *efs.DeleteMountTargetInput, opts ...request.Option) (*efs.DeleteMountTargetOutput, error)
	DeleteTagsWithContext(ctx context.Context, input *efs.DeleteTagsInput, opts ...request.Option) (*efs.DeleteTagsOutput, error)
	DescribeAccessPointsWithContext(ctx context.Context, input *efs.DescribeAccessPointsInput, opts ...request.Option) (*efs.DescribeAccessPointsOutput, error)
	DescribeAccessPointsPagesWithContext(ctx context.Context, input *efs.DescribeAccessPointsInput, cb func(*efs.DescribeAccessPointsOutput, bool) bool, opts ...request.Option) error
	DescribeFileSystemPolicyWithContext(ctx context.Context, input *efs.DescribeFileSystemPolicyInput, opts ...request.Option) (*efs.DescribeFileSystemPolicyOutput, error)
	DescribeFileSystemsWithContext(ctx context.Context, input *efs.DescribeFileSystemsInput, opts ...request.Option) (*efs.DescribeFileSystemsOutput, error)
	DescribeFileSystemsPagesWithContext(ctx context.Context, input *efs.DescribeFileSystemsInput, cb func(*efs.DescribeFileSystemsOutput, bool) bool, opts ...request.Option) error
	DescribeLifecycleConfigurationWithContext(ctx context.Context, input *efs.DescribeLifecycleConfigurationInput, opts ...request.Option) (*efs.DescribeLifecycleConfigurationOutput, error)
	DescribeMountTargetSecurityGroupsWithContext(ctx context.Context, input *efs.DescribeMountTargetSecurityGroupsInput, opts ...request.Option) (*efs.DescribeMountTargetSecurityGroupsOutput, error)
	DescribeMountTargetsWithContext(ctx context.Context, input *efs.DescribeMountTargetsInput, opts ...request.Option) (*efs.DescribeMountTargetsOutput, error)
	DescribeTagsWithContext(ctx context.Context, input *efs.DescribeTagsInput, opts ...request.Option) (*efs.DescribeTagsOutput, error)
	DescribeTagsPagesWithContext(ctx context.Context, input *efs.DescribeTagsInput, cb func(*efs.DescribeTagsOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *efs.ListTagsForResourceInput, opts ...request.Option) (*efs.ListTagsForResourceOutput, error)
	ListTagsForResourcePagesWithContext(ctx context.Context, input *efs.ListTagsForResourceInput, cb func(*efs.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error
	ModifyMountTargetSecurityGroupsWithContext(ctx context.Context, input *efs.ModifyMountTargetSecurityGroupsInput, opts ...request.Option) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
	PutFileSystemPolicyWithContext(ctx context.Context, input *efs.PutFileSystemPolicyInput, opts ...request.Option) (*efs.PutFileSystemPolicyOutput, error)
	PutLifecycleConfigurationWithContext(ctx context.Context, input *efs.PutLifecycleConfigurationInput, opts ...request.Option) (*efs.PutLifecycleConfigurationOutput, error)
	TagResourceWithContext(ctx context.Context, input *efs.TagResourceInput, opts ...request.Option) (*efs.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *efs.UntagResourceInput, opts ...request.Option) (*efs.UntagResourceOutput, error)
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

func (c *Client) CreateAccessPointWithContext(ctx context.Context, input *efs.CreateAccessPointInput, opts ...request.Option) (*efs.CreateAccessPointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "CreateAccessPoint",
		Input:   input,
		Output:  (*efs.CreateAccessPointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.CreateAccessPointWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.CreateAccessPointOutput), req.Error
}

func (c *Client) CreateFileSystemWithContext(ctx context.Context, input *efs.CreateFileSystemInput, opts ...request.Option) (*efs.FileSystemDescription, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "CreateFileSystem",
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
		Action:  "CreateMountTarget",
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
		Action:  "CreateTags",
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

func (c *Client) DeleteAccessPointWithContext(ctx context.Context, input *efs.DeleteAccessPointInput, opts ...request.Option) (*efs.DeleteAccessPointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DeleteAccessPoint",
		Input:   input,
		Output:  (*efs.DeleteAccessPointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DeleteAccessPointWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DeleteAccessPointOutput), req.Error
}

func (c *Client) DeleteFileSystemWithContext(ctx context.Context, input *efs.DeleteFileSystemInput, opts ...request.Option) (*efs.DeleteFileSystemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DeleteFileSystem",
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

func (c *Client) DeleteFileSystemPolicyWithContext(ctx context.Context, input *efs.DeleteFileSystemPolicyInput, opts ...request.Option) (*efs.DeleteFileSystemPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DeleteFileSystemPolicy",
		Input:   input,
		Output:  (*efs.DeleteFileSystemPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DeleteFileSystemPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DeleteFileSystemPolicyOutput), req.Error
}

func (c *Client) DeleteMountTargetWithContext(ctx context.Context, input *efs.DeleteMountTargetInput, opts ...request.Option) (*efs.DeleteMountTargetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DeleteMountTarget",
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
		Action:  "DeleteTags",
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

func (c *Client) DescribeAccessPointsWithContext(ctx context.Context, input *efs.DescribeAccessPointsInput, opts ...request.Option) (*efs.DescribeAccessPointsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeAccessPoints",
		Input:   input,
		Output:  (*efs.DescribeAccessPointsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DescribeAccessPointsWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DescribeAccessPointsOutput), req.Error
}

func (c *Client) DescribeAccessPointsPagesWithContext(ctx context.Context, input *efs.DescribeAccessPointsInput, cb func(*efs.DescribeAccessPointsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeAccessPoints",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.EFSAPI.DescribeAccessPointsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeFileSystemPolicyWithContext(ctx context.Context, input *efs.DescribeFileSystemPolicyInput, opts ...request.Option) (*efs.DescribeFileSystemPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeFileSystemPolicy",
		Input:   input,
		Output:  (*efs.DescribeFileSystemPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.DescribeFileSystemPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.DescribeFileSystemPolicyOutput), req.Error
}

func (c *Client) DescribeFileSystemsWithContext(ctx context.Context, input *efs.DescribeFileSystemsInput, opts ...request.Option) (*efs.DescribeFileSystemsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeFileSystems",
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

func (c *Client) DescribeFileSystemsPagesWithContext(ctx context.Context, input *efs.DescribeFileSystemsInput, cb func(*efs.DescribeFileSystemsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeFileSystems",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.EFSAPI.DescribeFileSystemsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeLifecycleConfigurationWithContext(ctx context.Context, input *efs.DescribeLifecycleConfigurationInput, opts ...request.Option) (*efs.DescribeLifecycleConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeLifecycleConfiguration",
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
		Action:  "DescribeMountTargetSecurityGroups",
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
		Action:  "DescribeMountTargets",
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
		Action:  "DescribeTags",
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

func (c *Client) DescribeTagsPagesWithContext(ctx context.Context, input *efs.DescribeTagsInput, cb func(*efs.DescribeTagsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "DescribeTags",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.EFSAPI.DescribeTagsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *efs.ListTagsForResourceInput, opts ...request.Option) (*efs.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*efs.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTagsForResourcePagesWithContext(ctx context.Context, input *efs.ListTagsForResourceInput, cb func(*efs.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "ListTagsForResource",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.EFSAPI.ListTagsForResourcePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ModifyMountTargetSecurityGroupsWithContext(ctx context.Context, input *efs.ModifyMountTargetSecurityGroupsInput, opts ...request.Option) (*efs.ModifyMountTargetSecurityGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "ModifyMountTargetSecurityGroups",
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

func (c *Client) PutFileSystemPolicyWithContext(ctx context.Context, input *efs.PutFileSystemPolicyInput, opts ...request.Option) (*efs.PutFileSystemPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "PutFileSystemPolicy",
		Input:   input,
		Output:  (*efs.PutFileSystemPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.PutFileSystemPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.PutFileSystemPolicyOutput), req.Error
}

func (c *Client) PutLifecycleConfigurationWithContext(ctx context.Context, input *efs.PutLifecycleConfigurationInput, opts ...request.Option) (*efs.PutLifecycleConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "PutLifecycleConfiguration",
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

func (c *Client) TagResourceWithContext(ctx context.Context, input *efs.TagResourceInput, opts ...request.Option) (*efs.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "TagResource",
		Input:   input,
		Output:  (*efs.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *efs.UntagResourceInput, opts ...request.Option) (*efs.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*efs.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EFSAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*efs.UntagResourceOutput), req.Error
}

func (c *Client) UpdateFileSystemWithContext(ctx context.Context, input *efs.UpdateFileSystemInput, opts ...request.Option) (*efs.UpdateFileSystemOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "efs",
		Action:  "UpdateFileSystem",
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
