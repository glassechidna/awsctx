// Code generated by internal/generate/main.go. DO NOT EDIT.

package route53resolverctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/route53resolver/route53resolveriface"
	"github.com/glassechidna/awsctx"
)

type Route53Resolver interface {
	AssociateResolverEndpointIpAddressWithContext(ctx context.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput, opts ...request.Option) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error)
	AssociateResolverRuleWithContext(ctx context.Context, input *route53resolver.AssociateResolverRuleInput, opts ...request.Option) (*route53resolver.AssociateResolverRuleOutput, error)
	CreateResolverEndpointWithContext(ctx context.Context, input *route53resolver.CreateResolverEndpointInput, opts ...request.Option) (*route53resolver.CreateResolverEndpointOutput, error)
	CreateResolverRuleWithContext(ctx context.Context, input *route53resolver.CreateResolverRuleInput, opts ...request.Option) (*route53resolver.CreateResolverRuleOutput, error)
	DeleteResolverEndpointWithContext(ctx context.Context, input *route53resolver.DeleteResolverEndpointInput, opts ...request.Option) (*route53resolver.DeleteResolverEndpointOutput, error)
	DeleteResolverRuleWithContext(ctx context.Context, input *route53resolver.DeleteResolverRuleInput, opts ...request.Option) (*route53resolver.DeleteResolverRuleOutput, error)
	DisassociateResolverEndpointIpAddressWithContext(ctx context.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput, opts ...request.Option) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error)
	DisassociateResolverRuleWithContext(ctx context.Context, input *route53resolver.DisassociateResolverRuleInput, opts ...request.Option) (*route53resolver.DisassociateResolverRuleOutput, error)
	GetResolverEndpointWithContext(ctx context.Context, input *route53resolver.GetResolverEndpointInput, opts ...request.Option) (*route53resolver.GetResolverEndpointOutput, error)
	GetResolverRuleWithContext(ctx context.Context, input *route53resolver.GetResolverRuleInput, opts ...request.Option) (*route53resolver.GetResolverRuleOutput, error)
	GetResolverRuleAssociationWithContext(ctx context.Context, input *route53resolver.GetResolverRuleAssociationInput, opts ...request.Option) (*route53resolver.GetResolverRuleAssociationOutput, error)
	GetResolverRulePolicyWithContext(ctx context.Context, input *route53resolver.GetResolverRulePolicyInput, opts ...request.Option) (*route53resolver.GetResolverRulePolicyOutput, error)
	ListResolverEndpointIpAddressesWithContext(ctx context.Context, input *route53resolver.ListResolverEndpointIpAddressesInput, opts ...request.Option) (*route53resolver.ListResolverEndpointIpAddressesOutput, error)
	ListResolverEndpointsWithContext(ctx context.Context, input *route53resolver.ListResolverEndpointsInput, opts ...request.Option) (*route53resolver.ListResolverEndpointsOutput, error)
	ListResolverRuleAssociationsWithContext(ctx context.Context, input *route53resolver.ListResolverRuleAssociationsInput, opts ...request.Option) (*route53resolver.ListResolverRuleAssociationsOutput, error)
	ListResolverRulesWithContext(ctx context.Context, input *route53resolver.ListResolverRulesInput, opts ...request.Option) (*route53resolver.ListResolverRulesOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *route53resolver.ListTagsForResourceInput, opts ...request.Option) (*route53resolver.ListTagsForResourceOutput, error)
	PutResolverRulePolicyWithContext(ctx context.Context, input *route53resolver.PutResolverRulePolicyInput, opts ...request.Option) (*route53resolver.PutResolverRulePolicyOutput, error)
	TagResourceWithContext(ctx context.Context, input *route53resolver.TagResourceInput, opts ...request.Option) (*route53resolver.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *route53resolver.UntagResourceInput, opts ...request.Option) (*route53resolver.UntagResourceOutput, error)
	UpdateResolverEndpointWithContext(ctx context.Context, input *route53resolver.UpdateResolverEndpointInput, opts ...request.Option) (*route53resolver.UpdateResolverEndpointOutput, error)
	UpdateResolverRuleWithContext(ctx context.Context, input *route53resolver.UpdateResolverRuleInput, opts ...request.Option) (*route53resolver.UpdateResolverRuleOutput, error)
}

