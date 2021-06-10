// Code generated by internal/generate/main.go. DO NOT EDIT.

package ramctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/aws/aws-sdk-go/service/ram/ramiface"
	"github.com/glassechidna/awsctx"
)

type RAM interface {
	AcceptResourceShareInvitationWithContext(ctx context.Context, input *ram.AcceptResourceShareInvitationInput, opts ...request.Option) (*ram.AcceptResourceShareInvitationOutput, error)
	AssociateResourceShareWithContext(ctx context.Context, input *ram.AssociateResourceShareInput, opts ...request.Option) (*ram.AssociateResourceShareOutput, error)
	AssociateResourceSharePermissionWithContext(ctx context.Context, input *ram.AssociateResourceSharePermissionInput, opts ...request.Option) (*ram.AssociateResourceSharePermissionOutput, error)
	CreateResourceShareWithContext(ctx context.Context, input *ram.CreateResourceShareInput, opts ...request.Option) (*ram.CreateResourceShareOutput, error)
	DeleteResourceShareWithContext(ctx context.Context, input *ram.DeleteResourceShareInput, opts ...request.Option) (*ram.DeleteResourceShareOutput, error)
	DisassociateResourceShareWithContext(ctx context.Context, input *ram.DisassociateResourceShareInput, opts ...request.Option) (*ram.DisassociateResourceShareOutput, error)
	DisassociateResourceSharePermissionWithContext(ctx context.Context, input *ram.DisassociateResourceSharePermissionInput, opts ...request.Option) (*ram.DisassociateResourceSharePermissionOutput, error)
	EnableSharingWithAwsOrganizationWithContext(ctx context.Context, input *ram.EnableSharingWithAwsOrganizationInput, opts ...request.Option) (*ram.EnableSharingWithAwsOrganizationOutput, error)
	GetPermissionWithContext(ctx context.Context, input *ram.GetPermissionInput, opts ...request.Option) (*ram.GetPermissionOutput, error)
	GetResourcePoliciesWithContext(ctx context.Context, input *ram.GetResourcePoliciesInput, opts ...request.Option) (*ram.GetResourcePoliciesOutput, error)
	GetResourcePoliciesPagesWithContext(ctx context.Context, input *ram.GetResourcePoliciesInput, cb func(*ram.GetResourcePoliciesOutput, bool) bool, opts ...request.Option) error
	GetResourceShareAssociationsWithContext(ctx context.Context, input *ram.GetResourceShareAssociationsInput, opts ...request.Option) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareAssociationsPagesWithContext(ctx context.Context, input *ram.GetResourceShareAssociationsInput, cb func(*ram.GetResourceShareAssociationsOutput, bool) bool, opts ...request.Option) error
	GetResourceShareInvitationsWithContext(ctx context.Context, input *ram.GetResourceShareInvitationsInput, opts ...request.Option) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShareInvitationsPagesWithContext(ctx context.Context, input *ram.GetResourceShareInvitationsInput, cb func(*ram.GetResourceShareInvitationsOutput, bool) bool, opts ...request.Option) error
	GetResourceSharesWithContext(ctx context.Context, input *ram.GetResourceSharesInput, opts ...request.Option) (*ram.GetResourceSharesOutput, error)
	GetResourceSharesPagesWithContext(ctx context.Context, input *ram.GetResourceSharesInput, cb func(*ram.GetResourceSharesOutput, bool) bool, opts ...request.Option) error
	ListPendingInvitationResourcesWithContext(ctx context.Context, input *ram.ListPendingInvitationResourcesInput, opts ...request.Option) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPendingInvitationResourcesPagesWithContext(ctx context.Context, input *ram.ListPendingInvitationResourcesInput, cb func(*ram.ListPendingInvitationResourcesOutput, bool) bool, opts ...request.Option) error
	ListPermissionsWithContext(ctx context.Context, input *ram.ListPermissionsInput, opts ...request.Option) (*ram.ListPermissionsOutput, error)
	ListPermissionsPagesWithContext(ctx context.Context, input *ram.ListPermissionsInput, cb func(*ram.ListPermissionsOutput, bool) bool, opts ...request.Option) error
	ListPrincipalsWithContext(ctx context.Context, input *ram.ListPrincipalsInput, opts ...request.Option) (*ram.ListPrincipalsOutput, error)
	ListPrincipalsPagesWithContext(ctx context.Context, input *ram.ListPrincipalsInput, cb func(*ram.ListPrincipalsOutput, bool) bool, opts ...request.Option) error
	ListResourceSharePermissionsWithContext(ctx context.Context, input *ram.ListResourceSharePermissionsInput, opts ...request.Option) (*ram.ListResourceSharePermissionsOutput, error)
	ListResourceSharePermissionsPagesWithContext(ctx context.Context, input *ram.ListResourceSharePermissionsInput, cb func(*ram.ListResourceSharePermissionsOutput, bool) bool, opts ...request.Option) error
	ListResourceTypesWithContext(ctx context.Context, input *ram.ListResourceTypesInput, opts ...request.Option) (*ram.ListResourceTypesOutput, error)
	ListResourceTypesPagesWithContext(ctx context.Context, input *ram.ListResourceTypesInput, cb func(*ram.ListResourceTypesOutput, bool) bool, opts ...request.Option) error
	ListResourcesWithContext(ctx context.Context, input *ram.ListResourcesInput, opts ...request.Option) (*ram.ListResourcesOutput, error)
	ListResourcesPagesWithContext(ctx context.Context, input *ram.ListResourcesInput, cb func(*ram.ListResourcesOutput, bool) bool, opts ...request.Option) error
	PromoteResourceShareCreatedFromPolicyWithContext(ctx context.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput, opts ...request.Option) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error)
	RejectResourceShareInvitationWithContext(ctx context.Context, input *ram.RejectResourceShareInvitationInput, opts ...request.Option) (*ram.RejectResourceShareInvitationOutput, error)
	TagResourceWithContext(ctx context.Context, input *ram.TagResourceInput, opts ...request.Option) (*ram.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *ram.UntagResourceInput, opts ...request.Option) (*ram.UntagResourceOutput, error)
	UpdateResourceShareWithContext(ctx context.Context, input *ram.UpdateResourceShareInput, opts ...request.Option) (*ram.UpdateResourceShareOutput, error)
}

