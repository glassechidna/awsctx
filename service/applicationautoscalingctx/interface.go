// Code generated by internal/generate/main.go. DO NOT EDIT.

package applicationautoscalingctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/glassechidna/awsctx"
)

type ApplicationAutoScaling interface {
	DeleteScalingPolicyWithContext(ctx context.Context, input *applicationautoscaling.DeleteScalingPolicyInput, opts ...request.Option) (*applicationautoscaling.DeleteScalingPolicyOutput, error)
	DeleteScheduledActionWithContext(ctx context.Context, input *applicationautoscaling.DeleteScheduledActionInput, opts ...request.Option) (*applicationautoscaling.DeleteScheduledActionOutput, error)
	DeregisterScalableTargetWithContext(ctx context.Context, input *applicationautoscaling.DeregisterScalableTargetInput, opts ...request.Option) (*applicationautoscaling.DeregisterScalableTargetOutput, error)
	DescribeScalableTargetsWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalableTargetsInput, opts ...request.Option) (*applicationautoscaling.DescribeScalableTargetsOutput, error)
	DescribeScalableTargetsPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalableTargetsInput, cb func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool, opts ...request.Option) error
	DescribeScalingActivitiesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingActivitiesInput, opts ...request.Option) (*applicationautoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingActivitiesInput, cb func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool, opts ...request.Option) error
	DescribeScalingPoliciesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingPoliciesInput, opts ...request.Option) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)
	DescribeScalingPoliciesPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingPoliciesInput, cb func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool, opts ...request.Option) error
	DescribeScheduledActionsWithContext(ctx context.Context, input *applicationautoscaling.DescribeScheduledActionsInput, opts ...request.Option) (*applicationautoscaling.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScheduledActionsInput, cb func(*applicationautoscaling.DescribeScheduledActionsOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *applicationautoscaling.ListTagsForResourceInput, opts ...request.Option) (*applicationautoscaling.ListTagsForResourceOutput, error)
	PutScalingPolicyWithContext(ctx context.Context, input *applicationautoscaling.PutScalingPolicyInput, opts ...request.Option) (*applicationautoscaling.PutScalingPolicyOutput, error)
	PutScheduledActionWithContext(ctx context.Context, input *applicationautoscaling.PutScheduledActionInput, opts ...request.Option) (*applicationautoscaling.PutScheduledActionOutput, error)
	RegisterScalableTargetWithContext(ctx context.Context, input *applicationautoscaling.RegisterScalableTargetInput, opts ...request.Option) (*applicationautoscaling.RegisterScalableTargetOutput, error)
	TagResourceWithContext(ctx context.Context, input *applicationautoscaling.TagResourceInput, opts ...request.Option) (*applicationautoscaling.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *applicationautoscaling.UntagResourceInput, opts ...request.Option) (*applicationautoscaling.UntagResourceOutput, error)
}

type Client struct {
	applicationautoscalingiface.ApplicationAutoScalingAPI
	Contexter awsctx.Contexter
}

func New(base applicationautoscalingiface.ApplicationAutoScalingAPI, ctxer awsctx.Contexter) ApplicationAutoScaling {
	return &Client{
		ApplicationAutoScalingAPI: base,
		Contexter: ctxer,
	}
}

var _ ApplicationAutoScaling = (*applicationautoscaling.ApplicationAutoScaling)(nil)
var _ ApplicationAutoScaling = (*Client)(nil)

func (c *Client) DeleteScalingPolicyWithContext(ctx context.Context, input *applicationautoscaling.DeleteScalingPolicyInput, opts ...request.Option) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DeleteScalingPolicy",
		Input:   input,
		Output:  (*applicationautoscaling.DeleteScalingPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.DeleteScalingPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.DeleteScalingPolicyOutput), req.Error
}

func (c *Client) DeleteScheduledActionWithContext(ctx context.Context, input *applicationautoscaling.DeleteScheduledActionInput, opts ...request.Option) (*applicationautoscaling.DeleteScheduledActionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DeleteScheduledAction",
		Input:   input,
		Output:  (*applicationautoscaling.DeleteScheduledActionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.DeleteScheduledActionWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.DeleteScheduledActionOutput), req.Error
}

