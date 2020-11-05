// Code generated by internal/generate/main.go. DO NOT EDIT.

package eventbridgectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"github.com/aws/aws-sdk-go/service/eventbridge/eventbridgeiface"
	"github.com/glassechidna/awsctx"
)

type EventBridge interface {
	ActivateEventSourceWithContext(ctx context.Context, input *eventbridge.ActivateEventSourceInput, opts ...request.Option) (*eventbridge.ActivateEventSourceOutput, error)
	CancelReplayWithContext(ctx context.Context, input *eventbridge.CancelReplayInput, opts ...request.Option) (*eventbridge.CancelReplayOutput, error)
	CreateArchiveWithContext(ctx context.Context, input *eventbridge.CreateArchiveInput, opts ...request.Option) (*eventbridge.CreateArchiveOutput, error)
	CreateEventBusWithContext(ctx context.Context, input *eventbridge.CreateEventBusInput, opts ...request.Option) (*eventbridge.CreateEventBusOutput, error)
	CreatePartnerEventSourceWithContext(ctx context.Context, input *eventbridge.CreatePartnerEventSourceInput, opts ...request.Option) (*eventbridge.CreatePartnerEventSourceOutput, error)
	DeactivateEventSourceWithContext(ctx context.Context, input *eventbridge.DeactivateEventSourceInput, opts ...request.Option) (*eventbridge.DeactivateEventSourceOutput, error)
	DeleteArchiveWithContext(ctx context.Context, input *eventbridge.DeleteArchiveInput, opts ...request.Option) (*eventbridge.DeleteArchiveOutput, error)
	DeleteEventBusWithContext(ctx context.Context, input *eventbridge.DeleteEventBusInput, opts ...request.Option) (*eventbridge.DeleteEventBusOutput, error)
	DeletePartnerEventSourceWithContext(ctx context.Context, input *eventbridge.DeletePartnerEventSourceInput, opts ...request.Option) (*eventbridge.DeletePartnerEventSourceOutput, error)
	DeleteRuleWithContext(ctx context.Context, input *eventbridge.DeleteRuleInput, opts ...request.Option) (*eventbridge.DeleteRuleOutput, error)
	DescribeArchiveWithContext(ctx context.Context, input *eventbridge.DescribeArchiveInput, opts ...request.Option) (*eventbridge.DescribeArchiveOutput, error)
	DescribeEventBusWithContext(ctx context.Context, input *eventbridge.DescribeEventBusInput, opts ...request.Option) (*eventbridge.DescribeEventBusOutput, error)
	DescribeEventSourceWithContext(ctx context.Context, input *eventbridge.DescribeEventSourceInput, opts ...request.Option) (*eventbridge.DescribeEventSourceOutput, error)
	DescribePartnerEventSourceWithContext(ctx context.Context, input *eventbridge.DescribePartnerEventSourceInput, opts ...request.Option) (*eventbridge.DescribePartnerEventSourceOutput, error)
	DescribeReplayWithContext(ctx context.Context, input *eventbridge.DescribeReplayInput, opts ...request.Option) (*eventbridge.DescribeReplayOutput, error)
	DescribeRuleWithContext(ctx context.Context, input *eventbridge.DescribeRuleInput, opts ...request.Option) (*eventbridge.DescribeRuleOutput, error)
	DisableRuleWithContext(ctx context.Context, input *eventbridge.DisableRuleInput, opts ...request.Option) (*eventbridge.DisableRuleOutput, error)
	EnableRuleWithContext(ctx context.Context, input *eventbridge.EnableRuleInput, opts ...request.Option) (*eventbridge.EnableRuleOutput, error)
	ListArchivesWithContext(ctx context.Context, input *eventbridge.ListArchivesInput, opts ...request.Option) (*eventbridge.ListArchivesOutput, error)
	ListEventBusesWithContext(ctx context.Context, input *eventbridge.ListEventBusesInput, opts ...request.Option) (*eventbridge.ListEventBusesOutput, error)
	ListEventSourcesWithContext(ctx context.Context, input *eventbridge.ListEventSourcesInput, opts ...request.Option) (*eventbridge.ListEventSourcesOutput, error)
	ListPartnerEventSourceAccountsWithContext(ctx context.Context, input *eventbridge.ListPartnerEventSourceAccountsInput, opts ...request.Option) (*eventbridge.ListPartnerEventSourceAccountsOutput, error)
	ListPartnerEventSourcesWithContext(ctx context.Context, input *eventbridge.ListPartnerEventSourcesInput, opts ...request.Option) (*eventbridge.ListPartnerEventSourcesOutput, error)
	ListReplaysWithContext(ctx context.Context, input *eventbridge.ListReplaysInput, opts ...request.Option) (*eventbridge.ListReplaysOutput, error)
	ListRuleNamesByTargetWithContext(ctx context.Context, input *eventbridge.ListRuleNamesByTargetInput, opts ...request.Option) (*eventbridge.ListRuleNamesByTargetOutput, error)
	ListRulesWithContext(ctx context.Context, input *eventbridge.ListRulesInput, opts ...request.Option) (*eventbridge.ListRulesOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *eventbridge.ListTagsForResourceInput, opts ...request.Option) (*eventbridge.ListTagsForResourceOutput, error)
	ListTargetsByRuleWithContext(ctx context.Context, input *eventbridge.ListTargetsByRuleInput, opts ...request.Option) (*eventbridge.ListTargetsByRuleOutput, error)
	PutEventsWithContext(ctx context.Context, input *eventbridge.PutEventsInput, opts ...request.Option) (*eventbridge.PutEventsOutput, error)
	PutPartnerEventsWithContext(ctx context.Context, input *eventbridge.PutPartnerEventsInput, opts ...request.Option) (*eventbridge.PutPartnerEventsOutput, error)
	PutPermissionWithContext(ctx context.Context, input *eventbridge.PutPermissionInput, opts ...request.Option) (*eventbridge.PutPermissionOutput, error)
	PutRuleWithContext(ctx context.Context, input *eventbridge.PutRuleInput, opts ...request.Option) (*eventbridge.PutRuleOutput, error)
	PutTargetsWithContext(ctx context.Context, input *eventbridge.PutTargetsInput, opts ...request.Option) (*eventbridge.PutTargetsOutput, error)
	RemovePermissionWithContext(ctx context.Context, input *eventbridge.RemovePermissionInput, opts ...request.Option) (*eventbridge.RemovePermissionOutput, error)
	RemoveTargetsWithContext(ctx context.Context, input *eventbridge.RemoveTargetsInput, opts ...request.Option) (*eventbridge.RemoveTargetsOutput, error)
	StartReplayWithContext(ctx context.Context, input *eventbridge.StartReplayInput, opts ...request.Option) (*eventbridge.StartReplayOutput, error)
	TagResourceWithContext(ctx context.Context, input *eventbridge.TagResourceInput, opts ...request.Option) (*eventbridge.TagResourceOutput, error)
	TestEventPatternWithContext(ctx context.Context, input *eventbridge.TestEventPatternInput, opts ...request.Option) (*eventbridge.TestEventPatternOutput, error)
	UntagResourceWithContext(ctx context.Context, input *eventbridge.UntagResourceInput, opts ...request.Option) (*eventbridge.UntagResourceOutput, error)
	UpdateArchiveWithContext(ctx context.Context, input *eventbridge.UpdateArchiveInput, opts ...request.Option) (*eventbridge.UpdateArchiveOutput, error)
}

