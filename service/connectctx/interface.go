// Code generated by internal/generate/main.go. DO NOT EDIT.

package connectctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/aws/aws-sdk-go/service/connect/connectiface"
	"github.com/glassechidna/awsctx"
)

type Connect interface {
	CreateUserWithContext(ctx context.Context, input *connect.CreateUserInput, opts ...request.Option) (*connect.CreateUserOutput, error)
	DeleteUserWithContext(ctx context.Context, input *connect.DeleteUserInput, opts ...request.Option) (*connect.DeleteUserOutput, error)
	DescribeUserWithContext(ctx context.Context, input *connect.DescribeUserInput, opts ...request.Option) (*connect.DescribeUserOutput, error)
	DescribeUserHierarchyGroupWithContext(ctx context.Context, input *connect.DescribeUserHierarchyGroupInput, opts ...request.Option) (*connect.DescribeUserHierarchyGroupOutput, error)
	DescribeUserHierarchyStructureWithContext(ctx context.Context, input *connect.DescribeUserHierarchyStructureInput, opts ...request.Option) (*connect.DescribeUserHierarchyStructureOutput, error)
	GetContactAttributesWithContext(ctx context.Context, input *connect.GetContactAttributesInput, opts ...request.Option) (*connect.GetContactAttributesOutput, error)
	GetCurrentMetricDataWithContext(ctx context.Context, input *connect.GetCurrentMetricDataInput, opts ...request.Option) (*connect.GetCurrentMetricDataOutput, error)
	GetFederationTokenWithContext(ctx context.Context, input *connect.GetFederationTokenInput, opts ...request.Option) (*connect.GetFederationTokenOutput, error)
	GetMetricDataWithContext(ctx context.Context, input *connect.GetMetricDataInput, opts ...request.Option) (*connect.GetMetricDataOutput, error)
	ListRoutingProfilesWithContext(ctx context.Context, input *connect.ListRoutingProfilesInput, opts ...request.Option) (*connect.ListRoutingProfilesOutput, error)
	ListSecurityProfilesWithContext(ctx context.Context, input *connect.ListSecurityProfilesInput, opts ...request.Option) (*connect.ListSecurityProfilesOutput, error)
	ListUserHierarchyGroupsWithContext(ctx context.Context, input *connect.ListUserHierarchyGroupsInput, opts ...request.Option) (*connect.ListUserHierarchyGroupsOutput, error)
	ListUsersWithContext(ctx context.Context, input *connect.ListUsersInput, opts ...request.Option) (*connect.ListUsersOutput, error)
	StartOutboundVoiceContactWithContext(ctx context.Context, input *connect.StartOutboundVoiceContactInput, opts ...request.Option) (*connect.StartOutboundVoiceContactOutput, error)
	StopContactWithContext(ctx context.Context, input *connect.StopContactInput, opts ...request.Option) (*connect.StopContactOutput, error)
	UpdateContactAttributesWithContext(ctx context.Context, input *connect.UpdateContactAttributesInput, opts ...request.Option) (*connect.UpdateContactAttributesOutput, error)
	UpdateUserHierarchyWithContext(ctx context.Context, input *connect.UpdateUserHierarchyInput, opts ...request.Option) (*connect.UpdateUserHierarchyOutput, error)
	UpdateUserIdentityInfoWithContext(ctx context.Context, input *connect.UpdateUserIdentityInfoInput, opts ...request.Option) (*connect.UpdateUserIdentityInfoOutput, error)
	UpdateUserPhoneConfigWithContext(ctx context.Context, input *connect.UpdateUserPhoneConfigInput, opts ...request.Option) (*connect.UpdateUserPhoneConfigOutput, error)
	UpdateUserRoutingProfileWithContext(ctx context.Context, input *connect.UpdateUserRoutingProfileInput, opts ...request.Option) (*connect.UpdateUserRoutingProfileOutput, error)
	UpdateUserSecurityProfilesWithContext(ctx context.Context, input *connect.UpdateUserSecurityProfilesInput, opts ...request.Option) (*connect.UpdateUserSecurityProfilesOutput, error)
}

type Client struct {
	connectiface.ConnectAPI
	Contexter awsctx.Contexter
}

