// Code generated by internal/generate/main.go. DO NOT EDIT.

package cognitoidentityctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentity/cognitoidentityiface"
	"github.com/glassechidna/awsctx"
)

type CognitoIdentity interface {
	CreateIdentityPoolWithContext(ctx context.Context, input *cognitoidentity.CreateIdentityPoolInput, opts ...request.Option) (*cognitoidentity.IdentityPool, error)
	DeleteIdentitiesWithContext(ctx context.Context, input *cognitoidentity.DeleteIdentitiesInput, opts ...request.Option) (*cognitoidentity.DeleteIdentitiesOutput, error)
	DeleteIdentityPoolWithContext(ctx context.Context, input *cognitoidentity.DeleteIdentityPoolInput, opts ...request.Option) (*cognitoidentity.DeleteIdentityPoolOutput, error)
	DescribeIdentityWithContext(ctx context.Context, input *cognitoidentity.DescribeIdentityInput, opts ...request.Option) (*cognitoidentity.IdentityDescription, error)
	DescribeIdentityPoolWithContext(ctx context.Context, input *cognitoidentity.DescribeIdentityPoolInput, opts ...request.Option) (*cognitoidentity.IdentityPool, error)
	GetCredentialsForIdentityWithContext(ctx context.Context, input *cognitoidentity.GetCredentialsForIdentityInput, opts ...request.Option) (*cognitoidentity.GetCredentialsForIdentityOutput, error)
	GetIdWithContext(ctx context.Context, input *cognitoidentity.GetIdInput, opts ...request.Option) (*cognitoidentity.GetIdOutput, error)
	GetIdentityPoolRolesWithContext(ctx context.Context, input *cognitoidentity.GetIdentityPoolRolesInput, opts ...request.Option) (*cognitoidentity.GetIdentityPoolRolesOutput, error)
	GetOpenIdTokenWithContext(ctx context.Context, input *cognitoidentity.GetOpenIdTokenInput, opts ...request.Option) (*cognitoidentity.GetOpenIdTokenOutput, error)
	GetOpenIdTokenForDeveloperIdentityWithContext(ctx context.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, opts ...request.Option) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)
	ListIdentitiesWithContext(ctx context.Context, input *cognitoidentity.ListIdentitiesInput, opts ...request.Option) (*cognitoidentity.ListIdentitiesOutput, error)
	ListIdentityPoolsWithContext(ctx context.Context, input *cognitoidentity.ListIdentityPoolsInput, opts ...request.Option) (*cognitoidentity.ListIdentityPoolsOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *cognitoidentity.ListTagsForResourceInput, opts ...request.Option) (*cognitoidentity.ListTagsForResourceOutput, error)
	LookupDeveloperIdentityWithContext(ctx context.Context, input *cognitoidentity.LookupDeveloperIdentityInput, opts ...request.Option) (*cognitoidentity.LookupDeveloperIdentityOutput, error)
	MergeDeveloperIdentitiesWithContext(ctx context.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput, opts ...request.Option) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)
	SetIdentityPoolRolesWithContext(ctx context.Context, input *cognitoidentity.SetIdentityPoolRolesInput, opts ...request.Option) (*cognitoidentity.SetIdentityPoolRolesOutput, error)
	TagResourceWithContext(ctx context.Context, input *cognitoidentity.TagResourceInput, opts ...request.Option) (*cognitoidentity.TagResourceOutput, error)
	UnlinkDeveloperIdentityWithContext(ctx context.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput, opts ...request.Option) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)
	UnlinkIdentityWithContext(ctx context.Context, input *cognitoidentity.UnlinkIdentityInput, opts ...request.Option) (*cognitoidentity.UnlinkIdentityOutput, error)
	UntagResourceWithContext(ctx context.Context, input *cognitoidentity.UntagResourceInput, opts ...request.Option) (*cognitoidentity.UntagResourceOutput, error)
	UpdateIdentityPoolWithContext(ctx context.Context, input *cognitoidentity.IdentityPool, opts ...request.Option) (*cognitoidentity.IdentityPool, error)
}