type Client struct {
	eventbridgeiface.EventBridgeAPI
	Contexter awsctx.Contexter
}

func New(base eventbridgeiface.EventBridgeAPI, ctxer awsctx.Contexter) EventBridge {
	return &Client{
		EventBridgeAPI: base,
		Contexter: ctxer,
	}
}

var _ EventBridge = (*eventbridge.EventBridge)(nil)
var _ EventBridge = (*Client)(nil)

func (c *Client) ActivateEventSourceWithContext(ctx context.Context, input *eventbridge.ActivateEventSourceInput, opts ...request.Option) (*eventbridge.ActivateEventSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ActivateEventSource",
		Input:   input,
		Output:  (*eventbridge.ActivateEventSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ActivateEventSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ActivateEventSourceOutput), req.Error
}

func (c *Client) CancelReplayWithContext(ctx context.Context, input *eventbridge.CancelReplayInput, opts ...request.Option) (*eventbridge.CancelReplayOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "CancelReplay",
		Input:   input,
		Output:  (*eventbridge.CancelReplayOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.CancelReplayWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.CancelReplayOutput), req.Error
}

func (c *Client) CreateArchiveWithContext(ctx context.Context, input *eventbridge.CreateArchiveInput, opts ...request.Option) (*eventbridge.CreateArchiveOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "CreateArchive",
		Input:   input,
		Output:  (*eventbridge.CreateArchiveOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.CreateArchiveWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.CreateArchiveOutput), req.Error
}

