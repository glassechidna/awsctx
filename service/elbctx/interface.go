// Code generated by internal/generate/main.go. DO NOT EDIT.

package elbctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/glassechidna/awsctx"
)

type ELB interface {
	AddTagsWithContext(ctx context.Context, input *elb.AddTagsInput, opts ...request.Option) (*elb.AddTagsOutput, error)
	ApplySecurityGroupsToLoadBalancerWithContext(ctx context.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput, opts ...request.Option) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error)
	AttachLoadBalancerToSubnetsWithContext(ctx context.Context, input *elb.AttachLoadBalancerToSubnetsInput, opts ...request.Option) (*elb.AttachLoadBalancerToSubnetsOutput, error)
	ConfigureHealthCheckWithContext(ctx context.Context, input *elb.ConfigureHealthCheckInput, opts ...request.Option) (*elb.ConfigureHealthCheckOutput, error)
	CreateAppCookieStickinessPolicyWithContext(ctx context.Context, input *elb.CreateAppCookieStickinessPolicyInput, opts ...request.Option) (*elb.CreateAppCookieStickinessPolicyOutput, error)
	CreateLBCookieStickinessPolicyWithContext(ctx context.Context, input *elb.CreateLBCookieStickinessPolicyInput, opts ...request.Option) (*elb.CreateLBCookieStickinessPolicyOutput, error)
	CreateLoadBalancerWithContext(ctx context.Context, input *elb.CreateLoadBalancerInput, opts ...request.Option) (*elb.CreateLoadBalancerOutput, error)
	CreateLoadBalancerListenersWithContext(ctx context.Context, input *elb.CreateLoadBalancerListenersInput, opts ...request.Option) (*elb.CreateLoadBalancerListenersOutput, error)
	CreateLoadBalancerPolicyWithContext(ctx context.Context, input *elb.CreateLoadBalancerPolicyInput, opts ...request.Option) (*elb.CreateLoadBalancerPolicyOutput, error)
	DeleteLoadBalancerWithContext(ctx context.Context, input *elb.DeleteLoadBalancerInput, opts ...request.Option) (*elb.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerListenersWithContext(ctx context.Context, input *elb.DeleteLoadBalancerListenersInput, opts ...request.Option) (*elb.DeleteLoadBalancerListenersOutput, error)
	DeleteLoadBalancerPolicyWithContext(ctx context.Context, input *elb.DeleteLoadBalancerPolicyInput, opts ...request.Option) (*elb.DeleteLoadBalancerPolicyOutput, error)
	DeregisterInstancesFromLoadBalancerWithContext(ctx context.Context, input *elb.DeregisterInstancesFromLoadBalancerInput, opts ...request.Option) (*elb.DeregisterInstancesFromLoadBalancerOutput, error)
	DescribeAccountLimitsWithContext(ctx context.Context, input *elb.DescribeAccountLimitsInput, opts ...request.Option) (*elb.DescribeAccountLimitsOutput, error)
	DescribeInstanceHealthWithContext(ctx context.Context, input *elb.DescribeInstanceHealthInput, opts ...request.Option) (*elb.DescribeInstanceHealthOutput, error)
	DescribeLoadBalancerAttributesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerAttributesInput, opts ...request.Option) (*elb.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerPoliciesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerPoliciesInput, opts ...request.Option) (*elb.DescribeLoadBalancerPoliciesOutput, error)
	DescribeLoadBalancerPolicyTypesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerPolicyTypesInput, opts ...request.Option) (*elb.DescribeLoadBalancerPolicyTypesOutput, error)
	DescribeLoadBalancersWithContext(ctx context.Context, input *elb.DescribeLoadBalancersInput, opts ...request.Option) (*elb.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersPagesWithContext(ctx context.Context, input *elb.DescribeLoadBalancersInput, cb func(*elb.DescribeLoadBalancersOutput, bool) bool, opts ...request.Option) error
	DescribeTagsWithContext(ctx context.Context, input *elb.DescribeTagsInput, opts ...request.Option) (*elb.DescribeTagsOutput, error)
	DetachLoadBalancerFromSubnetsWithContext(ctx context.Context, input *elb.DetachLoadBalancerFromSubnetsInput, opts ...request.Option) (*elb.DetachLoadBalancerFromSubnetsOutput, error)
	DisableAvailabilityZonesForLoadBalancerWithContext(ctx context.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput, opts ...request.Option) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error)
	EnableAvailabilityZonesForLoadBalancerWithContext(ctx context.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput, opts ...request.Option) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error)
	ModifyLoadBalancerAttributesWithContext(ctx context.Context, input *elb.ModifyLoadBalancerAttributesInput, opts ...request.Option) (*elb.ModifyLoadBalancerAttributesOutput, error)
	RegisterInstancesWithLoadBalancerWithContext(ctx context.Context, input *elb.RegisterInstancesWithLoadBalancerInput, opts ...request.Option) (*elb.RegisterInstancesWithLoadBalancerOutput, error)
	RemoveTagsWithContext(ctx context.Context, input *elb.RemoveTagsInput, opts ...request.Option) (*elb.RemoveTagsOutput, error)
	SetLoadBalancerListenerSSLCertificateWithContext(ctx context.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput, opts ...request.Option) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error)
	SetLoadBalancerPoliciesForBackendServerWithContext(ctx context.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput, opts ...request.Option) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error)
	SetLoadBalancerPoliciesOfListenerWithContext(ctx context.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput, opts ...request.Option) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error)
}

