// Code generated by internal/generate/main.go. DO NOT EDIT.

package servicediscoveryctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/servicediscovery"
	"github.com/aws/aws-sdk-go/service/servicediscovery/servicediscoveryiface"
	"github.com/glassechidna/awsctx"
)

type ServiceDiscovery interface {
	CreateHttpNamespaceWithContext(ctx context.Context, input *servicediscovery.CreateHttpNamespaceInput, opts ...request.Option) (*servicediscovery.CreateHttpNamespaceOutput, error)
	CreatePrivateDnsNamespaceWithContext(ctx context.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput, opts ...request.Option) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error)
	CreatePublicDnsNamespaceWithContext(ctx context.Context, input *servicediscovery.CreatePublicDnsNamespaceInput, opts ...request.Option) (*servicediscovery.CreatePublicDnsNamespaceOutput, error)
	CreateServiceWithContext(ctx context.Context, input *servicediscovery.CreateServiceInput, opts ...request.Option) (*servicediscovery.CreateServiceOutput, error)
	DeleteNamespaceWithContext(ctx context.Context, input *servicediscovery.DeleteNamespaceInput, opts ...request.Option) (*servicediscovery.DeleteNamespaceOutput, error)
	DeleteServiceWithContext(ctx context.Context, input *servicediscovery.DeleteServiceInput, opts ...request.Option) (*servicediscovery.DeleteServiceOutput, error)
	DeregisterInstanceWithContext(ctx context.Context, input *servicediscovery.DeregisterInstanceInput, opts ...request.Option) (*servicediscovery.DeregisterInstanceOutput, error)
	DiscoverInstancesWithContext(ctx context.Context, input *servicediscovery.DiscoverInstancesInput, opts ...request.Option) (*servicediscovery.DiscoverInstancesOutput, error)
	GetInstanceWithContext(ctx context.Context, input *servicediscovery.GetInstanceInput, opts ...request.Option) (*servicediscovery.GetInstanceOutput, error)
	GetInstancesHealthStatusWithContext(ctx context.Context, input *servicediscovery.GetInstancesHealthStatusInput, opts ...request.Option) (*servicediscovery.GetInstancesHealthStatusOutput, error)
	GetNamespaceWithContext(ctx context.Context, input *servicediscovery.GetNamespaceInput, opts ...request.Option) (*servicediscovery.GetNamespaceOutput, error)
	GetOperationWithContext(ctx context.Context, input *servicediscovery.GetOperationInput, opts ...request.Option) (*servicediscovery.GetOperationOutput, error)
	GetServiceWithContext(ctx context.Context, input *servicediscovery.GetServiceInput, opts ...request.Option) (*servicediscovery.GetServiceOutput, error)
	ListInstancesWithContext(ctx context.Context, input *servicediscovery.ListInstancesInput, opts ...request.Option) (*servicediscovery.ListInstancesOutput, error)
	ListNamespacesWithContext(ctx context.Context, input *servicediscovery.ListNamespacesInput, opts ...request.Option) (*servicediscovery.ListNamespacesOutput, error)
	ListOperationsWithContext(ctx context.Context, input *servicediscovery.ListOperationsInput, opts ...request.Option) (*servicediscovery.ListOperationsOutput, error)
	ListServicesWithContext(ctx context.Context, input *servicediscovery.ListServicesInput, opts ...request.Option) (*servicediscovery.ListServicesOutput, error)
	RegisterInstanceWithContext(ctx context.Context, input *servicediscovery.RegisterInstanceInput, opts ...request.Option) (*servicediscovery.RegisterInstanceOutput, error)
	UpdateInstanceCustomHealthStatusWithContext(ctx context.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput, opts ...request.Option) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error)
	UpdateServiceWithContext(ctx context.Context, input *servicediscovery.UpdateServiceInput, opts ...request.Option) (*servicediscovery.UpdateServiceOutput, error)
}

type Client struct {
	servicediscoveryiface.ServiceDiscoveryAPI
	Contexter awsctx.Contexter
}

