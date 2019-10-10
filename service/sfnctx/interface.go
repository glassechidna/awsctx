// Code generated by internal/generate/main.go. DO NOT EDIT.

package sfnctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
	"github.com/glassechidna/awsctx"
)

type SFN interface {
	CreateActivityWithContext(ctx context.Context, input *sfn.CreateActivityInput, opts ...request.Option) (*sfn.CreateActivityOutput, error)
	CreateStateMachineWithContext(ctx context.Context, input *sfn.CreateStateMachineInput, opts ...request.Option) (*sfn.CreateStateMachineOutput, error)
	DeleteActivityWithContext(ctx context.Context, input *sfn.DeleteActivityInput, opts ...request.Option) (*sfn.DeleteActivityOutput, error)
	DeleteStateMachineWithContext(ctx context.Context, input *sfn.DeleteStateMachineInput, opts ...request.Option) (*sfn.DeleteStateMachineOutput, error)
	DescribeActivityWithContext(ctx context.Context, input *sfn.DescribeActivityInput, opts ...request.Option) (*sfn.DescribeActivityOutput, error)
	DescribeExecutionWithContext(ctx context.Context, input *sfn.DescribeExecutionInput, opts ...request.Option) (*sfn.DescribeExecutionOutput, error)
	DescribeStateMachineWithContext(ctx context.Context, input *sfn.DescribeStateMachineInput, opts ...request.Option) (*sfn.DescribeStateMachineOutput, error)
	DescribeStateMachineForExecutionWithContext(ctx context.Context, input *sfn.DescribeStateMachineForExecutionInput, opts ...request.Option) (*sfn.DescribeStateMachineForExecutionOutput, error)
	GetActivityTaskWithContext(ctx context.Context, input *sfn.GetActivityTaskInput, opts ...request.Option) (*sfn.GetActivityTaskOutput, error)
	GetExecutionHistoryWithContext(ctx context.Context, input *sfn.GetExecutionHistoryInput, opts ...request.Option) (*sfn.GetExecutionHistoryOutput, error)
	ListActivitiesWithContext(ctx context.Context, input *sfn.ListActivitiesInput, opts ...request.Option) (*sfn.ListActivitiesOutput, error)
	ListExecutionsWithContext(ctx context.Context, input *sfn.ListExecutionsInput, opts ...request.Option) (*sfn.ListExecutionsOutput, error)
	ListStateMachinesWithContext(ctx context.Context, input *sfn.ListStateMachinesInput, opts ...request.Option) (*sfn.ListStateMachinesOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *sfn.ListTagsForResourceInput, opts ...request.Option) (*sfn.ListTagsForResourceOutput, error)
	SendTaskFailureWithContext(ctx context.Context, input *sfn.SendTaskFailureInput, opts ...request.Option) (*sfn.SendTaskFailureOutput, error)
	SendTaskHeartbeatWithContext(ctx context.Context, input *sfn.SendTaskHeartbeatInput, opts ...request.Option) (*sfn.SendTaskHeartbeatOutput, error)
	SendTaskSuccessWithContext(ctx context.Context, input *sfn.SendTaskSuccessInput, opts ...request.Option) (*sfn.SendTaskSuccessOutput, error)
	StartExecutionWithContext(ctx context.Context, input *sfn.StartExecutionInput, opts ...request.Option) (*sfn.StartExecutionOutput, error)
	StopExecutionWithContext(ctx context.Context, input *sfn.StopExecutionInput, opts ...request.Option) (*sfn.StopExecutionOutput, error)
	TagResourceWithContext(ctx context.Context, input *sfn.TagResourceInput, opts ...request.Option) (*sfn.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *sfn.UntagResourceInput, opts ...request.Option) (*sfn.UntagResourceOutput, error)
	UpdateStateMachineWithContext(ctx context.Context, input *sfn.UpdateStateMachineInput, opts ...request.Option) (*sfn.UpdateStateMachineOutput, error)
}

type Client struct {
	sfniface.SFNAPI
	Contexter awsctx.Contexter
}

func New(base sfniface.SFNAPI, ctxer awsctx.Contexter) SFN {
	return &Client{
		SFNAPI: base,
		Contexter: ctxer,
	}
}

var _ SFN = (*sfn.SFN)(nil)
var _ SFN = (*Client)(nil)

func (c *Client) CreateActivityWithContext(ctx context.Context, input *sfn.CreateActivityInput, opts ...request.Option) (*sfn.CreateActivityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "CreateActivity",
		Input:   input,
		Output:  (*sfn.CreateActivityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.CreateActivityWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.CreateActivityOutput), req.Error
}

func (c *Client) CreateStateMachineWithContext(ctx context.Context, input *sfn.CreateStateMachineInput, opts ...request.Option) (*sfn.CreateStateMachineOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "CreateStateMachine",
		Input:   input,
		Output:  (*sfn.CreateStateMachineOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.CreateStateMachineWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.CreateStateMachineOutput), req.Error
}

func (c *Client) DeleteActivityWithContext(ctx context.Context, input *sfn.DeleteActivityInput, opts ...request.Option) (*sfn.DeleteActivityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "DeleteActivity",
		Input:   input,
		Output:  (*sfn.DeleteActivityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.DeleteActivityWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.DeleteActivityOutput), req.Error
}

func (c *Client) DeleteStateMachineWithContext(ctx context.Context, input *sfn.DeleteStateMachineInput, opts ...request.Option) (*sfn.DeleteStateMachineOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "DeleteStateMachine",
		Input:   input,
		Output:  (*sfn.DeleteStateMachineOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.DeleteStateMachineWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.DeleteStateMachineOutput), req.Error
}

