// Code generated by internal/generate/main.go. DO NOT EDIT.

package ecrctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/glassechidna/awsctx"
)

type ECR interface {
	BatchCheckLayerAvailabilityWithContext(ctx context.Context, input *ecr.BatchCheckLayerAvailabilityInput, opts ...request.Option) (*ecr.BatchCheckLayerAvailabilityOutput, error)
	BatchDeleteImageWithContext(ctx context.Context, input *ecr.BatchDeleteImageInput, opts ...request.Option) (*ecr.BatchDeleteImageOutput, error)
	BatchGetImageWithContext(ctx context.Context, input *ecr.BatchGetImageInput, opts ...request.Option) (*ecr.BatchGetImageOutput, error)
	CompleteLayerUploadWithContext(ctx context.Context, input *ecr.CompleteLayerUploadInput, opts ...request.Option) (*ecr.CompleteLayerUploadOutput, error)
	CreateRepositoryWithContext(ctx context.Context, input *ecr.CreateRepositoryInput, opts ...request.Option) (*ecr.CreateRepositoryOutput, error)
	DeleteLifecyclePolicyWithContext(ctx context.Context, input *ecr.DeleteLifecyclePolicyInput, opts ...request.Option) (*ecr.DeleteLifecyclePolicyOutput, error)
	DeleteRepositoryWithContext(ctx context.Context, input *ecr.DeleteRepositoryInput, opts ...request.Option) (*ecr.DeleteRepositoryOutput, error)
	DeleteRepositoryPolicyWithContext(ctx context.Context, input *ecr.DeleteRepositoryPolicyInput, opts ...request.Option) (*ecr.DeleteRepositoryPolicyOutput, error)
	DescribeImageScanFindingsWithContext(ctx context.Context, input *ecr.DescribeImageScanFindingsInput, opts ...request.Option) (*ecr.DescribeImageScanFindingsOutput, error)
	DescribeImageScanFindingsPagesWithContext(ctx context.Context, input *ecr.DescribeImageScanFindingsInput, cb func(*ecr.DescribeImageScanFindingsOutput, bool) bool, opts ...request.Option) error
	DescribeImagesWithContext(ctx context.Context, input *ecr.DescribeImagesInput, opts ...request.Option) (*ecr.DescribeImagesOutput, error)
	DescribeImagesPagesWithContext(ctx context.Context, input *ecr.DescribeImagesInput, cb func(*ecr.DescribeImagesOutput, bool) bool, opts ...request.Option) error
	DescribeRepositoriesWithContext(ctx context.Context, input *ecr.DescribeRepositoriesInput, opts ...request.Option) (*ecr.DescribeRepositoriesOutput, error)
	DescribeRepositoriesPagesWithContext(ctx context.Context, input *ecr.DescribeRepositoriesInput, cb func(*ecr.DescribeRepositoriesOutput, bool) bool, opts ...request.Option) error
	GetAuthorizationTokenWithContext(ctx context.Context, input *ecr.GetAuthorizationTokenInput, opts ...request.Option) (*ecr.GetAuthorizationTokenOutput, error)
	GetDownloadUrlForLayerWithContext(ctx context.Context, input *ecr.GetDownloadUrlForLayerInput, opts ...request.Option) (*ecr.GetDownloadUrlForLayerOutput, error)
	GetLifecyclePolicyWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyInput, opts ...request.Option) (*ecr.GetLifecyclePolicyOutput, error)
	GetLifecyclePolicyPreviewWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyPreviewInput, opts ...request.Option) (*ecr.GetLifecyclePolicyPreviewOutput, error)
	GetLifecyclePolicyPreviewPagesWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyPreviewInput, cb func(*ecr.GetLifecyclePolicyPreviewOutput, bool) bool, opts ...request.Option) error
	GetRepositoryPolicyWithContext(ctx context.Context, input *ecr.GetRepositoryPolicyInput, opts ...request.Option) (*ecr.GetRepositoryPolicyOutput, error)
	InitiateLayerUploadWithContext(ctx context.Context, input *ecr.InitiateLayerUploadInput, opts ...request.Option) (*ecr.InitiateLayerUploadOutput, error)
	ListImagesWithContext(ctx context.Context, input *ecr.ListImagesInput, opts ...request.Option) (*ecr.ListImagesOutput, error)
	ListImagesPagesWithContext(ctx context.Context, input *ecr.ListImagesInput, cb func(*ecr.ListImagesOutput, bool) bool, opts ...request.Option) error
	ListTagsForResourceWithContext(ctx context.Context, input *ecr.ListTagsForResourceInput, opts ...request.Option) (*ecr.ListTagsForResourceOutput, error)
	PutImageWithContext(ctx context.Context, input *ecr.PutImageInput, opts ...request.Option) (*ecr.PutImageOutput, error)
	PutImageScanningConfigurationWithContext(ctx context.Context, input *ecr.PutImageScanningConfigurationInput, opts ...request.Option) (*ecr.PutImageScanningConfigurationOutput, error)
	PutImageTagMutabilityWithContext(ctx context.Context, input *ecr.PutImageTagMutabilityInput, opts ...request.Option) (*ecr.PutImageTagMutabilityOutput, error)
	PutLifecyclePolicyWithContext(ctx context.Context, input *ecr.PutLifecyclePolicyInput, opts ...request.Option) (*ecr.PutLifecyclePolicyOutput, error)
	SetRepositoryPolicyWithContext(ctx context.Context, input *ecr.SetRepositoryPolicyInput, opts ...request.Option) (*ecr.SetRepositoryPolicyOutput, error)
	StartImageScanWithContext(ctx context.Context, input *ecr.StartImageScanInput, opts ...request.Option) (*ecr.StartImageScanOutput, error)
	StartLifecyclePolicyPreviewWithContext(ctx context.Context, input *ecr.StartLifecyclePolicyPreviewInput, opts ...request.Option) (*ecr.StartLifecyclePolicyPreviewOutput, error)
	TagResourceWithContext(ctx context.Context, input *ecr.TagResourceInput, opts ...request.Option) (*ecr.TagResourceOutput, error)
	UntagResourceWithContext(ctx context.Context, input *ecr.UntagResourceInput, opts ...request.Option) (*ecr.UntagResourceOutput, error)
	UploadLayerPartWithContext(ctx context.Context, input *ecr.UploadLayerPartInput, opts ...request.Option) (*ecr.UploadLayerPartOutput, error)
}