type Client struct {
	cognitoidentityiface.CognitoIdentityAPI
	Contexter awsctx.Contexter
}

func New(base cognitoidentityiface.CognitoIdentityAPI, ctxer awsctx.Contexter) CognitoIdentity {
	return &Client{
		CognitoIdentityAPI: base,
		Contexter: ctxer,
	}
}

var _ CognitoIdentity = (*cognitoidentity.CognitoIdentity)(nil)
var _ CognitoIdentity = (*Client)(nil)

func (c *Client) CreateIdentityPoolWithContext(ctx context.Context, input *cognitoidentity.CreateIdentityPoolInput, opts ...request.Option) (*cognitoidentity.IdentityPool, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "CreateIdentityPool",
		Input:   input,
		Output:  (*cognitoidentity.IdentityPool)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.CreateIdentityPoolWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.IdentityPool), req.Error
}

func (c *Client) DeleteIdentitiesWithContext(ctx context.Context, input *cognitoidentity.DeleteIdentitiesInput, opts ...request.Option) (*cognitoidentity.DeleteIdentitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "DeleteIdentities",
		Input:   input,
		Output:  (*cognitoidentity.DeleteIdentitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.DeleteIdentitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.DeleteIdentitiesOutput), req.Error
}

func (c *Client) DeleteIdentityPoolWithContext(ctx context.Context, input *cognitoidentity.DeleteIdentityPoolInput, opts ...request.Option) (*cognitoidentity.DeleteIdentityPoolOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "DeleteIdentityPool",
		Input:   input,
		Output:  (*cognitoidentity.DeleteIdentityPoolOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.DeleteIdentityPoolWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.DeleteIdentityPoolOutput), req.Error
}

func (c *Client) DescribeIdentityWithContext(ctx context.Context, input *cognitoidentity.DescribeIdentityInput, opts ...request.Option) (*cognitoidentity.IdentityDescription, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "DescribeIdentity",
		Input:   input,
		Output:  (*cognitoidentity.IdentityDescription)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.DescribeIdentityWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.IdentityDescription), req.Error
}

func (c *Client) DescribeIdentityPoolWithContext(ctx context.Context, input *cognitoidentity.DescribeIdentityPoolInput, opts ...request.Option) (*cognitoidentity.IdentityPool, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "DescribeIdentityPool",
		Input:   input,
		Output:  (*cognitoidentity.IdentityPool)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.DescribeIdentityPoolWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.IdentityPool), req.Error
}

func (c *Client) GetCredentialsForIdentityWithContext(ctx context.Context, input *cognitoidentity.GetCredentialsForIdentityInput, opts ...request.Option) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "GetCredentialsForIdentity",
		Input:   input,
		Output:  (*cognitoidentity.GetCredentialsForIdentityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.GetCredentialsForIdentityWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.GetCredentialsForIdentityOutput), req.Error
}

func (c *Client) GetIdWithContext(ctx context.Context, input *cognitoidentity.GetIdInput, opts ...request.Option) (*cognitoidentity.GetIdOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "GetId",
		Input:   input,
		Output:  (*cognitoidentity.GetIdOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.GetIdWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.GetIdOutput), req.Error
}

func (c *Client) GetIdentityPoolRolesWithContext(ctx context.Context, input *cognitoidentity.GetIdentityPoolRolesInput, opts ...request.Option) (*cognitoidentity.GetIdentityPoolRolesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "GetIdentityPoolRoles",
		Input:   input,
		Output:  (*cognitoidentity.GetIdentityPoolRolesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.GetIdentityPoolRolesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.GetIdentityPoolRolesOutput), req.Error
}

func (c *Client) GetOpenIdTokenWithContext(ctx context.Context, input *cognitoidentity.GetOpenIdTokenInput, opts ...request.Option) (*cognitoidentity.GetOpenIdTokenOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "GetOpenIdToken",
		Input:   input,
		Output:  (*cognitoidentity.GetOpenIdTokenOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.GetOpenIdTokenWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.GetOpenIdTokenOutput), req.Error
}