type Client struct {
	ramiface.RAMAPI
	Contexter awsctx.Contexter
}

func New(base ramiface.RAMAPI, ctxer awsctx.Contexter) RAM {
	return &Client{
		RAMAPI: base,
		Contexter: ctxer,
	}
}

var _ RAM = (*ram.RAM)(nil)
var _ RAM = (*Client)(nil)

func (c *Client) AcceptResourceShareInvitationWithContext(ctx context.Context, input *ram.AcceptResourceShareInvitationInput, opts ...request.Option) (*ram.AcceptResourceShareInvitationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "AcceptResourceShareInvitation",
		Input:   input,
		Output:  (*ram.AcceptResourceShareInvitationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.AcceptResourceShareInvitationWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.AcceptResourceShareInvitationOutput), req.Error
}

func (c *Client) AssociateResourceShareWithContext(ctx context.Context, input *ram.AssociateResourceShareInput, opts ...request.Option) (*ram.AssociateResourceShareOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "AssociateResourceShare",
		Input:   input,
		Output:  (*ram.AssociateResourceShareOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.AssociateResourceShareWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.AssociateResourceShareOutput), req.Error
}

func (c *Client) AssociateResourceSharePermissionWithContext(ctx context.Context, input *ram.AssociateResourceSharePermissionInput, opts ...request.Option) (*ram.AssociateResourceSharePermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "AssociateResourceSharePermission",
		Input:   input,
		Output:  (*ram.AssociateResourceSharePermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.AssociateResourceSharePermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.AssociateResourceSharePermissionOutput), req.Error
}

func (c *Client) CreateResourceShareWithContext(ctx context.Context, input *ram.CreateResourceShareInput, opts ...request.Option) (*ram.CreateResourceShareOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "CreateResourceShare",
		Input:   input,
		Output:  (*ram.CreateResourceShareOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.CreateResourceShareWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.CreateResourceShareOutput), req.Error
}

func (c *Client) DeleteResourceShareWithContext(ctx context.Context, input *ram.DeleteResourceShareInput, opts ...request.Option) (*ram.DeleteResourceShareOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "DeleteResourceShare",
		Input:   input,
		Output:  (*ram.DeleteResourceShareOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.DeleteResourceShareWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.DeleteResourceShareOutput), req.Error
}