type Client struct {
	route53resolveriface.Route53ResolverAPI
	Contexter awsctx.Contexter
}

func New(base route53resolveriface.Route53ResolverAPI, ctxer awsctx.Contexter) Route53Resolver {
	return &Client{
		Route53ResolverAPI: base,
		Contexter: ctxer,
	}
}

var _ Route53Resolver = (*route53resolver.Route53Resolver)(nil)
var _ Route53Resolver = (*Client)(nil)

func (c *Client) AssociateResolverEndpointIpAddressWithContext(ctx context.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput, opts ...request.Option) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "AssociateResolverEndpointIpAddress",
		Input:   input,
		Output:  (*route53resolver.AssociateResolverEndpointIpAddressOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.AssociateResolverEndpointIpAddressWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.AssociateResolverEndpointIpAddressOutput), req.Error
}

func (c *Client) AssociateResolverRuleWithContext(ctx context.Context, input *route53resolver.AssociateResolverRuleInput, opts ...request.Option) (*route53resolver.AssociateResolverRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "AssociateResolverRule",
		Input:   input,
		Output:  (*route53resolver.AssociateResolverRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.AssociateResolverRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.AssociateResolverRuleOutput), req.Error
}

func (c *Client) CreateResolverEndpointWithContext(ctx context.Context, input *route53resolver.CreateResolverEndpointInput, opts ...request.Option) (*route53resolver.CreateResolverEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "CreateResolverEndpoint",
		Input:   input,
		Output:  (*route53resolver.CreateResolverEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.CreateResolverEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.CreateResolverEndpointOutput), req.Error
}

func (c *Client) CreateResolverRuleWithContext(ctx context.Context, input *route53resolver.CreateResolverRuleInput, opts ...request.Option) (*route53resolver.CreateResolverRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "CreateResolverRule",
		Input:   input,
		Output:  (*route53resolver.CreateResolverRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.CreateResolverRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.CreateResolverRuleOutput), req.Error
}

func (c *Client) DeleteResolverEndpointWithContext(ctx context.Context, input *route53resolver.DeleteResolverEndpointInput, opts ...request.Option) (*route53resolver.DeleteResolverEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "DeleteResolverEndpoint",
		Input:   input,
		Output:  (*route53resolver.DeleteResolverEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.DeleteResolverEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.DeleteResolverEndpointOutput), req.Error
}

func (c *Client) DeleteResolverRuleWithContext(ctx context.Context, input *route53resolver.DeleteResolverRuleInput, opts ...request.Option) (*route53resolver.DeleteResolverRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "DeleteResolverRule",
		Input:   input,
		Output:  (*route53resolver.DeleteResolverRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.DeleteResolverRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.DeleteResolverRuleOutput), req.Error
}

func (c *Client) DisassociateResolverEndpointIpAddressWithContext(ctx context.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput, opts ...request.Option) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "DisassociateResolverEndpointIpAddress",
		Input:   input,
		Output:  (*route53resolver.DisassociateResolverEndpointIpAddressOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.DisassociateResolverEndpointIpAddressWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.DisassociateResolverEndpointIpAddressOutput), req.Error
}

func (c *Client) DisassociateResolverRuleWithContext(ctx context.Context, input *route53resolver.DisassociateResolverRuleInput, opts ...request.Option) (*route53resolver.DisassociateResolverRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "DisassociateResolverRule",
		Input:   input,
		Output:  (*route53resolver.DisassociateResolverRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.DisassociateResolverRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.DisassociateResolverRuleOutput), req.Error
}

func (c *Client) GetResolverEndpointWithContext(ctx context.Context, input *route53resolver.GetResolverEndpointInput, opts ...request.Option) (*route53resolver.GetResolverEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "GetResolverEndpoint",
		Input:   input,
		Output:  (*route53resolver.GetResolverEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.GetResolverEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.GetResolverEndpointOutput), req.Error
}

func (c *Client) GetResolverRuleWithContext(ctx context.Context, input *route53resolver.GetResolverRuleInput, opts ...request.Option) (*route53resolver.GetResolverRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "GetResolverRule",
		Input:   input,
		Output:  (*route53resolver.GetResolverRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.GetResolverRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.GetResolverRuleOutput), req.Error
}

