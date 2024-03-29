// Code generated by internal/generate/main.go. DO NOT EDIT.

package batchctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/aws/aws-sdk-go/service/batch/batchiface"
	"github.com/glassechidna/awsctx"
)

type Batch interface {
	CancelJobWithContext(ctx context.Context, input *batch.CancelJobInput, opts ...request.Option) (*batch.CancelJobOutput, error)
	CreateComputeEnvironmentWithContext(ctx context.Context, input *batch.CreateComputeEnvironmentInput, opts ...request.Option) (*batch.CreateComputeEnvironmentOutput, error)
	CreateJobQueueWithContext(ctx context.Context, input *batch.CreateJobQueueInput, opts ...request.Option) (*batch.CreateJobQueueOutput, error)
	CreateSchedulingPolicyWithContext(ctx context.Context, input *batch.CreateSchedulingPolicyInput, opts ...request.Option) (*batch.CreateSchedulingPolicyOutput, error)
	DeleteComputeEnvironmentWithContext(ctx context.Context, input *batch.DeleteComputeEnvironmentInput, opts ...request.Option) (*batch.DeleteComputeEnvironmentOutput, error)
	DeleteJobQueueWithContext(ctx context.Context, input *batch.DeleteJobQueueInput, opts ...request.Option) (*batch.DeleteJobQueueOutput, error)
	DeleteSchedulingPolicyWithContext(ctx context.Context, input *batch.DeleteSchedulingPolicyInput, opts ...request.Option) (*batch.DeleteSchedulingPolicyOutput, error)
	DeregisterJobDefinitionWithContext(ctx context.Context, input *batch.DeregisterJobDefinitionInput, opts ...request.Option) (*batch.DeregisterJobDefinitionOutput, error)
	DescribeComputeEnvironmentsWithContext(ctx context.Context, input *batch.DescribeComputeEnvironmentsInput, opts ...request.Option) (*batch.DescribeComputeEnvironmentsOutput, error)
	DescribeComputeEnvironmentsPagesWithContext(ctx context.Context, input *batch.DescribeComputeEnvironmentsInput, cb func(*batch.DescribeComputeEnvironmentsOutput, bool) bool, opts ...request.Option) error
	DescribeJobDefinitionsWithContext(ctx context.Context, input *batch.DescribeJobDefinitionsInput, opts ...request.Option) (*batch.DescribeJobDefinitionsOutput, error)
	DescribeJobDefinitionsPagesWithContext(ctx context.Context, input *batch.DescribeJobDefinitionsInput, cb func(*batch.DescribeJobDefinitionsOutput, bool) bool, opts ...request.Option) error
	DescribeJobQueuesWithContext(ctx context.Context, input *batch.DescribeJobQueuesInput, opts ...request.Option) (*batch.DescribeJobQueuesOutput, error)
	DescribeJobQueuesPagesWithContext(ctx context.Context, input *batch.DescribeJobQueuesInput, cb func(*batch.DescribeJobQueuesOutput, bool) bool, opts ...request.Option) error
	DescribeJobsWithContext(ctx context.Context, input *batch.DescribeJobsInput, opts ...request.Option) (*batch.DescribeJobsOutput, error)
	DescribeSchedulingPoliciesWithContext(ctx context.Context, input *batch.DescribeSchedulingPoliciesInput, opts ...request.Option) (*batch.DescribeSchedulingPoliciesOutput, error)
	ListJobsWithContext(ctx context.Context, input *batch.ListJobsInput, opts ...request.Option) (*batch.ListJobsOutput, error)
	ListJobsPagesWithContext(ctx context.Context, input *batch.ListJobsInput, cb func(*batch.ListJobsOutput, bool) bool, opts ...request.Option) error
	ListSchedulingPoliciesWithContext(ctx context.Context, input *batch.ListSchedulingPoliciesInput, opts ...request.Option) (*batch.ListSchedulingPoliciesOutput, error)
	ListSchedulingPoliciesPagesWithContext(ctx context.Context, input *batch.ListSchedulingPoliciesInput, cb func(*batch.ListSchedulingPoliciesOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *batch.ListTagsForResourceInput, opts ...request.Option) (*batch.ListTagsForResourceOutput, error)
	RegisterJobDefinitionWithContext(ctx context.Context, input *batch.RegisterJobDefinitionInput, opts ...request.Option) (*batch.RegisterJobDefinitionOutput, error)
	SubmitJobWithContext(ctx context.Context, input *batch.SubmitJobInput, opts ...request.Option) (*batch.SubmitJobOutput, error)
	TagResourceWithContext(ctx context.Context, input *batch.TagResourceInput, opts ...request.Option) (*batch.TagResourceOutput, error)
	TerminateJobWithContext(ctx context.Context, input *batch.TerminateJobInput, opts ...request.Option) (*batch.TerminateJobOutput, error)
	UntagResourceWithContext(ctx context.Context, input *batch.UntagResourceInput, opts ...request.Option) (*batch.UntagResourceOutput, error)
	UpdateComputeEnvironmentWithContext(ctx context.Context, input *batch.UpdateComputeEnvironmentInput, opts ...request.Option) (*batch.UpdateComputeEnvironmentOutput, error)
	UpdateJobQueueWithContext(ctx context.Context, input *batch.UpdateJobQueueInput, opts ...request.Option) (*batch.UpdateJobQueueOutput, error)
	UpdateSchedulingPolicyWithContext(ctx context.Context, input *batch.UpdateSchedulingPolicyInput, opts ...request.Option) (*batch.UpdateSchedulingPolicyOutput, error)
}

type Client struct {
	batchiface.BatchAPI
	Contexter awsctx.Contexter
}

func New(base batchiface.BatchAPI, ctxer awsctx.Contexter) Batch {
	return &Client{
		BatchAPI: base,
		Contexter: ctxer,
	}
}

var _ Batch = (*batch.Batch)(nil)
var _ Batch = (*Client)(nil)

func (c *Client) CancelJobWithContext(ctx context.Context, input *batch.CancelJobInput, opts ...request.Option) (*batch.CancelJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "CancelJob",
		Input:   input,
		Output:  (*batch.CancelJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.CancelJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.CancelJobOutput), req.Error
}

func (c *Client) CreateComputeEnvironmentWithContext(ctx context.Context, input *batch.CreateComputeEnvironmentInput, opts ...request.Option) (*batch.CreateComputeEnvironmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "CreateComputeEnvironment",
		Input:   input,
		Output:  (*batch.CreateComputeEnvironmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.CreateComputeEnvironmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.CreateComputeEnvironmentOutput), req.Error
}

func (c *Client) CreateJobQueueWithContext(ctx context.Context, input *batch.CreateJobQueueInput, opts ...request.Option) (*batch.CreateJobQueueOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "CreateJobQueue",
		Input:   input,
		Output:  (*batch.CreateJobQueueOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.CreateJobQueueWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.CreateJobQueueOutput), req.Error
}

func (c *Client) CreateSchedulingPolicyWithContext(ctx context.Context, input *batch.CreateSchedulingPolicyInput, opts ...request.Option) (*batch.CreateSchedulingPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "CreateSchedulingPolicy",
		Input:   input,
		Output:  (*batch.CreateSchedulingPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.CreateSchedulingPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.CreateSchedulingPolicyOutput), req.Error
}

func (c *Client) DeleteComputeEnvironmentWithContext(ctx context.Context, input *batch.DeleteComputeEnvironmentInput, opts ...request.Option) (*batch.DeleteComputeEnvironmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DeleteComputeEnvironment",
		Input:   input,
		Output:  (*batch.DeleteComputeEnvironmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DeleteComputeEnvironmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DeleteComputeEnvironmentOutput), req.Error
}

func (c *Client) DeleteJobQueueWithContext(ctx context.Context, input *batch.DeleteJobQueueInput, opts ...request.Option) (*batch.DeleteJobQueueOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DeleteJobQueue",
		Input:   input,
		Output:  (*batch.DeleteJobQueueOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DeleteJobQueueWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DeleteJobQueueOutput), req.Error
}

func (c *Client) DeleteSchedulingPolicyWithContext(ctx context.Context, input *batch.DeleteSchedulingPolicyInput, opts ...request.Option) (*batch.DeleteSchedulingPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DeleteSchedulingPolicy",
		Input:   input,
		Output:  (*batch.DeleteSchedulingPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DeleteSchedulingPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DeleteSchedulingPolicyOutput), req.Error
}

func (c *Client) DeregisterJobDefinitionWithContext(ctx context.Context, input *batch.DeregisterJobDefinitionInput, opts ...request.Option) (*batch.DeregisterJobDefinitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DeregisterJobDefinition",
		Input:   input,
		Output:  (*batch.DeregisterJobDefinitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DeregisterJobDefinitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DeregisterJobDefinitionOutput), req.Error
}

func (c *Client) DescribeComputeEnvironmentsWithContext(ctx context.Context, input *batch.DescribeComputeEnvironmentsInput, opts ...request.Option) (*batch.DescribeComputeEnvironmentsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DescribeComputeEnvironments",
		Input:   input,
		Output:  (*batch.DescribeComputeEnvironmentsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DescribeComputeEnvironmentsWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DescribeComputeEnvironmentsOutput), req.Error
}

func (c *Client) DescribeComputeEnvironmentsPagesWithContext(ctx context.Context, input *batch.DescribeComputeEnvironmentsInput, cb func(*batch.DescribeComputeEnvironmentsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DescribeComputeEnvironments",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.BatchAPI.DescribeComputeEnvironmentsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeJobDefinitionsWithContext(ctx context.Context, input *batch.DescribeJobDefinitionsInput, opts ...request.Option) (*batch.DescribeJobDefinitionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DescribeJobDefinitions",
		Input:   input,
		Output:  (*batch.DescribeJobDefinitionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DescribeJobDefinitionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DescribeJobDefinitionsOutput), req.Error
}

func (c *Client) DescribeJobDefinitionsPagesWithContext(ctx context.Context, input *batch.DescribeJobDefinitionsInput, cb func(*batch.DescribeJobDefinitionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DescribeJobDefinitions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.BatchAPI.DescribeJobDefinitionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeJobQueuesWithContext(ctx context.Context, input *batch.DescribeJobQueuesInput, opts ...request.Option) (*batch.DescribeJobQueuesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DescribeJobQueues",
		Input:   input,
		Output:  (*batch.DescribeJobQueuesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DescribeJobQueuesWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DescribeJobQueuesOutput), req.Error
}

func (c *Client) DescribeJobQueuesPagesWithContext(ctx context.Context, input *batch.DescribeJobQueuesInput, cb func(*batch.DescribeJobQueuesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DescribeJobQueues",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.BatchAPI.DescribeJobQueuesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeJobsWithContext(ctx context.Context, input *batch.DescribeJobsInput, opts ...request.Option) (*batch.DescribeJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DescribeJobs",
		Input:   input,
		Output:  (*batch.DescribeJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DescribeJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DescribeJobsOutput), req.Error
}

func (c *Client) DescribeSchedulingPoliciesWithContext(ctx context.Context, input *batch.DescribeSchedulingPoliciesInput, opts ...request.Option) (*batch.DescribeSchedulingPoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "DescribeSchedulingPolicies",
		Input:   input,
		Output:  (*batch.DescribeSchedulingPoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.DescribeSchedulingPoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.DescribeSchedulingPoliciesOutput), req.Error
}

func (c *Client) ListJobsWithContext(ctx context.Context, input *batch.ListJobsInput, opts ...request.Option) (*batch.ListJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "ListJobs",
		Input:   input,
		Output:  (*batch.ListJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.ListJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.ListJobsOutput), req.Error
}

func (c *Client) ListJobsPagesWithContext(ctx context.Context, input *batch.ListJobsInput, cb func(*batch.ListJobsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "ListJobs",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.BatchAPI.ListJobsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListSchedulingPoliciesWithContext(ctx context.Context, input *batch.ListSchedulingPoliciesInput, opts ...request.Option) (*batch.ListSchedulingPoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "ListSchedulingPolicies",
		Input:   input,
		Output:  (*batch.ListSchedulingPoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.ListSchedulingPoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.ListSchedulingPoliciesOutput), req.Error
}

func (c *Client) ListSchedulingPoliciesPagesWithContext(ctx context.Context, input *batch.ListSchedulingPoliciesInput, cb func(*batch.ListSchedulingPoliciesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "ListSchedulingPolicies",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.BatchAPI.ListSchedulingPoliciesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *batch.ListTagsForResourceInput, opts ...request.Option) (*batch.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*batch.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.ListTagsForResourceOutput), req.Error
}

func (c *Client) RegisterJobDefinitionWithContext(ctx context.Context, input *batch.RegisterJobDefinitionInput, opts ...request.Option) (*batch.RegisterJobDefinitionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "RegisterJobDefinition",
		Input:   input,
		Output:  (*batch.RegisterJobDefinitionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.RegisterJobDefinitionWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.RegisterJobDefinitionOutput), req.Error
}

func (c *Client) SubmitJobWithContext(ctx context.Context, input *batch.SubmitJobInput, opts ...request.Option) (*batch.SubmitJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "SubmitJob",
		Input:   input,
		Output:  (*batch.SubmitJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.SubmitJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.SubmitJobOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *batch.TagResourceInput, opts ...request.Option) (*batch.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "TagResource",
		Input:   input,
		Output:  (*batch.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.TagResourceOutput), req.Error
}

func (c *Client) TerminateJobWithContext(ctx context.Context, input *batch.TerminateJobInput, opts ...request.Option) (*batch.TerminateJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "TerminateJob",
		Input:   input,
		Output:  (*batch.TerminateJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.TerminateJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.TerminateJobOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *batch.UntagResourceInput, opts ...request.Option) (*batch.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*batch.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.UntagResourceOutput), req.Error
}

func (c *Client) UpdateComputeEnvironmentWithContext(ctx context.Context, input *batch.UpdateComputeEnvironmentInput, opts ...request.Option) (*batch.UpdateComputeEnvironmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "UpdateComputeEnvironment",
		Input:   input,
		Output:  (*batch.UpdateComputeEnvironmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.UpdateComputeEnvironmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.UpdateComputeEnvironmentOutput), req.Error
}

func (c *Client) UpdateJobQueueWithContext(ctx context.Context, input *batch.UpdateJobQueueInput, opts ...request.Option) (*batch.UpdateJobQueueOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "UpdateJobQueue",
		Input:   input,
		Output:  (*batch.UpdateJobQueueOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.UpdateJobQueueWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.UpdateJobQueueOutput), req.Error
}

func (c *Client) UpdateSchedulingPolicyWithContext(ctx context.Context, input *batch.UpdateSchedulingPolicyInput, opts ...request.Option) (*batch.UpdateSchedulingPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "batch",
		Action:  "UpdateSchedulingPolicy",
		Input:   input,
		Output:  (*batch.UpdateSchedulingPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.BatchAPI.UpdateSchedulingPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*batch.UpdateSchedulingPolicyOutput), req.Error
}
