// Code generated by internal/generate/main.go. DO NOT EDIT.

package snsctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/glassechidna/awsctx"
)

type SNS interface {
	AddPermissionWithContext(ctx context.Context, input *sns.AddPermissionInput, opts ...request.Option) (*sns.AddPermissionOutput, error)
	CheckIfPhoneNumberIsOptedOutWithContext(ctx context.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput, opts ...request.Option) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error)
	ConfirmSubscriptionWithContext(ctx context.Context, input *sns.ConfirmSubscriptionInput, opts ...request.Option) (*sns.ConfirmSubscriptionOutput, error)
	CreatePlatformApplicationWithContext(ctx context.Context, input *sns.CreatePlatformApplicationInput, opts ...request.Option) (*sns.CreatePlatformApplicationOutput, error)
	CreatePlatformEndpointWithContext(ctx context.Context, input *sns.CreatePlatformEndpointInput, opts ...request.Option) (*sns.CreatePlatformEndpointOutput, error)
	CreateTopicWithContext(ctx context.Context, input *sns.CreateTopicInput, opts ...request.Option) (*sns.CreateTopicOutput, error)
	DeleteEndpointWithContext(ctx context.Context, input *sns.DeleteEndpointInput, opts ...request.Option) (*sns.DeleteEndpointOutput, error)
	DeletePlatformApplicationWithContext(ctx context.Context, input *sns.DeletePlatformApplicationInput, opts ...request.Option) (*sns.DeletePlatformApplicationOutput, error)
	DeleteTopicWithContext(ctx context.Context, input *sns.DeleteTopicInput, opts ...request.Option) (*sns.DeleteTopicOutput, error)
	GetEndpointAttributesWithContext(ctx context.Context, input *sns.GetEndpointAttributesInput, opts ...request.Option) (*sns.GetEndpointAttributesOutput, error)
	GetPlatformApplicationAttributesWithContext(ctx context.Context, input *sns.GetPlatformApplicationAttributesInput, opts ...request.Option) (*sns.GetPlatformApplicationAttributesOutput, error)
	GetSMSAttributesWithContext(ctx context.Context, input *sns.GetSMSAttributesInput, opts ...request.Option) (*sns.GetSMSAttributesOutput, error)
	GetSubscriptionAttributesWithContext(ctx context.Context, input *sns.GetSubscriptionAttributesInput, opts ...request.Option) (*sns.GetSubscriptionAttributesOutput, error)
	GetTopicAttributesWithContext(ctx context.Context, input *sns.GetTopicAttributesInput, opts ...request.Option) (*sns.GetTopicAttributesOutput, error)
	ListEndpointsByPlatformApplicationWithContext(ctx context.Context, input *sns.ListEndpointsByPlatformApplicationInput, opts ...request.Option) (*sns.ListEndpointsByPlatformApplicationOutput, error)
	ListPhoneNumbersOptedOutWithContext(ctx context.Context, input *sns.ListPhoneNumbersOptedOutInput, opts ...request.Option) (*sns.ListPhoneNumbersOptedOutOutput, error)
	ListPlatformApplicationsWithContext(ctx context.Context, input *sns.ListPlatformApplicationsInput, opts ...request.Option) (*sns.ListPlatformApplicationsOutput, error)
	ListSubscriptionsWithContext(ctx context.Context, input *sns.ListSubscriptionsInput, opts ...request.Option) (*sns.ListSubscriptionsOutput, error)
	ListSubscriptionsByTopicWithContext(ctx context.Context, input *sns.ListSubscriptionsByTopicInput, opts ...request.Option) (*sns.ListSubscriptionsByTopicOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *sns.ListTagsForResourceInput, opts ...request.Option) (*sns.ListTagsForResourceOutput, error)
	ListTopicsWithContext(ctx context.Context, input *sns.ListTopicsInput, opts ...request.Option) (*sns.ListTopicsOutput, error)
	OptInPhoneNumberWithContext(ctx context.Context, input *sns.OptInPhoneNumberInput, opts ...request.Option) (*sns.OptInPhoneNumberOutput, error)
	PublishWithContext(ctx context.Context, input *sns.PublishInput, opts ...request.Option) (*sns.PublishOutput, error)
	RemovePermissionWithContext(ctx context.Context, input *sns.RemovePermissionInput, opts ...request.Option) (*sns.RemovePermissionOutput, error)
	SetEndpointAttributesWithContext(ctx context.Context, input *sns.SetEndpointAttributesInput, opts ...request.Option) (*sns.SetEndpointAttributesOutput, error)
	SetPlatformApplicationAttributesWithContext(ctx context.Context, input *sns.SetPlatformApplicationAttributesInput, opts ...request.Option) (*sns.SetPlatformApplicationAttributesOutput, error)
	SetSMSAttributesWithContext(ctx context.Context, input *sns.SetSMSAttributesInput, opts ...request.Option) (*sns.SetSMSAttributesOutput, error)
	SetSubscriptionAttributesWithContext(ctx context.Context, input *sns.SetSubscriptionAttributesInput, opts ...request.Option) (*sns.SetSubscriptionAttributesOutput, error)
	SetTopicAttributesWithContext(ctx context.Context, input *sns.SetTopicAttributesInput, opts ...request.Option) (*sns.SetTopicAttributesOutput, error)
	SubscribeWithContext(ctx context.Context, input *sns.SubscribeInput, opts ...request.Option) (*sns.SubscribeOutput, error)
	TagResourceWithContext(ctx context.Context, input *sns.TagResourceInput, opts ...request.Option) (*sns.TagResourceOutput, error)
	UnsubscribeWithContext(ctx context.Context, input *sns.UnsubscribeInput, opts ...request.Option) (*sns.UnsubscribeOutput, error)
	UntagResourceWithContext(ctx context.Context, input *sns.UntagResourceInput, opts ...request.Option) (*sns.UntagResourceOutput, error)
}