type Client struct {
	elbiface.ELBAPI
	Contexter awsctx.Contexter
}

func New(base elbiface.ELBAPI, ctxer awsctx.Contexter) ELB {
	return &Client{
		ELBAPI: base,
		Contexter: ctxer,
	}
}

var _ ELB = (*elb.ELB)(nil)
var _ ELB = (*Client)(nil)

func (c *Client) AddTagsWithContext(ctx context.Context, input *elb.AddTagsInput, opts ...request.Option) (*elb.AddTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "AddTags",
		Input:   input,
		Output:  (*elb.AddTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.AddTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.AddTagsOutput), req.Error
}

func (c *Client) ApplySecurityGroupsToLoadBalancerWithContext(ctx context.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput, opts ...request.Option) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "ApplySecurityGroupsToLoadBalancer",
		Input:   input,
		Output:  (*elb.ApplySecurityGroupsToLoadBalancerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.ApplySecurityGroupsToLoadBalancerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.ApplySecurityGroupsToLoadBalancerOutput), req.Error
}

func (c *Client) AttachLoadBalancerToSubnetsWithContext(ctx context.Context, input *elb.AttachLoadBalancerToSubnetsInput, opts ...request.Option) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "AttachLoadBalancerToSubnets",
		Input:   input,
		Output:  (*elb.AttachLoadBalancerToSubnetsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.AttachLoadBalancerToSubnetsWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.AttachLoadBalancerToSubnetsOutput), req.Error
}

func (c *Client) ConfigureHealthCheckWithContext(ctx context.Context, input *elb.ConfigureHealthCheckInput, opts ...request.Option) (*elb.ConfigureHealthCheckOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "ConfigureHealthCheck",
		Input:   input,
		Output:  (*elb.ConfigureHealthCheckOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.ConfigureHealthCheckWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.ConfigureHealthCheckOutput), req.Error
}

func (c *Client) CreateAppCookieStickinessPolicyWithContext(ctx context.Context, input *elb.CreateAppCookieStickinessPolicyInput, opts ...request.Option) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "CreateAppCookieStickinessPolicy",
		Input:   input,
		Output:  (*elb.CreateAppCookieStickinessPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.CreateAppCookieStickinessPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.CreateAppCookieStickinessPolicyOutput), req.Error
}

func (c *Client) CreateLBCookieStickinessPolicyWithContext(ctx context.Context, input *elb.CreateLBCookieStickinessPolicyInput, opts ...request.Option) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "CreateLBCookieStickinessPolicy",
		Input:   input,
		Output:  (*elb.CreateLBCookieStickinessPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.CreateLBCookieStickinessPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.CreateLBCookieStickinessPolicyOutput), req.Error
}

func (c *Client) CreateLoadBalancerWithContext(ctx context.Context, input *elb.CreateLoadBalancerInput, opts ...request.Option) (*elb.CreateLoadBalancerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "CreateLoadBalancer",
		Input:   input,
		Output:  (*elb.CreateLoadBalancerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.CreateLoadBalancerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.CreateLoadBalancerOutput), req.Error
}

func (c *Client) CreateLoadBalancerListenersWithContext(ctx context.Context, input *elb.CreateLoadBalancerListenersInput, opts ...request.Option) (*elb.CreateLoadBalancerListenersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "CreateLoadBalancerListeners",
		Input:   input,
		Output:  (*elb.CreateLoadBalancerListenersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.CreateLoadBalancerListenersWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.CreateLoadBalancerListenersOutput), req.Error
}

