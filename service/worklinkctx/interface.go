// Code generated by internal/generate/main.go. DO NOT EDIT.

package worklinkctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/worklink/worklinkiface"
	"github.com/glassechidna/awsctx"
)

type WorkLink interface {
	AssociateDomainWithContext(ctx context.Context, input *worklink.AssociateDomainInput, opts ...request.Option) (*worklink.AssociateDomainOutput, error)
	AssociateWebsiteAuthorizationProviderWithContext(ctx context.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput, opts ...request.Option) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error)
	AssociateWebsiteCertificateAuthorityWithContext(ctx context.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput, opts ...request.Option) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error)
	CreateFleetWithContext(ctx context.Context, input *worklink.CreateFleetInput, opts ...request.Option) (*worklink.CreateFleetOutput, error)
	DeleteFleetWithContext(ctx context.Context, input *worklink.DeleteFleetInput, opts ...request.Option) (*worklink.DeleteFleetOutput, error)
	DescribeAuditStreamConfigurationWithContext(ctx context.Context, input *worklink.DescribeAuditStreamConfigurationInput, opts ...request.Option) (*worklink.DescribeAuditStreamConfigurationOutput, error)
	DescribeCompanyNetworkConfigurationWithContext(ctx context.Context, input *worklink.DescribeCompanyNetworkConfigurationInput, opts ...request.Option) (*worklink.DescribeCompanyNetworkConfigurationOutput, error)
	DescribeDeviceWithContext(ctx context.Context, input *worklink.DescribeDeviceInput, opts ...request.Option) (*worklink.DescribeDeviceOutput, error)
	DescribeDevicePolicyConfigurationWithContext(ctx context.Context, input *worklink.DescribeDevicePolicyConfigurationInput, opts ...request.Option) (*worklink.DescribeDevicePolicyConfigurationOutput, error)
	DescribeDomainWithContext(ctx context.Context, input *worklink.DescribeDomainInput, opts ...request.Option) (*worklink.DescribeDomainOutput, error)
	DescribeFleetMetadataWithContext(ctx context.Context, input *worklink.DescribeFleetMetadataInput, opts ...request.Option) (*worklink.DescribeFleetMetadataOutput, error)
	DescribeIdentityProviderConfigurationWithContext(ctx context.Context, input *worklink.DescribeIdentityProviderConfigurationInput, opts ...request.Option) (*worklink.DescribeIdentityProviderConfigurationOutput, error)
	DescribeWebsiteCertificateAuthorityWithContext(ctx context.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput, opts ...request.Option) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error)
	DisassociateDomainWithContext(ctx context.Context, input *worklink.DisassociateDomainInput, opts ...request.Option) (*worklink.DisassociateDomainOutput, error)
	DisassociateWebsiteAuthorizationProviderWithContext(ctx context.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput, opts ...request.Option) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error)
	DisassociateWebsiteCertificateAuthorityWithContext(ctx context.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput, opts ...request.Option) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error)
	ListDevicesWithContext(ctx context.Context, input *worklink.ListDevicesInput, opts ...request.Option) (*worklink.ListDevicesOutput, error)
	ListDomainsWithContext(ctx context.Context, input *worklink.ListDomainsInput, opts ...request.Option) (*worklink.ListDomainsOutput, error)
	ListFleetsWithContext(ctx context.Context, input *worklink.ListFleetsInput, opts ...request.Option) (*worklink.ListFleetsOutput, error)
	ListWebsiteAuthorizationProvidersWithContext(ctx context.Context, input *worklink.ListWebsiteAuthorizationProvidersInput, opts ...request.Option) (*worklink.ListWebsiteAuthorizationProvidersOutput, error)
	ListWebsiteCertificateAuthoritiesWithContext(ctx context.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput, opts ...request.Option) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error)
	RestoreDomainAccessWithContext(ctx context.Context, input *worklink.RestoreDomainAccessInput, opts ...request.Option) (*worklink.RestoreDomainAccessOutput, error)
	RevokeDomainAccessWithContext(ctx context.Context, input *worklink.RevokeDomainAccessInput, opts ...request.Option) (*worklink.RevokeDomainAccessOutput, error)
	SignOutUserWithContext(ctx context.Context, input *worklink.SignOutUserInput, opts ...request.Option) (*worklink.SignOutUserOutput, error)
	UpdateAuditStreamConfigurationWithContext(ctx context.Context, input *worklink.UpdateAuditStreamConfigurationInput, opts ...request.Option) (*worklink.UpdateAuditStreamConfigurationOutput, error)
	UpdateCompanyNetworkConfigurationWithContext(ctx context.Context, input *worklink.UpdateCompanyNetworkConfigurationInput, opts ...request.Option) (*worklink.UpdateCompanyNetworkConfigurationOutput, error)
	UpdateDevicePolicyConfigurationWithContext(ctx context.Context, input *worklink.UpdateDevicePolicyConfigurationInput, opts ...request.Option) (*worklink.UpdateDevicePolicyConfigurationOutput, error)
	UpdateDomainMetadataWithContext(ctx context.Context, input *worklink.UpdateDomainMetadataInput, opts ...request.Option) (*worklink.UpdateDomainMetadataOutput, error)
	UpdateFleetMetadataWithContext(ctx context.Context, input *worklink.UpdateFleetMetadataInput, opts ...request.Option) (*worklink.UpdateFleetMetadataOutput, error)
	UpdateIdentityProviderConfigurationWithContext(ctx context.Context, input *worklink.UpdateIdentityProviderConfigurationInput, opts ...request.Option) (*worklink.UpdateIdentityProviderConfigurationOutput, error)
}

