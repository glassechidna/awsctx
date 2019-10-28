// Code generated by internal/generate/main.go. DO NOT EDIT.

package healthctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/health/healthiface"
	"github.com/glassechidna/awsctx"
)

type Health interface {
	DescribeAffectedEntitiesWithContext(ctx context.Context, input *health.DescribeAffectedEntitiesInput, opts ...request.Option) (*health.DescribeAffectedEntitiesOutput, error)
	DescribeAffectedEntitiesPagesWithContext(ctx context.Context, input *health.DescribeAffectedEntitiesInput, cb func(*health.DescribeAffectedEntitiesOutput, bool) bool, opts ...request.Option) error
	DescribeEntityAggregatesWithContext(ctx context.Context, input *health.DescribeEntityAggregatesInput, opts ...request.Option) (*health.DescribeEntityAggregatesOutput, error)
	DescribeEventAggregatesWithContext(ctx context.Context, input *health.DescribeEventAggregatesInput, opts ...request.Option) (*health.DescribeEventAggregatesOutput, error)
	DescribeEventAggregatesPagesWithContext(ctx context.Context, input *health.DescribeEventAggregatesInput, cb func(*health.DescribeEventAggregatesOutput, bool) bool, opts ...request.Option) error
	DescribeEventDetailsWithContext(ctx context.Context, input *health.DescribeEventDetailsInput, opts ...request.Option) (*health.DescribeEventDetailsOutput, error)
	DescribeEventTypesWithContext(ctx context.Context, input *health.DescribeEventTypesInput, opts ...request.Option) (*health.DescribeEventTypesOutput, error)
	DescribeEventTypesPagesWithContext(ctx context.Context, input *health.DescribeEventTypesInput, cb func(*health.DescribeEventTypesOutput, bool) bool, opts ...request.Option) error
	DescribeEventsWithContext(ctx context.Context, input *health.DescribeEventsInput, opts ...request.Option) (*health.DescribeEventsOutput, error)
	DescribeEventsPagesWithContext(ctx context.Context, input *health.DescribeEventsInput, cb func(*health.DescribeEventsOutput, bool) bool, opts ...request.Option) error
}

type Client struct {
	healthiface.HealthAPI
	Contexter awsctx.Contexter
}

func New(base healthiface.HealthAPI, ctxer awsctx.Contexter) Health {
	return &Client{
		HealthAPI: base,
		Contexter: ctxer,
	}
}

var _ Health = (*health.Health)(nil)
var _ Health = (*Client)(nil)

func (c *Client) DescribeAffectedEntitiesWithContext(ctx context.Context, input *health.DescribeAffectedEntitiesInput, opts ...request.Option) (*health.DescribeAffectedEntitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeAffectedEntities",
		Input:   input,
		Output:  (*health.DescribeAffectedEntitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.HealthAPI.DescribeAffectedEntitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*health.DescribeAffectedEntitiesOutput), req.Error
}

func (c *Client) DescribeAffectedEntitiesPagesWithContext(ctx context.Context, input *health.DescribeAffectedEntitiesInput, cb func(*health.DescribeAffectedEntitiesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeAffectedEntities",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.HealthAPI.DescribeAffectedEntitiesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeEntityAggregatesWithContext(ctx context.Context, input *health.DescribeEntityAggregatesInput, opts ...request.Option) (*health.DescribeEntityAggregatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEntityAggregates",
		Input:   input,
		Output:  (*health.DescribeEntityAggregatesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.HealthAPI.DescribeEntityAggregatesWithContext(ctx, input, opts...)
	})

	return req.Output.(*health.DescribeEntityAggregatesOutput), req.Error
}

func (c *Client) DescribeEventAggregatesWithContext(ctx context.Context, input *health.DescribeEventAggregatesInput, opts ...request.Option) (*health.DescribeEventAggregatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEventAggregates",
		Input:   input,
		Output:  (*health.DescribeEventAggregatesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.HealthAPI.DescribeEventAggregatesWithContext(ctx, input, opts...)
	})

	return req.Output.(*health.DescribeEventAggregatesOutput), req.Error
}

func (c *Client) DescribeEventAggregatesPagesWithContext(ctx context.Context, input *health.DescribeEventAggregatesInput, cb func(*health.DescribeEventAggregatesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEventAggregates",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.HealthAPI.DescribeEventAggregatesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeEventDetailsWithContext(ctx context.Context, input *health.DescribeEventDetailsInput, opts ...request.Option) (*health.DescribeEventDetailsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEventDetails",
		Input:   input,
		Output:  (*health.DescribeEventDetailsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.HealthAPI.DescribeEventDetailsWithContext(ctx, input, opts...)
	})

	return req.Output.(*health.DescribeEventDetailsOutput), req.Error
}

func (c *Client) DescribeEventTypesWithContext(ctx context.Context, input *health.DescribeEventTypesInput, opts ...request.Option) (*health.DescribeEventTypesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEventTypes",
		Input:   input,
		Output:  (*health.DescribeEventTypesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.HealthAPI.DescribeEventTypesWithContext(ctx, input, opts...)
	})

	return req.Output.(*health.DescribeEventTypesOutput), req.Error
}

func (c *Client) DescribeEventTypesPagesWithContext(ctx context.Context, input *health.DescribeEventTypesInput, cb func(*health.DescribeEventTypesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEventTypes",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.HealthAPI.DescribeEventTypesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeEventsWithContext(ctx context.Context, input *health.DescribeEventsInput, opts ...request.Option) (*health.DescribeEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEvents",
		Input:   input,
		Output:  (*health.DescribeEventsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.HealthAPI.DescribeEventsWithContext(ctx, input, opts...)
	})

	return req.Output.(*health.DescribeEventsOutput), req.Error
}

func (c *Client) DescribeEventsPagesWithContext(ctx context.Context, input *health.DescribeEventsInput, cb func(*health.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEvents",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.HealthAPI.DescribeEventsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}
