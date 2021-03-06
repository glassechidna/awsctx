// Code generated by internal/generate/main.go. DO NOT EDIT.

package transferctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/aws/aws-sdk-go/service/transfer/transferiface"
	"github.com/glassechidna/awsctx"
)

type Transfer interface {
	CreateServerWithContext(ctx context.Context, input *transfer.CreateServerInput, opts ...request.Option) (*transfer.CreateServerOutput, error)
	CreateUserWithContext(ctx context.Context, input *transfer.CreateUserInput, opts ...request.Option) (*transfer.CreateUserOutput, error)
	DeleteServerWithContext(ctx context.Context, input *transfer.DeleteServerInput, opts ...request.Option) (*transfer.DeleteServerOutput, error)
	DeleteSshPublicKeyWithContext(ctx context.Context, input *transfer.DeleteSshPublicKeyInput, opts ...request.Option) (*transfer.DeleteSshPublicKeyOutput, error)
	DeleteUserWithContext(ctx context.Context, input *transfer.DeleteUserInput, opts ...request.Option) (*transfer.DeleteUserOutput, error)
	DescribeSecurityPolicyWithContext(ctx context.Context, input *transfer.DescribeSecurityPolicyInput, opts ...request.Option) (*transfer.DescribeSecurityPolicyOutput, error)
	DescribeServerWithContext(ctx context.Context, input *transfer.DescribeServerInput, opts ...request.Option) (*transfer.DescribeServerOutput, error)
	DescribeUserWithContext(ctx context.Context, input *transfer.DescribeUserInput, opts ...request.Option) (*transfer.DescribeUserOutput, error)
	ImportSshPublicKeyWithContext(ctx context.Context, input *transfer.ImportSshPublicKeyInput, opts ...request.Option) (*transfer.ImportSshPublicKeyOutput, error)
	ListSecurityPoliciesWithContext(ctx context.Context, input *transfer.ListSecurityPoliciesInput, opts ...request.Option) (*transfer.ListSecurityPoliciesOutput, error)
	ListSecurityPoliciesPagesWithContext(ctx context.Context, input *transfer.ListSecurityPoliciesInput, cb func(*transfer.ListSecurityPoliciesOutput, bool) bool, opts ...request.Option) error
	ListServersWithContext(ctx context.Context, input *transfer.ListServersInput, opts ...request.Option) (*transfer.ListServersOutput, error)
	ListServersPagesWithContext(ctx context.Context, input *transfer.ListServersInput, cb func(*transfer.ListServersOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *transfer.ListTagsForResourceInput, opts ...request.Option) (*transfer.ListTagsForResourceOutput, error)
	ListTagsForResourcePagesWithContext(ctx context.Context, input *transfer.ListTagsForResourceInput, cb func(*transfer.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error
	ListUsersWithContext(ctx context.Context, input *transfer.ListUsersInput, opts ...request.Option) (*transfer.ListUsersOutput, error)
	ListUsersPagesWithContext(ctx context.Context, input *transfer.ListUsersInput, cb func(*transfer.ListUsersOutput, bool) bool, opts ...request.Option) error
	StartServerWithContext(ctx context.Context, input *transfer.StartServerInput, opts ...request.Option) (*transfer.StartServerOutput, error)
	StopServerWithContext(ctx context.Context, input *transfer.StopServerInput, opts ...request.Option) (*transfer.StopServerOutput, error)
	TagResourceWithContext(ctx context.Context, input *transfer.TagResourceInput, opts ...request.Option) (*transfer.TagResourceOutput, error)
	TestIdentityProviderWithContext(ctx context.Context, input *transfer.TestIdentityProviderInput, opts ...request.Option) (*transfer.TestIdentityProviderOutput, error)
	UntagResourceWithContext(ctx context.Context, input *transfer.UntagResourceInput, opts ...request.Option) (*transfer.UntagResourceOutput, error)
	UpdateServerWithContext(ctx context.Context, input *transfer.UpdateServerInput, opts ...request.Option) (*transfer.UpdateServerOutput, error)
	UpdateUserWithContext(ctx context.Context, input *transfer.UpdateUserInput, opts ...request.Option) (*transfer.UpdateUserOutput, error)
}

type Client struct {
	transferiface.TransferAPI
	Contexter awsctx.Contexter
}

func New(base transferiface.TransferAPI, ctxer awsctx.Contexter) Transfer {
	return &Client{
		TransferAPI: base,
		Contexter: ctxer,
	}
}

var _ Transfer = (*transfer.Transfer)(nil)
var _ Transfer = (*Client)(nil)

func (c *Client) CreateServerWithContext(ctx context.Context, input *transfer.CreateServerInput, opts ...request.Option) (*transfer.CreateServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "CreateServer",
		Input:   input,
		Output:  (*transfer.CreateServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.CreateServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.CreateServerOutput), req.Error
}

func (c *Client) CreateUserWithContext(ctx context.Context, input *transfer.CreateUserInput, opts ...request.Option) (*transfer.CreateUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "CreateUser",
		Input:   input,
		Output:  (*transfer.CreateUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.CreateUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.CreateUserOutput), req.Error
}

func (c *Client) DeleteServerWithContext(ctx context.Context, input *transfer.DeleteServerInput, opts ...request.Option) (*transfer.DeleteServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "DeleteServer",
		Input:   input,
		Output:  (*transfer.DeleteServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.DeleteServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.DeleteServerOutput), req.Error
}

func (c *Client) DeleteSshPublicKeyWithContext(ctx context.Context, input *transfer.DeleteSshPublicKeyInput, opts ...request.Option) (*transfer.DeleteSshPublicKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "DeleteSshPublicKey",
		Input:   input,
		Output:  (*transfer.DeleteSshPublicKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.DeleteSshPublicKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.DeleteSshPublicKeyOutput), req.Error
}

func (c *Client) DeleteUserWithContext(ctx context.Context, input *transfer.DeleteUserInput, opts ...request.Option) (*transfer.DeleteUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "DeleteUser",
		Input:   input,
		Output:  (*transfer.DeleteUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.DeleteUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.DeleteUserOutput), req.Error
}

func (c *Client) DescribeSecurityPolicyWithContext(ctx context.Context, input *transfer.DescribeSecurityPolicyInput, opts ...request.Option) (*transfer.DescribeSecurityPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "DescribeSecurityPolicy",
		Input:   input,
		Output:  (*transfer.DescribeSecurityPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.DescribeSecurityPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.DescribeSecurityPolicyOutput), req.Error
}

func (c *Client) DescribeServerWithContext(ctx context.Context, input *transfer.DescribeServerInput, opts ...request.Option) (*transfer.DescribeServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "DescribeServer",
		Input:   input,
		Output:  (*transfer.DescribeServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.DescribeServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.DescribeServerOutput), req.Error
}

func (c *Client) DescribeUserWithContext(ctx context.Context, input *transfer.DescribeUserInput, opts ...request.Option) (*transfer.DescribeUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "DescribeUser",
		Input:   input,
		Output:  (*transfer.DescribeUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.DescribeUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.DescribeUserOutput), req.Error
}

func (c *Client) ImportSshPublicKeyWithContext(ctx context.Context, input *transfer.ImportSshPublicKeyInput, opts ...request.Option) (*transfer.ImportSshPublicKeyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ImportSshPublicKey",
		Input:   input,
		Output:  (*transfer.ImportSshPublicKeyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.ImportSshPublicKeyWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.ImportSshPublicKeyOutput), req.Error
}

func (c *Client) ListSecurityPoliciesWithContext(ctx context.Context, input *transfer.ListSecurityPoliciesInput, opts ...request.Option) (*transfer.ListSecurityPoliciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ListSecurityPolicies",
		Input:   input,
		Output:  (*transfer.ListSecurityPoliciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.ListSecurityPoliciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.ListSecurityPoliciesOutput), req.Error
}

func (c *Client) ListSecurityPoliciesPagesWithContext(ctx context.Context, input *transfer.ListSecurityPoliciesInput, cb func(*transfer.ListSecurityPoliciesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ListSecurityPolicies",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.TransferAPI.ListSecurityPoliciesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListServersWithContext(ctx context.Context, input *transfer.ListServersInput, opts ...request.Option) (*transfer.ListServersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ListServers",
		Input:   input,
		Output:  (*transfer.ListServersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.ListServersWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.ListServersOutput), req.Error
}

func (c *Client) ListServersPagesWithContext(ctx context.Context, input *transfer.ListServersInput, cb func(*transfer.ListServersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ListServers",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.TransferAPI.ListServersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *transfer.ListTagsForResourceInput, opts ...request.Option) (*transfer.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*transfer.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListTagsForResourcePagesWithContext(ctx context.Context, input *transfer.ListTagsForResourceInput, cb func(*transfer.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ListTagsForResource",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.TransferAPI.ListTagsForResourcePagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListUsersWithContext(ctx context.Context, input *transfer.ListUsersInput, opts ...request.Option) (*transfer.ListUsersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ListUsers",
		Input:   input,
		Output:  (*transfer.ListUsersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.ListUsersWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.ListUsersOutput), req.Error
}

func (c *Client) ListUsersPagesWithContext(ctx context.Context, input *transfer.ListUsersInput, cb func(*transfer.ListUsersOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "ListUsers",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.TransferAPI.ListUsersPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) StartServerWithContext(ctx context.Context, input *transfer.StartServerInput, opts ...request.Option) (*transfer.StartServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "StartServer",
		Input:   input,
		Output:  (*transfer.StartServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.StartServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.StartServerOutput), req.Error
}

func (c *Client) StopServerWithContext(ctx context.Context, input *transfer.StopServerInput, opts ...request.Option) (*transfer.StopServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "StopServer",
		Input:   input,
		Output:  (*transfer.StopServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.StopServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.StopServerOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *transfer.TagResourceInput, opts ...request.Option) (*transfer.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "TagResource",
		Input:   input,
		Output:  (*transfer.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.TagResourceOutput), req.Error
}

func (c *Client) TestIdentityProviderWithContext(ctx context.Context, input *transfer.TestIdentityProviderInput, opts ...request.Option) (*transfer.TestIdentityProviderOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "TestIdentityProvider",
		Input:   input,
		Output:  (*transfer.TestIdentityProviderOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.TestIdentityProviderWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.TestIdentityProviderOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *transfer.UntagResourceInput, opts ...request.Option) (*transfer.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*transfer.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.UntagResourceOutput), req.Error
}

func (c *Client) UpdateServerWithContext(ctx context.Context, input *transfer.UpdateServerInput, opts ...request.Option) (*transfer.UpdateServerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "UpdateServer",
		Input:   input,
		Output:  (*transfer.UpdateServerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.UpdateServerWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.UpdateServerOutput), req.Error
}

func (c *Client) UpdateUserWithContext(ctx context.Context, input *transfer.UpdateUserInput, opts ...request.Option) (*transfer.UpdateUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "transfer",
		Action:  "UpdateUser",
		Input:   input,
		Output:  (*transfer.UpdateUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.TransferAPI.UpdateUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*transfer.UpdateUserOutput), req.Error
}