func (c *Client) CreateEventBusWithContext(ctx context.Context, input *eventbridge.CreateEventBusInput, opts ...request.Option) (*eventbridge.CreateEventBusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "CreateEventBus",
		Input:   input,
		Output:  (*eventbridge.CreateEventBusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.CreateEventBusWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.CreateEventBusOutput), req.Error
}

func (c *Client) CreatePartnerEventSourceWithContext(ctx context.Context, input *eventbridge.CreatePartnerEventSourceInput, opts ...request.Option) (*eventbridge.CreatePartnerEventSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "CreatePartnerEventSource",
		Input:   input,
		Output:  (*eventbridge.CreatePartnerEventSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.CreatePartnerEventSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.CreatePartnerEventSourceOutput), req.Error
}

func (c *Client) DeactivateEventSourceWithContext(ctx context.Context, input *eventbridge.DeactivateEventSourceInput, opts ...request.Option) (*eventbridge.DeactivateEventSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DeactivateEventSource",
		Input:   input,
		Output:  (*eventbridge.DeactivateEventSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DeactivateEventSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DeactivateEventSourceOutput), req.Error
}

func (c *Client) DeleteArchiveWithContext(ctx context.Context, input *eventbridge.DeleteArchiveInput, opts ...request.Option) (*eventbridge.DeleteArchiveOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DeleteArchive",
		Input:   input,
		Output:  (*eventbridge.DeleteArchiveOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DeleteArchiveWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DeleteArchiveOutput), req.Error
}

func (c *Client) DeleteEventBusWithContext(ctx context.Context, input *eventbridge.DeleteEventBusInput, opts ...request.Option) (*eventbridge.DeleteEventBusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DeleteEventBus",
		Input:   input,
		Output:  (*eventbridge.DeleteEventBusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DeleteEventBusWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DeleteEventBusOutput), req.Error
}

func (c *Client) DeletePartnerEventSourceWithContext(ctx context.Context, input *eventbridge.DeletePartnerEventSourceInput, opts ...request.Option) (*eventbridge.DeletePartnerEventSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DeletePartnerEventSource",
		Input:   input,
		Output:  (*eventbridge.DeletePartnerEventSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DeletePartnerEventSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DeletePartnerEventSourceOutput), req.Error
}

func (c *Client) DeleteRuleWithContext(ctx context.Context, input *eventbridge.DeleteRuleInput, opts ...request.Option) (*eventbridge.DeleteRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DeleteRule",
		Input:   input,
		Output:  (*eventbridge.DeleteRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DeleteRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DeleteRuleOutput), req.Error
}

func (c *Client) DescribeArchiveWithContext(ctx context.Context, input *eventbridge.DescribeArchiveInput, opts ...request.Option) (*eventbridge.DescribeArchiveOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DescribeArchive",
		Input:   input,
		Output:  (*eventbridge.DescribeArchiveOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DescribeArchiveWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DescribeArchiveOutput), req.Error
}

func (c *Client) DescribeEventBusWithContext(ctx context.Context, input *eventbridge.DescribeEventBusInput, opts ...request.Option) (*eventbridge.DescribeEventBusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DescribeEventBus",
		Input:   input,
		Output:  (*eventbridge.DescribeEventBusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DescribeEventBusWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DescribeEventBusOutput), req.Error
}

func (c *Client) DescribeEventSourceWithContext(ctx context.Context, input *eventbridge.DescribeEventSourceInput, opts ...request.Option) (*eventbridge.DescribeEventSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DescribeEventSource",
		Input:   input,
		Output:  (*eventbridge.DescribeEventSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DescribeEventSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DescribeEventSourceOutput), req.Error
}

func (c *Client) DescribePartnerEventSourceWithContext(ctx context.Context, input *eventbridge.DescribePartnerEventSourceInput, opts ...request.Option) (*eventbridge.DescribePartnerEventSourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DescribePartnerEventSource",
		Input:   input,
		Output:  (*eventbridge.DescribePartnerEventSourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DescribePartnerEventSourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DescribePartnerEventSourceOutput), req.Error
}