func (c *Client) DisassociateResourceShareWithContext(ctx context.Context, input *ram.DisassociateResourceShareInput, opts ...request.Option) (*ram.DisassociateResourceShareOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "DisassociateResourceShare",
		Input:   input,
		Output:  (*ram.DisassociateResourceShareOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.DisassociateResourceShareWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.DisassociateResourceShareOutput), req.Error
}

func (c *Client) DisassociateResourceSharePermissionWithContext(ctx context.Context, input *ram.DisassociateResourceSharePermissionInput, opts ...request.Option) (*ram.DisassociateResourceSharePermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "DisassociateResourceSharePermission",
		Input:   input,
		Output:  (*ram.DisassociateResourceSharePermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.DisassociateResourceSharePermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.DisassociateResourceSharePermissionOutput), req.Error
}

func (c *Client) EnableSharingWithAwsOrganizationWithContext(ctx context.Context, input *ram.EnableSharingWithAwsOrganizationInput, opts ...request.Option) (*ram.EnableSharingWithAwsOrganizationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "EnableSharingWithAwsOrganization",
		Input:   input,
		Output:  (*ram.EnableSharingWithAwsOrganizationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.EnableSharingWithAwsOrganizationWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.EnableSharingWithAwsOrganizationOutput), req.Error
}

func (c *Client) GetPermissionWithContext(ctx context.Context, input *ram.GetPermissionInput, opts ...request.Option) (*ram.GetPermissionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetPermission",
		Input:   input,
		Output:  (*ram.GetPermissionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.GetPermissionWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.GetPermissionOutput), req.Error
}

func (c *Client) GetResourcePoliciesWithContext(ctx context.Context, input *ram.GetResourcePoliciesInput, opts ...request.Option) (*ram.GetResourcePoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetResourcePolicies",
		Input:   input,
		Output:  (*ram.GetResourcePoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.GetResourcePoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.GetResourcePoliciesOutput), req.Error
}