type Client struct {
	ecriface.ECRAPI
	Contexter awsctx.Contexter
}

func New(base ecriface.ECRAPI, ctxer awsctx.Contexter) ECR {
	return &Client{
		ECRAPI: base,
		Contexter: ctxer,
	}
}

var _ ECR = (*ecr.ECR)(nil)
var _ ECR = (*Client)(nil)

func (c *Client) BatchCheckLayerAvailabilityWithContext(ctx context.Context, input *ecr.BatchCheckLayerAvailabilityInput, opts ...request.Option) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "BatchCheckLayerAvailability",
		Input:   input,
		Output:  (*ecr.BatchCheckLayerAvailabilityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.BatchCheckLayerAvailabilityWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.BatchCheckLayerAvailabilityOutput), req.Error
}

func (c *Client) BatchDeleteImageWithContext(ctx context.Context, input *ecr.BatchDeleteImageInput, opts ...request.Option) (*ecr.BatchDeleteImageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "BatchDeleteImage",
		Input:   input,
		Output:  (*ecr.BatchDeleteImageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.BatchDeleteImageWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.BatchDeleteImageOutput), req.Error
}

func (c *Client) BatchGetImageWithContext(ctx context.Context, input *ecr.BatchGetImageInput, opts ...request.Option) (*ecr.BatchGetImageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "BatchGetImage",
		Input:   input,
		Output:  (*ecr.BatchGetImageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.BatchGetImageWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.BatchGetImageOutput), req.Error
}

func (c *Client) CompleteLayerUploadWithContext(ctx context.Context, input *ecr.CompleteLayerUploadInput, opts ...request.Option) (*ecr.CompleteLayerUploadOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "CompleteLayerUpload",
		Input:   input,
		Output:  (*ecr.CompleteLayerUploadOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.CompleteLayerUploadWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.CompleteLayerUploadOutput), req.Error
}

func (c *Client) CreateRepositoryWithContext(ctx context.Context, input *ecr.CreateRepositoryInput, opts ...request.Option) (*ecr.CreateRepositoryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "CreateRepository",
		Input:   input,
		Output:  (*ecr.CreateRepositoryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.CreateRepositoryWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.CreateRepositoryOutput), req.Error
}

func (c *Client) DeleteLifecyclePolicyWithContext(ctx context.Context, input *ecr.DeleteLifecyclePolicyInput, opts ...request.Option) (*ecr.DeleteLifecyclePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DeleteLifecyclePolicy",
		Input:   input,
		Output:  (*ecr.DeleteLifecyclePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.DeleteLifecyclePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.DeleteLifecyclePolicyOutput), req.Error
}

func (c *Client) DeleteRepositoryWithContext(ctx context.Context, input *ecr.DeleteRepositoryInput, opts ...request.Option) (*ecr.DeleteRepositoryOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DeleteRepository",
		Input:   input,
		Output:  (*ecr.DeleteRepositoryOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.DeleteRepositoryWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.DeleteRepositoryOutput), req.Error
}