func (c *Client) DescribeReplayWithContext(ctx context.Context, input *eventbridge.DescribeReplayInput, opts ...request.Option) (*eventbridge.DescribeReplayOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DescribeReplay",
		Input:   input,
		Output:  (*eventbridge.DescribeReplayOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DescribeReplayWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DescribeReplayOutput), req.Error
}

func (c *Client) DescribeRuleWithContext(ctx context.Context, input *eventbridge.DescribeRuleInput, opts ...request.Option) (*eventbridge.DescribeRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DescribeRule",
		Input:   input,
		Output:  (*eventbridge.DescribeRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DescribeRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DescribeRuleOutput), req.Error
}

func (c *Client) DisableRuleWithContext(ctx context.Context, input *eventbridge.DisableRuleInput, opts ...request.Option) (*eventbridge.DisableRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "DisableRule",
		Input:   input,
		Output:  (*eventbridge.DisableRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.DisableRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.DisableRuleOutput), req.Error
}

func (c *Client) EnableRuleWithContext(ctx context.Context, input *eventbridge.EnableRuleInput, opts ...request.Option) (*eventbridge.EnableRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "EnableRule",
		Input:   input,
		Output:  (*eventbridge.EnableRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.EnableRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.EnableRuleOutput), req.Error
}

func (c *Client) ListArchivesWithContext(ctx context.Context, input *eventbridge.ListArchivesInput, opts ...request.Option) (*eventbridge.ListArchivesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListArchives",
		Input:   input,
		Output:  (*eventbridge.ListArchivesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListArchivesWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListArchivesOutput), req.Error
}

func (c *Client) ListEventBusesWithContext(ctx context.Context, input *eventbridge.ListEventBusesInput, opts ...request.Option) (*eventbridge.ListEventBusesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListEventBuses",
		Input:   input,
		Output:  (*eventbridge.ListEventBusesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListEventBusesWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListEventBusesOutput), req.Error
}

func (c *Client) ListEventSourcesWithContext(ctx context.Context, input *eventbridge.ListEventSourcesInput, opts ...request.Option) (*eventbridge.ListEventSourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListEventSources",
		Input:   input,
		Output:  (*eventbridge.ListEventSourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListEventSourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListEventSourcesOutput), req.Error
}

func (c *Client) ListPartnerEventSourceAccountsWithContext(ctx context.Context, input *eventbridge.ListPartnerEventSourceAccountsInput, opts ...request.Option) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListPartnerEventSourceAccounts",
		Input:   input,
		Output:  (*eventbridge.ListPartnerEventSourceAccountsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListPartnerEventSourceAccountsWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListPartnerEventSourceAccountsOutput), req.Error
}

func (c *Client) ListPartnerEventSourcesWithContext(ctx context.Context, input *eventbridge.ListPartnerEventSourcesInput, opts ...request.Option) (*eventbridge.ListPartnerEventSourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListPartnerEventSources",
		Input:   input,
		Output:  (*eventbridge.ListPartnerEventSourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListPartnerEventSourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListPartnerEventSourcesOutput), req.Error
}

func (c *Client) ListReplaysWithContext(ctx context.Context, input *eventbridge.ListReplaysInput, opts ...request.Option) (*eventbridge.ListReplaysOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListReplays",
		Input:   input,
		Output:  (*eventbridge.ListReplaysOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListReplaysWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListReplaysOutput), req.Error
}

func (c *Client) ListRuleNamesByTargetWithContext(ctx context.Context, input *eventbridge.ListRuleNamesByTargetInput, opts ...request.Option) (*eventbridge.ListRuleNamesByTargetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListRuleNamesByTarget",
		Input:   input,
		Output:  (*eventbridge.ListRuleNamesByTargetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListRuleNamesByTargetWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListRuleNamesByTargetOutput), req.Error
}

