// Code generated by internal/generate/main.go. DO NOT EDIT.

package datasyncctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/aws/aws-sdk-go/service/datasync/datasynciface"
	"github.com/glassechidna/awsctx"
)

type DataSync interface {
	CancelTaskExecutionWithContext(ctx context.Context, input *datasync.CancelTaskExecutionInput, opts ...request.Option) (*datasync.CancelTaskExecutionOutput, error)
	CreateAgentWithContext(ctx context.Context, input *datasync.CreateAgentInput, opts ...request.Option) (*datasync.CreateAgentOutput, error)
	CreateLocationEfsWithContext(ctx context.Context, input *datasync.CreateLocationEfsInput, opts ...request.Option) (*datasync.CreateLocationEfsOutput, error)
	CreateLocationFsxWindowsWithContext(ctx context.Context, input *datasync.CreateLocationFsxWindowsInput, opts ...request.Option) (*datasync.CreateLocationFsxWindowsOutput, error)
	CreateLocationNfsWithContext(ctx context.Context, input *datasync.CreateLocationNfsInput, opts ...request.Option) (*datasync.CreateLocationNfsOutput, error)
	CreateLocationObjectStorageWithContext(ctx context.Context, input *datasync.CreateLocationObjectStorageInput, opts ...request.Option) (*datasync.CreateLocationObjectStorageOutput, error)
	CreateLocationS3WithContext(ctx context.Context, input *datasync.CreateLocationS3Input, opts ...request.Option) (*datasync.CreateLocationS3Output, error)
	CreateLocationSmbWithContext(ctx context.Context, input *datasync.CreateLocationSmbInput, opts ...request.Option) (*datasync.CreateLocationSmbOutput, error)
	CreateTaskWithContext(ctx context.Context, input *datasync.CreateTaskInput, opts ...request.Option) (*datasync.CreateTaskOutput, error)
	DeleteAgentWithContext(ctx context.Context, input *datasync.DeleteAgentInput, opts ...request.Option) (*datasync.DeleteAgentOutput, error)
	DeleteLocationWithContext(ctx context.Context, input *datasync.DeleteLocationInput, opts ...request.Option) (*datasync.DeleteLocationOutput, error)
	DeleteTaskWithContext(ctx context.Context, input *datasync.DeleteTaskInput, opts ...request.Option) (*datasync.DeleteTaskOutput, error)
	DescribeAgentWithContext(ctx context.Context, input *datasync.DescribeAgentInput, opts ...request.Option) (*datasync.DescribeAgentOutput, error)
	DescribeLocationEfsWithContext(ctx context.Context, input *datasync.DescribeLocationEfsInput, opts ...request.Option) (*datasync.DescribeLocationEfsOutput, error)
	DescribeLocationFsxWindowsWithContext(ctx context.Context, input *datasync.DescribeLocationFsxWindowsInput, opts ...request.Option) (*datasync.DescribeLocationFsxWindowsOutput, error)
	DescribeLocationNfsWithContext(ctx context.Context, input *datasync.DescribeLocationNfsInput, opts ...request.Option) (*datasync.DescribeLocationNfsOutput, error)
	DescribeLocationObjectStorageWithContext(ctx context.Context, input *datasync.DescribeLocationObjectStorageInput, opts ...request.Option) (*datasync.DescribeLocationObjectStorageOutput, error)
	DescribeLocationS3WithContext(ctx context.Context, input *datasync.DescribeLocationS3Input, opts ...request.Option) (*datasync.DescribeLocationS3Output, error)
	DescribeLocationSmbWithContext(ctx context.Context, input *datasync.DescribeLocationSmbInput, opts ...request.Option) (*datasync.DescribeLocationSmbOutput, error)
	DescribeTaskWithContext(ctx context.Context, input *datasync.DescribeTaskInput, opts ...request.Option) (*datasync.DescribeTaskOutput, error)
	DescribeTaskExecutionWithContext(ctx context.Context, input *datasync.DescribeTaskExecutionInput, opts ...request.Option) (*datasync.DescribeTaskExecutionOutput, error)
	ListAgentsWithContext(ctx context.Context, input *datasync.ListAgentsInput, opts ...request.Option) (*datasync.ListAgentsOutput, error)
	ListAgentsPagesWithContext(ctx context.Context, input *datasync.ListAgentsInput, cb func(*datasync.ListAgentsOutput, bool) bool, opts ...request.Option) error
	ListLocationsWithContext(ctx context.Context, input *datasync.ListLocationsInput, opts ...request.Option) (*datasync.ListLocationsOutput, error)
	ListLocationsPagesWithContext(ctx context.Context, input *datasync.ListLocationsInput, cb func(*datasync.ListLocationsOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *datasync.ListTagsForResourceInput, opts ...request.Option) (*datasync.ListTagsForResourceOutput, error)
	ListTagsForResourcePagesWithContext(ctx context.Context, input *datasync.ListTagsForResourceInput, cb func(*datasync.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error
	ListTaskExecutionsWithContext(ctx context.Context, input *datasync.ListTaskExecutionsInput, opts ...request.Option) (*datasync.ListTaskExecutionsOutput, error)
	ListTaskExecutionsPagesWithContext(ctx context.Context, input *datasync.ListTaskExecutionsInput, cb func(*datasync.ListTaskExecutionsOutput, bool) bool, opts ...request.Option) error
	ListTasksWithContext(ctx context.Context, input *datasync.ListTasksInput, opts ...request.Option) (*datasync.ListTasksOutput, error)
	ListTasksPagesWithContext(ctx context.Context, input *datasync.ListTasksInput, cb func(*datasync.ListTasksOutput, bool) bool, opts ...request.Option) error
	StartTaskExecutionWithContext(ctx context.Context, input *datasync.StartTaskExecutionInput, opts ...request.Option) (*datasync.StartTaskExecutionOutput, error)
	TagResourceWithContext(ctx context.Context, input *datasync.TagResourceInput, opts ...request.Option) (*datasync.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *datasync.UntagResourceInput, opts ...request.Option) (*datasync.UntagResourceOutput, error)
	UpdateAgentWithContext(ctx context.Context, input *datasync.UpdateAgentInput, opts ...request.Option) (*datasync.UpdateAgentOutput, error)
	UpdateLocationNfsWithContext(ctx context.Context, input *datasync.UpdateLocationNfsInput, opts ...request.Option) (*datasync.UpdateLocationNfsOutput, error)
	UpdateLocationObjectStorageWithContext(ctx context.Context, input *datasync.UpdateLocationObjectStorageInput, opts ...request.Option) (*datasync.UpdateLocationObjectStorageOutput, error)
	UpdateLocationSmbWithContext(ctx context.Context, input *datasync.UpdateLocationSmbInput, opts ...request.Option) (*datasync.UpdateLocationSmbOutput, error)
	UpdateTaskWithContext(ctx context.Context, input *datasync.UpdateTaskInput, opts ...request.Option) (*datasync.UpdateTaskOutput, error)
	UpdateTaskExecutionWithContext(ctx context.Context, input *datasync.UpdateTaskExecutionInput, opts ...request.Option) (*datasync.UpdateTaskExecutionOutput, error)
}

type Client struct {
	datasynciface.DataSyncAPI
	Contexter awsctx.Contexter
}

func New(base datasynciface.DataSyncAPI, ctxer awsctx.Contexter) DataSync {
	return &Client{
		DataSyncAPI: base,
		Contexter: ctxer,
	}
}

var _ DataSync = (*datasync.DataSync)(nil)
var _ DataSync = (*Client)(nil)

func (c *Client) CancelTaskExecutionWithContext(ctx context.Context, input *datasync.CancelTaskExecutionInput, opts ...request.Option) (*datasync.CancelTaskExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CancelTaskExecution",
		Input:   input,
		Output:  (*datasync.CancelTaskExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CancelTaskExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CancelTaskExecutionOutput), req.Error
}

func (c *Client) CreateAgentWithContext(ctx context.Context, input *datasync.CreateAgentInput, opts ...request.Option) (*datasync.CreateAgentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CreateAgent",
		Input:   input,
		Output:  (*datasync.CreateAgentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CreateAgentWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CreateAgentOutput), req.Error
}

func (c *Client) CreateLocationEfsWithContext(ctx context.Context, input *datasync.CreateLocationEfsInput, opts ...request.Option) (*datasync.CreateLocationEfsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CreateLocationEfs",
		Input:   input,
		Output:  (*datasync.CreateLocationEfsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CreateLocationEfsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CreateLocationEfsOutput), req.Error
}

func (c *Client) CreateLocationFsxWindowsWithContext(ctx context.Context, input *datasync.CreateLocationFsxWindowsInput, opts ...request.Option) (*datasync.CreateLocationFsxWindowsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CreateLocationFsxWindows",
		Input:   input,
		Output:  (*datasync.CreateLocationFsxWindowsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CreateLocationFsxWindowsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CreateLocationFsxWindowsOutput), req.Error
}

func (c *Client) CreateLocationNfsWithContext(ctx context.Context, input *datasync.CreateLocationNfsInput, opts ...request.Option) (*datasync.CreateLocationNfsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CreateLocationNfs",
		Input:   input,
		Output:  (*datasync.CreateLocationNfsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CreateLocationNfsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CreateLocationNfsOutput), req.Error
}

func (c *Client) CreateLocationObjectStorageWithContext(ctx context.Context, input *datasync.CreateLocationObjectStorageInput, opts ...request.Option) (*datasync.CreateLocationObjectStorageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CreateLocationObjectStorage",
		Input:   input,
		Output:  (*datasync.CreateLocationObjectStorageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CreateLocationObjectStorageWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CreateLocationObjectStorageOutput), req.Error
}

func (c *Client) CreateLocationS3WithContext(ctx context.Context, input *datasync.CreateLocationS3Input, opts ...request.Option) (*datasync.CreateLocationS3Output, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CreateLocationS3",
		Input:   input,
		Output:  (*datasync.CreateLocationS3Output)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CreateLocationS3WithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CreateLocationS3Output), req.Error
}

func (c *Client) CreateLocationSmbWithContext(ctx context.Context, input *datasync.CreateLocationSmbInput, opts ...request.Option) (*datasync.CreateLocationSmbOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CreateLocationSmb",
		Input:   input,
		Output:  (*datasync.CreateLocationSmbOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CreateLocationSmbWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CreateLocationSmbOutput), req.Error
}

func (c *Client) CreateTaskWithContext(ctx context.Context, input *datasync.CreateTaskInput, opts ...request.Option) (*datasync.CreateTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "CreateTask",
		Input:   input,
		Output:  (*datasync.CreateTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.CreateTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.CreateTaskOutput), req.Error
}

func (c *Client) DeleteAgentWithContext(ctx context.Context, input *datasync.DeleteAgentInput, opts ...request.Option) (*datasync.DeleteAgentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DeleteAgent",
		Input:   input,
		Output:  (*datasync.DeleteAgentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DeleteAgentWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DeleteAgentOutput), req.Error
}

func (c *Client) DeleteLocationWithContext(ctx context.Context, input *datasync.DeleteLocationInput, opts ...request.Option) (*datasync.DeleteLocationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DeleteLocation",
		Input:   input,
		Output:  (*datasync.DeleteLocationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DeleteLocationWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DeleteLocationOutput), req.Error
}

func (c *Client) DeleteTaskWithContext(ctx context.Context, input *datasync.DeleteTaskInput, opts ...request.Option) (*datasync.DeleteTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DeleteTask",
		Input:   input,
		Output:  (*datasync.DeleteTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DeleteTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DeleteTaskOutput), req.Error
}

func (c *Client) DescribeAgentWithContext(ctx context.Context, input *datasync.DescribeAgentInput, opts ...request.Option) (*datasync.DescribeAgentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeAgent",
		Input:   input,
		Output:  (*datasync.DescribeAgentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeAgentWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeAgentOutput), req.Error
}

func (c *Client) DescribeLocationEfsWithContext(ctx context.Context, input *datasync.DescribeLocationEfsInput, opts ...request.Option) (*datasync.DescribeLocationEfsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeLocationEfs",
		Input:   input,
		Output:  (*datasync.DescribeLocationEfsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeLocationEfsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeLocationEfsOutput), req.Error
}

func (c *Client) DescribeLocationFsxWindowsWithContext(ctx context.Context, input *datasync.DescribeLocationFsxWindowsInput, opts ...request.Option) (*datasync.DescribeLocationFsxWindowsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeLocationFsxWindows",
		Input:   input,
		Output:  (*datasync.DescribeLocationFsxWindowsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeLocationFsxWindowsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeLocationFsxWindowsOutput), req.Error
}

func (c *Client) DescribeLocationNfsWithContext(ctx context.Context, input *datasync.DescribeLocationNfsInput, opts ...request.Option) (*datasync.DescribeLocationNfsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeLocationNfs",
		Input:   input,
		Output:  (*datasync.DescribeLocationNfsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeLocationNfsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeLocationNfsOutput), req.Error
}

func (c *Client) DescribeLocationObjectStorageWithContext(ctx context.Context, input *datasync.DescribeLocationObjectStorageInput, opts ...request.Option) (*datasync.DescribeLocationObjectStorageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeLocationObjectStorage",
		Input:   input,
		Output:  (*datasync.DescribeLocationObjectStorageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeLocationObjectStorageWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeLocationObjectStorageOutput), req.Error
}

func (c *Client) DescribeLocationS3WithContext(ctx context.Context, input *datasync.DescribeLocationS3Input, opts ...request.Option) (*datasync.DescribeLocationS3Output, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeLocationS3",
		Input:   input,
		Output:  (*datasync.DescribeLocationS3Output)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeLocationS3WithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeLocationS3Output), req.Error
}

func (c *Client) DescribeLocationSmbWithContext(ctx context.Context, input *datasync.DescribeLocationSmbInput, opts ...request.Option) (*datasync.DescribeLocationSmbOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeLocationSmb",
		Input:   input,
		Output:  (*datasync.DescribeLocationSmbOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeLocationSmbWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeLocationSmbOutput), req.Error
}

func (c *Client) DescribeTaskWithContext(ctx context.Context, input *datasync.DescribeTaskInput, opts ...request.Option) (*datasync.DescribeTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeTask",
		Input:   input,
		Output:  (*datasync.DescribeTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeTaskOutput), req.Error
}

func (c *Client) DescribeTaskExecutionWithContext(ctx context.Context, input *datasync.DescribeTaskExecutionInput, opts ...request.Option) (*datasync.DescribeTaskExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "DescribeTaskExecution",
		Input:   input,
		Output:  (*datasync.DescribeTaskExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.DescribeTaskExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.DescribeTaskExecutionOutput), req.Error
}

func (c *Client) ListAgentsWithContext(ctx context.Context, input *datasync.ListAgentsInput, opts ...request.Option) (*datasync.ListAgentsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListAgents",
		Input:   input,
		Output:  (*datasync.ListAgentsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.ListAgentsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.ListAgentsOutput), req.Error
}

func (c *Client) ListAgentsPagesWithContext(ctx context.Context, input *datasync.ListAgentsInput, cb func(*datasync.ListAgentsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListAgents",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.DataSyncAPI.ListAgentsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListLocationsWithContext(ctx context.Context, input *datasync.ListLocationsInput, opts ...request.Option) (*datasync.ListLocationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListLocations",
		Input:   input,
		Output:  (*datasync.ListLocationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.ListLocationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.ListLocationsOutput), req.Error
}

func (c *Client) ListLocationsPagesWithContext(ctx context.Context, input *datasync.ListLocationsInput, cb func(*datasync.ListLocationsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListLocations",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.DataSyncAPI.ListLocationsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *datasync.ListTagsForResourceInput, opts ...request.Option) (*datasync.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*datasync.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTagsForResourcePagesWithContext(ctx context.Context, input *datasync.ListTagsForResourceInput, cb func(*datasync.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListTagsForResource",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.DataSyncAPI.ListTagsForResourcePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTaskExecutionsWithContext(ctx context.Context, input *datasync.ListTaskExecutionsInput, opts ...request.Option) (*datasync.ListTaskExecutionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListTaskExecutions",
		Input:   input,
		Output:  (*datasync.ListTaskExecutionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.ListTaskExecutionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.ListTaskExecutionsOutput), req.Error
}

func (c *Client) ListTaskExecutionsPagesWithContext(ctx context.Context, input *datasync.ListTaskExecutionsInput, cb func(*datasync.ListTaskExecutionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListTaskExecutions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.DataSyncAPI.ListTaskExecutionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTasksWithContext(ctx context.Context, input *datasync.ListTasksInput, opts ...request.Option) (*datasync.ListTasksOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListTasks",
		Input:   input,
		Output:  (*datasync.ListTasksOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.ListTasksWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.ListTasksOutput), req.Error
}

func (c *Client) ListTasksPagesWithContext(ctx context.Context, input *datasync.ListTasksInput, cb func(*datasync.ListTasksOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "ListTasks",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.DataSyncAPI.ListTasksPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) StartTaskExecutionWithContext(ctx context.Context, input *datasync.StartTaskExecutionInput, opts ...request.Option) (*datasync.StartTaskExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "StartTaskExecution",
		Input:   input,
		Output:  (*datasync.StartTaskExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.StartTaskExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.StartTaskExecutionOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *datasync.TagResourceInput, opts ...request.Option) (*datasync.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "TagResource",
		Input:   input,
		Output:  (*datasync.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *datasync.UntagResourceInput, opts ...request.Option) (*datasync.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*datasync.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.UntagResourceOutput), req.Error
}

func (c *Client) UpdateAgentWithContext(ctx context.Context, input *datasync.UpdateAgentInput, opts ...request.Option) (*datasync.UpdateAgentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "UpdateAgent",
		Input:   input,
		Output:  (*datasync.UpdateAgentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.UpdateAgentWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.UpdateAgentOutput), req.Error
}

func (c *Client) UpdateLocationNfsWithContext(ctx context.Context, input *datasync.UpdateLocationNfsInput, opts ...request.Option) (*datasync.UpdateLocationNfsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "UpdateLocationNfs",
		Input:   input,
		Output:  (*datasync.UpdateLocationNfsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.UpdateLocationNfsWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.UpdateLocationNfsOutput), req.Error
}

func (c *Client) UpdateLocationObjectStorageWithContext(ctx context.Context, input *datasync.UpdateLocationObjectStorageInput, opts ...request.Option) (*datasync.UpdateLocationObjectStorageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "UpdateLocationObjectStorage",
		Input:   input,
		Output:  (*datasync.UpdateLocationObjectStorageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.UpdateLocationObjectStorageWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.UpdateLocationObjectStorageOutput), req.Error
}

func (c *Client) UpdateLocationSmbWithContext(ctx context.Context, input *datasync.UpdateLocationSmbInput, opts ...request.Option) (*datasync.UpdateLocationSmbOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "UpdateLocationSmb",
		Input:   input,
		Output:  (*datasync.UpdateLocationSmbOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.UpdateLocationSmbWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.UpdateLocationSmbOutput), req.Error
}

func (c *Client) UpdateTaskWithContext(ctx context.Context, input *datasync.UpdateTaskInput, opts ...request.Option) (*datasync.UpdateTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "UpdateTask",
		Input:   input,
		Output:  (*datasync.UpdateTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.UpdateTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.UpdateTaskOutput), req.Error
}

func (c *Client) UpdateTaskExecutionWithContext(ctx context.Context, input *datasync.UpdateTaskExecutionInput, opts ...request.Option) (*datasync.UpdateTaskExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "datasync",
		Action:  "UpdateTaskExecution",
		Input:   input,
		Output:  (*datasync.UpdateTaskExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.DataSyncAPI.UpdateTaskExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*datasync.UpdateTaskExecutionOutput), req.Error
}