type Client struct {
	worklinkiface.WorkLinkAPI
	Contexter awsctx.Contexter
}

func New(base worklinkiface.WorkLinkAPI, ctxer awsctx.Contexter) WorkLink {
	return &Client{
		WorkLinkAPI: base,
		Contexter: ctxer,
	}
}

var _ WorkLink = (*worklink.WorkLink)(nil)
var _ WorkLink = (*Client)(nil)

func (c *Client) AssociateDomainWithContext(ctx context.Context, input *worklink.AssociateDomainInput, opts ...request.Option) (*worklink.AssociateDomainOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "AssociateDomainWithContext",
		Input:   input,
		Output:  (*worklink.AssociateDomainOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.AssociateDomainWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.AssociateDomainOutput), req.Error
}

func (c *Client) AssociateWebsiteAuthorizationProviderWithContext(ctx context.Context, input *worklink.AssociateWebsiteAuthorizationProviderInput, opts ...request.Option) (*worklink.AssociateWebsiteAuthorizationProviderOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "AssociateWebsiteAuthorizationProviderWithContext",
		Input:   input,
		Output:  (*worklink.AssociateWebsiteAuthorizationProviderOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.AssociateWebsiteAuthorizationProviderWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.AssociateWebsiteAuthorizationProviderOutput), req.Error
}

func (c *Client) AssociateWebsiteCertificateAuthorityWithContext(ctx context.Context, input *worklink.AssociateWebsiteCertificateAuthorityInput, opts ...request.Option) (*worklink.AssociateWebsiteCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "AssociateWebsiteCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*worklink.AssociateWebsiteCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.AssociateWebsiteCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.AssociateWebsiteCertificateAuthorityOutput), req.Error
}

func (c *Client) CreateFleetWithContext(ctx context.Context, input *worklink.CreateFleetInput, opts ...request.Option) (*worklink.CreateFleetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "CreateFleetWithContext",
		Input:   input,
		Output:  (*worklink.CreateFleetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.CreateFleetWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.CreateFleetOutput), req.Error
}

func (c *Client) DeleteFleetWithContext(ctx context.Context, input *worklink.DeleteFleetInput, opts ...request.Option) (*worklink.DeleteFleetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DeleteFleetWithContext",
		Input:   input,
		Output:  (*worklink.DeleteFleetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DeleteFleetWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DeleteFleetOutput), req.Error
}

func (c *Client) DescribeAuditStreamConfigurationWithContext(ctx context.Context, input *worklink.DescribeAuditStreamConfigurationInput, opts ...request.Option) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DescribeAuditStreamConfigurationWithContext",
		Input:   input,
		Output:  (*worklink.DescribeAuditStreamConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DescribeAuditStreamConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DescribeAuditStreamConfigurationOutput), req.Error
}

func (c *Client) DescribeCompanyNetworkConfigurationWithContext(ctx context.Context, input *worklink.DescribeCompanyNetworkConfigurationInput, opts ...request.Option) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DescribeCompanyNetworkConfigurationWithContext",
		Input:   input,
		Output:  (*worklink.DescribeCompanyNetworkConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DescribeCompanyNetworkConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DescribeCompanyNetworkConfigurationOutput), req.Error
}

func (c *Client) DescribeDeviceWithContext(ctx context.Context, input *worklink.DescribeDeviceInput, opts ...request.Option) (*worklink.DescribeDeviceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DescribeDeviceWithContext",
		Input:   input,
		Output:  (*worklink.DescribeDeviceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DescribeDeviceWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DescribeDeviceOutput), req.Error
}

