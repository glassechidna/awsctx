// Code generated by internal/generate/main.go. DO NOT EDIT.

package workmailctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workmail"
	"github.com/aws/aws-sdk-go/service/workmail/workmailiface"
	"github.com/glassechidna/awsctx"
)

type WorkMail interface {
	AssociateDelegateToResourceWithContext(ctx context.Context, input *workmail.AssociateDelegateToResourceInput, opts ...request.Option) (*workmail.AssociateDelegateToResourceOutput, error)
	AssociateMemberToGroupWithContext(ctx context.Context, input *workmail.AssociateMemberToGroupInput, opts ...request.Option) (*workmail.AssociateMemberToGroupOutput, error)
	CreateAliasWithContext(ctx context.Context, input *workmail.CreateAliasInput, opts ...request.Option) (*workmail.CreateAliasOutput, error)
	CreateGroupWithContext(ctx context.Context, input *workmail.CreateGroupInput, opts ...request.Option) (*workmail.CreateGroupOutput, error)
	CreateResourceWithContext(ctx context.Context, input *workmail.CreateResourceInput, opts ...request.Option) (*workmail.CreateResourceOutput, error)
	CreateUserWithContext(ctx context.Context, input *workmail.CreateUserInput, opts ...request.Option) (*workmail.CreateUserOutput, error)
	DeleteAliasWithContext(ctx context.Context, input *workmail.DeleteAliasInput, opts ...request.Option) (*workmail.DeleteAliasOutput, error)
	DeleteGroupWithContext(ctx context.Context, input *workmail.DeleteGroupInput, opts ...request.Option) (*workmail.DeleteGroupOutput, error)
	DeleteMailboxPermissionsWithContext(ctx context.Context, input *workmail.DeleteMailboxPermissionsInput, opts ...request.Option) (*workmail.DeleteMailboxPermissionsOutput, error)
	DeleteResourceWithContext(ctx context.Context, input *workmail.DeleteResourceInput, opts ...request.Option) (*workmail.DeleteResourceOutput, error)
	DeleteUserWithContext(ctx context.Context, input *workmail.DeleteUserInput, opts ...request.Option) (*workmail.DeleteUserOutput, error)
	DeregisterFromWorkMailWithContext(ctx context.Context, input *workmail.DeregisterFromWorkMailInput, opts ...request.Option) (*workmail.DeregisterFromWorkMailOutput, error)
	DescribeGroupWithContext(ctx context.Context, input *workmail.DescribeGroupInput, opts ...request.Option) (*workmail.DescribeGroupOutput, error)
	DescribeOrganizationWithContext(ctx context.Context, input *workmail.DescribeOrganizationInput, opts ...request.Option) (*workmail.DescribeOrganizationOutput, error)
	DescribeResourceWithContext(ctx context.Context, input *workmail.DescribeResourceInput, opts ...request.Option) (*workmail.DescribeResourceOutput, error)
	DescribeUserWithContext(ctx context.Context, input *workmail.DescribeUserInput, opts ...request.Option) (*workmail.DescribeUserOutput, error)
	DisassociateDelegateFromResourceWithContext(ctx context.Context, input *workmail.DisassociateDelegateFromResourceInput, opts ...request.Option) (*workmail.DisassociateDelegateFromResourceOutput, error)
	DisassociateMemberFromGroupWithContext(ctx context.Context, input *workmail.DisassociateMemberFromGroupInput, opts ...request.Option) (*workmail.DisassociateMemberFromGroupOutput, error)
	GetMailboxDetailsWithContext(ctx context.Context, input *workmail.GetMailboxDetailsInput, opts ...request.Option) (*workmail.GetMailboxDetailsOutput, error)
	ListAliasesWithContext(ctx context.Context, input *workmail.ListAliasesInput, opts ...request.Option) (*workmail.ListAliasesOutput, error)
	ListAliasesPagesWithContext(ctx context.Context, input *workmail.ListAliasesInput, cb func(*workmail.ListAliasesOutput, bool) bool, opts ...request.Option) error
	ListGroupMembersWithContext(ctx context.Context, input *workmail.ListGroupMembersInput, opts ...request.Option) (*workmail.ListGroupMembersOutput, error)
	ListGroupMembersPagesWithContext(ctx context.Context, input *workmail.ListGroupMembersInput, cb func(*workmail.ListGroupMembersOutput, bool) bool, opts ...request.Option) error
	ListGroupsWithContext(ctx context.Context, input *workmail.ListGroupsInput, opts ...request.Option) (*workmail.ListGroupsOutput, error)
	ListGroupsPagesWithContext(ctx context.Context, input *workmail.ListGroupsInput, cb func(*workmail.ListGroupsOutput, bool) bool, opts ...request.Option) error
	ListMailboxPermissionsWithContext(ctx context.Context, input *workmail.ListMailboxPermissionsInput, opts ...request.Option) (*workmail.ListMailboxPermissionsOutput, error)
	ListMailboxPermissionsPagesWithContext(ctx context.Context, input *workmail.ListMailboxPermissionsInput, cb func(*workmail.ListMailboxPermissionsOutput, bool) bool, opts ...request.Option) error
	ListOrganizationsWithContext(ctx context.Context, input *workmail.ListOrganizationsInput, opts ...request.Option) (*workmail.ListOrganizationsOutput, error)
	ListOrganizationsPagesWithContext(ctx context.Context, input *workmail.ListOrganizationsInput, cb func(*workmail.ListOrganizationsOutput, bool) bool, opts ...request.Option) error
	ListResourceDelegatesWithContext(ctx context.Context, input *workmail.ListResourceDelegatesInput, opts ...request.Option) (*workmail.ListResourceDelegatesOutput, error)
	ListResourceDelegatesPagesWithContext(ctx context.Context, input *workmail.ListResourceDelegatesInput, cb func(*workmail.ListResourceDelegatesOutput, bool) bool, opts ...request.Option) error
	ListResourcesWithContext(ctx context.Context, input *workmail.ListResourcesInput, opts ...request.Option) (*workmail.ListResourcesOutput, error)
	ListResourcesPagesWithContext(ctx context.Context, input *workmail.ListResourcesInput, cb func(*workmail.ListResourcesOutput, bool) bool, opts ...request.Option) error
	ListUsersWithContext(ctx context.Context, input *workmail.ListUsersInput, opts ...request.Option) (*workmail.ListUsersOutput, error)
	ListUsersPagesWithContext(ctx context.Context, input *workmail.ListUsersInput, cb func(*workmail.ListUsersOutput, bool) bool, opts ...request.Option) error
	PutMailboxPermissionsWithContext(ctx context.Context, input *workmail.PutMailboxPermissionsInput, opts ...request.Option) (*workmail.PutMailboxPermissionsOutput, error)
	RegisterToWorkMailWithContext(ctx context.Context, input *workmail.RegisterToWorkMailInput, opts ...request.Option) (*workmail.RegisterToWorkMailOutput, error)
	ResetPasswordWithContext(ctx context.Context, input *workmail.ResetPasswordInput, opts ...request.Option) (*workmail.ResetPasswordOutput, error)
	UpdateMailboxQuotaWithContext(ctx context.Context, input *workmail.UpdateMailboxQuotaInput, opts ...request.Option) (*workmail.UpdateMailboxQuotaOutput, error)
	UpdatePrimaryEmailAddressWithContext(ctx context.Context, input *workmail.UpdatePrimaryEmailAddressInput, opts ...request.Option) (*workmail.UpdatePrimaryEmailAddressOutput, error)
	UpdateResourceWithContext(ctx context.Context, input *workmail.UpdateResourceInput, opts ...request.Option) (*workmail.UpdateResourceOutput, error)
}