type Client struct {
	snsiface.SNSAPI
	Contexter awsctx.Contexter
}

func New(base snsiface.SNSAPI, ctxer awsctx.Contexter) SNS {
	return &Client{
		SNSAPI: base,
		Contexter: ctxer,
	}
}

var _ SNS = (*sns.SNS)(nil)
var _ SNS = (*Client)(nil)

func (c *Client) AddPermissionWithContext(ctx context.Context, input *sns.AddPermissionInput, opts ...request.Option) (*sns.AddPermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "AddPermission",
		Input:   input,
		Output:  (*sns.AddPermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.AddPermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.AddPermissionOutput), req.Error
}

func (c *Client) CheckIfPhoneNumberIsOptedOutWithContext(ctx context.Context, input *sns.CheckIfPhoneNumberIsOptedOutInput, opts ...request.Option) (*sns.CheckIfPhoneNumberIsOptedOutOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "CheckIfPhoneNumberIsOptedOut",
		Input:   input,
		Output:  (*sns.CheckIfPhoneNumberIsOptedOutOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.CheckIfPhoneNumberIsOptedOutWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.CheckIfPhoneNumberIsOptedOutOutput), req.Error
}

func (c *Client) ConfirmSubscriptionWithContext(ctx context.Context, input *sns.ConfirmSubscriptionInput, opts ...request.Option) (*sns.ConfirmSubscriptionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "ConfirmSubscription",
		Input:   input,
		Output:  (*sns.ConfirmSubscriptionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.ConfirmSubscriptionWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.ConfirmSubscriptionOutput), req.Error
}

func (c *Client) CreatePlatformApplicationWithContext(ctx context.Context, input *sns.CreatePlatformApplicationInput, opts ...request.Option) (*sns.CreatePlatformApplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "CreatePlatformApplication",
		Input:   input,
		Output:  (*sns.CreatePlatformApplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.CreatePlatformApplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.CreatePlatformApplicationOutput), req.Error
}

func (c *Client) CreatePlatformEndpointWithContext(ctx context.Context, input *sns.CreatePlatformEndpointInput, opts ...request.Option) (*sns.CreatePlatformEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "CreatePlatformEndpoint",
		Input:   input,
		Output:  (*sns.CreatePlatformEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.CreatePlatformEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.CreatePlatformEndpointOutput), req.Error
}

func (c *Client) CreateTopicWithContext(ctx context.Context, input *sns.CreateTopicInput, opts ...request.Option) (*sns.CreateTopicOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "CreateTopic",
		Input:   input,
		Output:  (*sns.CreateTopicOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.CreateTopicWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.CreateTopicOutput), req.Error
}

