// Code generated by internal/generate/main.go. DO NOT EDIT.

package serverlessapplicationrepositoryctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go/service/serverlessapplicationrepository/serverlessapplicationrepositoryiface"
	"github.com/glassechidna/awsctx"
)

type ServerlessApplicationRepository interface {
	CreateApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.CreateApplicationRequest, opts ...request.Option) (*serverlessapplicationrepository.CreateApplicationOutput, error)
	CreateApplicationVersionWithContext(ctx context.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest, opts ...request.Option) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error)
	CreateCloudFormationChangeSetWithContext(ctx context.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest, opts ...request.Option) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error)
	CreateCloudFormationTemplateWithContext(ctx context.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput, opts ...request.Option) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error)
	DeleteApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.DeleteApplicationInput, opts ...request.Option) (*serverlessapplicationrepository.DeleteApplicationOutput, error)
	GetApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.GetApplicationInput, opts ...request.Option) (*serverlessapplicationrepository.GetApplicationOutput, error)
	GetApplicationPolicyWithContext(ctx context.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput, opts ...request.Option) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error)
	GetCloudFormationTemplateWithContext(ctx context.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput, opts ...request.Option) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error)
	ListApplicationDependenciesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error)
	ListApplicationDependenciesPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput, cb func(*serverlessapplicationrepository.ListApplicationDependenciesOutput, bool) bool, opts ...request.Option) error
	ListApplicationVersionsWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error)
	ListApplicationVersionsPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput, cb func(*serverlessapplicationrepository.ListApplicationVersionsOutput, bool) bool, opts ...request.Option) error
	ListApplicationsWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationsInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationsOutput, error)
	ListApplicationsPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationsInput, cb func(*serverlessapplicationrepository.ListApplicationsOutput, bool) bool, opts ...request.Option) error
	PutApplicationPolicyWithContext(ctx context.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput, opts ...request.Option) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error)
	UnshareApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.UnshareApplicationInput, opts ...request.Option) (*serverlessapplicationrepository.UnshareApplicationOutput, error)
	UpdateApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.UpdateApplicationRequest, opts ...request.Option) (*serverlessapplicationrepository.UpdateApplicationOutput, error)
}

type Client struct {
	serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI
	Contexter awsctx.Contexter
}

func New(base serverlessapplicationrepositoryiface.ServerlessApplicationRepositoryAPI, ctxer awsctx.Contexter) ServerlessApplicationRepository {
	return &Client{
		ServerlessApplicationRepositoryAPI: base,
		Contexter: ctxer,
	}
}

var _ ServerlessApplicationRepository = (*serverlessapplicationrepository.ServerlessApplicationRepository)(nil)
var _ ServerlessApplicationRepository = (*Client)(nil)

func (c *Client) CreateApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.CreateApplicationRequest, opts ...request.Option) (*serverlessapplicationrepository.CreateApplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "CreateApplication",
		Input:   input,
		Output:  (*serverlessapplicationrepository.CreateApplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.CreateApplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.CreateApplicationOutput), req.Error
}

func (c *Client) CreateApplicationVersionWithContext(ctx context.Context, input *serverlessapplicationrepository.CreateApplicationVersionRequest, opts ...request.Option) (*serverlessapplicationrepository.CreateApplicationVersionOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "CreateApplicationVersion",
		Input:   input,
		Output:  (*serverlessapplicationrepository.CreateApplicationVersionOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.CreateApplicationVersionWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.CreateApplicationVersionOutput), req.Error
}

func (c *Client) CreateCloudFormationChangeSetWithContext(ctx context.Context, input *serverlessapplicationrepository.CreateCloudFormationChangeSetRequest, opts ...request.Option) (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "CreateCloudFormationChangeSet",
		Input:   input,
		Output:  (*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.CreateCloudFormationChangeSetWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.CreateCloudFormationChangeSetOutput), req.Error
}

func (c *Client) CreateCloudFormationTemplateWithContext(ctx context.Context, input *serverlessapplicationrepository.CreateCloudFormationTemplateInput, opts ...request.Option) (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "CreateCloudFormationTemplate",
		Input:   input,
		Output:  (*serverlessapplicationrepository.CreateCloudFormationTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.CreateCloudFormationTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.CreateCloudFormationTemplateOutput), req.Error
}