func (c *Client) GetResolverRuleAssociationWithContext(ctx context.Context, input *route53resolver.GetResolverRuleAssociationInput, opts ...request.Option) (*route53resolver.GetResolverRuleAssociationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "GetResolverRuleAssociation",
		Input:   input,
		Output:  (*route53resolver.GetResolverRuleAssociationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.GetResolverRuleAssociationWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.GetResolverRuleAssociationOutput), req.Error
}

func (c *Client) GetResolverRulePolicyWithContext(ctx context.Context, input *route53resolver.GetResolverRulePolicyInput, opts ...request.Option) (*route53resolver.GetResolverRulePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "GetResolverRulePolicy",
		Input:   input,
		Output:  (*route53resolver.GetResolverRulePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.GetResolverRulePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.GetResolverRulePolicyOutput), req.Error
}

func (c *Client) ListResolverEndpointIpAddressesWithContext(ctx context.Context, input *route53resolver.ListResolverEndpointIpAddressesInput, opts ...request.Option) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "ListResolverEndpointIpAddresses",
		Input:   input,
		Output:  (*route53resolver.ListResolverEndpointIpAddressesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.ListResolverEndpointIpAddressesWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.ListResolverEndpointIpAddressesOutput), req.Error
}

func (c *Client) ListResolverEndpointsWithContext(ctx context.Context, input *route53resolver.ListResolverEndpointsInput, opts ...request.Option) (*route53resolver.ListResolverEndpointsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "ListResolverEndpoints",
		Input:   input,
		Output:  (*route53resolver.ListResolverEndpointsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.ListResolverEndpointsWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.ListResolverEndpointsOutput), req.Error
}

func (c *Client) ListResolverRuleAssociationsWithContext(ctx context.Context, input *route53resolver.ListResolverRuleAssociationsInput, opts ...request.Option) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "ListResolverRuleAssociations",
		Input:   input,
		Output:  (*route53resolver.ListResolverRuleAssociationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.ListResolverRuleAssociationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.ListResolverRuleAssociationsOutput), req.Error
}

func (c *Client) ListResolverRulesWithContext(ctx context.Context, input *route53resolver.ListResolverRulesInput, opts ...request.Option) (*route53resolver.ListResolverRulesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "ListResolverRules",
		Input:   input,
		Output:  (*route53resolver.ListResolverRulesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.ListResolverRulesWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.ListResolverRulesOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *route53resolver.ListTagsForResourceInput, opts ...request.Option) (*route53resolver.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*route53resolver.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.ListTagsForResourceOutput), req.Error
}

func (c *Client) PutResolverRulePolicyWithContext(ctx context.Context, input *route53resolver.PutResolverRulePolicyInput, opts ...request.Option) (*route53resolver.PutResolverRulePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "PutResolverRulePolicy",
		Input:   input,
		Output:  (*route53resolver.PutResolverRulePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.PutResolverRulePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.PutResolverRulePolicyOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *route53resolver.TagResourceInput, opts ...request.Option) (*route53resolver.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "TagResource",
		Input:   input,
		Output:  (*route53resolver.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *route53resolver.UntagResourceInput, opts ...request.Option) (*route53resolver.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*route53resolver.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.UntagResourceOutput), req.Error
}

func (c *Client) UpdateResolverEndpointWithContext(ctx context.Context, input *route53resolver.UpdateResolverEndpointInput, opts ...request.Option) (*route53resolver.UpdateResolverEndpointOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "UpdateResolverEndpoint",
		Input:   input,
		Output:  (*route53resolver.UpdateResolverEndpointOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.UpdateResolverEndpointWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.UpdateResolverEndpointOutput), req.Error
}

func (c *Client) UpdateResolverRuleWithContext(ctx context.Context, input *route53resolver.UpdateResolverRuleInput, opts ...request.Option) (*route53resolver.UpdateResolverRuleOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "route53resolver",
		Action:  "UpdateResolverRule",
		Input:   input,
		Output:  (*route53resolver.UpdateResolverRuleOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.Route53ResolverAPI.UpdateResolverRuleWithContext(ctx, input, opts...)
	})

	return req.Output.(*route53resolver.UpdateResolverRuleOutput), req.Error
}