func (c *Client) GetResourcePoliciesPagesWithContext(ctx context.Context, input *ram.GetResourcePoliciesInput, cb func(*ram.GetResourcePoliciesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetResourcePolicies",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.GetResourcePoliciesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetResourceShareAssociationsWithContext(ctx context.Context, input *ram.GetResourceShareAssociationsInput, opts ...request.Option) (*ram.GetResourceShareAssociationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetResourceShareAssociations",
		Input:   input,
		Output:  (*ram.GetResourceShareAssociationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.GetResourceShareAssociationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.GetResourceShareAssociationsOutput), req.Error
}

func (c *Client) GetResourceShareAssociationsPagesWithContext(ctx context.Context, input *ram.GetResourceShareAssociationsInput, cb func(*ram.GetResourceShareAssociationsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetResourceShareAssociations",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.GetResourceShareAssociationsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetResourceShareInvitationsWithContext(ctx context.Context, input *ram.GetResourceShareInvitationsInput, opts ...request.Option) (*ram.GetResourceShareInvitationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetResourceShareInvitations",
		Input:   input,
		Output:  (*ram.GetResourceShareInvitationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.GetResourceShareInvitationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.GetResourceShareInvitationsOutput), req.Error
}

func (c *Client) GetResourceShareInvitationsPagesWithContext(ctx context.Context, input *ram.GetResourceShareInvitationsInput, cb func(*ram.GetResourceShareInvitationsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetResourceShareInvitations",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.GetResourceShareInvitationsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetResourceSharesWithContext(ctx context.Context, input *ram.GetResourceSharesInput, opts ...request.Option) (*ram.GetResourceSharesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetResourceShares",
		Input:   input,
		Output:  (*ram.GetResourceSharesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.GetResourceSharesWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.GetResourceSharesOutput), req.Error
}

func (c *Client) GetResourceSharesPagesWithContext(ctx context.Context, input *ram.GetResourceSharesInput, cb func(*ram.GetResourceSharesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "GetResourceShares",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.GetResourceSharesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListPendingInvitationResourcesWithContext(ctx context.Context, input *ram.ListPendingInvitationResourcesInput, opts ...request.Option) (*ram.ListPendingInvitationResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListPendingInvitationResources",
		Input:   input,
		Output:  (*ram.ListPendingInvitationResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.ListPendingInvitationResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.ListPendingInvitationResourcesOutput), req.Error
}

func (c *Client) ListPendingInvitationResourcesPagesWithContext(ctx context.Context, input *ram.ListPendingInvitationResourcesInput, cb func(*ram.ListPendingInvitationResourcesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListPendingInvitationResources",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.ListPendingInvitationResourcesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListPermissionsWithContext(ctx context.Context, input *ram.ListPermissionsInput, opts ...request.Option) (*ram.ListPermissionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListPermissions",
		Input:   input,
		Output:  (*ram.ListPermissionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.ListPermissionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.ListPermissionsOutput), req.Error
}

func (c *Client) ListPermissionsPagesWithContext(ctx context.Context, input *ram.ListPermissionsInput, cb func(*ram.ListPermissionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListPermissions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.ListPermissionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListPrincipalsWithContext(ctx context.Context, input *ram.ListPrincipalsInput, opts ...request.Option) (*ram.ListPrincipalsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListPrincipals",
		Input:   input,
		Output:  (*ram.ListPrincipalsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.ListPrincipalsWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.ListPrincipalsOutput), req.Error
}

func (c *Client) ListPrincipalsPagesWithContext(ctx context.Context, input *ram.ListPrincipalsInput, cb func(*ram.ListPrincipalsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListPrincipals",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.ListPrincipalsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListResourceSharePermissionsWithContext(ctx context.Context, input *ram.ListResourceSharePermissionsInput, opts ...request.Option) (*ram.ListResourceSharePermissionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListResourceSharePermissions",
		Input:   input,
		Output:  (*ram.ListResourceSharePermissionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.ListResourceSharePermissionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.ListResourceSharePermissionsOutput), req.Error
}

func (c *Client) ListResourceSharePermissionsPagesWithContext(ctx context.Context, input *ram.ListResourceSharePermissionsInput, cb func(*ram.ListResourceSharePermissionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListResourceSharePermissions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.ListResourceSharePermissionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListResourceTypesWithContext(ctx context.Context, input *ram.ListResourceTypesInput, opts ...request.Option) (*ram.ListResourceTypesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListResourceTypes",
		Input:   input,
		Output:  (*ram.ListResourceTypesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.ListResourceTypesWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.ListResourceTypesOutput), req.Error
}

func (c *Client) ListResourceTypesPagesWithContext(ctx context.Context, input *ram.ListResourceTypesInput, cb func(*ram.ListResourceTypesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListResourceTypes",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.ListResourceTypesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListResourcesWithContext(ctx context.Context, input *ram.ListResourcesInput, opts ...request.Option) (*ram.ListResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListResources",
		Input:   input,
		Output:  (*ram.ListResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.ListResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.ListResourcesOutput), req.Error
}

func (c *Client) ListResourcesPagesWithContext(ctx context.Context, input *ram.ListResourcesInput, cb func(*ram.ListResourcesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "ListResources",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.RAMAPI.ListResourcesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) PromoteResourceShareCreatedFromPolicyWithContext(ctx context.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput, opts ...request.Option) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "PromoteResourceShareCreatedFromPolicy",
		Input:   input,
		Output:  (*ram.PromoteResourceShareCreatedFromPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.PromoteResourceShareCreatedFromPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.PromoteResourceShareCreatedFromPolicyOutput), req.Error
}

func (c *Client) RejectResourceShareInvitationWithContext(ctx context.Context, input *ram.RejectResourceShareInvitationInput, opts ...request.Option) (*ram.RejectResourceShareInvitationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "RejectResourceShareInvitation",
		Input:   input,
		Output:  (*ram.RejectResourceShareInvitationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.RejectResourceShareInvitationWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.RejectResourceShareInvitationOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *ram.TagResourceInput, opts ...request.Option) (*ram.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "TagResource",
		Input:   input,
		Output:  (*ram.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *ram.UntagResourceInput, opts ...request.Option) (*ram.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*ram.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.UntagResourceOutput), req.Error
}

func (c *Client) UpdateResourceShareWithContext(ctx context.Context, input *ram.UpdateResourceShareInput, opts ...request.Option) (*ram.UpdateResourceShareOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ram",
		Action:  "UpdateResourceShare",
		Input:   input,
		Output:  (*ram.UpdateResourceShareOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.RAMAPI.UpdateResourceShareWithContext(ctx, input, opts...)
	})

	return req.Output.(*ram.UpdateResourceShareOutput), req.Error
}