func (c *Client) CreateLoadBalancerPolicyWithContext(ctx context.Context, input *elb.CreateLoadBalancerPolicyInput, opts ...request.Option) (*elb.CreateLoadBalancerPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "CreateLoadBalancerPolicy",
		Input:   input,
		Output:  (*elb.CreateLoadBalancerPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.CreateLoadBalancerPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.CreateLoadBalancerPolicyOutput), req.Error
}

func (c *Client) DeleteLoadBalancerWithContext(ctx context.Context, input *elb.DeleteLoadBalancerInput, opts ...request.Option) (*elb.DeleteLoadBalancerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DeleteLoadBalancer",
		Input:   input,
		Output:  (*elb.DeleteLoadBalancerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DeleteLoadBalancerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DeleteLoadBalancerOutput), req.Error
}

func (c *Client) DeleteLoadBalancerListenersWithContext(ctx context.Context, input *elb.DeleteLoadBalancerListenersInput, opts ...request.Option) (*elb.DeleteLoadBalancerListenersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DeleteLoadBalancerListeners",
		Input:   input,
		Output:  (*elb.DeleteLoadBalancerListenersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DeleteLoadBalancerListenersWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DeleteLoadBalancerListenersOutput), req.Error
}

func (c *Client) DeleteLoadBalancerPolicyWithContext(ctx context.Context, input *elb.DeleteLoadBalancerPolicyInput, opts ...request.Option) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DeleteLoadBalancerPolicy",
		Input:   input,
		Output:  (*elb.DeleteLoadBalancerPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DeleteLoadBalancerPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DeleteLoadBalancerPolicyOutput), req.Error
}

func (c *Client) DeregisterInstancesFromLoadBalancerWithContext(ctx context.Context, input *elb.DeregisterInstancesFromLoadBalancerInput, opts ...request.Option) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DeregisterInstancesFromLoadBalancer",
		Input:   input,
		Output:  (*elb.DeregisterInstancesFromLoadBalancerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DeregisterInstancesFromLoadBalancerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DeregisterInstancesFromLoadBalancerOutput), req.Error
}

func (c *Client) DescribeAccountLimitsWithContext(ctx context.Context, input *elb.DescribeAccountLimitsInput, opts ...request.Option) (*elb.DescribeAccountLimitsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DescribeAccountLimits",
		Input:   input,
		Output:  (*elb.DescribeAccountLimitsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DescribeAccountLimitsWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DescribeAccountLimitsOutput), req.Error
}

func (c *Client) DescribeInstanceHealthWithContext(ctx context.Context, input *elb.DescribeInstanceHealthInput, opts ...request.Option) (*elb.DescribeInstanceHealthOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DescribeInstanceHealth",
		Input:   input,
		Output:  (*elb.DescribeInstanceHealthOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DescribeInstanceHealthWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DescribeInstanceHealthOutput), req.Error
}

func (c *Client) DescribeLoadBalancerAttributesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerAttributesInput, opts ...request.Option) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DescribeLoadBalancerAttributes",
		Input:   input,
		Output:  (*elb.DescribeLoadBalancerAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DescribeLoadBalancerAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DescribeLoadBalancerAttributesOutput), req.Error
}

func (c *Client) DescribeLoadBalancerPoliciesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerPoliciesInput, opts ...request.Option) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DescribeLoadBalancerPolicies",
		Input:   input,
		Output:  (*elb.DescribeLoadBalancerPoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DescribeLoadBalancerPoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DescribeLoadBalancerPoliciesOutput), req.Error
}

func (c *Client) DescribeLoadBalancerPolicyTypesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerPolicyTypesInput, opts ...request.Option) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DescribeLoadBalancerPolicyTypes",
		Input:   input,
		Output:  (*elb.DescribeLoadBalancerPolicyTypesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DescribeLoadBalancerPolicyTypesWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DescribeLoadBalancerPolicyTypesOutput), req.Error
}

func (c *Client) DescribeLoadBalancersWithContext(ctx context.Context, input *elb.DescribeLoadBalancersInput, opts ...request.Option) (*elb.DescribeLoadBalancersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DescribeLoadBalancers",
		Input:   input,
		Output:  (*elb.DescribeLoadBalancersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DescribeLoadBalancersWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DescribeLoadBalancersOutput), req.Error
}