func New(base connectiface.ConnectAPI, ctxer awsctx.Contexter) Connect {
	return &Client{
		ConnectAPI: base,
		Contexter: ctxer,
	}
}

var _ Connect = (*connect.Connect)(nil)
var _ Connect = (*Client)(nil)

func (c *Client) CreateUserWithContext(ctx context.Context, input *connect.CreateUserInput, opts ...request.Option) (*connect.CreateUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "CreateUser",
		Input:   input,
		Output:  (*connect.CreateUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.CreateUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.CreateUserOutput), req.Error
}

func (c *Client) DeleteUserWithContext(ctx context.Context, input *connect.DeleteUserInput, opts ...request.Option) (*connect.DeleteUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "DeleteUser",
		Input:   input,
		Output:  (*connect.DeleteUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.DeleteUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.DeleteUserOutput), req.Error
}

func (c *Client) DescribeUserWithContext(ctx context.Context, input *connect.DescribeUserInput, opts ...request.Option) (*connect.DescribeUserOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "DescribeUser",
		Input:   input,
		Output:  (*connect.DescribeUserOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.DescribeUserWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.DescribeUserOutput), req.Error
}

func (c *Client) DescribeUserHierarchyGroupWithContext(ctx context.Context, input *connect.DescribeUserHierarchyGroupInput, opts ...request.Option) (*connect.DescribeUserHierarchyGroupOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "DescribeUserHierarchyGroup",
		Input:   input,
		Output:  (*connect.DescribeUserHierarchyGroupOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.DescribeUserHierarchyGroupWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.DescribeUserHierarchyGroupOutput), req.Error
}

func (c *Client) DescribeUserHierarchyStructureWithContext(ctx context.Context, input *connect.DescribeUserHierarchyStructureInput, opts ...request.Option) (*connect.DescribeUserHierarchyStructureOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "DescribeUserHierarchyStructure",
		Input:   input,
		Output:  (*connect.DescribeUserHierarchyStructureOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.DescribeUserHierarchyStructureWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.DescribeUserHierarchyStructureOutput), req.Error
}

func (c *Client) GetContactAttributesWithContext(ctx context.Context, input *connect.GetContactAttributesInput, opts ...request.Option) (*connect.GetContactAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "GetContactAttributes",
		Input:   input,
		Output:  (*connect.GetContactAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.GetContactAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.GetContactAttributesOutput), req.Error
}

func (c *Client) GetCurrentMetricDataWithContext(ctx context.Context, input *connect.GetCurrentMetricDataInput, opts ...request.Option) (*connect.GetCurrentMetricDataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "GetCurrentMetricData",
		Input:   input,
		Output:  (*connect.GetCurrentMetricDataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.GetCurrentMetricDataWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.GetCurrentMetricDataOutput), req.Error
}

func (c *Client) GetFederationTokenWithContext(ctx context.Context, input *connect.GetFederationTokenInput, opts ...request.Option) (*connect.GetFederationTokenOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "GetFederationToken",
		Input:   input,
		Output:  (*connect.GetFederationTokenOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.GetFederationTokenWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.GetFederationTokenOutput), req.Error
}

func (c *Client) GetMetricDataWithContext(ctx context.Context, input *connect.GetMetricDataInput, opts ...request.Option) (*connect.GetMetricDataOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "GetMetricData",
		Input:   input,
		Output:  (*connect.GetMetricDataOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.GetMetricDataWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.GetMetricDataOutput), req.Error
}

func (c *Client) ListRoutingProfilesWithContext(ctx context.Context, input *connect.ListRoutingProfilesInput, opts ...request.Option) (*connect.ListRoutingProfilesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "ListRoutingProfiles",
		Input:   input,
		Output:  (*connect.ListRoutingProfilesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.ListRoutingProfilesWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.ListRoutingProfilesOutput), req.Error
}

func (c *Client) ListSecurityProfilesWithContext(ctx context.Context, input *connect.ListSecurityProfilesInput, opts ...request.Option) (*connect.ListSecurityProfilesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "ListSecurityProfiles",
		Input:   input,
		Output:  (*connect.ListSecurityProfilesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.ListSecurityProfilesWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.ListSecurityProfilesOutput), req.Error
}