func (c *Client) DeleteRepositoryPolicyWithContext(ctx context.Context, input *ecr.DeleteRepositoryPolicyInput, opts ...request.Option) (*ecr.DeleteRepositoryPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DeleteRepositoryPolicy",
		Input:   input,
		Output:  (*ecr.DeleteRepositoryPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.DeleteRepositoryPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.DeleteRepositoryPolicyOutput), req.Error
}

func (c *Client) DescribeImageScanFindingsWithContext(ctx context.Context, input *ecr.DescribeImageScanFindingsInput, opts ...request.Option) (*ecr.DescribeImageScanFindingsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DescribeImageScanFindings",
		Input:   input,
		Output:  (*ecr.DescribeImageScanFindingsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.DescribeImageScanFindingsWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.DescribeImageScanFindingsOutput), req.Error
}

func (c *Client) DescribeImageScanFindingsPagesWithContext(ctx context.Context, input *ecr.DescribeImageScanFindingsInput, cb func(*ecr.DescribeImageScanFindingsOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DescribeImageScanFindings",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ECRAPI.DescribeImageScanFindingsPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeImagesWithContext(ctx context.Context, input *ecr.DescribeImagesInput, opts ...request.Option) (*ecr.DescribeImagesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DescribeImages",
		Input:   input,
		Output:  (*ecr.DescribeImagesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.DescribeImagesWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.DescribeImagesOutput), req.Error
}

func (c *Client) DescribeImagesPagesWithContext(ctx context.Context, input *ecr.DescribeImagesInput, cb func(*ecr.DescribeImagesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DescribeImages",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ECRAPI.DescribeImagesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) DescribeRepositoriesWithContext(ctx context.Context, input *ecr.DescribeRepositoriesInput, opts ...request.Option) (*ecr.DescribeRepositoriesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DescribeRepositories",
		Input:   input,
		Output:  (*ecr.DescribeRepositoriesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.DescribeRepositoriesWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.DescribeRepositoriesOutput), req.Error
}

func (c *Client) DescribeRepositoriesPagesWithContext(ctx context.Context, input *ecr.DescribeRepositoriesInput, cb func(*ecr.DescribeRepositoriesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "DescribeRepositories",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ECRAPI.DescribeRepositoriesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetAuthorizationTokenWithContext(ctx context.Context, input *ecr.GetAuthorizationTokenInput, opts ...request.Option) (*ecr.GetAuthorizationTokenOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "GetAuthorizationToken",
		Input:   input,
		Output:  (*ecr.GetAuthorizationTokenOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.GetAuthorizationTokenWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.GetAuthorizationTokenOutput), req.Error
}

func (c *Client) GetDownloadUrlForLayerWithContext(ctx context.Context, input *ecr.GetDownloadUrlForLayerInput, opts ...request.Option) (*ecr.GetDownloadUrlForLayerOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "GetDownloadUrlForLayer",
		Input:   input,
		Output:  (*ecr.GetDownloadUrlForLayerOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.GetDownloadUrlForLayerWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.GetDownloadUrlForLayerOutput), req.Error
}

func (c *Client) GetLifecyclePolicyWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyInput, opts ...request.Option) (*ecr.GetLifecyclePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "GetLifecyclePolicy",
		Input:   input,
		Output:  (*ecr.GetLifecyclePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.GetLifecyclePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.GetLifecyclePolicyOutput), req.Error
}

func (c *Client) GetLifecyclePolicyPreviewWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyPreviewInput, opts ...request.Option) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "GetLifecyclePolicyPreview",
		Input:   input,
		Output:  (*ecr.GetLifecyclePolicyPreviewOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.GetLifecyclePolicyPreviewWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.GetLifecyclePolicyPreviewOutput), req.Error
}

func (c *Client) GetLifecyclePolicyPreviewPagesWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyPreviewInput, cb func(*ecr.GetLifecyclePolicyPreviewOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "GetLifecyclePolicyPreview",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ECRAPI.GetLifecyclePolicyPreviewPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) GetRepositoryPolicyWithContext(ctx context.Context, input *ecr.GetRepositoryPolicyInput, opts ...request.Option) (*ecr.GetRepositoryPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "GetRepositoryPolicy",
		Input:   input,
		Output:  (*ecr.GetRepositoryPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.GetRepositoryPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.GetRepositoryPolicyOutput), req.Error
}

func (c *Client) InitiateLayerUploadWithContext(ctx context.Context, input *ecr.InitiateLayerUploadInput, opts ...request.Option) (*ecr.InitiateLayerUploadOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "InitiateLayerUpload",
		Input:   input,
		Output:  (*ecr.InitiateLayerUploadOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.InitiateLayerUploadWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.InitiateLayerUploadOutput), req.Error
}

func (c *Client) ListImagesWithContext(ctx context.Context, input *ecr.ListImagesInput, opts ...request.Option) (*ecr.ListImagesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "ListImages",
		Input:   input,
		Output:  (*ecr.ListImagesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.ListImagesWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.ListImagesOutput), req.Error
}

func (c *Client) ListImagesPagesWithContext(ctx context.Context, input *ecr.ListImagesInput, cb func(*ecr.ListImagesOutput, bool) bool, opts ...request.Option) error {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "ListImages",
		Input:   input,
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Error = c.ECRAPI.ListImagesPagesWithContext(ctx, input, cb, opts...)
	})

	return req.Error
}