func (c *Client) DeleteEndpointWithContext(ctx context.Context, input *sns.DeleteEndpointInput, opts ...request.Option) (*sns.DeleteEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "DeleteEndpoint",
		Input:   input,
		Output:  (*sns.DeleteEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.DeleteEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.DeleteEndpointOutput), req.Error
}

func (c *Client) DeletePlatformApplicationWithContext(ctx context.Context, input *sns.DeletePlatformApplicationInput, opts ...request.Option) (*sns.DeletePlatformApplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "DeletePlatformApplication",
		Input:   input,
		Output:  (*sns.DeletePlatformApplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.DeletePlatformApplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.DeletePlatformApplicationOutput), req.Error
}

func (c *Client) DeleteTopicWithContext(ctx context.Context, input *sns.DeleteTopicInput, opts ...request.Option) (*sns.DeleteTopicOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "DeleteTopic",
		Input:   input,
		Output:  (*sns.DeleteTopicOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.DeleteTopicWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.DeleteTopicOutput), req.Error
}

func (c *Client) GetEndpointAttributesWithContext(ctx context.Context, input *sns.GetEndpointAttributesInput, opts ...request.Option) (*sns.GetEndpointAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "GetEndpointAttributes",
		Input:   input,
		Output:  (*sns.GetEndpointAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.GetEndpointAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.GetEndpointAttributesOutput), req.Error
}

func (c *Client) GetPlatformApplicationAttributesWithContext(ctx context.Context, input *sns.GetPlatformApplicationAttributesInput, opts ...request.Option) (*sns.GetPlatformApplicationAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "GetPlatformApplicationAttributes",
		Input:   input,
		Output:  (*sns.GetPlatformApplicationAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.GetPlatformApplicationAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.GetPlatformApplicationAttributesOutput), req.Error
}

func (c *Client) GetSMSAttributesWithContext(ctx context.Context, input *sns.GetSMSAttributesInput, opts ...request.Option) (*sns.GetSMSAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "GetSMSAttributes",
		Input:   input,
		Output:  (*sns.GetSMSAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.GetSMSAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.GetSMSAttributesOutput), req.Error
}

func (c *Client) GetSubscriptionAttributesWithContext(ctx context.Context, input *sns.GetSubscriptionAttributesInput, opts ...request.Option) (*sns.GetSubscriptionAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "GetSubscriptionAttributes",
		Input:   input,
		Output:  (*sns.GetSubscriptionAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.GetSubscriptionAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.GetSubscriptionAttributesOutput), req.Error
}

func (c *Client) GetTopicAttributesWithContext(ctx context.Context, input *sns.GetTopicAttributesInput, opts ...request.Option) (*sns.GetTopicAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "GetTopicAttributes",
		Input:   input,
		Output:  (*sns.GetTopicAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.GetTopicAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.GetTopicAttributesOutput), req.Error
}

func (c *Client) ListEndpointsByPlatformApplicationWithContext(ctx context.Context, input *sns.ListEndpointsByPlatformApplicationInput, opts ...request.Option) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "ListEndpointsByPlatformApplication",
		Input:   input,
		Output:  (*sns.ListEndpointsByPlatformApplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.ListEndpointsByPlatformApplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.ListEndpointsByPlatformApplicationOutput), req.Error
}

func (c *Client) ListPhoneNumbersOptedOutWithContext(ctx context.Context, input *sns.ListPhoneNumbersOptedOutInput, opts ...request.Option) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "ListPhoneNumbersOptedOut",
		Input:   input,
		Output:  (*sns.ListPhoneNumbersOptedOutOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.ListPhoneNumbersOptedOutWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.ListPhoneNumbersOptedOutOutput), req.Error
}

func (c *Client) ListPlatformApplicationsWithContext(ctx context.Context, input *sns.ListPlatformApplicationsInput, opts ...request.Option) (*sns.ListPlatformApplicationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "ListPlatformApplications",
		Input:   input,
		Output:  (*sns.ListPlatformApplicationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.ListPlatformApplicationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.ListPlatformApplicationsOutput), req.Error
}

func (c *Client) ListSubscriptionsWithContext(ctx context.Context, input *sns.ListSubscriptionsInput, opts ...request.Option) (*sns.ListSubscriptionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "ListSubscriptions",
		Input:   input,
		Output:  (*sns.ListSubscriptionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.ListSubscriptionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.ListSubscriptionsOutput), req.Error
}

