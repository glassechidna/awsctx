// Code generated by internal/generate/main.go. DO NOT EDIT.

package codebuildctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"
	"github.com/glassechidna/awsctx"
)

type CodeBuild interface {
	BatchDeleteBuildsWithContext(ctx context.Context, input *codebuild.BatchDeleteBuildsInput, opts ...request.Option) (*codebuild.BatchDeleteBuildsOutput, error)
	BatchGetBuildsWithContext(ctx context.Context, input *codebuild.BatchGetBuildsInput, opts ...request.Option) (*codebuild.BatchGetBuildsOutput, error)
	BatchGetProjectsWithContext(ctx context.Context, input *codebuild.BatchGetProjectsInput, opts ...request.Option) (*codebuild.BatchGetProjectsOutput, error)
	CreateProjectWithContext(ctx context.Context, input *codebuild.CreateProjectInput, opts ...request.Option) (*codebuild.CreateProjectOutput, error)
	CreateWebhookWithContext(ctx context.Context, input *codebuild.CreateWebhookInput, opts ...request.Option) (*codebuild.CreateWebhookOutput, error)
	DeleteProjectWithContext(ctx context.Context, input *codebuild.DeleteProjectInput, opts ...request.Option) (*codebuild.DeleteProjectOutput, error)
	DeleteSourceCredentialsWithContext(ctx context.Context, input *codebuild.DeleteSourceCredentialsInput, opts ...request.Option) (*codebuild.DeleteSourceCredentialsOutput, error)
	DeleteWebhookWithContext(ctx context.Context, input *codebuild.DeleteWebhookInput, opts ...request.Option) (*codebuild.DeleteWebhookOutput, error)
	ImportSourceCredentialsWithContext(ctx context.Context, input *codebuild.ImportSourceCredentialsInput, opts ...request.Option) (*codebuild.ImportSourceCredentialsOutput, error)
	InvalidateProjectCacheWithContext(ctx context.Context, input *codebuild.InvalidateProjectCacheInput, opts ...request.Option) (*codebuild.InvalidateProjectCacheOutput, error)
	ListBuildsWithContext(ctx context.Context, input *codebuild.ListBuildsInput, opts ...request.Option) (*codebuild.ListBuildsOutput, error)
	ListBuildsForProjectWithContext(ctx context.Context, input *codebuild.ListBuildsForProjectInput, opts ...request.Option) (*codebuild.ListBuildsForProjectOutput, error)
	ListCuratedEnvironmentImagesWithContext(ctx context.Context, input *codebuild.ListCuratedEnvironmentImagesInput, opts ...request.Option) (*codebuild.ListCuratedEnvironmentImagesOutput, error)
	ListProjectsWithContext(ctx context.Context, input *codebuild.ListProjectsInput, opts ...request.Option) (*codebuild.ListProjectsOutput, error)
	ListSourceCredentialsWithContext(ctx context.Context, input *codebuild.ListSourceCredentialsInput, opts ...request.Option) (*codebuild.ListSourceCredentialsOutput, error)
	StartBuildWithContext(ctx context.Context, input *codebuild.StartBuildInput, opts ...request.Option) (*codebuild.StartBuildOutput, error)
	StopBuildWithContext(ctx context.Context, input *codebuild.StopBuildInput, opts ...request.Option) (*codebuild.StopBuildOutput, error)
	UpdateProjectWithContext(ctx context.Context, input *codebuild.UpdateProjectInput, opts ...request.Option) (*codebuild.UpdateProjectOutput, error)
	UpdateWebhookWithContext(ctx context.Context, input *codebuild.UpdateWebhookInput, opts ...request.Option) (*codebuild.UpdateWebhookOutput, error)
}

type Client struct {
	codebuildiface.CodeBuildAPI
	Contexter awsctx.Contexter
}

func New(base codebuildiface.CodeBuildAPI, ctxer awsctx.Contexter) CodeBuild {
	return &Client{
		CodeBuildAPI: base,
		Contexter: ctxer,
	}
}

var _ CodeBuild = (*codebuild.CodeBuild)(nil)
var _ CodeBuild = (*Client)(nil)

func (c *Client) BatchDeleteBuildsWithContext(ctx context.Context, input *codebuild.BatchDeleteBuildsInput, opts ...request.Option) (*codebuild.BatchDeleteBuildsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "BatchDeleteBuildsWithContext",
		Input:   input,
		Output:  (*codebuild.BatchDeleteBuildsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.BatchDeleteBuildsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.BatchDeleteBuildsOutput), req.Error
}

func (c *Client) BatchGetBuildsWithContext(ctx context.Context, input *codebuild.BatchGetBuildsInput, opts ...request.Option) (*codebuild.BatchGetBuildsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "BatchGetBuildsWithContext",
		Input:   input,
		Output:  (*codebuild.BatchGetBuildsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.BatchGetBuildsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.BatchGetBuildsOutput), req.Error
}

func (c *Client) BatchGetProjectsWithContext(ctx context.Context, input *codebuild.BatchGetProjectsInput, opts ...request.Option) (*codebuild.BatchGetProjectsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "BatchGetProjectsWithContext",
		Input:   input,
		Output:  (*codebuild.BatchGetProjectsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.BatchGetProjectsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.BatchGetProjectsOutput), req.Error
}

func (c *Client) CreateProjectWithContext(ctx context.Context, input *codebuild.CreateProjectInput, opts ...request.Option) (*codebuild.CreateProjectOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "CreateProjectWithContext",
		Input:   input,
		Output:  (*codebuild.CreateProjectOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.CreateProjectWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.CreateProjectOutput), req.Error
}

