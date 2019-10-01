// Code generated by internal/generate/main.go. DO NOT EDIT.

package amplifyctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"github.com/glassechidna/awsctx"
)

type Amplify interface {
	CreateAppWithContext(ctx context.Context, input *amplify.CreateAppInput, opts ...request.Option) (*amplify.CreateAppOutput, error)
	CreateBranchWithContext(ctx context.Context, input *amplify.CreateBranchInput, opts ...request.Option) (*amplify.CreateBranchOutput, error)
	CreateDeploymentWithContext(ctx context.Context, input *amplify.CreateDeploymentInput, opts ...request.Option) (*amplify.CreateDeploymentOutput, error)
	CreateDomainAssociationWithContext(ctx context.Context, input *amplify.CreateDomainAssociationInput, opts ...request.Option) (*amplify.CreateDomainAssociationOutput, error)
	CreateWebhookWithContext(ctx context.Context, input *amplify.CreateWebhookInput, opts ...request.Option) (*amplify.CreateWebhookOutput, error)
	DeleteAppWithContext(ctx context.Context, input *amplify.DeleteAppInput, opts ...request.Option) (*amplify.DeleteAppOutput, error)
	DeleteBranchWithContext(ctx context.Context, input *amplify.DeleteBranchInput, opts ...request.Option) (*amplify.DeleteBranchOutput, error)
	DeleteDomainAssociationWithContext(ctx context.Context, input *amplify.DeleteDomainAssociationInput, opts ...request.Option) (*amplify.DeleteDomainAssociationOutput, error)
	DeleteJobWithContext(ctx context.Context, input *amplify.DeleteJobInput, opts ...request.Option) (*amplify.DeleteJobOutput, error)
	DeleteWebhookWithContext(ctx context.Context, input *amplify.DeleteWebhookInput, opts ...request.Option) (*amplify.DeleteWebhookOutput, error)
	GetAppWithContext(ctx context.Context, input *amplify.GetAppInput, opts ...request.Option) (*amplify.GetAppOutput, error)
	GetBranchWithContext(ctx context.Context, input *amplify.GetBranchInput, opts ...request.Option) (*amplify.GetBranchOutput, error)
	GetDomainAssociationWithContext(ctx context.Context, input *amplify.GetDomainAssociationInput, opts ...request.Option) (*amplify.GetDomainAssociationOutput, error)
	GetJobWithContext(ctx context.Context, input *amplify.GetJobInput, opts ...request.Option) (*amplify.GetJobOutput, error)
	GetWebhookWithContext(ctx context.Context, input *amplify.GetWebhookInput, opts ...request.Option) (*amplify.GetWebhookOutput, error)
	ListAppsWithContext(ctx context.Context, input *amplify.ListAppsInput, opts ...request.Option) (*amplify.ListAppsOutput, error)
	ListBranchesWithContext(ctx context.Context, input *amplify.ListBranchesInput, opts ...request.Option) (*amplify.ListBranchesOutput, error)
	ListDomainAssociationsWithContext(ctx context.Context, input *amplify.ListDomainAssociationsInput, opts ...request.Option) (*amplify.ListDomainAssociationsOutput, error)
	ListJobsWithContext(ctx context.Context, input *amplify.ListJobsInput, opts ...request.Option) (*amplify.ListJobsOutput, error)
	ListTagsForResourceWithContext(ctx context.Context, input *amplify.ListTagsForResourceInput, opts ...request.Option) (*amplify.ListTagsForResourceOutput, error)
	ListWebhooksWithContext(ctx context.Context, input *amplify.ListWebhooksInput, opts ...request.Option) (*amplify.ListWebhooksOutput, error)
	StartDeploymentWithContext(ctx context.Context, input *amplify.StartDeploymentInput, opts ...request.Option) (*amplify.StartDeploymentOutput, error)
	StartJobWithContext(ctx context.Context, input *amplify.StartJobInput, opts ...request.Option) (*amplify.StartJobOutput, error)
	StopJobWithContext(ctx context.Context, input *amplify.StopJobInput, opts ...request.Option) (*amplify.StopJobOutput, error)
	TagResourceWithContext(ctx context.Context, input *amplify.TagResourceInput, opts ...request.Option) (*amplify.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *amplify.UntagResourceInput, opts ...request.Option) (*amplify.UntagResourceOutput, error)
	UpdateAppWithContext(ctx context.Context, input *amplify.UpdateAppInput, opts ...request.Option) (*amplify.UpdateAppOutput, error)
	UpdateBranchWithContext(ctx context.Context, input *amplify.UpdateBranchInput, opts ...request.Option) (*amplify.UpdateBranchOutput, error)
	UpdateDomainAssociationWithContext(ctx context.Context, input *amplify.UpdateDomainAssociationInput, opts ...request.Option) (*amplify.UpdateDomainAssociationOutput, error)
	UpdateWebhookWithContext(ctx context.Context, input *amplify.UpdateWebhookInput, opts ...request.Option) (*amplify.UpdateWebhookOutput, error)
}