func (c *Client) ListTagsForResourceWithContext(ctx context.Context, input *ecr.ListTagsForResourceInput, opts ...request.Option) (*ecr.ListTagsForResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "ListTagsForResource",
		Input:   input,
		Output:  (*ecr.ListTagsForResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.ListTagsForResourceOutput), req.Error
}

func (c *Client) PutImageWithContext(ctx context.Context, input *ecr.PutImageInput, opts ...request.Option) (*ecr.PutImageOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "PutImage",
		Input:   input,
		Output:  (*ecr.PutImageOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.PutImageWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.PutImageOutput), req.Error
}

func (c *Client) PutImageScanningConfigurationWithContext(ctx context.Context, input *ecr.PutImageScanningConfigurationInput, opts ...request.Option) (*ecr.PutImageScanningConfigurationOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "PutImageScanningConfiguration",
		Input:   input,
		Output:  (*ecr.PutImageScanningConfigurationOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.PutImageScanningConfigurationWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.PutImageScanningConfigurationOutput), req.Error
}

func (c *Client) PutImageTagMutabilityWithContext(ctx context.Context, input *ecr.PutImageTagMutabilityInput, opts ...request.Option) (*ecr.PutImageTagMutabilityOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "PutImageTagMutability",
		Input:   input,
		Output:  (*ecr.PutImageTagMutabilityOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.PutImageTagMutabilityWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.PutImageTagMutabilityOutput), req.Error
}

func (c *Client) PutLifecyclePolicyWithContext(ctx context.Context, input *ecr.PutLifecyclePolicyInput, opts ...request.Option) (*ecr.PutLifecyclePolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "PutLifecyclePolicy",
		Input:   input,
		Output:  (*ecr.PutLifecyclePolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.PutLifecyclePolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.PutLifecyclePolicyOutput), req.Error
}

func (c *Client) SetRepositoryPolicyWithContext(ctx context.Context, input *ecr.SetRepositoryPolicyInput, opts ...request.Option) (*ecr.SetRepositoryPolicyOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "SetRepositoryPolicy",
		Input:   input,
		Output:  (*ecr.SetRepositoryPolicyOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.SetRepositoryPolicyWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.SetRepositoryPolicyOutput), req.Error
}

func (c *Client) StartImageScanWithContext(ctx context.Context, input *ecr.StartImageScanInput, opts ...request.Option) (*ecr.StartImageScanOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "StartImageScan",
		Input:   input,
		Output:  (*ecr.StartImageScanOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.StartImageScanWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.StartImageScanOutput), req.Error
}

func (c *Client) StartLifecyclePolicyPreviewWithContext(ctx context.Context, input *ecr.StartLifecyclePolicyPreviewInput, opts ...request.Option) (*ecr.StartLifecyclePolicyPreviewOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "StartLifecyclePolicyPreview",
		Input:   input,
		Output:  (*ecr.StartLifecyclePolicyPreviewOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.StartLifecyclePolicyPreviewWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.StartLifecyclePolicyPreviewOutput), req.Error
}

func (c *Client) TagResourceWithContext(ctx context.Context, input *ecr.TagResourceInput, opts ...request.Option) (*ecr.TagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "TagResource",
		Input:   input,
		Output:  (*ecr.TagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.TagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.TagResourceOutput), req.Error
}

func (c *Client) UntagResourceWithContext(ctx context.Context, input *ecr.UntagResourceInput, opts ...request.Option) (*ecr.UntagResourceOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "UntagResource",
		Input:   input,
		Output:  (*ecr.UntagResourceOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.UntagResourceWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.UntagResourceOutput), req.Error
}

func (c *Client) UploadLayerPartWithContext(ctx context.Context, input *ecr.UploadLayerPartInput, opts ...request.Option) (*ecr.UploadLayerPartOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "ecr",
		Action:  "UploadLayerPart",
		Input:   input,
		Output:  (*ecr.UploadLayerPartOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.ECRAPI.UploadLayerPartWithContext(ctx, input, opts...)
	})

	return req.Output.(*ecr.UploadLayerPartOutput), req.Error
}