type Client struct {
	workmailiface.WorkMailAPI
	Contexter awsctx.Contexter
}

func New(base workmailiface.WorkMailAPI, ctxer awsctx.Contexter) WorkMail {
	return &Client{
		WorkMailAPI: base,
		Contexter: ctxer,
	}
}

var _ WorkMail = (*workmail.WorkMail)(nil)
var _ WorkMail = (*Client)(nil)

func (c *Client) AssociateDelegateToResourceWithContext(ctx context.Context, input *workmail.AssociateDelegateToResourceInput, opts ...request.Option) (*workmail.AssociateDelegateToResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "AssociateDelegateToResource",
		Input:   input,
		Output:  (*workmail.AssociateDelegateToResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.AssociateDelegateToResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.AssociateDelegateToResourceOutput), req.Error
}

func (c *Client) AssociateMemberToGroupWithContext(ctx context.Context, input *workmail.AssociateMemberToGroupInput, opts ...request.Option) (*workmail.AssociateMemberToGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "AssociateMemberToGroup",
		Input:   input,
		Output:  (*workmail.AssociateMemberToGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.AssociateMemberToGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.AssociateMemberToGroupOutput), req.Error
}

func (c *Client) CreateAliasWithContext(ctx context.Context, input *workmail.CreateAliasInput, opts ...request.Option) (*workmail.CreateAliasOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "CreateAlias",
		Input:   input,
		Output:  (*workmail.CreateAliasOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.CreateAliasWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.CreateAliasOutput), req.Error
}