func (c *Client) ListRulesWithContext(ctx context.Context, input *eventbridge.ListRulesInput, opts ...request.Option) (*eventbridge.ListRulesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListRules",
		Input:   input,
		Output:  (*eventbridge.ListRulesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListRulesWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListRulesOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *eventbridge.ListTagsForResourceInput, opts ...request.Option) (*eventbridge.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*eventbridge.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTargetsByRuleWithContext(ctx context.Context, input *eventbridge.ListTargetsByRuleInput, opts ...request.Option) (*eventbridge.ListTargetsByRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "ListTargetsByRule",
		Input:   input,
		Output:  (*eventbridge.ListTargetsByRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.ListTargetsByRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.ListTargetsByRuleOutput), req.Error
}

func (c *Client) PutEventsWithContext(ctx context.Context, input *eventbridge.PutEventsInput, opts ...request.Option) (*eventbridge.PutEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "PutEvents",
		Input:   input,
		Output:  (*eventbridge.PutEventsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.PutEventsWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.PutEventsOutput), req.Error
}

func (c *Client) PutPartnerEventsWithContext(ctx context.Context, input *eventbridge.PutPartnerEventsInput, opts ...request.Option) (*eventbridge.PutPartnerEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "PutPartnerEvents",
		Input:   input,
		Output:  (*eventbridge.PutPartnerEventsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.PutPartnerEventsWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.PutPartnerEventsOutput), req.Error
}

func (c *Client) PutPermissionWithContext(ctx context.Context, input *eventbridge.PutPermissionInput, opts ...request.Option) (*eventbridge.PutPermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "PutPermission",
		Input:   input,
		Output:  (*eventbridge.PutPermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.PutPermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.PutPermissionOutput), req.Error
}

func (c *Client) PutRuleWithContext(ctx context.Context, input *eventbridge.PutRuleInput, opts ...request.Option) (*eventbridge.PutRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "PutRule",
		Input:   input,
		Output:  (*eventbridge.PutRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.PutRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.PutRuleOutput), req.Error
}

func (c *Client) PutTargetsWithContext(ctx context.Context, input *eventbridge.PutTargetsInput, opts ...request.Option) (*eventbridge.PutTargetsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "PutTargets",
		Input:   input,
		Output:  (*eventbridge.PutTargetsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.PutTargetsWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.PutTargetsOutput), req.Error
}

func (c *Client) RemovePermissionWithContext(ctx context.Context, input *eventbridge.RemovePermissionInput, opts ...request.Option) (*eventbridge.RemovePermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "RemovePermission",
		Input:   input,
		Output:  (*eventbridge.RemovePermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.RemovePermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.RemovePermissionOutput), req.Error
}

func (c *Client) RemoveTargetsWithContext(ctx context.Context, input *eventbridge.RemoveTargetsInput, opts ...request.Option) (*eventbridge.RemoveTargetsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "RemoveTargets",
		Input:   input,
		Output:  (*eventbridge.RemoveTargetsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.RemoveTargetsWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.RemoveTargetsOutput), req.Error
}

func (c *Client) StartReplayWithContext(ctx context.Context, input *eventbridge.StartReplayInput, opts ...request.Option) (*eventbridge.StartReplayOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "StartReplay",
		Input:   input,
		Output:  (*eventbridge.StartReplayOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.StartReplayWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.StartReplayOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *eventbridge.TagResourceInput, opts ...request.Option) (*eventbridge.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "TagResource",
		Input:   input,
		Output:  (*eventbridge.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.TagResourceOutput), req.Error
}

func (c *Client) TestEventPatternWithContext(ctx context.Context, input *eventbridge.TestEventPatternInput, opts ...request.Option) (*eventbridge.TestEventPatternOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "TestEventPattern",
		Input:   input,
		Output:  (*eventbridge.TestEventPatternOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.TestEventPatternWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.TestEventPatternOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *eventbridge.UntagResourceInput, opts ...request.Option) (*eventbridge.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*eventbridge.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.UntagResourceOutput), req.Error
}

func (c *Client) UpdateArchiveWithContext(ctx context.Context, input *eventbridge.UpdateArchiveInput, opts ...request.Option) (*eventbridge.UpdateArchiveOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "eventbridge",
		Action:  "UpdateArchive",
		Input:   input,
		Output:  (*eventbridge.UpdateArchiveOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.EventBridgeAPI.UpdateArchiveWithContext(ctx, input, opts...)
	})

	return req.Output.(*eventbridge.UpdateArchiveOutput), req.Error
}