func (c *Client) DeregisterScalableTargetWithContext(ctx context.Context, input *applicationautoscaling.DeregisterScalableTargetInput, opts ...request.Option) (*applicationautoscaling.DeregisterScalableTargetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DeregisterScalableTarget",
		Input:   input,
		Output:  (*applicationautoscaling.DeregisterScalableTargetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.DeregisterScalableTargetWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.DeregisterScalableTargetOutput), req.Error
}

func (c *Client) DescribeScalableTargetsWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalableTargetsInput, opts ...request.Option) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DescribeScalableTargets",
		Input:   input,
		Output:  (*applicationautoscaling.DescribeScalableTargetsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.DescribeScalableTargetsWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.DescribeScalableTargetsOutput), req.Error
}

func (c *Client) DescribeScalableTargetsPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalableTargetsInput, cb func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DescribeScalableTargets",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ApplicationAutoScalingAPI.DescribeScalableTargetsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeScalingActivitiesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingActivitiesInput, opts ...request.Option) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DescribeScalingActivities",
		Input:   input,
		Output:  (*applicationautoscaling.DescribeScalingActivitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.DescribeScalingActivitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.DescribeScalingActivitiesOutput), req.Error
}

func (c *Client) DescribeScalingActivitiesPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingActivitiesInput, cb func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DescribeScalingActivities",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ApplicationAutoScalingAPI.DescribeScalingActivitiesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeScalingPoliciesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingPoliciesInput, opts ...request.Option) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DescribeScalingPolicies",
		Input:   input,
		Output:  (*applicationautoscaling.DescribeScalingPoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.DescribeScalingPoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.DescribeScalingPoliciesOutput), req.Error
}

func (c *Client) DescribeScalingPoliciesPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingPoliciesInput, cb func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DescribeScalingPolicies",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ApplicationAutoScalingAPI.DescribeScalingPoliciesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeScheduledActionsWithContext(ctx context.Context, input *applicationautoscaling.DescribeScheduledActionsInput, opts ...request.Option) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DescribeScheduledActions",
		Input:   input,
		Output:  (*applicationautoscaling.DescribeScheduledActionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.DescribeScheduledActionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.DescribeScheduledActionsOutput), req.Error
}

func (c *Client) DescribeScheduledActionsPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScheduledActionsInput, cb func(*applicationautoscaling.DescribeScheduledActionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "DescribeScheduledActions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ApplicationAutoScalingAPI.DescribeScheduledActionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *applicationautoscaling.ListTagsForResourceInput, opts ...request.Option) (*applicationautoscaling.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*applicationautoscaling.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.ListTagsForResourceOutput), req.Error
}

func (c *Client) PutScalingPolicyWithContext(ctx context.Context, input *applicationautoscaling.PutScalingPolicyInput, opts ...request.Option) (*applicationautoscaling.PutScalingPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "PutScalingPolicy",
		Input:   input,
		Output:  (*applicationautoscaling.PutScalingPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.PutScalingPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.PutScalingPolicyOutput), req.Error
}

func (c *Client) PutScheduledActionWithContext(ctx context.Context, input *applicationautoscaling.PutScheduledActionInput, opts ...request.Option) (*applicationautoscaling.PutScheduledActionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "PutScheduledAction",
		Input:   input,
		Output:  (*applicationautoscaling.PutScheduledActionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.PutScheduledActionWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.PutScheduledActionOutput), req.Error
}

func (c *Client) RegisterScalableTargetWithContext(ctx context.Context, input *applicationautoscaling.RegisterScalableTargetInput, opts ...request.Option) (*applicationautoscaling.RegisterScalableTargetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "RegisterScalableTarget",
		Input:   input,
		Output:  (*applicationautoscaling.RegisterScalableTargetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.RegisterScalableTargetWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.RegisterScalableTargetOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *applicationautoscaling.TagResourceInput, opts ...request.Option) (*applicationautoscaling.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "TagResource",
		Input:   input,
		Output:  (*applicationautoscaling.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *applicationautoscaling.UntagResourceInput, opts ...request.Option) (*applicationautoscaling.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "applicationautoscaling",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*applicationautoscaling.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ApplicationAutoScalingAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*applicationautoscaling.UntagResourceOutput), req.Error
}