func (c *Client) CreateGroupWithContext(ctx context.Context, input *workmail.CreateGroupInput, opts ...request.Option) (*workmail.CreateGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "CreateGroup",
		Input:   input,
		Output:  (*workmail.CreateGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.CreateGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.CreateGroupOutput), req.Error
}

func (c *Client) CreateResourceWithContext(ctx context.Context, input *workmail.CreateResourceInput, opts ...request.Option) (*workmail.CreateResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "CreateResource",
		Input:   input,
		Output:  (*workmail.CreateResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.CreateResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.CreateResourceOutput), req.Error
}

func (c *Client) CreateUserWithContext(ctx context.Context, input *workmail.CreateUserInput, opts ...request.Option) (*workmail.CreateUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "CreateUser",
		Input:   input,
		Output:  (*workmail.CreateUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.CreateUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.CreateUserOutput), req.Error
}

func (c *Client) DeleteAliasWithContext(ctx context.Context, input *workmail.DeleteAliasInput, opts ...request.Option) (*workmail.DeleteAliasOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DeleteAlias",
		Input:   input,
		Output:  (*workmail.DeleteAliasOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DeleteAliasWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DeleteAliasOutput), req.Error
}

func (c *Client) DeleteGroupWithContext(ctx context.Context, input *workmail.DeleteGroupInput, opts ...request.Option) (*workmail.DeleteGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DeleteGroup",
		Input:   input,
		Output:  (*workmail.DeleteGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DeleteGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DeleteGroupOutput), req.Error
}

func (c *Client) DeleteMailboxPermissionsWithContext(ctx context.Context, input *workmail.DeleteMailboxPermissionsInput, opts ...request.Option) (*workmail.DeleteMailboxPermissionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DeleteMailboxPermissions",
		Input:   input,
		Output:  (*workmail.DeleteMailboxPermissionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DeleteMailboxPermissionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DeleteMailboxPermissionsOutput), req.Error
}

func (c *Client) DeleteResourceWithContext(ctx context.Context, input *workmail.DeleteResourceInput, opts ...request.Option) (*workmail.DeleteResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DeleteResource",
		Input:   input,
		Output:  (*workmail.DeleteResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DeleteResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DeleteResourceOutput), req.Error
}

func (c *Client) DeleteUserWithContext(ctx context.Context, input *workmail.DeleteUserInput, opts ...request.Option) (*workmail.DeleteUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DeleteUser",
		Input:   input,
		Output:  (*workmail.DeleteUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DeleteUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DeleteUserOutput), req.Error
}

func (c *Client) DeregisterFromWorkMailWithContext(ctx context.Context, input *workmail.DeregisterFromWorkMailInput, opts ...request.Option) (*workmail.DeregisterFromWorkMailOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DeregisterFromWorkMail",
		Input:   input,
		Output:  (*workmail.DeregisterFromWorkMailOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DeregisterFromWorkMailWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DeregisterFromWorkMailOutput), req.Error
}

func (c *Client) DescribeGroupWithContext(ctx context.Context, input *workmail.DescribeGroupInput, opts ...request.Option) (*workmail.DescribeGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DescribeGroup",
		Input:   input,
		Output:  (*workmail.DescribeGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DescribeGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DescribeGroupOutput), req.Error
}

func (c *Client) DescribeOrganizationWithContext(ctx context.Context, input *workmail.DescribeOrganizationInput, opts ...request.Option) (*workmail.DescribeOrganizationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DescribeOrganization",
		Input:   input,
		Output:  (*workmail.DescribeOrganizationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DescribeOrganizationWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DescribeOrganizationOutput), req.Error
}