type Client struct {
	amplifyiface.AmplifyAPI
	Contexter awsctx.Contexter
}

func New(base amplifyiface.AmplifyAPI, ctxer awsctx.Contexter) Amplify {
	return &Client{
		AmplifyAPI: base,
		Contexter: ctxer,
	}
}

var _ Amplify = (*amplify.Amplify)(nil)
var _ Amplify = (*Client)(nil)

func (c *Client) CreateAppWithContext(ctx context.Context, input *amplify.CreateAppInput, opts ...request.Option) (*amplify.CreateAppOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "CreateAppWithContext",
		Input:   input,
		Output:  (*amplify.CreateAppOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.CreateAppWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.CreateAppOutput), req.Error
}

func (c *Client) CreateBranchWithContext(ctx context.Context, input *amplify.CreateBranchInput, opts ...request.Option) (*amplify.CreateBranchOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "CreateBranchWithContext",
		Input:   input,
		Output:  (*amplify.CreateBranchOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.CreateBranchWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.CreateBranchOutput), req.Error
}

func (c *Client) CreateDeploymentWithContext(ctx context.Context, input *amplify.CreateDeploymentInput, opts ...request.Option) (*amplify.CreateDeploymentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "CreateDeploymentWithContext",
		Input:   input,
		Output:  (*amplify.CreateDeploymentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.CreateDeploymentWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.CreateDeploymentOutput), req.Error
}

func (c *Client) CreateDomainAssociationWithContext(ctx context.Context, input *amplify.CreateDomainAssociationInput, opts ...request.Option) (*amplify.CreateDomainAssociationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "CreateDomainAssociationWithContext",
		Input:   input,
		Output:  (*amplify.CreateDomainAssociationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.CreateDomainAssociationWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.CreateDomainAssociationOutput), req.Error
}

func (c *Client) CreateWebhookWithContext(ctx context.Context, input *amplify.CreateWebhookInput, opts ...request.Option) (*amplify.CreateWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "CreateWebhookWithContext",
		Input:   input,
		Output:  (*amplify.CreateWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.CreateWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.CreateWebhookOutput), req.Error
}

func (c *Client) DeleteAppWithContext(ctx context.Context, input *amplify.DeleteAppInput, opts ...request.Option) (*amplify.DeleteAppOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "DeleteAppWithContext",
		Input:   input,
		Output:  (*amplify.DeleteAppOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.DeleteAppWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.DeleteAppOutput), req.Error
}

func (c *Client) DeleteBranchWithContext(ctx context.Context, input *amplify.DeleteBranchInput, opts ...request.Option) (*amplify.DeleteBranchOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "DeleteBranchWithContext",
		Input:   input,
		Output:  (*amplify.DeleteBranchOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.DeleteBranchWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.DeleteBranchOutput), req.Error
}

func (c *Client) DeleteDomainAssociationWithContext(ctx context.Context, input *amplify.DeleteDomainAssociationInput, opts ...request.Option) (*amplify.DeleteDomainAssociationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "DeleteDomainAssociationWithContext",
		Input:   input,
		Output:  (*amplify.DeleteDomainAssociationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.DeleteDomainAssociationWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.DeleteDomainAssociationOutput), req.Error
}

func (c *Client) DeleteJobWithContext(ctx context.Context, input *amplify.DeleteJobInput, opts ...request.Option) (*amplify.DeleteJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "DeleteJobWithContext",
		Input:   input,
		Output:  (*amplify.DeleteJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.DeleteJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.DeleteJobOutput), req.Error
}

func (c *Client) DeleteWebhookWithContext(ctx context.Context, input *amplify.DeleteWebhookInput, opts ...request.Option) (*amplify.DeleteWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "DeleteWebhookWithContext",
		Input:   input,
		Output:  (*amplify.DeleteWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.DeleteWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.DeleteWebhookOutput), req.Error
}

func (c *Client) GetAppWithContext(ctx context.Context, input *amplify.GetAppInput, opts ...request.Option) (*amplify.GetAppOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "GetAppWithContext",
		Input:   input,
		Output:  (*amplify.GetAppOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.GetAppWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.GetAppOutput), req.Error
}

func (c *Client) GetBranchWithContext(ctx context.Context, input *amplify.GetBranchInput, opts ...request.Option) (*amplify.GetBranchOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "GetBranchWithContext",
		Input:   input,
		Output:  (*amplify.GetBranchOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.GetBranchWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.GetBranchOutput), req.Error
}

func (c *Client) GetDomainAssociationWithContext(ctx context.Context, input *amplify.GetDomainAssociationInput, opts ...request.Option) (*amplify.GetDomainAssociationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "GetDomainAssociationWithContext",
		Input:   input,
		Output:  (*amplify.GetDomainAssociationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.GetDomainAssociationWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.GetDomainAssociationOutput), req.Error
}

func (c *Client) GetJobWithContext(ctx context.Context, input *amplify.GetJobInput, opts ...request.Option) (*amplify.GetJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "GetJobWithContext",
		Input:   input,
		Output:  (*amplify.GetJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.GetJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.GetJobOutput), req.Error
}

