// Code generated by internal/generate/main.go. DO NOT EDIT.

package personalizeeventsctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/personalizeevents"
	"github.com/aws/aws-sdk-go/service/personalizeevents/personalizeeventsiface"
	"github.com/glassechidna/awsctx"
)

type PersonalizeEvents interface {
	PutEventsWithContext(ctx context.Context, input *personalizeevents.PutEventsInput, opts ...request.Option) (*personalizeevents.PutEventsOutput, error)
	PutItemsWithContext(ctx context.Context, input *personalizeevents.PutItemsInput, opts ...request.Option) (*personalizeevents.PutItemsOutput, error)
	PutUsersWithContext(ctx context.Context, input *personalizeevents.PutUsersInput, opts ...request.Option) (*personalizeevents.PutUsersOutput, error)
}

type Client struct {
	personalizeeventsiface.PersonalizeEventsAPI
	Contexter awsctx.Contexter
}

func New(base personalizeeventsiface.PersonalizeEventsAPI, ctxer awsctx.Contexter) PersonalizeEvents {
	return &Client{
		PersonalizeEventsAPI: base,
		Contexter: ctxer,
	}
}

var _ PersonalizeEvents = (*personalizeevents.PersonalizeEvents)(nil)
var _ PersonalizeEvents = (*Client)(nil)

func (c *Client) PutEventsWithContext(ctx context.Context, input *personalizeevents.PutEventsInput, opts ...request.Option) (*personalizeevents.PutEventsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "personalizeevents",
		Action:  "PutEvents",
		Input:   input,
		Output:  (*personalizeevents.PutEventsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PersonalizeEventsAPI.PutEventsWithContext(ctx, input, opts...)
	})

	return req.Output.(*personalizeevents.PutEventsOutput), req.Error
}

func (c *Client) PutItemsWithContext(ctx context.Context, input *personalizeevents.PutItemsInput, opts ...request.Option) (*personalizeevents.PutItemsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "personalizeevents",
		Action:  "PutItems",
		Input:   input,
		Output:  (*personalizeevents.PutItemsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PersonalizeEventsAPI.PutItemsWithContext(ctx, input, opts...)
	})

	return req.Output.(*personalizeevents.PutItemsOutput), req.Error
}

func (c *Client) PutUsersWithContext(ctx context.Context, input *personalizeevents.PutUsersInput, opts ...request.Option) (*personalizeevents.PutUsersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "personalizeevents",
		Action:  "PutUsers",
		Input:   input,
		Output:  (*personalizeevents.PutUsersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.PersonalizeEventsAPI.PutUsersWithContext(ctx, input, opts...)
	})

	return req.Output.(*personalizeevents.PutUsersOutput), req.Error
}