func (c *Client) ListSubscriptionsByTopicWithContext(ctx context.Context, input *sns.ListSubscriptionsByTopicInput, opts ...request.Option) (*sns.ListSubscriptionsByTopicOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "ListSubscriptionsByTopic",
		Input:   input,
		Output:  (*sns.ListSubscriptionsByTopicOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.ListSubscriptionsByTopicWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.ListSubscriptionsByTopicOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *sns.ListTagsForResourceInput, opts ...request.Option) (*sns.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*sns.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTopicsWithContext(ctx context.Context, input *sns.ListTopicsInput, opts ...request.Option) (*sns.ListTopicsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "ListTopics",
		Input:   input,
		Output:  (*sns.ListTopicsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.ListTopicsWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.ListTopicsOutput), req.Error
}

func (c *Client) OptInPhoneNumberWithContext(ctx context.Context, input *sns.OptInPhoneNumberInput, opts ...request.Option) (*sns.OptInPhoneNumberOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "OptInPhoneNumber",
		Input:   input,
		Output:  (*sns.OptInPhoneNumberOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.OptInPhoneNumberWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.OptInPhoneNumberOutput), req.Error
}

func (c *Client) PublishWithContext(ctx context.Context, input *sns.PublishInput, opts ...request.Option) (*sns.PublishOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "Publish",
		Input:   input,
		Output:  (*sns.PublishOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.PublishWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.PublishOutput), req.Error
}

func (c *Client) RemovePermissionWithContext(ctx context.Context, input *sns.RemovePermissionInput, opts ...request.Option) (*sns.RemovePermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "RemovePermission",
		Input:   input,
		Output:  (*sns.RemovePermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.RemovePermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.RemovePermissionOutput), req.Error
}

func (c *Client) SetEndpointAttributesWithContext(ctx context.Context, input *sns.SetEndpointAttributesInput, opts ...request.Option) (*sns.SetEndpointAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "SetEndpointAttributes",
		Input:   input,
		Output:  (*sns.SetEndpointAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.SetEndpointAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.SetEndpointAttributesOutput), req.Error
}

func (c *Client) SetPlatformApplicationAttributesWithContext(ctx context.Context, input *sns.SetPlatformApplicationAttributesInput, opts ...request.Option) (*sns.SetPlatformApplicationAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "SetPlatformApplicationAttributes",
		Input:   input,
		Output:  (*sns.SetPlatformApplicationAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.SetPlatformApplicationAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.SetPlatformApplicationAttributesOutput), req.Error
}

func (c *Client) SetSMSAttributesWithContext(ctx context.Context, input *sns.SetSMSAttributesInput, opts ...request.Option) (*sns.SetSMSAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "SetSMSAttributes",
		Input:   input,
		Output:  (*sns.SetSMSAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.SetSMSAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.SetSMSAttributesOutput), req.Error
}

func (c *Client) SetSubscriptionAttributesWithContext(ctx context.Context, input *sns.SetSubscriptionAttributesInput, opts ...request.Option) (*sns.SetSubscriptionAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "SetSubscriptionAttributes",
		Input:   input,
		Output:  (*sns.SetSubscriptionAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.SetSubscriptionAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.SetSubscriptionAttributesOutput), req.Error
}

func (c *Client) SetTopicAttributesWithContext(ctx context.Context, input *sns.SetTopicAttributesInput, opts ...request.Option) (*sns.SetTopicAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "SetTopicAttributes",
		Input:   input,
		Output:  (*sns.SetTopicAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.SetTopicAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.SetTopicAttributesOutput), req.Error
}

func (c *Client) SubscribeWithContext(ctx context.Context, input *sns.SubscribeInput, opts ...request.Option) (*sns.SubscribeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "Subscribe",
		Input:   input,
		Output:  (*sns.SubscribeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.SubscribeWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.SubscribeOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *sns.TagResourceInput, opts ...request.Option) (*sns.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "TagResource",
		Input:   input,
		Output:  (*sns.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.TagResourceOutput), req.Error
}

func (c *Client) UnsubscribeWithContext(ctx context.Context, input *sns.UnsubscribeInput, opts ...request.Option) (*sns.UnsubscribeOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "Unsubscribe",
		Input:   input,
		Output:  (*sns.UnsubscribeOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.UnsubscribeWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.UnsubscribeOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *sns.UntagResourceInput, opts ...request.Option) (*sns.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "sns",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*sns.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SNSAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*sns.UntagResourceOutput), req.Error
}