func (c *Client) DescribeActivityWithContext(ctx context.Context, input *sfn.DescribeActivityInput, opts ...request.Option) (*sfn.DescribeActivityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "DescribeActivity",
		Input:   input,
		Output:  (*sfn.DescribeActivityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.DescribeActivityWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.DescribeActivityOutput), req.Error
}

func (c *Client) DescribeExecutionWithContext(ctx context.Context, input *sfn.DescribeExecutionInput, opts ...request.Option) (*sfn.DescribeExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "DescribeExecution",
		Input:   input,
		Output:  (*sfn.DescribeExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.DescribeExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.DescribeExecutionOutput), req.Error
}

func (c *Client) DescribeStateMachineWithContext(ctx context.Context, input *sfn.DescribeStateMachineInput, opts ...request.Option) (*sfn.DescribeStateMachineOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "DescribeStateMachine",
		Input:   input,
		Output:  (*sfn.DescribeStateMachineOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.DescribeStateMachineWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.DescribeStateMachineOutput), req.Error
}

func (c *Client) DescribeStateMachineForExecutionWithContext(ctx context.Context, input *sfn.DescribeStateMachineForExecutionInput, opts ...request.Option) (*sfn.DescribeStateMachineForExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "DescribeStateMachineForExecution",
		Input:   input,
		Output:  (*sfn.DescribeStateMachineForExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.DescribeStateMachineForExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.DescribeStateMachineForExecutionOutput), req.Error
}

func (c *Client) GetActivityTaskWithContext(ctx context.Context, input *sfn.GetActivityTaskInput, opts ...request.Option) (*sfn.GetActivityTaskOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "GetActivityTask",
		Input:   input,
		Output:  (*sfn.GetActivityTaskOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.GetActivityTaskWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.GetActivityTaskOutput), req.Error
}

func (c *Client) GetExecutionHistoryWithContext(ctx context.Context, input *sfn.GetExecutionHistoryInput, opts ...request.Option) (*sfn.GetExecutionHistoryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "GetExecutionHistory",
		Input:   input,
		Output:  (*sfn.GetExecutionHistoryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.GetExecutionHistoryWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.GetExecutionHistoryOutput), req.Error
}

func (c *Client) ListActivitiesWithContext(ctx context.Context, input *sfn.ListActivitiesInput, opts ...request.Option) (*sfn.ListActivitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "ListActivities",
		Input:   input,
		Output:  (*sfn.ListActivitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.ListActivitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.ListActivitiesOutput), req.Error
}

func (c *Client) ListExecutionsWithContext(ctx context.Context, input *sfn.ListExecutionsInput, opts ...request.Option) (*sfn.ListExecutionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "ListExecutions",
		Input:   input,
		Output:  (*sfn.ListExecutionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.ListExecutionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.ListExecutionsOutput), req.Error
}

func (c *Client) ListStateMachinesWithContext(ctx context.Context, input *sfn.ListStateMachinesInput, opts ...request.Option) (*sfn.ListStateMachinesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "ListStateMachines",
		Input:   input,
		Output:  (*sfn.ListStateMachinesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.ListStateMachinesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.ListStateMachinesOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *sfn.ListTagsForResourceInput, opts ...request.Option) (*sfn.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*sfn.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.ListTagsForResourceOutput), req.Error
}

func (c *Client) SendTaskFailureWithContext(ctx context.Context, input *sfn.SendTaskFailureInput, opts ...request.Option) (*sfn.SendTaskFailureOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "SendTaskFailure",
		Input:   input,
		Output:  (*sfn.SendTaskFailureOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.SendTaskFailureWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.SendTaskFailureOutput), req.Error
}

func (c *Client) SendTaskHeartbeatWithContext(ctx context.Context, input *sfn.SendTaskHeartbeatInput, opts ...request.Option) (*sfn.SendTaskHeartbeatOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "SendTaskHeartbeat",
		Input:   input,
		Output:  (*sfn.SendTaskHeartbeatOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.SendTaskHeartbeatWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.SendTaskHeartbeatOutput), req.Error
}

func (c *Client) SendTaskSuccessWithContext(ctx context.Context, input *sfn.SendTaskSuccessInput, opts ...request.Option) (*sfn.SendTaskSuccessOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "SendTaskSuccess",
		Input:   input,
		Output:  (*sfn.SendTaskSuccessOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.SendTaskSuccessWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.SendTaskSuccessOutput), req.Error
}

func (c *Client) StartExecutionWithContext(ctx context.Context, input *sfn.StartExecutionInput, opts ...request.Option) (*sfn.StartExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "StartExecution",
		Input:   input,
		Output:  (*sfn.StartExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.StartExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.StartExecutionOutput), req.Error
}

func (c *Client) StopExecutionWithContext(ctx context.Context, input *sfn.StopExecutionInput, opts ...request.Option) (*sfn.StopExecutionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "StopExecution",
		Input:   input,
		Output:  (*sfn.StopExecutionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.StopExecutionWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.StopExecutionOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *sfn.TagResourceInput, opts ...request.Option) (*sfn.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "TagResource",
		Input:   input,
		Output:  (*sfn.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *sfn.UntagResourceInput, opts ...request.Option) (*sfn.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*sfn.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.UntagResourceOutput), req.Error
}

func (c *Client) UpdateStateMachineWithContext(ctx context.Context, input *sfn.UpdateStateMachineInput, opts ...request.Option) (*sfn.UpdateStateMachineOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sfn",
		Action:  "UpdateStateMachine",
		Input:   input,
		Output:  (*sfn.UpdateStateMachineOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SFNAPI.UpdateStateMachineWithContext(ctx, input, opts...)
	})

	return req.Output.(*sfn.UpdateStateMachineOutput), req.Error
}