func (c *Client) DescribeResourceWithContext(ctx context.Context, input *workmail.DescribeResourceInput, opts ...request.Option) (*workmail.DescribeResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DescribeResource",
		Input:   input,
		Output:  (*workmail.DescribeResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DescribeResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DescribeResourceOutput), req.Error
}

func (c *Client) DescribeUserWithContext(ctx context.Context, input *workmail.DescribeUserInput, opts ...request.Option) (*workmail.DescribeUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DescribeUser",
		Input:   input,
		Output:  (*workmail.DescribeUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DescribeUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DescribeUserOutput), req.Error
}

func (c *Client) DisassociateDelegateFromResourceWithContext(ctx context.Context, input *workmail.DisassociateDelegateFromResourceInput, opts ...request.Option) (*workmail.DisassociateDelegateFromResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DisassociateDelegateFromResource",
		Input:   input,
		Output:  (*workmail.DisassociateDelegateFromResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DisassociateDelegateFromResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DisassociateDelegateFromResourceOutput), req.Error
}

func (c *Client) DisassociateMemberFromGroupWithContext(ctx context.Context, input *workmail.DisassociateMemberFromGroupInput, opts ...request.Option) (*workmail.DisassociateMemberFromGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "DisassociateMemberFromGroup",
		Input:   input,
		Output:  (*workmail.DisassociateMemberFromGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.DisassociateMemberFromGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.DisassociateMemberFromGroupOutput), req.Error
}

func (c *Client) GetMailboxDetailsWithContext(ctx context.Context, input *workmail.GetMailboxDetailsInput, opts ...request.Option) (*workmail.GetMailboxDetailsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "GetMailboxDetails",
		Input:   input,
		Output:  (*workmail.GetMailboxDetailsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.GetMailboxDetailsWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.GetMailboxDetailsOutput), req.Error
}

func (c *Client) ListAliasesWithContext(ctx context.Context, input *workmail.ListAliasesInput, opts ...request.Option) (*workmail.ListAliasesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListAliases",
		Input:   input,
		Output:  (*workmail.ListAliasesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ListAliasesWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ListAliasesOutput), req.Error
}