func New(base servicediscoveryiface.ServiceDiscoveryAPI, ctxer awsctx.Contexter) ServiceDiscovery {
	return &Client{
		ServiceDiscoveryAPI: base,
		Contexter: ctxer,
	}
}

var _ ServiceDiscovery = (*servicediscovery.ServiceDiscovery)(nil)
var _ ServiceDiscovery = (*Client)(nil)

func (c *Client) CreateHttpNamespaceWithContext(ctx context.Context, input *servicediscovery.CreateHttpNamespaceInput, opts ...request.Option) (*servicediscovery.CreateHttpNamespaceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "CreateHttpNamespaceWithContext",
		Input:   input,
		Output:  (*servicediscovery.CreateHttpNamespaceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.CreateHttpNamespaceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.CreateHttpNamespaceOutput), req.Error
}

func (c *Client) CreatePrivateDnsNamespaceWithContext(ctx context.Context, input *servicediscovery.CreatePrivateDnsNamespaceInput, opts ...request.Option) (*servicediscovery.CreatePrivateDnsNamespaceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "CreatePrivateDnsNamespaceWithContext",
		Input:   input,
		Output:  (*servicediscovery.CreatePrivateDnsNamespaceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.CreatePrivateDnsNamespaceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.CreatePrivateDnsNamespaceOutput), req.Error
}

func (c *Client) CreatePublicDnsNamespaceWithContext(ctx context.Context, input *servicediscovery.CreatePublicDnsNamespaceInput, opts ...request.Option) (*servicediscovery.CreatePublicDnsNamespaceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "CreatePublicDnsNamespaceWithContext",
		Input:   input,
		Output:  (*servicediscovery.CreatePublicDnsNamespaceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.CreatePublicDnsNamespaceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.CreatePublicDnsNamespaceOutput), req.Error
}

func (c *Client) CreateServiceWithContext(ctx context.Context, input *servicediscovery.CreateServiceInput, opts ...request.Option) (*servicediscovery.CreateServiceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "CreateServiceWithContext",
		Input:   input,
		Output:  (*servicediscovery.CreateServiceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.CreateServiceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.CreateServiceOutput), req.Error
}

func (c *Client) DeleteNamespaceWithContext(ctx context.Context, input *servicediscovery.DeleteNamespaceInput, opts ...request.Option) (*servicediscovery.DeleteNamespaceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "DeleteNamespaceWithContext",
		Input:   input,
		Output:  (*servicediscovery.DeleteNamespaceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.DeleteNamespaceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.DeleteNamespaceOutput), req.Error
}

func (c *Client) DeleteServiceWithContext(ctx context.Context, input *servicediscovery.DeleteServiceInput, opts ...request.Option) (*servicediscovery.DeleteServiceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "DeleteServiceWithContext",
		Input:   input,
		Output:  (*servicediscovery.DeleteServiceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.DeleteServiceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.DeleteServiceOutput), req.Error
}

func (c *Client) DeregisterInstanceWithContext(ctx context.Context, input *servicediscovery.DeregisterInstanceInput, opts ...request.Option) (*servicediscovery.DeregisterInstanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "DeregisterInstanceWithContext",
		Input:   input,
		Output:  (*servicediscovery.DeregisterInstanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.DeregisterInstanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.DeregisterInstanceOutput), req.Error
}

func (c *Client) DiscoverInstancesWithContext(ctx context.Context, input *servicediscovery.DiscoverInstancesInput, opts ...request.Option) (*servicediscovery.DiscoverInstancesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "DiscoverInstancesWithContext",
		Input:   input,
		Output:  (*servicediscovery.DiscoverInstancesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.DiscoverInstancesWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.DiscoverInstancesOutput), req.Error
}

func (c *Client) GetInstanceWithContext(ctx context.Context, input *servicediscovery.GetInstanceInput, opts ...request.Option) (*servicediscovery.GetInstanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "GetInstanceWithContext",
		Input:   input,
		Output:  (*servicediscovery.GetInstanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.GetInstanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.GetInstanceOutput), req.Error
}

