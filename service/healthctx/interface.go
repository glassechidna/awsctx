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
	DescribeEntityAggregatesWithContext(ctx context.Context, input *health.DescribeEntityAggregatesInput, opts ...request.Option) (*health.DescribeEntityAggregatesOutput, error)
	DescribeEventAggregatesWithContext(ctx context.Context, input *health.DescribeEventAggregatesInput, opts ...request.Option) (*health.DescribeEventAggregatesOutput, error)
	DescribeEventDetailsWithContext(ctx context.Context, input *health.DescribeEventDetailsInput, opts ...request.Option) (*health.DescribeEventDetailsOutput, error)
	DescribeEventTypesWithContext(ctx context.Context, input *health.DescribeEventTypesInput, opts ...request.Option) (*health.DescribeEventTypesOutput, error)
	DescribeEventsWithContext(ctx context.Context, input *health.DescribeEventsInput, opts ...request.Option) (*health.DescribeEventsOutput, error)
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
		Action:  "DescribeAffectedEntitiesWithContext",
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

func (c *Client) DescribeEntityAggregatesWithContext(ctx context.Context, input *health.DescribeEntityAggregatesInput, opts ...request.Option) (*health.DescribeEntityAggregatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEntityAggregatesWithContext",
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
		Action:  "DescribeEventAggregatesWithContext",
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

func (c *Client) DescribeEventDetailsWithContext(ctx context.Context, input *health.DescribeEventDetailsInput, opts ...request.Option) (*health.DescribeEventDetailsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEventDetailsWithContext",
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
		Action:  "DescribeEventTypesWithContext",
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

func (c *Client) DescribeEventsWithContext(ctx context.Context, input *health.DescribeEventsInput, opts ...request.Option) (*health.DescribeEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "health",
		Action:  "DescribeEventsWithContext",
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
