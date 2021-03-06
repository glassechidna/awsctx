// Code generated by internal/generate/main.go. DO NOT EDIT.

package qldbsessionctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/qldbsession"
	"github.com/aws/aws-sdk-go/service/qldbsession/qldbsessioniface"
	"github.com/glassechidna/awsctx"
)

type QLDBSession interface {
	SendCommandWithContext(ctx context.Context, input *qldbsession.SendCommandInput, opts ...request.Option) (*qldbsession.SendCommandOutput, error)
}

type Client struct {
	qldbsessioniface.QLDBSessionAPI
	Contexter awsctx.Contexter
}

func New(base qldbsessioniface.QLDBSessionAPI, ctxer awsctx.Contexter) QLDBSession {
	return &Client{
		QLDBSessionAPI: base,
		Contexter: ctxer,
	}
}

var _ QLDBSession = (*qldbsession.QLDBSession)(nil)
var _ QLDBSession = (*Client)(nil)

func (c *Client) SendCommandWithContext(ctx context.Context, input *qldbsession.SendCommandInput, opts ...request.Option) (*qldbsession.SendCommandOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "qldbsession",
		Action:  "SendCommand",
		Input:   input,
		Output:  (*qldbsession.SendCommandOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.QLDBSessionAPI.SendCommandWithContext(ctx, input, opts...)
	})

	return req.Output.(*qldbsession.SendCommandOutput), req.Error
}