func (c *Client) GetInstancesHealthStatusWithContext(ctx context.Context, input *servicediscovery.GetInstancesHealthStatusInput, opts ...request.Option) (*servicediscovery.GetInstancesHealthStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "GetInstancesHealthStatusWithContext",
		Input:   input,
		Output:  (*servicediscovery.GetInstancesHealthStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.GetInstancesHealthStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.GetInstancesHealthStatusOutput), req.Error
}

func (c *Client) GetNamespaceWithContext(ctx context.Context, input *servicediscovery.GetNamespaceInput, opts ...request.Option) (*servicediscovery.GetNamespaceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "GetNamespaceWithContext",
		Input:   input,
		Output:  (*servicediscovery.GetNamespaceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.GetNamespaceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.GetNamespaceOutput), req.Error
}

func (c *Client) GetOperationWithContext(ctx context.Context, input *servicediscovery.GetOperationInput, opts ...request.Option) (*servicediscovery.GetOperationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "GetOperationWithContext",
		Input:   input,
		Output:  (*servicediscovery.GetOperationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.GetOperationWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.GetOperationOutput), req.Error
}

func (c *Client) GetServiceWithContext(ctx context.Context, input *servicediscovery.GetServiceInput, opts ...request.Option) (*servicediscovery.GetServiceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "GetServiceWithContext",
		Input:   input,
		Output:  (*servicediscovery.GetServiceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.GetServiceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.GetServiceOutput), req.Error
}

func (c *Client) ListInstancesWithContext(ctx context.Context, input *servicediscovery.ListInstancesInput, opts ...request.Option) (*servicediscovery.ListInstancesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "ListInstancesWithContext",
		Input:   input,
		Output:  (*servicediscovery.ListInstancesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.ListInstancesWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.ListInstancesOutput), req.Error
}

func (c *Client) ListNamespacesWithContext(ctx context.Context, input *servicediscovery.ListNamespacesInput, opts ...request.Option) (*servicediscovery.ListNamespacesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "ListNamespacesWithContext",
		Input:   input,
		Output:  (*servicediscovery.ListNamespacesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.ListNamespacesWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.ListNamespacesOutput), req.Error
}

func (c *Client) ListOperationsWithContext(ctx context.Context, input *servicediscovery.ListOperationsInput, opts ...request.Option) (*servicediscovery.ListOperationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "ListOperationsWithContext",
		Input:   input,
		Output:  (*servicediscovery.ListOperationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.ListOperationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.ListOperationsOutput), req.Error
}

func (c *Client) ListServicesWithContext(ctx context.Context, input *servicediscovery.ListServicesInput, opts ...request.Option) (*servicediscovery.ListServicesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "ListServicesWithContext",
		Input:   input,
		Output:  (*servicediscovery.ListServicesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.ListServicesWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.ListServicesOutput), req.Error
}

func (c *Client) RegisterInstanceWithContext(ctx context.Context, input *servicediscovery.RegisterInstanceInput, opts ...request.Option) (*servicediscovery.RegisterInstanceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "RegisterInstanceWithContext",
		Input:   input,
		Output:  (*servicediscovery.RegisterInstanceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.RegisterInstanceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.RegisterInstanceOutput), req.Error
}

func (c *Client) UpdateInstanceCustomHealthStatusWithContext(ctx context.Context, input *servicediscovery.UpdateInstanceCustomHealthStatusInput, opts ...request.Option) (*servicediscovery.UpdateInstanceCustomHealthStatusOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "UpdateInstanceCustomHealthStatusWithContext",
		Input:   input,
		Output:  (*servicediscovery.UpdateInstanceCustomHealthStatusOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.UpdateInstanceCustomHealthStatusWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.UpdateInstanceCustomHealthStatusOutput), req.Error
}

func (c *Client) UpdateServiceWithContext(ctx context.Context, input *servicediscovery.UpdateServiceInput, opts ...request.Option) (*servicediscovery.UpdateServiceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "servicediscovery",
		Action:  "UpdateServiceWithContext",
		Input:   input,
		Output:  (*servicediscovery.UpdateServiceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServiceDiscoveryAPI.UpdateServiceWithContext(ctx, input, opts...)
	})

	return req.Output.(*servicediscovery.UpdateServiceOutput), req.Error
}
