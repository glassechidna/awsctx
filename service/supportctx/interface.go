// Code generated by internal/generate/main.go. DO NOT EDIT.

package supportctx

import (
	"context"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/aws/aws-sdk-go/service/support/supportiface"
	"github.com/glassechidna/awsctx"
)

type Support interface {
	AddAttachmentsToSetWithContext(ctx context.Context, input *support.AddAttachmentsToSetInput, opts ...request.Option) (*support.AddAttachmentsToSetOutput, error)
	AddCommunicationToCaseWithContext(ctx context.Context, input *support.AddCommunicationToCaseInput, opts ...request.Option) (*support.AddCommunicationToCaseOutput, error)
	CreateCaseWithContext(ctx context.Context, input *support.CreateCaseInput, opts ...request.Option) (*support.CreateCaseOutput, error)
	DescribeAttachmentWithContext(ctx context.Context, input *support.DescribeAttachmentInput, opts ...request.Option) (*support.DescribeAttachmentOutput, error)
	DescribeCasesWithContext(ctx context.Context, input *support.DescribeCasesInput, opts ...request.Option) (*support.DescribeCasesOutput, error)
	DescribeCommunicationsWithContext(ctx context.Context, input *support.DescribeCommunicationsInput, opts ...request.Option) (*support.DescribeCommunicationsOutput, error)
	DescribeServicesWithContext(ctx context.Context, input *support.DescribeServicesInput, opts ...request.Option) (*support.DescribeServicesOutput, error)
	DescribeSeverityLevelsWithContext(ctx context.Context, input *support.DescribeSeverityLevelsInput, opts ...request.Option) (*support.DescribeSeverityLevelsOutput, error)
	DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)
	DescribeTrustedAdvisorCheckResultWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckResultInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckResultOutput, error)
	DescribeTrustedAdvisorCheckSummariesWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)
	DescribeTrustedAdvisorChecksWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorChecksInput, opts ...request.Option) (*support.DescribeTrustedAdvisorChecksOutput, error)
	RefreshTrustedAdvisorCheckWithContext(ctx context.Context, input *support.RefreshTrustedAdvisorCheckInput, opts ...request.Option) (*support.RefreshTrustedAdvisorCheckOutput, error)
	ResolveCaseWithContext(ctx context.Context, input *support.ResolveCaseInput, opts ...request.Option) (*support.ResolveCaseOutput, error)
}

type Client struct {
	supportiface.SupportAPI
	Contexter awsctx.Contexter
}

func New(base supportiface.SupportAPI, ctxer awsctx.Contexter) Support {
	return &Client{
		SupportAPI: base,
		Contexter: ctxer,
	}
}

var _ Support = (*support.Support)(nil)
var _ Support = (*Client)(nil)

func (c *Client) AddAttachmentsToSetWithContext(ctx context.Context, input *support.AddAttachmentsToSetInput, opts ...request.Option) (*support.AddAttachmentsToSetOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "AddAttachmentsToSetWithContext",
		Input:   input,
		Output:  (*support.AddAttachmentsToSetOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.AddAttachmentsToSetWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.AddAttachmentsToSetOutput), req.Error
}

func (c *Client) AddCommunicationToCaseWithContext(ctx context.Context, input *support.AddCommunicationToCaseInput, opts ...request.Option) (*support.AddCommunicationToCaseOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "AddCommunicationToCaseWithContext",
		Input:   input,
		Output:  (*support.AddCommunicationToCaseOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.AddCommunicationToCaseWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.AddCommunicationToCaseOutput), req.Error
}

func (c *Client) CreateCaseWithContext(ctx context.Context, input *support.CreateCaseInput, opts ...request.Option) (*support.CreateCaseOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "CreateCaseWithContext",
		Input:   input,
		Output:  (*support.CreateCaseOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.CreateCaseWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.CreateCaseOutput), req.Error
}

func (c *Client) DescribeAttachmentWithContext(ctx context.Context, input *support.DescribeAttachmentInput, opts ...request.Option) (*support.DescribeAttachmentOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeAttachmentWithContext",
		Input:   input,
		Output:  (*support.DescribeAttachmentOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeAttachmentWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeAttachmentOutput), req.Error
}