func (c *Client) DeleteApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.DeleteApplicationInput, opts ...request.Option) (*serverlessapplicationrepository.DeleteApplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "DeleteApplication",
		Input:   input,
		Output:  (*serverlessapplicationrepository.DeleteApplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.DeleteApplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.DeleteApplicationOutput), req.Error
}

func (c *Client) GetApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.GetApplicationInput, opts ...request.Option) (*serverlessapplicationrepository.GetApplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "GetApplication",
		Input:   input,
		Output:  (*serverlessapplicationrepository.GetApplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.GetApplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.GetApplicationOutput), req.Error
}

func (c *Client) GetApplicationPolicyWithContext(ctx context.Context, input *serverlessapplicationrepository.GetApplicationPolicyInput, opts ...request.Option) (*serverlessapplicationrepository.GetApplicationPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "GetApplicationPolicy",
		Input:   input,
		Output:  (*serverlessapplicationrepository.GetApplicationPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.GetApplicationPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.GetApplicationPolicyOutput), req.Error
}

func (c *Client) GetCloudFormationTemplateWithContext(ctx context.Context, input *serverlessapplicationrepository.GetCloudFormationTemplateInput, opts ...request.Option) (*serverlessapplicationrepository.GetCloudFormationTemplateOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "GetCloudFormationTemplate",
		Input:   input,
		Output:  (*serverlessapplicationrepository.GetCloudFormationTemplateOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.GetCloudFormationTemplateWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.GetCloudFormationTemplateOutput), req.Error
}

func (c *Client) ListApplicationDependenciesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationDependenciesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "ListApplicationDependencies",
		Input:   input,
		Output:  (*serverlessapplicationrepository.ListApplicationDependenciesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.ListApplicationDependenciesWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.ListApplicationDependenciesOutput), req.Error
}

func (c *Client) ListApplicationDependenciesPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationDependenciesInput, cb func(*serverlessapplicationrepository.ListApplicationDependenciesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "ListApplicationDependencies",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServerlessApplicationRepositoryAPI.ListApplicationDependenciesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListApplicationVersionsWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationVersionsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "ListApplicationVersions",
		Input:   input,
		Output:  (*serverlessapplicationrepository.ListApplicationVersionsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.ListApplicationVersionsWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.ListApplicationVersionsOutput), req.Error
}

func (c *Client) ListApplicationVersionsPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationVersionsInput, cb func(*serverlessapplicationrepository.ListApplicationVersionsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "ListApplicationVersions",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServerlessApplicationRepositoryAPI.ListApplicationVersionsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListApplicationsWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationsInput, opts ...request.Option) (*serverlessapplicationrepository.ListApplicationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "ListApplications",
		Input:   input,
		Output:  (*serverlessapplicationrepository.ListApplicationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.ListApplicationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.ListApplicationsOutput), req.Error
}

func (c *Client) ListApplicationsPagesWithContext(ctx context.Context, input *serverlessapplicationrepository.ListApplicationsInput, cb func(*serverlessapplicationrepository.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "ListApplications",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ServerlessApplicationRepositoryAPI.ListApplicationsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) PutApplicationPolicyWithContext(ctx context.Context, input *serverlessapplicationrepository.PutApplicationPolicyInput, opts ...request.Option) (*serverlessapplicationrepository.PutApplicationPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "PutApplicationPolicy",
		Input:   input,
		Output:  (*serverlessapplicationrepository.PutApplicationPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.PutApplicationPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.PutApplicationPolicyOutput), req.Error
}

func (c *Client) UnshareApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.UnshareApplicationInput, opts ...request.Option) (*serverlessapplicationrepository.UnshareApplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "UnshareApplication",
		Input:   input,
		Output:  (*serverlessapplicationrepository.UnshareApplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.UnshareApplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.UnshareApplicationOutput), req.Error
}

func (c *Client) UpdateApplicationWithContext(ctx context.Context, input *serverlessapplicationrepository.UpdateApplicationRequest, opts ...request.Option) (*serverlessapplicationrepository.UpdateApplicationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "serverlessapplicationrepository",
		Action:  "UpdateApplication",
		Input:   input,
		Output:  (*serverlessapplicationrepository.UpdateApplicationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ServerlessApplicationRepositoryAPI.UpdateApplicationWithContext(ctx, input, opts...)
	})

	return req.Output.(*serverlessapplicationrepository.UpdateApplicationOutput), req.Error
}