func (c *Client) CreateWebhookWithContext(ctx context.Context, input *codebuild.CreateWebhookInput, opts ...request.Option) (*codebuild.CreateWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "CreateWebhookWithContext",
		Input:   input,
		Output:  (*codebuild.CreateWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.CreateWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.CreateWebhookOutput), req.Error
}

func (c *Client) DeleteProjectWithContext(ctx context.Context, input *codebuild.DeleteProjectInput, opts ...request.Option) (*codebuild.DeleteProjectOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "DeleteProjectWithContext",
		Input:   input,
		Output:  (*codebuild.DeleteProjectOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.DeleteProjectWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.DeleteProjectOutput), req.Error
}

func (c *Client) DeleteSourceCredentialsWithContext(ctx context.Context, input *codebuild.DeleteSourceCredentialsInput, opts ...request.Option) (*codebuild.DeleteSourceCredentialsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "DeleteSourceCredentialsWithContext",
		Input:   input,
		Output:  (*codebuild.DeleteSourceCredentialsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.DeleteSourceCredentialsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.DeleteSourceCredentialsOutput), req.Error
}

func (c *Client) DeleteWebhookWithContext(ctx context.Context, input *codebuild.DeleteWebhookInput, opts ...request.Option) (*codebuild.DeleteWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "DeleteWebhookWithContext",
		Input:   input,
		Output:  (*codebuild.DeleteWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.DeleteWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.DeleteWebhookOutput), req.Error
}

func (c *Client) ImportSourceCredentialsWithContext(ctx context.Context, input *codebuild.ImportSourceCredentialsInput, opts ...request.Option) (*codebuild.ImportSourceCredentialsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "ImportSourceCredentialsWithContext",
		Input:   input,
		Output:  (*codebuild.ImportSourceCredentialsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.ImportSourceCredentialsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.ImportSourceCredentialsOutput), req.Error
}

func (c *Client) InvalidateProjectCacheWithContext(ctx context.Context, input *codebuild.InvalidateProjectCacheInput, opts ...request.Option) (*codebuild.InvalidateProjectCacheOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "InvalidateProjectCacheWithContext",
		Input:   input,
		Output:  (*codebuild.InvalidateProjectCacheOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.InvalidateProjectCacheWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.InvalidateProjectCacheOutput), req.Error
}

func (c *Client) ListBuildsWithContext(ctx context.Context, input *codebuild.ListBuildsInput, opts ...request.Option) (*codebuild.ListBuildsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "ListBuildsWithContext",
		Input:   input,
		Output:  (*codebuild.ListBuildsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.ListBuildsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.ListBuildsOutput), req.Error
}

func (c *Client) ListBuildsForProjectWithContext(ctx context.Context, input *codebuild.ListBuildsForProjectInput, opts ...request.Option) (*codebuild.ListBuildsForProjectOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "ListBuildsForProjectWithContext",
		Input:   input,
		Output:  (*codebuild.ListBuildsForProjectOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.ListBuildsForProjectWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.ListBuildsForProjectOutput), req.Error
}

func (c *Client) ListCuratedEnvironmentImagesWithContext(ctx context.Context, input *codebuild.ListCuratedEnvironmentImagesInput, opts ...request.Option) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "ListCuratedEnvironmentImagesWithContext",
		Input:   input,
		Output:  (*codebuild.ListCuratedEnvironmentImagesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.ListCuratedEnvironmentImagesWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.ListCuratedEnvironmentImagesOutput), req.Error
}

func (c *Client) ListProjectsWithContext(ctx context.Context, input *codebuild.ListProjectsInput, opts ...request.Option) (*codebuild.ListProjectsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "ListProjectsWithContext",
		Input:   input,
		Output:  (*codebuild.ListProjectsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.ListProjectsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.ListProjectsOutput), req.Error
}

func (c *Client) ListSourceCredentialsWithContext(ctx context.Context, input *codebuild.ListSourceCredentialsInput, opts ...request.Option) (*codebuild.ListSourceCredentialsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "ListSourceCredentialsWithContext",
		Input:   input,
		Output:  (*codebuild.ListSourceCredentialsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.ListSourceCredentialsWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.ListSourceCredentialsOutput), req.Error
}

func (c *Client) StartBuildWithContext(ctx context.Context, input *codebuild.StartBuildInput, opts ...request.Option) (*codebuild.StartBuildOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "StartBuildWithContext",
		Input:   input,
		Output:  (*codebuild.StartBuildOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.StartBuildWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.StartBuildOutput), req.Error
}

func (c *Client) StopBuildWithContext(ctx context.Context, input *codebuild.StopBuildInput, opts ...request.Option) (*codebuild.StopBuildOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "StopBuildWithContext",
		Input:   input,
		Output:  (*codebuild.StopBuildOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.StopBuildWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.StopBuildOutput), req.Error
}

func (c *Client) UpdateProjectWithContext(ctx context.Context, input *codebuild.UpdateProjectInput, opts ...request.Option) (*codebuild.UpdateProjectOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "UpdateProjectWithContext",
		Input:   input,
		Output:  (*codebuild.UpdateProjectOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.UpdateProjectWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.UpdateProjectOutput), req.Error
}

func (c *Client) UpdateWebhookWithContext(ctx context.Context, input *codebuild.UpdateWebhookInput, opts ...request.Option) (*codebuild.UpdateWebhookOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "codebuild",
		Action:  "UpdateWebhookWithContext",
		Input:   input,
		Output:  (*codebuild.UpdateWebhookOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.CodeBuildAPI.UpdateWebhookWithContext(ctx, input, opts...)
	})

	return req.Output.(*codebuild.UpdateWebhookOutput), req.Error
}