func (c *Client) ListAliasesPagesWithContext(ctx context.Context, input *workmail.ListAliasesInput, cb func(*workmail.ListAliasesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListAliases",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.WorkMailAPI.ListAliasesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListGroupMembersWithContext(ctx context.Context, input *workmail.ListGroupMembersInput, opts ...request.Option) (*workmail.ListGroupMembersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListGroupMembers",
		Input:   input,
		Output:  (*workmail.ListGroupMembersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ListGroupMembersWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ListGroupMembersOutput), req.Error
}

func (c *Client) ListGroupMembersPagesWithContext(ctx context.Context, input *workmail.ListGroupMembersInput, cb func(*workmail.ListGroupMembersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListGroupMembers",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.WorkMailAPI.ListGroupMembersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListGroupsWithContext(ctx context.Context, input *workmail.ListGroupsInput, opts ...request.Option) (*workmail.ListGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListGroups",
		Input:   input,
		Output:  (*workmail.ListGroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ListGroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ListGroupsOutput), req.Error
}

func (c *Client) ListGroupsPagesWithContext(ctx context.Context, input *workmail.ListGroupsInput, cb func(*workmail.ListGroupsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListGroups",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.WorkMailAPI.ListGroupsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListMailboxPermissionsWithContext(ctx context.Context, input *workmail.ListMailboxPermissionsInput, opts ...request.Option) (*workmail.ListMailboxPermissionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListMailboxPermissions",
		Input:   input,
		Output:  (*workmail.ListMailboxPermissionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ListMailboxPermissionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ListMailboxPermissionsOutput), req.Error
}

func (c *Client) ListMailboxPermissionsPagesWithContext(ctx context.Context, input *workmail.ListMailboxPermissionsInput, cb func(*workmail.ListMailboxPermissionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListMailboxPermissions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.WorkMailAPI.ListMailboxPermissionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListOrganizationsWithContext(ctx context.Context, input *workmail.ListOrganizationsInput, opts ...request.Option) (*workmail.ListOrganizationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListOrganizations",
		Input:   input,
		Output:  (*workmail.ListOrganizationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ListOrganizationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ListOrganizationsOutput), req.Error
}

func (c *Client) ListOrganizationsPagesWithContext(ctx context.Context, input *workmail.ListOrganizationsInput, cb func(*workmail.ListOrganizationsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListOrganizations",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.WorkMailAPI.ListOrganizationsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListResourceDelegatesWithContext(ctx context.Context, input *workmail.ListResourceDelegatesInput, opts ...request.Option) (*workmail.ListResourceDelegatesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListResourceDelegates",
		Input:   input,
		Output:  (*workmail.ListResourceDelegatesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ListResourceDelegatesWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ListResourceDelegatesOutput), req.Error
}

func (c *Client) ListResourceDelegatesPagesWithContext(ctx context.Context, input *workmail.ListResourceDelegatesInput, cb func(*workmail.ListResourceDelegatesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListResourceDelegates",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.WorkMailAPI.ListResourceDelegatesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListResourcesWithContext(ctx context.Context, input *workmail.ListResourcesInput, opts ...request.Option) (*workmail.ListResourcesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListResources",
		Input:   input,
		Output:  (*workmail.ListResourcesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ListResourcesWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ListResourcesOutput), req.Error
}

func (c *Client) ListResourcesPagesWithContext(ctx context.Context, input *workmail.ListResourcesInput, cb func(*workmail.ListResourcesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListResources",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.WorkMailAPI.ListResourcesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListUsersWithContext(ctx context.Context, input *workmail.ListUsersInput, opts ...request.Option) (*workmail.ListUsersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListUsers",
		Input:   input,
		Output:  (*workmail.ListUsersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ListUsersWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ListUsersOutput), req.Error
}

func (c *Client) ListUsersPagesWithContext(ctx context.Context, input *workmail.ListUsersInput, cb func(*workmail.ListUsersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ListUsers",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.WorkMailAPI.ListUsersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) PutMailboxPermissionsWithContext(ctx context.Context, input *workmail.PutMailboxPermissionsInput, opts ...request.Option) (*workmail.PutMailboxPermissionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "PutMailboxPermissions",
		Input:   input,
		Output:  (*workmail.PutMailboxPermissionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.PutMailboxPermissionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.PutMailboxPermissionsOutput), req.Error
}

func (c *Client) RegisterToWorkMailWithContext(ctx context.Context, input *workmail.RegisterToWorkMailInput, opts ...request.Option) (*workmail.RegisterToWorkMailOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "RegisterToWorkMail",
		Input:   input,
		Output:  (*workmail.RegisterToWorkMailOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.RegisterToWorkMailWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.RegisterToWorkMailOutput), req.Error
}

func (c *Client) ResetPasswordWithContext(ctx context.Context, input *workmail.ResetPasswordInput, opts ...request.Option) (*workmail.ResetPasswordOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "ResetPassword",
		Input:   input,
		Output:  (*workmail.ResetPasswordOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.ResetPasswordWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.ResetPasswordOutput), req.Error
}

func (c *Client) UpdateMailboxQuotaWithContext(ctx context.Context, input *workmail.UpdateMailboxQuotaInput, opts ...request.Option) (*workmail.UpdateMailboxQuotaOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "UpdateMailboxQuota",
		Input:   input,
		Output:  (*workmail.UpdateMailboxQuotaOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.UpdateMailboxQuotaWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.UpdateMailboxQuotaOutput), req.Error
}

func (c *Client) UpdatePrimaryEmailAddressWithContext(ctx context.Context, input *workmail.UpdatePrimaryEmailAddressInput, opts ...request.Option) (*workmail.UpdatePrimaryEmailAddressOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "UpdatePrimaryEmailAddress",
		Input:   input,
		Output:  (*workmail.UpdatePrimaryEmailAddressOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.UpdatePrimaryEmailAddressWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.UpdatePrimaryEmailAddressOutput), req.Error
}

func (c *Client) UpdateResourceWithContext(ctx context.Context, input *workmail.UpdateResourceInput, opts ...request.Option) (*workmail.UpdateResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "workmail",
		Action:  "UpdateResource",
		Input:   input,
		Output:  (*workmail.UpdateResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.WorkMailAPI.UpdateResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*workmail.UpdateResourceOutput), req.Error
}