func (c *Client) ListUserHierarchyGroupsWithContext(ctx context.Context, input *connect.ListUserHierarchyGroupsInput, opts ...request.Option) (*connect.ListUserHierarchyGroupsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "ListUserHierarchyGroups",
		Input:   input,
		Output:  (*connect.ListUserHierarchyGroupsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.ListUserHierarchyGroupsWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.ListUserHierarchyGroupsOutput), req.Error
}

func (c *Client) ListUsersWithContext(ctx context.Context, input *connect.ListUsersInput, opts ...request.Option) (*connect.ListUsersOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "ListUsers",
		Input:   input,
		Output:  (*connect.ListUsersOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.ListUsersWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.ListUsersOutput), req.Error
}

func (c *Client) StartOutboundVoiceContactWithContext(ctx context.Context, input *connect.StartOutboundVoiceContactInput, opts ...request.Option) (*connect.StartOutboundVoiceContactOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "StartOutboundVoiceContact",
		Input:   input,
		Output:  (*connect.StartOutboundVoiceContactOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.StartOutboundVoiceContactWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.StartOutboundVoiceContactOutput), req.Error
}

func (c *Client) StopContactWithContext(ctx context.Context, input *connect.StopContactInput, opts ...request.Option) (*connect.StopContactOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "StopContact",
		Input:   input,
		Output:  (*connect.StopContactOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.StopContactWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.StopContactOutput), req.Error
}

func (c *Client) UpdateContactAttributesWithContext(ctx context.Context, input *connect.UpdateContactAttributesInput, opts ...request.Option) (*connect.UpdateContactAttributesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "UpdateContactAttributes",
		Input:   input,
		Output:  (*connect.UpdateContactAttributesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.UpdateContactAttributesWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.UpdateContactAttributesOutput), req.Error
}

func (c *Client) UpdateUserHierarchyWithContext(ctx context.Context, input *connect.UpdateUserHierarchyInput, opts ...request.Option) (*connect.UpdateUserHierarchyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "UpdateUserHierarchy",
		Input:   input,
		Output:  (*connect.UpdateUserHierarchyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.UpdateUserHierarchyWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.UpdateUserHierarchyOutput), req.Error
}

func (c *Client) UpdateUserIdentityInfoWithContext(ctx context.Context, input *connect.UpdateUserIdentityInfoInput, opts ...request.Option) (*connect.UpdateUserIdentityInfoOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "UpdateUserIdentityInfo",
		Input:   input,
		Output:  (*connect.UpdateUserIdentityInfoOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.UpdateUserIdentityInfoWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.UpdateUserIdentityInfoOutput), req.Error
}

func (c *Client) UpdateUserPhoneConfigWithContext(ctx context.Context, input *connect.UpdateUserPhoneConfigInput, opts ...request.Option) (*connect.UpdateUserPhoneConfigOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "UpdateUserPhoneConfig",
		Input:   input,
		Output:  (*connect.UpdateUserPhoneConfigOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.UpdateUserPhoneConfigWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.UpdateUserPhoneConfigOutput), req.Error
}

func (c *Client) UpdateUserRoutingProfileWithContext(ctx context.Context, input *connect.UpdateUserRoutingProfileInput, opts ...request.Option) (*connect.UpdateUserRoutingProfileOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "UpdateUserRoutingProfile",
		Input:   input,
		Output:  (*connect.UpdateUserRoutingProfileOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.UpdateUserRoutingProfileWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.UpdateUserRoutingProfileOutput), req.Error
}

func (c *Client) UpdateUserSecurityProfilesWithContext(ctx context.Context, input *connect.UpdateUserSecurityProfilesInput, opts ...request.Option) (*connect.UpdateUserSecurityProfilesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "connect",
		Action:  "UpdateUserSecurityProfiles",
		Input:   input,
		Output:  (*connect.UpdateUserSecurityProfilesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ConnectAPI.UpdateUserSecurityProfilesWithContext(ctx, input, opts...)
	})

	return req.Output.(*connect.UpdateUserSecurityProfilesOutput), req.Error
}
