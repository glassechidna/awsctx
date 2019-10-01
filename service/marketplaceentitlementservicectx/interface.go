// Code generated by internal/generate/main.go. DO NOT EDIT.

package marketplaceentitlementservicectx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice/marketplaceentitlementserviceiface"
	"github.com/glassechidna/awsctx"
)

type MarketplaceEntitlementService interface {
	GetEntitlementsWithContext(ctx context.Context, input *marketplaceentitlementservice.GetEntitlementsInput, opts ...request.Option) (*marketplaceentitlementservice.GetEntitlementsOutput, error)
}

type Client struct {
	marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI
	Contexter awsctx.Contexter
}

func New(base marketplaceentitlementserviceiface.MarketplaceEntitlementServiceAPI, ctxer awsctx.Contexter) MarketplaceEntitlementService {
	return &Client{
		MarketplaceEntitlementServiceAPI: base,
		Contexter: ctxer,
	}
}

var _ MarketplaceEntitlementService = (*marketplaceentitlementservice.MarketplaceEntitlementService)(nil)
var _ MarketplaceEntitlementService = (*Client)(nil)

func (c *Client) GetEntitlementsWithContext(ctx context.Context, input *marketplaceentitlementservice.GetEntitlementsInput, opts ...request.Option) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "marketplaceentitlementservice",
		Action:  "GetEntitlementsWithContext",
		Input:   input,
		Output:  (*marketplaceentitlementservice.GetEntitlementsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.MarketplaceEntitlementServiceAPI.GetEntitlementsWithContext(ctx, input, opts...)
	})

	return req.Output.(*marketplaceentitlementservice.GetEntitlementsOutput), req.Error
}