func (c *Client) GetWebhookWithContext(ctx context.Context, input *amplify.GetWebhookInput, opts ...request.Option) (*amplify.GetWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "GetWebhookWithContext",
		Input:   input,
		Output:  (*amplify.GetWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.GetWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.GetWebhookOutput), req.Error
}

func (c *Client) ListAppsWithContext(ctx context.Context, input *amplify.ListAppsInput, opts ...request.Option) (*amplify.ListAppsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "ListAppsWithContext",
		Input:   input,
		Output:  (*amplify.ListAppsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.ListAppsWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.ListAppsOutput), req.Error
}

func (c *Client) ListBranchesWithContext(ctx context.Context, input *amplify.ListBranchesInput, opts ...request.Option) (*amplify.ListBranchesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "ListBranchesWithContext",
		Input:   input,
		Output:  (*amplify.ListBranchesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.ListBranchesWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.ListBranchesOutput), req.Error
}

func (c *Client) ListDomainAssociationsWithContext(ctx context.Context, input *amplify.ListDomainAssociationsInput, opts ...request.Option) (*amplify.ListDomainAssociationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "ListDomainAssociationsWithContext",
		Input:   input,
		Output:  (*amplify.ListDomainAssociationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.ListDomainAssociationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.ListDomainAssociationsOutput), req.Error
}

func (c *Client) ListJobsWithContext(ctx context.Context, input *amplify.ListJobsInput, opts ...request.Option) (*amplify.ListJobsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "ListJobsWithContext",
		Input:   input,
		Output:  (*amplify.ListJobsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.ListJobsWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.ListJobsOutput), req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *amplify.ListTagsForResourceInput, opts ...request.Option) (*amplify.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "ListTagsForResourceWithContext",
		Input:   input,
		Output:  (*amplify.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.ListTagsForResourceOutput), req.Error
}

func (c *Client) ListWebhooksWithContext(ctx context.Context, input *amplify.ListWebhooksInput, opts ...request.Option) (*amplify.ListWebhooksOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "ListWebhooksWithContext",
		Input:   input,
		Output:  (*amplify.ListWebhooksOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.ListWebhooksWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.ListWebhooksOutput), req.Error
}

func (c *Client) StartDeploymentWithContext(ctx context.Context, input *amplify.StartDeploymentInput, opts ...request.Option) (*amplify.StartDeploymentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "StartDeploymentWithContext",
		Input:   input,
		Output:  (*amplify.StartDeploymentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.StartDeploymentWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.StartDeploymentOutput), req.Error
}

func (c *Client) StartJobWithContext(ctx context.Context, input *amplify.StartJobInput, opts ...request.Option) (*amplify.StartJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "StartJobWithContext",
		Input:   input,
		Output:  (*amplify.StartJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.StartJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.StartJobOutput), req.Error
}

func (c *Client) StopJobWithContext(ctx context.Context, input *amplify.StopJobInput, opts ...request.Option) (*amplify.StopJobOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "StopJobWithContext",
		Input:   input,
		Output:  (*amplify.StopJobOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.StopJobWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.StopJobOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *amplify.TagResourceInput, opts ...request.Option) (*amplify.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "TagResourceWithContext",
		Input:   input,
		Output:  (*amplify.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *amplify.UntagResourceInput, opts ...request.Option) (*amplify.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "UntagResourceWithContext",
		Input:   input,
		Output:  (*amplify.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.UntagResourceOutput), req.Error
}

func (c *Client) UpdateAppWithContext(ctx context.Context, input *amplify.UpdateAppInput, opts ...request.Option) (*amplify.UpdateAppOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "UpdateAppWithContext",
		Input:   input,
		Output:  (*amplify.UpdateAppOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.UpdateAppWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.UpdateAppOutput), req.Error
}

func (c *Client) UpdateBranchWithContext(ctx context.Context, input *amplify.UpdateBranchInput, opts ...request.Option) (*amplify.UpdateBranchOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "UpdateBranchWithContext",
		Input:   input,
		Output:  (*amplify.UpdateBranchOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.UpdateBranchWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.UpdateBranchOutput), req.Error
}

func (c *Client) UpdateDomainAssociationWithContext(ctx context.Context, input *amplify.UpdateDomainAssociationInput, opts ...request.Option) (*amplify.UpdateDomainAssociationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "UpdateDomainAssociationWithContext",
		Input:   input,
		Output:  (*amplify.UpdateDomainAssociationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.UpdateDomainAssociationWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.UpdateDomainAssociationOutput), req.Error
}

func (c *Client) UpdateWebhookWithContext(ctx context.Context, input *amplify.UpdateWebhookInput, opts ...request.Option) (*amplify.UpdateWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "amplify",
		Action:  "UpdateWebhookWithContext",
		Input:   input,
		Output:  (*amplify.UpdateWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.AmplifyAPI.UpdateWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*amplify.UpdateWebhookOutput), req.Error
}