func (c *Client) DescribeDevicePolicyConfigurationWithContext(ctx context.Context, input *worklink.DescribeDevicePolicyConfigurationInput, opts ...request.Option) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DescribeDevicePolicyConfigurationWithContext",
		Input:   input,
		Output:  (*worklink.DescribeDevicePolicyConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DescribeDevicePolicyConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DescribeDevicePolicyConfigurationOutput), req.Error
}

func (c *Client) DescribeDomainWithContext(ctx context.Context, input *worklink.DescribeDomainInput, opts ...request.Option) (*worklink.DescribeDomainOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DescribeDomainWithContext",
		Input:   input,
		Output:  (*worklink.DescribeDomainOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DescribeDomainWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DescribeDomainOutput), req.Error
}

func (c *Client) DescribeFleetMetadataWithContext(ctx context.Context, input *worklink.DescribeFleetMetadataInput, opts ...request.Option) (*worklink.DescribeFleetMetadataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DescribeFleetMetadataWithContext",
		Input:   input,
		Output:  (*worklink.DescribeFleetMetadataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DescribeFleetMetadataWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DescribeFleetMetadataOutput), req.Error
}

func (c *Client) DescribeIdentityProviderConfigurationWithContext(ctx context.Context, input *worklink.DescribeIdentityProviderConfigurationInput, opts ...request.Option) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DescribeIdentityProviderConfigurationWithContext",
		Input:   input,
		Output:  (*worklink.DescribeIdentityProviderConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DescribeIdentityProviderConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DescribeIdentityProviderConfigurationOutput), req.Error
}

func (c *Client) DescribeWebsiteCertificateAuthorityWithContext(ctx context.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput, opts ...request.Option) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DescribeWebsiteCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*worklink.DescribeWebsiteCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DescribeWebsiteCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DescribeWebsiteCertificateAuthorityOutput), req.Error
}

func (c *Client) DisassociateDomainWithContext(ctx context.Context, input *worklink.DisassociateDomainInput, opts ...request.Option) (*worklink.DisassociateDomainOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DisassociateDomainWithContext",
		Input:   input,
		Output:  (*worklink.DisassociateDomainOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DisassociateDomainWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DisassociateDomainOutput), req.Error
}

func (c *Client) DisassociateWebsiteAuthorizationProviderWithContext(ctx context.Context, input *worklink.DisassociateWebsiteAuthorizationProviderInput, opts ...request.Option) (*worklink.DisassociateWebsiteAuthorizationProviderOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DisassociateWebsiteAuthorizationProviderWithContext",
		Input:   input,
		Output:  (*worklink.DisassociateWebsiteAuthorizationProviderOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DisassociateWebsiteAuthorizationProviderWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DisassociateWebsiteAuthorizationProviderOutput), req.Error
}

func (c *Client) DisassociateWebsiteCertificateAuthorityWithContext(ctx context.Context, input *worklink.DisassociateWebsiteCertificateAuthorityInput, opts ...request.Option) (*worklink.DisassociateWebsiteCertificateAuthorityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "DisassociateWebsiteCertificateAuthorityWithContext",
		Input:   input,
		Output:  (*worklink.DisassociateWebsiteCertificateAuthorityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.DisassociateWebsiteCertificateAuthorityWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.DisassociateWebsiteCertificateAuthorityOutput), req.Error
}

func (c *Client) ListDevicesWithContext(ctx context.Context, input *worklink.ListDevicesInput, opts ...request.Option) (*worklink.ListDevicesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "ListDevicesWithContext",
		Input:   input,
		Output:  (*worklink.ListDevicesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.ListDevicesWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.ListDevicesOutput), req.Error
}

func (c *Client) ListDomainsWithContext(ctx context.Context, input *worklink.ListDomainsInput, opts ...request.Option) (*worklink.ListDomainsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "ListDomainsWithContext",
		Input:   input,
		Output:  (*worklink.ListDomainsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.ListDomainsWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.ListDomainsOutput), req.Error
}

func (c *Client) ListFleetsWithContext(ctx context.Context, input *worklink.ListFleetsInput, opts ...request.Option) (*worklink.ListFleetsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "ListFleetsWithContext",
		Input:   input,
		Output:  (*worklink.ListFleetsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.ListFleetsWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.ListFleetsOutput), req.Error
}