func (c *Client) DescribeCasesWithContext(ctx context.Context, input *support.DescribeCasesInput, opts ...request.Option) (*support.DescribeCasesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeCasesWithContext",
		Input:   input,
		Output:  (*support.DescribeCasesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeCasesWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeCasesOutput), req.Error
}

func (c *Client) DescribeCommunicationsWithContext(ctx context.Context, input *support.DescribeCommunicationsInput, opts ...request.Option) (*support.DescribeCommunicationsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeCommunicationsWithContext",
		Input:   input,
		Output:  (*support.DescribeCommunicationsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeCommunicationsWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeCommunicationsOutput), req.Error
}

func (c *Client) DescribeServicesWithContext(ctx context.Context, input *support.DescribeServicesInput, opts ...request.Option) (*support.DescribeServicesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeServicesWithContext",
		Input:   input,
		Output:  (*support.DescribeServicesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeServicesWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeServicesOutput), req.Error
}

func (c *Client) DescribeSeverityLevelsWithContext(ctx context.Context, input *support.DescribeSeverityLevelsInput, opts ...request.Option) (*support.DescribeSeverityLevelsOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeSeverityLevelsWithContext",
		Input:   input,
		Output:  (*support.DescribeSeverityLevelsOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeSeverityLevelsWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeSeverityLevelsOutput), req.Error
}

func (c *Client) DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeTrustedAdvisorCheckRefreshStatusesWithContext",
		Input:   input,
		Output:  (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput), req.Error
}

func (c *Client) DescribeTrustedAdvisorCheckResultWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckResultInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeTrustedAdvisorCheckResultWithContext",
		Input:   input,
		Output:  (*support.DescribeTrustedAdvisorCheckResultOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeTrustedAdvisorCheckResultWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeTrustedAdvisorCheckResultOutput), req.Error
}

func (c *Client) DescribeTrustedAdvisorCheckSummariesWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeTrustedAdvisorCheckSummariesWithContext",
		Input:   input,
		Output:  (*support.DescribeTrustedAdvisorCheckSummariesOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeTrustedAdvisorCheckSummariesWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeTrustedAdvisorCheckSummariesOutput), req.Error
}

func (c *Client) DescribeTrustedAdvisorChecksWithContext(ctx context.Context, input *support.DescribeTrustedAdvisorChecksInput, opts ...request.Option) (*support.DescribeTrustedAdvisorChecksOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "DescribeTrustedAdvisorChecksWithContext",
		Input:   input,
		Output:  (*support.DescribeTrustedAdvisorChecksOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.DescribeTrustedAdvisorChecksWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.DescribeTrustedAdvisorChecksOutput), req.Error
}

func (c *Client) RefreshTrustedAdvisorCheckWithContext(ctx context.Context, input *support.RefreshTrustedAdvisorCheckInput, opts ...request.Option) (*support.RefreshTrustedAdvisorCheckOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "RefreshTrustedAdvisorCheckWithContext",
		Input:   input,
		Output:  (*support.RefreshTrustedAdvisorCheckOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.RefreshTrustedAdvisorCheckWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.RefreshTrustedAdvisorCheckOutput), req.Error
}

func (c *Client) ResolveCaseWithContext(ctx context.Context, input *support.ResolveCaseInput, opts ...request.Option) (*support.ResolveCaseOutput, error) {
	req := &awsctx.AwsRequest{
		Service: "support",
		Action:  "ResolveCaseWithContext",
		Input:   input,
		Output:  (*support.ResolveCaseOutput)(nil),
		Error:   nil,
	}

	ctxer := c.Contexter
	if ctxer == nil {
		ctxer = awsctx.NoopContexter
	}

	ctxer.WrapContext(ctx, req, func(ctx context.Context) {
		req.Output, req.Error = c.SupportAPI.ResolveCaseWithContext(ctx, input, opts...)
	})

	return req.Output.(*support.ResolveCaseOutput), req.Error
}