func (c *Client) DescribeLoadBalancersPagesWithContext(ctx context.Context, input *elb.DescribeLoadBalancersInput, cb func(*elb.DescribeLoadBalancersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DescribeLoadBalancers",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ELBAPI.DescribeLoadBalancersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeTagsWithContext(ctx context.Context, input *elb.DescribeTagsInput, opts ...request.Option) (*elb.DescribeTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DescribeTags",
		Input:   input,
		Output:  (*elb.DescribeTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DescribeTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DescribeTagsOutput), req.Error
}

func (c *Client) DetachLoadBalancerFromSubnetsWithContext(ctx context.Context, input *elb.DetachLoadBalancerFromSubnetsInput, opts ...request.Option) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DetachLoadBalancerFromSubnets",
		Input:   input,
		Output:  (*elb.DetachLoadBalancerFromSubnetsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DetachLoadBalancerFromSubnetsWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DetachLoadBalancerFromSubnetsOutput), req.Error
}

func (c *Client) DisableAvailabilityZonesForLoadBalancerWithContext(ctx context.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput, opts ...request.Option) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "DisableAvailabilityZonesForLoadBalancer",
		Input:   input,
		Output:  (*elb.DisableAvailabilityZonesForLoadBalancerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.DisableAvailabilityZonesForLoadBalancerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.DisableAvailabilityZonesForLoadBalancerOutput), req.Error
}

func (c *Client) EnableAvailabilityZonesForLoadBalancerWithContext(ctx context.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput, opts ...request.Option) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "EnableAvailabilityZonesForLoadBalancer",
		Input:   input,
		Output:  (*elb.EnableAvailabilityZonesForLoadBalancerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.EnableAvailabilityZonesForLoadBalancerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.EnableAvailabilityZonesForLoadBalancerOutput), req.Error
}

func (c *Client) ModifyLoadBalancerAttributesWithContext(ctx context.Context, input *elb.ModifyLoadBalancerAttributesInput, opts ...request.Option) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "ModifyLoadBalancerAttributes",
		Input:   input,
		Output:  (*elb.ModifyLoadBalancerAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.ModifyLoadBalancerAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.ModifyLoadBalancerAttributesOutput), req.Error
}

func (c *Client) RegisterInstancesWithLoadBalancerWithContext(ctx context.Context, input *elb.RegisterInstancesWithLoadBalancerInput, opts ...request.Option) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "RegisterInstancesWithLoadBalancer",
		Input:   input,
		Output:  (*elb.RegisterInstancesWithLoadBalancerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.RegisterInstancesWithLoadBalancerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.RegisterInstancesWithLoadBalancerOutput), req.Error
}

func (c *Client) RemoveTagsWithContext(ctx context.Context, input *elb.RemoveTagsInput, opts ...request.Option) (*elb.RemoveTagsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "RemoveTags",
		Input:   input,
		Output:  (*elb.RemoveTagsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.RemoveTagsWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.RemoveTagsOutput), req.Error
}

func (c *Client) SetLoadBalancerListenerSSLCertificateWithContext(ctx context.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput, opts ...request.Option) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "SetLoadBalancerListenerSSLCertificate",
		Input:   input,
		Output:  (*elb.SetLoadBalancerListenerSSLCertificateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.SetLoadBalancerListenerSSLCertificateWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.SetLoadBalancerListenerSSLCertificateOutput), req.Error
}

func (c *Client) SetLoadBalancerPoliciesForBackendServerWithContext(ctx context.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput, opts ...request.Option) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "SetLoadBalancerPoliciesForBackendServer",
		Input:   input,
		Output:  (*elb.SetLoadBalancerPoliciesForBackendServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.SetLoadBalancerPoliciesForBackendServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.SetLoadBalancerPoliciesForBackendServerOutput), req.Error
}

func (c *Client) SetLoadBalancerPoliciesOfListenerWithContext(ctx context.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput, opts ...request.Option) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "elb",
		Action:  "SetLoadBalancerPoliciesOfListener",
		Input:   input,
		Output:  (*elb.SetLoadBalancerPoliciesOfListenerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ELBAPI.SetLoadBalancerPoliciesOfListenerWithContext(ctx, input, opts...)
	})

	return req.Output.(*elb.SetLoadBalancerPoliciesOfListenerOutput), req.Error
}