func (c *Client) ListWebsiteAuthorizationProvidersWithContext(ctx context.Context, input *worklink.ListWebsiteAuthorizationProvidersInput, opts ...request.Option) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "ListWebsiteAuthorizationProvidersWithContext",
		Input:   input,
		Output:  (*worklink.ListWebsiteAuthorizationProvidersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.ListWebsiteAuthorizationProvidersWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.ListWebsiteAuthorizationProvidersOutput), req.Error
}

func (c *Client) ListWebsiteCertificateAuthoritiesWithContext(ctx context.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput, opts ...request.Option) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "ListWebsiteCertificateAuthoritiesWithContext",
		Input:   input,
		Output:  (*worklink.ListWebsiteCertificateAuthoritiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.ListWebsiteCertificateAuthoritiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.ListWebsiteCertificateAuthoritiesOutput), req.Error
}

func (c *Client) RestoreDomainAccessWithContext(ctx context.Context, input *worklink.RestoreDomainAccessInput, opts ...request.Option) (*worklink.RestoreDomainAccessOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "RestoreDomainAccessWithContext",
		Input:   input,
		Output:  (*worklink.RestoreDomainAccessOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.RestoreDomainAccessWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.RestoreDomainAccessOutput), req.Error
}

func (c *Client) RevokeDomainAccessWithContext(ctx context.Context, input *worklink.RevokeDomainAccessInput, opts ...request.Option) (*worklink.RevokeDomainAccessOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "RevokeDomainAccessWithContext",
		Input:   input,
		Output:  (*worklink.RevokeDomainAccessOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.RevokeDomainAccessWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.RevokeDomainAccessOutput), req.Error
}

func (c *Client) SignOutUserWithContext(ctx context.Context, input *worklink.SignOutUserInput, opts ...request.Option) (*worklink.SignOutUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "SignOutUserWithContext",
		Input:   input,
		Output:  (*worklink.SignOutUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.SignOutUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.SignOutUserOutput), req.Error
}

func (c *Client) UpdateAuditStreamConfigurationWithContext(ctx context.Context, input *worklink.UpdateAuditStreamConfigurationInput, opts ...request.Option) (*worklink.UpdateAuditStreamConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "UpdateAuditStreamConfigurationWithContext",
		Input:   input,
		Output:  (*worklink.UpdateAuditStreamConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.UpdateAuditStreamConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.UpdateAuditStreamConfigurationOutput), req.Error
}

func (c *Client) UpdateCompanyNetworkConfigurationWithContext(ctx context.Context, input *worklink.UpdateCompanyNetworkConfigurationInput, opts ...request.Option) (*worklink.UpdateCompanyNetworkConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "UpdateCompanyNetworkConfigurationWithContext",
		Input:   input,
		Output:  (*worklink.UpdateCompanyNetworkConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.UpdateCompanyNetworkConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.UpdateCompanyNetworkConfigurationOutput), req.Error
}

func (c *Client) UpdateDevicePolicyConfigurationWithContext(ctx context.Context, input *worklink.UpdateDevicePolicyConfigurationInput, opts ...request.Option) (*worklink.UpdateDevicePolicyConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "UpdateDevicePolicyConfigurationWithContext",
		Input:   input,
		Output:  (*worklink.UpdateDevicePolicyConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.UpdateDevicePolicyConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.UpdateDevicePolicyConfigurationOutput), req.Error
}

func (c *Client) UpdateDomainMetadataWithContext(ctx context.Context, input *worklink.UpdateDomainMetadataInput, opts ...request.Option) (*worklink.UpdateDomainMetadataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "UpdateDomainMetadataWithContext",
		Input:   input,
		Output:  (*worklink.UpdateDomainMetadataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.UpdateDomainMetadataWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.UpdateDomainMetadataOutput), req.Error
}

func (c *Client) UpdateFleetMetadataWithContext(ctx context.Context, input *worklink.UpdateFleetMetadataInput, opts ...request.Option) (*worklink.UpdateFleetMetadataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "UpdateFleetMetadataWithContext",
		Input:   input,
		Output:  (*worklink.UpdateFleetMetadataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.UpdateFleetMetadataWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.UpdateFleetMetadataOutput), req.Error
}

func (c *Client) UpdateIdentityProviderConfigurationWithContext(ctx context.Context, input *worklink.UpdateIdentityProviderConfigurationInput, opts ...request.Option) (*worklink.UpdateIdentityProviderConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "worklink",
		Action:  "UpdateIdentityProviderConfigurationWithContext",
		Input:   input,
		Output:  (*worklink.UpdateIdentityProviderConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkLinkAPI.UpdateIdentityProviderConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*worklink.UpdateIdentityProviderConfigurationOutput), req.Error
}