func (c *Client) GetOpenIdTokenForDeveloperIdentityWithContext(ctx context.Context, input *cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput, opts ...request.Option) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "GetOpenIdTokenForDeveloperIdentity",
		Input:   input,
		Output:  (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.GetOpenIdTokenForDeveloperIdentityWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput), req.Error
}

func (c *Client) ListIdentitiesWithContext(ctx context.Context, input *cognitoidentity.ListIdentitiesInput, opts ...request.Option) (*cognitoidentity.ListIdentitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "ListIdentities",
		Input:   input,
		Output:  (*cognitoidentity.ListIdentitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.ListIdentitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.ListIdentitiesOutput), req.Error
}

func (c *Client) ListIdentityPoolsWithContext(ctx context.Context, input *cognitoidentity.ListIdentityPoolsInput, opts ...request.Option) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "ListIdentityPools",
		Input:   input,
		Output:  (*cognitoidentity.ListIdentityPoolsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.ListIdentityPoolsWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.ListIdentityPoolsOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *cognitoidentity.ListTagsForResourceInput, opts ...request.Option) (*cognitoidentity.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*cognitoidentity.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.ListTagsForResourceOutput), req.Error
}

func (c *Client) LookupDeveloperIdentityWithContext(ctx context.Context, input *cognitoidentity.LookupDeveloperIdentityInput, opts ...request.Option) (*cognitoidentity.LookupDeveloperIdentityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "LookupDeveloperIdentity",
		Input:   input,
		Output:  (*cognitoidentity.LookupDeveloperIdentityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.LookupDeveloperIdentityWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.LookupDeveloperIdentityOutput), req.Error
}

func (c *Client) MergeDeveloperIdentitiesWithContext(ctx context.Context, input *cognitoidentity.MergeDeveloperIdentitiesInput, opts ...request.Option) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "MergeDeveloperIdentities",
		Input:   input,
		Output:  (*cognitoidentity.MergeDeveloperIdentitiesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.MergeDeveloperIdentitiesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.MergeDeveloperIdentitiesOutput), req.Error
}

func (c *Client) SetIdentityPoolRolesWithContext(ctx context.Context, input *cognitoidentity.SetIdentityPoolRolesInput, opts ...request.Option) (*cognitoidentity.SetIdentityPoolRolesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "SetIdentityPoolRoles",
		Input:   input,
		Output:  (*cognitoidentity.SetIdentityPoolRolesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.SetIdentityPoolRolesWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.SetIdentityPoolRolesOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *cognitoidentity.TagResourceInput, opts ...request.Option) (*cognitoidentity.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "TagResource",
		Input:   input,
		Output:  (*cognitoidentity.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.TagResourceOutput), req.Error
}

func (c *Client) UnlinkDeveloperIdentityWithContext(ctx context.Context, input *cognitoidentity.UnlinkDeveloperIdentityInput, opts ...request.Option) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "UnlinkDeveloperIdentity",
		Input:   input,
		Output:  (*cognitoidentity.UnlinkDeveloperIdentityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.UnlinkDeveloperIdentityWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.UnlinkDeveloperIdentityOutput), req.Error
}

func (c *Client) UnlinkIdentityWithContext(ctx context.Context, input *cognitoidentity.UnlinkIdentityInput, opts ...request.Option) (*cognitoidentity.UnlinkIdentityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "UnlinkIdentity",
		Input:   input,
		Output:  (*cognitoidentity.UnlinkIdentityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.UnlinkIdentityWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.UnlinkIdentityOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *cognitoidentity.UntagResourceInput, opts ...request.Option) (*cognitoidentity.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*cognitoidentity.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.UntagResourceOutput), req.Error
}

func (c *Client) UpdateIdentityPoolWithContext(ctx context.Context, input *cognitoidentity.IdentityPool, opts ...request.Option) (*cognitoidentity.IdentityPool, error) {
	req := &awsctx.AwsRequest{
		Service: "cognitoidentity",
		Action:  "UpdateIdentityPool",
		Input:   input,
		Output:  (*cognitoidentity.IdentityPool)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CognitoIdentityAPI.UpdateIdentityPoolWithContext(ctx, input, opts...)
	})

	return req.Output.(*cognitoidentity.IdentityPool), req.Error
}
